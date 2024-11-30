package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PullOffsetTabMgr struct {
	*_BaseMgr
}

// PullOffsetTabMgr open func
func PullOffsetTabMgr(db *gorm.DB) *_PullOffsetTabMgr {
	if db == nil {
		panic(fmt.Errorf("PullOffsetTabMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PullOffsetTabMgr{_BaseMgr: &_BaseMgr{DB: db.Table("pull_offset_tab"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PullOffsetTabMgr) GetTableName() string {
	return "pull_offset_tab"
}

// Reset 重置gorm会话
func (obj *_PullOffsetTabMgr) Reset() *_PullOffsetTabMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PullOffsetTabMgr) Get() (result PullOffsetTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PullOffsetTabMgr) Gets() (results []*PullOffsetTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PullOffsetTabMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PullOffsetTabMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOffset offset获取
func (obj *_PullOffsetTabMgr) WithOffset(offset int64) Option {
	return optionFunc(func(o *options) { o.query["offset"] = offset })
}

// GetByOption 功能选项模式获取
func (obj *_PullOffsetTabMgr) GetByOption(opts ...Option) (result PullOffsetTab, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PullOffsetTabMgr) GetByOptions(opts ...Option) (results []*PullOffsetTab, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_PullOffsetTabMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]PullOffsetTab, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).Where(options.query)
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
func (obj *_PullOffsetTabMgr) GetFromID(id int) (result PullOffsetTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_PullOffsetTabMgr) GetBatchFromID(ids []int) (results []*PullOffsetTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOffset 通过offset获取内容
func (obj *_PullOffsetTabMgr) GetFromOffset(offset int64) (results []*PullOffsetTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).Where("`offset` = ?", offset).Find(&results).Error

	return
}

// GetBatchFromOffset 批量查找
func (obj *_PullOffsetTabMgr) GetBatchFromOffset(offsets []int64) (results []*PullOffsetTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).Where("`offset` IN (?)", offsets).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PullOffsetTabMgr) FetchByPrimaryKey(id int) (result PullOffsetTab, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PullOffsetTab{}).Where("`id` = ?", id).First(&result).Error

	return
}
