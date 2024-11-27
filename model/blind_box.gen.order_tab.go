package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _OrderTabMgr struct {
	*_BaseMgr
}

// OrderTabMgr open func
func OrderTabMgr(db *gorm.DB) *_OrderTabMgr {
	if db == nil {
		panic(fmt.Errorf("OrderTabMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderTabMgr{_BaseMgr: &_BaseMgr{DB: db.Table("order_tab"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderTabMgr) GetTableName() string {
	return "order_tab"
}

// Reset 重置gorm会话
func (obj *_OrderTabMgr) Reset() *_OrderTabMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_OrderTabMgr) Get() (result OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderTabMgr) Gets() (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_OrderTabMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithOrderID order_id获取
func (obj *_OrderTabMgr) WithOrderID(orderID int64) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithSeller seller获取
func (obj *_OrderTabMgr) WithSeller(seller string) Option {
	return optionFunc(func(o *options) { o.query["seller"] = seller })
}

// WithNftAddr nft_addr获取
func (obj *_OrderTabMgr) WithNftAddr(nftAddr string) Option {
	return optionFunc(func(o *options) { o.query["nft_addr"] = nftAddr })
}

// WithTokenID token_id获取
func (obj *_OrderTabMgr) WithTokenID(tokenID int64) Option {
	return optionFunc(func(o *options) { o.query["token_id"] = tokenID })
}

// WithPrice price获取
func (obj *_OrderTabMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithAmount amount获取
func (obj *_OrderTabMgr) WithAmount(amount int64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithState state获取 1-active 2-inactive
func (obj *_OrderTabMgr) WithState(state int8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithCreateTimestamp create_timestamp获取 row created block timestamp
func (obj *_OrderTabMgr) WithCreateTimestamp(createTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["create_timestamp"] = createTimestamp })
}

// WithUpdateTimestamp update_timestamp获取 row updated block timestamp
func (obj *_OrderTabMgr) WithUpdateTimestamp(updateTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["update_timestamp"] = updateTimestamp })
}

// WithCreateTime create_time获取 db row create time
func (obj *_OrderTabMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 db row update time
func (obj *_OrderTabMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_OrderTabMgr) GetByOption(opts ...Option) (result OrderTab, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_OrderTabMgr) GetByOptions(opts ...Option) (results []*OrderTab, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_OrderTabMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]OrderTab, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromOrderID 通过order_id获取内容
func (obj *_OrderTabMgr) GetFromOrderID(orderID int64) (result OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// GetBatchFromOrderID 批量查找
func (obj *_OrderTabMgr) GetBatchFromOrderID(orderIDs []int64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromSeller 通过seller获取内容
func (obj *_OrderTabMgr) GetFromSeller(seller string) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`seller` = ?", seller).Find(&results).Error

	return
}

// GetBatchFromSeller 批量查找
func (obj *_OrderTabMgr) GetBatchFromSeller(sellers []string) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`seller` IN (?)", sellers).Find(&results).Error

	return
}

// GetFromNftAddr 通过nft_addr获取内容
func (obj *_OrderTabMgr) GetFromNftAddr(nftAddr string) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`nft_addr` = ?", nftAddr).Find(&results).Error

	return
}

// GetBatchFromNftAddr 批量查找
func (obj *_OrderTabMgr) GetBatchFromNftAddr(nftAddrs []string) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`nft_addr` IN (?)", nftAddrs).Find(&results).Error

	return
}

// GetFromTokenID 通过token_id获取内容
func (obj *_OrderTabMgr) GetFromTokenID(tokenID int64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`token_id` = ?", tokenID).Find(&results).Error

	return
}

// GetBatchFromTokenID 批量查找
func (obj *_OrderTabMgr) GetBatchFromTokenID(tokenIDs []int64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`token_id` IN (?)", tokenIDs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容
func (obj *_OrderTabMgr) GetFromPrice(price float64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找
func (obj *_OrderTabMgr) GetBatchFromPrice(prices []float64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容
func (obj *_OrderTabMgr) GetFromAmount(amount int64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`amount` = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量查找
func (obj *_OrderTabMgr) GetBatchFromAmount(amounts []int64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`amount` IN (?)", amounts).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 1-active 2-inactive
func (obj *_OrderTabMgr) GetFromState(state int8) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 1-active 2-inactive
func (obj *_OrderTabMgr) GetBatchFromState(states []int8) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromCreateTimestamp 通过create_timestamp获取内容 row created block timestamp
func (obj *_OrderTabMgr) GetFromCreateTimestamp(createTimestamp int64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`create_timestamp` = ?", createTimestamp).Find(&results).Error

	return
}

// GetBatchFromCreateTimestamp 批量查找 row created block timestamp
func (obj *_OrderTabMgr) GetBatchFromCreateTimestamp(createTimestamps []int64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`create_timestamp` IN (?)", createTimestamps).Find(&results).Error

	return
}

// GetFromUpdateTimestamp 通过update_timestamp获取内容 row updated block timestamp
func (obj *_OrderTabMgr) GetFromUpdateTimestamp(updateTimestamp int64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`update_timestamp` = ?", updateTimestamp).Find(&results).Error

	return
}

// GetBatchFromUpdateTimestamp 批量查找 row updated block timestamp
func (obj *_OrderTabMgr) GetBatchFromUpdateTimestamp(updateTimestamps []int64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`update_timestamp` IN (?)", updateTimestamps).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 db row create time
func (obj *_OrderTabMgr) GetFromCreateTime(createTime time.Time) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 db row create time
func (obj *_OrderTabMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 db row update time
func (obj *_OrderTabMgr) GetFromUpdateTime(updateTime time.Time) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 db row update time
func (obj *_OrderTabMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_OrderTabMgr) FetchByPrimaryKey(orderID int64) (result OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// FetchIndexByIDxSeller  获取多个内容
func (obj *_OrderTabMgr) FetchIndexByIDxSeller(seller string) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`seller` = ?", seller).Find(&results).Error

	return
}

// FetchIndexByIDxNftAddrTokenID  获取多个内容
func (obj *_OrderTabMgr) FetchIndexByIDxNftAddrTokenID(nftAddr string, tokenID int64) (results []*OrderTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderTab{}).Where("`nft_addr` = ? AND `token_id` = ?", nftAddr, tokenID).Find(&results).Error

	return
}
