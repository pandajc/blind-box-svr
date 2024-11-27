package biz

import (
	nftmarket "blind-box-svr/contracts"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

var jsonStr string = `
	{
		"code": 200,
		"success": true,
		"message": "OK",
		"contracts": {
			"0xD60d331B1999824C84A0A702D07368Fde493dF94": {
			  "creator": "0x186b4735E114d2666Ea2CAD0Ec3B30ac7b386447",
			  "metadata": "{\"output\": {\"abi\": [{\"name\": \"TokenIdOrderExists\", \"type\": \"error\", \"inputs\": [{\"name\": \"tokenId\", \"type\": \"uint256\", \"internalType\": \"uint256\"}]}, {\"name\": \"List\", \"type\": \"event\", \"inputs\": [{\"name\": \"seller\", \"type\": \"address\", \"indexed\": true, \"internalType\": \"address\"}, {\"name\": \"nftAddr\", \"type\": \"address\", \"indexed\": true, \"internalType\": \"address\"}, {\"name\": \"orderId\", \"type\": \"uint256\", \"indexed\": true, \"internalType\": \"uint256\"}, {\"name\": \"price\", \"type\": \"uint256\", \"indexed\": false, \"internalType\": \"uint256\"}, {\"name\": \"amount\", \"type\": \"uint256\", \"indexed\": false, \"internalType\": \"uint256\"}, {\"name\": \"tokenId\", \"type\": \"uint256\", \"indexed\": false, \"internalType\": \"uint256\"}], \"anonymous\": false}, {\"name\": \"Purchase\", \"type\": \"event\", \"inputs\": [{\"name\": \"buyer\", \"type\": \"address\", \"indexed\": true, \"internalType\": \"address\"}, {\"name\": \"nftAddr\", \"type\": \"address\", \"indexed\": true, \"internalType\": \"address\"}, {\"name\": \"orderId\", \"type\": \"uint256\", \"indexed\": true, \"internalType\": \"uint256\"}, {\"name\": \"price\", \"type\": \"uint256\", \"indexed\": false, \"internalType\": \"uint256\"}, {\"name\": \"amount\", \"type\": \"uint256\", \"indexed\": false, \"internalType\": \"uint256\"}], \"anonymous\": false}, {\"name\": \"Revoke\", \"type\": \"event\", \"inputs\": [{\"name\": \"seller\", \"type\": \"address\", \"indexed\": true, \"internalType\": \"address\"}, {\"name\": \"nftAddr\", \"type\": \"address\", \"indexed\": true, \"internalType\": \"address\"}, {\"name\": \"orderId\", \"type\": \"uint256\", \"indexed\": true, \"internalType\": \"uint256\"}], \"anonymous\": false}, {\"name\": \"UpdateAmount\", \"type\": \"event\", \"inputs\": [{\"name\": \"seller\", \"type\": \"address\", \"indexed\": true, \"internalType\": \"address\"}, {\"name\": \"nftAddr\", \"type\": \"address\", \"indexed\": true, \"internalType\": \"address\"}, {\"name\": \"orderId\", \"type\": \"uint256\", \"indexed\": true, \"internalType\": \"uint256\"}, {\"name\": \"newAmount\", \"type\": \"uint256\", \"indexed\": false, \"internalType\": \"uint256\"}], \"anonymous\": false}, {\"name\": \"UpdatePrice\", \"type\": \"event\", \"inputs\": [{\"name\": \"seller\", \"type\": \"address\", \"indexed\": true, \"internalType\": \"address\"}, {\"name\": \"nftAddr\", \"type\": \"address\", \"indexed\": true, \"internalType\": \"address\"}, {\"name\": \"orderId\", \"type\": \"uint256\", \"indexed\": true, \"internalType\": \"uint256\"}, {\"name\": \"newPrice\", \"type\": \"uint256\", \"indexed\": false, \"internalType\": \"uint256\"}], \"anonymous\": false}, {\"type\": \"fallback\", \"stateMutability\": \"payable\"}, {\"name\": \"get\", \"type\": \"function\", \"inputs\": [{\"name\": \"_tokenId\", \"type\": \"uint256\", \"internalType\": \"uint256\"}], \"outputs\": [{\"name\": \"\", \"type\": \"tuple\", \"components\": [{\"name\": \"token\", \"type\": \"address\", \"internalType\": \"address\"}, {\"name\": \"tokenId\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"owner\", \"type\": \"address\", \"internalType\": \"address\"}, {\"name\": \"price\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"amount\", \"type\": \"uint256\", \"internalType\": \"uint256\"}], \"internalType\": \"struct NFTMarket.Order\"}], \"stateMutability\": \"view\"}, {\"name\": \"list\", \"type\": \"function\", \"inputs\": [{\"name\": \"_nftAddr\", \"type\": \"address\", \"internalType\": \"address\"}, {\"name\": \"_tokenId\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"_price\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"_amount\", \"type\": \"uint256\", \"internalType\": \"uint256\"}], \"outputs\": [{\"name\": \"\", \"type\": \"uint256\", \"internalType\": \"uint256\"}], \"stateMutability\": \"nonpayable\"}, {\"name\": \"onERC1155BatchReceived\", \"type\": \"function\", \"inputs\": [{\"name\": \"\", \"type\": \"address\", \"internalType\": \"address\"}, {\"name\": \"\", \"type\": \"address\", \"internalType\": \"address\"}, {\"name\": \"\", \"type\": \"uint256[]\", \"internalType\": \"uint256[]\"}, {\"name\": \"\", \"type\": \"uint256[]\", \"internalType\": \"uint256[]\"}, {\"name\": \"\", \"type\": \"bytes\", \"internalType\": \"bytes\"}], \"outputs\": [{\"name\": \"\", \"type\": \"bytes4\", \"internalType\": \"bytes4\"}], \"stateMutability\": \"nonpayable\"}, {\"name\": \"onERC1155Received\", \"type\": \"function\", \"inputs\": [{\"name\": \"\", \"type\": \"address\", \"internalType\": \"address\"}, {\"name\": \"\", \"type\": \"address\", \"internalType\": \"address\"}, {\"name\": \"\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"\", \"type\": \"bytes\", \"internalType\": \"bytes\"}], \"outputs\": [{\"name\": \"\", \"type\": \"bytes4\", \"internalType\": \"bytes4\"}], \"stateMutability\": \"nonpayable\"}, {\"name\": \"orders\", \"type\": \"function\", \"inputs\": [{\"name\": \"\", \"type\": \"uint256\", \"internalType\": \"uint256\"}], \"outputs\": [{\"name\": \"token\", \"type\": \"address\", \"internalType\": \"address\"}, {\"name\": \"tokenId\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"owner\", \"type\": \"address\", \"internalType\": \"address\"}, {\"name\": \"price\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"amount\", \"type\": \"uint256\", \"internalType\": \"uint256\"}], \"stateMutability\": \"view\"}, {\"name\": \"purchase\", \"type\": \"function\", \"inputs\": [{\"name\": \"_orderId\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"_amount\", \"type\": \"uint256\", \"internalType\": \"uint256\"}], \"outputs\": [], \"stateMutability\": \"payable\"}, {\"name\": \"revoke\", \"type\": \"function\", \"inputs\": [{\"name\": \"_orderId\", \"type\": \"uint256\", \"internalType\": \"uint256\"}], \"outputs\": [], \"stateMutability\": \"nonpayable\"}, {\"name\": \"supportsInterface\", \"type\": \"function\", \"inputs\": [{\"name\": \"interfaceId\", \"type\": \"bytes4\", \"internalType\": \"bytes4\"}], \"outputs\": [{\"name\": \"\", \"type\": \"bool\", \"internalType\": \"bool\"}], \"stateMutability\": \"view\"}, {\"name\": \"updateAmount\", \"type\": \"function\", \"inputs\": [{\"name\": \"_orderId\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"_amount\", \"type\": \"uint256\", \"internalType\": \"uint256\"}], \"outputs\": [], \"stateMutability\": \"nonpayable\"}, {\"name\": \"updatePrice\", \"type\": \"function\", \"inputs\": [{\"name\": \"_orderId\", \"type\": \"uint256\", \"internalType\": \"uint256\"}, {\"name\": \"_newPrice\", \"type\": \"uint256\", \"internalType\": \"uint256\"}], \"outputs\": [], \"stateMutability\": \"nonpayable\"}, {\"type\": \"receive\", \"stateMutability\": \"payable\"}], \"devdoc\": {\"kind\": \"dev\", \"methods\": {\"supportsInterface(bytes4)\": {\"details\": \"See {IERC165-supportsInterface}.\"}}, \"version\": 1}, \"userdoc\": {\"kind\": \"user\", \"methods\": {}, \"version\": 1}}, \"sources\": {\"src/NFTMarket.sol\": {\"urls\": [\"bzz-raw://0f44556eb93e6c18dc14a5d4f81a4f7ec726a3b12089111ae0a2ed4b50ae3f37\", \"dweb:/ipfs/Qmb4YCRRPS7bpc3HUE3vUvx7uGoA6RqHB89PgdjE98982K\"], \"license\": \"MIT\", \"keccak256\": \"0x5475ab79eea69dda04515d64d286f4e6efe507e2823eca505dd449099d0be58f\"}, \"lib/openzeppelin-contracts/contracts/token/ERC1155/IERC1155.sol\": {\"urls\": [\"bzz-raw://7ce608c19d5e917c60f9c8aa3e5f0eb05b326280ac0a235e8bb9a848a3a64a91\", \"dweb:/ipfs/QmdLPsWQJj7JvRae8MM13GEo4PBXaEFmD4b4heqcyMJNPG\"], \"license\": \"MIT\", \"keccak256\": \"0x68d6fdbeb467192c3627a46aa7bf5cbb73267363b740abc511f521a5a41a446e\"}, \"lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol\": {\"urls\": [\"bzz-raw://8084aa71a4cc7d2980972412a88fe4f114869faea3fefa5436431644eb5c0287\", \"dweb:/ipfs/Qmbqfs5dRdPvHVKY8kTaeyc65NdqXRQwRK7h9s5UJEhD1p\"], \"license\": \"MIT\", \"keccak256\": \"0xddce8e17e3d3f9ed818b4f4c4478a8262aab8b11ed322f1bf5ed705bb4bd97fa\"}, \"lib/openzeppelin-contracts/contracts/utils/introspection/IERC165.sol\": {\"urls\": [\"bzz-raw://f6fda447a62815e8064f47eff0dd1cf58d9207ad69b5d32280f8d7ed1d1e4621\", \"dweb:/ipfs/QmfDRc7pxfaXB2Dh9np5Uf29Na3pQ7tafRS684wd3GLjVL\"], \"license\": \"MIT\", \"keccak256\": \"0x79796192ec90263f21b464d5bc90b777a525971d3de8232be80d9c4f9fb353b8\"}, \"lib/openzeppelin-contracts/contracts/token/ERC1155/IERC1155Receiver.sol\": {\"urls\": [\"bzz-raw://d8cbb06152d82ebdd5ba1d33454e5759492040f309a82637c7e99c948a04fa20\", \"dweb:/ipfs/QmQQuLr6WSfLu97pMEh6XLefk99TSj9k5Qu1zXGPepwGiK\"], \"license\": \"MIT\", \"keccak256\": \"0x61a23d601c2ab69dd726ac55058604cbda98e1d728ba31a51c379a3f9eeea715\"}, \"lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Holder.sol\": {\"urls\": [\"bzz-raw://6cf8cc5d07cf8003255f9d766fe8188b9f6e33b6240e126a605f0d061566b23e\", \"dweb:/ipfs/Qmd7okDCSoUt1L4G9hmb5c4W8kWUnfpAa2jyBKUp4xKErd\"], \"license\": \"MIT\", \"keccak256\": \"0xe103e95f854ef0cd1bba5f469175f67cd332f5c2561941f165e3dd65cee94d6d\"}}, \"version\": 1, \"compiler\": {\"version\": \"0.8.28+commit.7893614a\"}, \"language\": \"Solidity\", \"settings\": {\"metadata\": {\"bytecodeHash\": \"ipfs\"}, \"libraries\": {}, \"optimizer\": {\"runs\": 200, \"enabled\": true}, \"evmVersion\": \"paris\", \"remappings\": [\":@openzeppelin/contracts/=lib/openzeppelin-contracts/contracts/\", \":ds-test/=lib/solmate/lib/ds-test/src/\", \":erc4626-tests/=lib/openzeppelin-contracts/lib/erc4626-tests/\", \":forge-std/=lib/forge-std/src/\", \":halmos-cheatcodes/=lib/openzeppelin-contracts/lib/halmos-cheatcodes/src/\", \":openzeppelin-contracts/=lib/openzeppelin-contracts/\", \":solmate/=lib/solmate/src/\"], \"compilationTarget\": {\"src/NFTMarket.sol\": \"NFTMarket\"}}}",
			  "address": "0xD60d331B1999824C84A0A702D07368Fde493dF94",
			  "calldata": "{}",
			  "fromTrace": false,
			  "verified": true,
			  "block": 331511156,
			  "supportedInterfaces": [
				"none"
			  ],
			  "transaction": "0x25ee356998c2952eb5f7d042c0cd17fbd545c8a9497fddf0845b41f6246d8b81",
			  "timestamp": 1732522910000
			}
		  },
		"results": [
		  {
			"blockHash": "0xcda2ccb0fcd9193d263354db819f0c100b6cd5317fa86c978e9716c823dba171",
			"logIndex": 1,
			"address": "0xD60d331B1999824C84A0A702D07368Fde493dF94",
			"data": "0x000000000000000000000000000000000000000000000000002bb2c8eabcc00000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000014",
			"removed": false,
			"topics": [
			  "0xb9f26d3c47b1e6cc5b8e24d075d969f3a79727a118e23cf23cdec7d26f74c682",
			  "0x000000000000000000000000186b4735e114d2666ea2cad0ec3b30ac7b386447",
			  "0x000000000000000000000000f3981b00662d7c43c805ce0ed65b521893b9c3e6",
			  "0x0000000000000000000000000000000000000000000000000000000000000001"
			],
			"blockNumber": 331526358,
			"transactionIndex": 0,
			"transactionHash": "0xd4d0aadf631cc3402f351e72e2264396ee201a9c9cbebea6363d44a01de20429",
			"timestamp": 1732530511000
		  }
		]
	  }
   `

func TestUnpackIntoInterface(t *testing.T) {

	fmt.Println(jsonStr)
	var data JSONData
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		t.Error(err)
		return
	}
	NewTelosScanBiz()
	fmt.Println(data)
	nftMarketList := &nftmarket.NftmarketList{}
	err = telosScanBiz.contractAbi.UnpackIntoInterface(&nftMarketList, "List", common.FromHex(data.Results[0].Data))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(nftMarketList)
}

func TestSubClient(t *testing.T) {
	// need wss protocal
	client, err := ethclient.Dial("wss://")
	if err != nil {
		log.Error(err)
	}
	contractAddress := common.HexToAddress(contractAddrStr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Errorf("call SubscribeFilterLogs err: %v", err)
		return
	}
	for {
		select {
		case err := <-sub.Err():
			log.Errorf("sub err: %v", err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}

func TestCallContract(t *testing.T) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Error(err)
	}
	address := common.HexToAddress(contractAddrStr)
	instantce, err := nftmarket.NewNftmarket(address, client)
	if err != nil {
		log.Errorf("NewNftmarket err: %v ", err)
		return
	}
	order, err := instantce.Get(nil, big.NewInt(1))
	if err != nil {
		log.Errorf("call Get err: %v", err)
		return
	}
	fmt.Println(order)

}

func TestHttpUtil(t *testing.T) {
	client := resty.New()
	var jsonData JSONData
	resp, err := client.R().SetPathParam("address", contractAddrStr).SetQueryParams(map[string]string{
		"offset": "0",
		"limit":  "10",
		"sort":   "ASC",
	}).SetHeader("Accept", "application/json").SetResult(&jsonData).
		Get(apiUrl + "/v1/contract/{address}/logs")
	if err != nil {
		log.Errorf("request err: %v", err)
		return
	}
	fmt.Println(resp)
	fmt.Println(jsonData)

}

func TestTicker(t *testing.T) {

	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			//从定时器中获取数据
			t := <-ticker.C
			time.Sleep(2 * time.Second)
			fmt.Println("当前时间为:", t)
		}
	}()

	// for {
	// 	time.Sleep(time.Second * 1)
	// }

}

func TestTelosScanBiz_startProcessEvents(t *testing.T) {
	NewTelosScanBiz()
	telosScanBiz.startProcessEvents()
	// for {
	// 	time.Sleep(time.Second * 1)
	// }
}
