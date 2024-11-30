package biz

// {
// 	"blockHash": "0xcda2ccb0fcd9193d263354db819f0c100b6cd5317fa86c978e9716c823dba171",
// 	"logIndex": 1,
// 	"address": "0xD60d331B1999824C84A0A702D07368Fde493dF94",
// 	"data": "0x000000000000000000000000000000000000000000000000002bb2c8eabcc00000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000014",
// 	"removed": false,
// 	"topics": [
// 	  "0xb9f26d3c47b1e6cc5b8e24d075d969f3a79727a118e23cf23cdec7d26f74c682",
// 	  "0x000000000000000000000000186b4735e114d2666ea2cad0ec3b30ac7b386447",
// 	  "0x000000000000000000000000f3981b00662d7c43c805ce0ed65b521893b9c3e6",
// 	  "0x0000000000000000000000000000000000000000000000000000000000000001"
// 	],
// 	"blockNumber": 331526358,
// 	"transactionIndex": 0,
// 	"transactionHash": "0xd4d0aadf631cc3402f351e72e2264396ee201a9c9cbebea6363d44a01de20429",
// 	"timestamp": 1732530511000
//   }

import (
	"blind-box-svr/common/constants"
	nftmarket "blind-box-svr/contracts"
	"blind-box-svr/dto/params"
	"blind-box-svr/global"
	"blind-box-svr/model"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"

	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
)

type JSONData struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results []*BlockLog `json:"results"`
}

type BlockLog struct {
	Address common.Address `json:"address" gencodec:"required"`
	Topics  []common.Hash  `json:"topics" gencodec:"required"`
	// Data        []byte         `json:"data" gencodec:"required"`
	Data        string      `json:"data" gencodec:"required"`
	BlockNumber uint64      `json:"blockNumber" rlp:"-"`
	TxHash      common.Hash `json:"transactionHash" gencodec:"required" rlp:"-"`
	TxIndex     uint        `json:"transactionIndex" rlp:"-"`
	BlockHash   common.Hash `json:"blockHash" rlp:"-"`
	Index       uint        `json:"logIndex" rlp:"-"`
	Removed     bool        `json:"removed" rlp:"-"`
	Timestamp   uint64      `json:"timestamp"`
}

type TelosScanBiz struct {
	eventDataProcessor map[string]func(*BlockLog, abi.ABI) error
	contractAbi        abi.ABI
	apiClient          *resty.Client
}

var logListSig string = calculateTopicHash("List(address,address,uint256,uint256,uint256,uint256)")
var logPurchaseSig string = calculateTopicHash("Purchase(address,address,uint256,uint256,uint256)")
var logRevokeSig string = calculateTopicHash("Revoke(address,address,uint256)")
var logUpdatePriceSig string = calculateTopicHash("UpdatePrice(address,address,uint256,uint256)")
var logUpdateAmountSig string = calculateTopicHash("UpdateAmount(address,address,uint256,uint256)")

const contractAddrStr string = "0xD60d331B1999824C84A0A702D07368Fde493dF94"
const rpcUrl string = "https://rpc.testnet.telos.net"
const apiUrl string = "https://api.testnet.teloscan.io"
const fetchApiIntervalSec int64 = 10

var telosScanBiz *TelosScanBiz

func GetTelosScanBiz() *TelosScanBiz {
	return telosScanBiz
}

// var rpcClient          *ethclient.Client
var nftMarket *nftmarket.Nftmarket

func initRpcClient() {
	rpcClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Errorf("ethclient.Dial err: %v", err)
		panic(err)
	}
	address := common.HexToAddress(contractAddrStr)
	nftMarket, err = nftmarket.NewNftmarket(address, rpcClient)
	if err != nil {
		log.Errorf("NewNftmarket err: %v ", err)
		panic(err)
	}
}

func hashToAddress(bytes []byte) common.Address {
	return common.Address(bytes[common.HashLength-common.AddressLength:])
}

func NewTelosScanBiz() {
	telosScanBiz = &TelosScanBiz{}
	abiABI, err := abi.JSON(strings.NewReader(nftmarket.NftmarketABI))
	if err != nil {
		log.Errorf("parse abi error %v", err)
		panic(err)
	}
	initRpcClient()
	telosScanBiz.eventDataProcessor = make(map[string]func(*BlockLog, abi.ABI) error)
	telosScanBiz.contractAbi = abiABI
	telosScanBiz.apiClient = resty.New()
	telosScanBiz.eventDataProcessor[logListSig] = processEventListFunc
	telosScanBiz.eventDataProcessor[logPurchaseSig] = processEventPurchaseFunc
	telosScanBiz.eventDataProcessor[logRevokeSig] = processEventRevokeFunc
	telosScanBiz.eventDataProcessor[logUpdatePriceSig] = processEventUpdatePriceFunc
	telosScanBiz.eventDataProcessor[logUpdateAmountSig] = processEventUpdateAmountFunc
	telosScanBiz.startProcessEvents()
}

func (b *TelosScanBiz) startProcessEvents() {
	ticker := time.NewTicker(time.Duration(fetchApiIntervalSec) * time.Second)
	go func() {
		for {
			t := <-ticker.C
			log.Infof("ticker invoke time: %v", t)
			// todo query last index
			if true {
				index := int64(0)
				b.batchProcessLog(index)

			}
		}
	}()

}

func (b *TelosScanBiz) batchProcessLog(index int64) {
	logs, err := b.requestContractEvents(index)
	if err != nil {
		log.Errorf("requestContractEvents err: %v", err)
	}
	for _, eventLog := range logs {
		b.processLog(eventLog)
	}
}

func (b *TelosScanBiz) requestContractEvents(offset int64) ([]*BlockLog, error) {

	jsonData := JSONData{}
	req := b.apiClient.R().SetPathParam("address", contractAddrStr).SetQueryParams(map[string]string{
		"offset": strconv.FormatInt(offset, 10),
		"limit":  "50",
		"sort":   "ASC",
	}).SetHeader("Accept", "application/json").SetResult(&jsonData)
	resp, err := req.Get(apiUrl + "/v1/contract/{address}/logs")
	if err != nil {
		log.Errorf("request err: %v", err)
		return nil, err
	}
	// log.Infof("request req: %v resp: %v", req, resp)
	if resp.IsSuccess() {
		if jsonData.Success {
			return jsonData.Results, nil
		} else {
			log.Errorf("telos api return failed, code: %v, msg: %v", jsonData.Code, jsonData.Message)
			return nil, err
		}
	} else {
		log.Errorf("request error, resp: %v, req: %v", resp, req)
		return nil, err
	}
}

func calculateTopicHash(eventAbi string) string {
	// eventSignatureBytes := []byte("Transfer(address,address,uint256)")
	return crypto.Keccak256Hash([]byte(eventAbi)).Hex()
}

func (b *TelosScanBiz) processLog(bl *BlockLog) {
	log.Infof("preparing process log: %v", bl)
	topics := bl.Topics
	if len(topics) == 0 {
		return
	}
	eventSigStr := topics[0].Hex()
	processor, ok := b.eventDataProcessor[eventSigStr]
	log.Infof("eventSigStr: %v, whether got processor: %v", eventSigStr, ok)
	if processor != nil {
		if err := processor(bl, b.contractAbi); err != nil {
			log.Errorf("processor %v return err: %v", eventSigStr, err)
		}
	}
}

func (b *TelosScanBiz) GetOrderList(param *params.OrderParam) ([]*model.OrderTab, error) {
	otm := model.OrderTabMgr(global.GetDB())
	gdb := otm.Where("nft_addr = ?", param.NftAddr)
	if len(param.Seller) > 0 {
		gdb.Where("seller = ?", param.Seller)
	}
	if param.NftType > 0 && param.MaxCardTokenId > 0 {
		if param.NftType == 1 {
			gdb.Where("token_id <= ?", param.MaxCardTokenId)
		} else if param.NftType == 2 {
			gdb.Where("token_id > ?", param.MaxCardTokenId)
		}
	}
	gdb.Where("amount > 0").Where(fmt.Sprintf("state = %v", constants.ACTIVE))
	limit := param.Limit
	if limit == 0 || limit > 50 {
		limit = 50
	}
	var orders []*model.OrderTab
	err := gdb.Limit(int(limit)).Offset(int(param.Offset)).Order("order_id desc").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

var processEventListFunc = func(bl *BlockLog, contractAbi abi.ABI) error {
	nftMarketList := nftmarket.NftmarketList{}
	//event List(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 price, uint256 amount, uint256 tokenId)
	topics := bl.Topics
	nftMarketList.Seller = hashToAddress(topics[1].Bytes())
	nftMarketList.NftAddr = hashToAddress(topics[2].Bytes())
	nftMarketList.OrderId = new(big.Int).SetBytes(topics[3].Bytes())
	err := contractAbi.UnpackIntoInterface(&nftMarketList, "List", common.FromHex(bl.Data))
	if err != nil {
		log.Errorf("processEventListFunc call UnpackIntoInterface err: %v", err)
		return err
	}
	log.Infof("nftMarketList: %v", nftMarketList)
	order := &model.OrderTab{
		OrderID:         nftMarketList.OrderId.Int64(),
		NftAddr:         nftMarketList.NftAddr.Hex(),
		Seller:          nftMarketList.Seller.Hex(),
		TokenID:         nftMarketList.TokenId.Int64(),
		Price:           float64(nftMarketList.Price.Int64()),
		Amount:          nftMarketList.Amount.Int64(),
		State:           constants.ACTIVE,
		CreateTimestamp: int64(bl.Timestamp),
		UpdateTimestamp: 0,
		CreateTime:      time.Now(),
		// UpdateTime:      nil,
	}
	err = model.OrderTabMgr(global.GetDB()).Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

var processEventPurchaseFunc = func(bl *BlockLog, contractAbi abi.ABI) error {
	nftMarketPurchase := nftmarket.NftmarketPurchase{}
	//event Purchase(address indexed buyer, address indexed nftAddr, uint256 indexed orderId, uint256 price, uint256 amount)
	topics := bl.Topics
	nftMarketPurchase.Buyer = hashToAddress(topics[1].Bytes())
	nftMarketPurchase.NftAddr = hashToAddress(topics[2].Bytes())
	nftMarketPurchase.OrderId = new(big.Int).SetBytes(topics[3].Bytes())
	err := contractAbi.UnpackIntoInterface(&nftMarketPurchase, "Purchase", common.FromHex(bl.Data))
	if err != nil {
		log.Errorf("processEventPurchaseFunc call UnpackIntoInterface err: %v", err)
		return err
	}
	log.Infof("nftMarketPurchase: %v", nftMarketPurchase)
	otm := model.OrderTabMgr(global.GetDB())

	// orderTab, err := otm.GetFromOrderID(nftMarketPurchase.OrderId.Int64())
	// if err != nil {
	// 	return err
	// }
	// if orderTab.OrderID == 0 {
	// 	return errors.New("orderId is 0")
	// }
	nftMarketOrder, err := nftMarket.Get(nil, nftMarketPurchase.OrderId)
	if err != nil {
		return err
	}
	cols := model.OrderTabColumns
	err = otm.Select(cols.Amount, cols.UpdateTimestamp, cols.UpdateTime).Where("order_id = ?", nftMarketPurchase.OrderId.Int64()).
		Updates(model.OrderTab{
			Amount:          nftMarketOrder.Amount.Int64(),
			UpdateTimestamp: int64(bl.Timestamp),
			UpdateTime:      time.Now(),
		}).Error
	if err != nil {
		return err
	}
	return nil
}

var processEventRevokeFunc = func(bl *BlockLog, contractAbi abi.ABI) error {
	nftMarketRevoke := nftmarket.NftmarketRevoke{}
	//event Revoke(address indexed seller, address indexed nftAddr, uint256 indexed orderId)
	topics := bl.Topics
	nftMarketRevoke.Seller = hashToAddress(topics[1].Bytes())
	nftMarketRevoke.NftAddr = hashToAddress(topics[2].Bytes())
	nftMarketRevoke.OrderId = new(big.Int).SetBytes(topics[3].Bytes())
	log.Infof("nftMarketRevoke: %v", nftMarketRevoke)
	cols := model.OrderTabColumns
	err := model.OrderTabMgr(global.GetDB()).Where("order_id = ?", nftMarketRevoke.OrderId.Int64()).
		Select(cols.State, cols.UpdateTimestamp, cols.UpdateTime).Updates(model.OrderTab{
		State:           constants.INACTIVE,
		UpdateTimestamp: int64(bl.Timestamp),
		UpdateTime:      time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

var processEventUpdatePriceFunc = func(bl *BlockLog, contractAbi abi.ABI) error {
	nftMarketUpdatePrice := nftmarket.NftmarketUpdatePrice{}
	// event UpdatePrice(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 newPrice);
	topics := bl.Topics
	nftMarketUpdatePrice.Seller = hashToAddress(topics[1].Bytes())
	nftMarketUpdatePrice.NftAddr = hashToAddress(topics[2].Bytes())
	nftMarketUpdatePrice.OrderId = new(big.Int).SetBytes(topics[3].Bytes())
	err := contractAbi.UnpackIntoInterface(&nftMarketUpdatePrice, "UpdatePrice", common.FromHex(bl.Data))
	if err != nil {
		log.Errorf("processEventUpdatePriceFunc call UnpackIntoInterface err: %v", err)
		return err
	}
	log.Infof("nftMarketUpdatePrice: %v", nftMarketUpdatePrice)
	cols := model.OrderTabColumns
	err = model.OrderTabMgr(global.GetDB()).Where("order_id = ?", nftMarketUpdatePrice.OrderId.Int64()).
		Select(cols.Price, cols.UpdateTimestamp, cols.UpdateTime).Updates(model.OrderTab{
		Price:           float64(nftMarketUpdatePrice.NewPrice.Int64()),
		UpdateTimestamp: int64(bl.Timestamp),
		UpdateTime:      time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

var processEventUpdateAmountFunc = func(bl *BlockLog, contractAbi abi.ABI) error {
	nftMarketUpdateAmount := nftmarket.NftmarketUpdateAmount{}
	// event UpdateAmount(address indexed seller, address indexed nftAddr, uint256 indexed orderId, uint256 newAmount);
	topics := bl.Topics
	nftMarketUpdateAmount.Seller = hashToAddress(topics[1].Bytes())
	nftMarketUpdateAmount.NftAddr = hashToAddress(topics[2].Bytes())
	nftMarketUpdateAmount.OrderId = new(big.Int).SetBytes(topics[3].Bytes())
	err := contractAbi.UnpackIntoInterface(&nftMarketUpdateAmount, "UpdateAmount", common.FromHex(bl.Data))
	if err != nil {
		log.Errorf("processEventUpdateAmountFunc call UnpackIntoInterface err: %v", err)
		return err
	}
	log.Infof("nftMarketUpdateAmount: %v", nftMarketUpdateAmount)
	cols := model.OrderTabColumns
	err = model.OrderTabMgr(global.GetDB()).Where("order_id = ?", nftMarketUpdateAmount.OrderId.Int64()).
		Select(cols.Price, cols.UpdateTimestamp, cols.UpdateTime).Updates(model.OrderTab{
		Amount:          nftMarketUpdateAmount.NewAmount.Int64(),
		UpdateTimestamp: int64(bl.Timestamp),
		UpdateTime:      time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}
