package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserTabMgr struct {
	*_BaseMgr
}

// UserTabMgr open func
func UserTabMgr(db *gorm.DB) *_UserTabMgr {
	if db == nil {
		panic(fmt.Errorf("UserTabMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserTabMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_tab"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserTabMgr) GetTableName() string {
	return "user_tab"
}

// Reset 重置gorm会话
func (obj *_UserTabMgr) Reset() *_UserTabMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserTabMgr) Get() (result UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserTabMgr) Gets() (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserTabMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UserTab{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserTabMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOpenid openid获取
func (obj *_UserTabMgr) WithOpenid(openid string) Option {
	return optionFunc(func(o *options) { o.query["openid"] = openid })
}

// WithMoney money获取
func (obj *_UserTabMgr) WithMoney(money int64) Option {
	return optionFunc(func(o *options) { o.query["money"] = money })
}

// WithGameTime game_time获取
func (obj *_UserTabMgr) WithGameTime(gameTime int64) Option {
	return optionFunc(func(o *options) { o.query["game_time"] = gameTime })
}

// WithState state获取 1-active 2-inactive
func (obj *_UserTabMgr) WithState(state int8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithCreateTime create_time获取
func (obj *_UserTabMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_UserTabMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithBizData biz_data获取
func (obj *_UserTabMgr) WithBizData(bizData string) Option {
	return optionFunc(func(o *options) { o.query["biz_data"] = bizData })
}

// GetByOption 功能选项模式获取
func (obj *_UserTabMgr) GetByOption(opts ...Option) (result UserTab, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserTabMgr) GetByOptions(opts ...Option) (results []*UserTab, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UserTabMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UserTab, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_UserTabMgr) GetFromID(id int64) (result UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UserTabMgr) GetBatchFromID(ids []int64) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOpenid 通过openid获取内容
func (obj *_UserTabMgr) GetFromOpenid(openid string) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`openid` = ?", openid).Find(&results).Error

	return
}

// GetBatchFromOpenid 批量查找
func (obj *_UserTabMgr) GetBatchFromOpenid(openids []string) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`openid` IN (?)", openids).Find(&results).Error

	return
}

// GetFromMoney 通过money获取内容
func (obj *_UserTabMgr) GetFromMoney(money int64) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`money` = ?", money).Find(&results).Error

	return
}

// GetBatchFromMoney 批量查找
func (obj *_UserTabMgr) GetBatchFromMoney(moneys []int64) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`money` IN (?)", moneys).Find(&results).Error

	return
}

// GetFromGameTime 通过game_time获取内容
func (obj *_UserTabMgr) GetFromGameTime(gameTime int64) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`game_time` = ?", gameTime).Find(&results).Error

	return
}

// GetBatchFromGameTime 批量查找
func (obj *_UserTabMgr) GetBatchFromGameTime(gameTimes []int64) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`game_time` IN (?)", gameTimes).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 1-active 2-inactive
func (obj *_UserTabMgr) GetFromState(state int8) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 1-active 2-inactive
func (obj *_UserTabMgr) GetBatchFromState(states []int8) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_UserTabMgr) GetFromCreateTime(createTime time.Time) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_UserTabMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_UserTabMgr) GetFromUpdateTime(updateTime time.Time) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_UserTabMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromBizData 通过biz_data获取内容
func (obj *_UserTabMgr) GetFromBizData(bizData string) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`biz_data` = ?", bizData).Find(&results).Error

	return
}

// GetBatchFromBizData 批量查找
func (obj *_UserTabMgr) GetBatchFromBizData(bizDatas []string) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`biz_data` IN (?)", bizDatas).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserTabMgr) FetchByPrimaryKey(id int64) (result UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIDxOpenid  获取多个内容
func (obj *_UserTabMgr) FetchIndexByIDxOpenid(openid string) (results []*UserTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserTab{}).Where("`openid` = ?", openid).Find(&results).Error

	return
}
