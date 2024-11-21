package search

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/searches"
	"gorm.io/gorm"
)

type StatusField struct {
	searches.Select
}

// 状态
func Status() *StatusField {
	field := &StatusField{}
	field.Name = "状态"

	return field
}

// 执行查询
func (p *StatusField) Apply(ctx *quark.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where("status = ?", value)
}

// 属性
func (p *StatusField) Options(ctx *quark.Context) interface{} {

	return []selectfield.Option{
		p.Option("未开始", 1),
		p.Option("进行中", 2),
		p.Option("已完成", 3),
	}
}
