package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _TStudentMgr struct {
	*_BaseMgr
}

// TStudentMgr open func
func TStudentMgr(db *gorm.DB) *_TStudentMgr {
	if db == nil {
		panic(fmt.Errorf("TStudentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TStudentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_student"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TStudentMgr) GetTableName() string {
	return "t_student"
}

// Reset 重置gorm会话
func (obj *_TStudentMgr) Reset() *_TStudentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TStudentMgr) Get() (result TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TStudentMgr) Gets() (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TStudentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TStudent{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TStudentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_TStudentMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAge age获取
func (obj *_TStudentMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithPhone phone获取
func (obj *_TStudentMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithCls cls获取
func (obj *_TStudentMgr) WithCls(cls string) Option {
	return optionFunc(func(o *options) { o.query["cls"] = cls })
}

// WithCreateTime create_time获取
func (obj *_TStudentMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_TStudentMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_TStudentMgr) GetByOption(opts ...Option) (result TStudent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TStudentMgr) GetByOptions(opts ...Option) (results []*TStudent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TStudentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TStudent, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where(options.query)
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
func (obj *_TStudentMgr) GetFromID(id int) (result TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TStudentMgr) GetBatchFromID(ids []int) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TStudentMgr) GetFromName(name string) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_TStudentMgr) GetBatchFromName(names []string) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容
func (obj *_TStudentMgr) GetFromAge(age int) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`age` = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量查找
func (obj *_TStudentMgr) GetBatchFromAge(ages []int) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`age` IN (?)", ages).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容
func (obj *_TStudentMgr) GetFromPhone(phone string) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找
func (obj *_TStudentMgr) GetBatchFromPhone(phones []string) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromCls 通过cls获取内容
func (obj *_TStudentMgr) GetFromCls(cls string) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`cls` = ?", cls).Find(&results).Error

	return
}

// GetBatchFromCls 批量查找
func (obj *_TStudentMgr) GetBatchFromCls(clss []string) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`cls` IN (?)", clss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_TStudentMgr) GetFromCreateTime(createTime time.Time) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_TStudentMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_TStudentMgr) GetFromUpdateTime(updateTime time.Time) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_TStudentMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TStudentMgr) FetchByPrimaryKey(id int) (result TStudent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TStudent{}).Where("`id` = ?", id).First(&result).Error

	return
}
