package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/action"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type PhotoshopTask struct {
	resource.Template
}

// 初始化
func (p *PhotoshopTask) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "任务"

	// 模型
	p.Model = &model.PhotoshopTask{}

	// 分页
	p.PerPage = 10

	return p
}

func (p *PhotoshopTask) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("client_name", "客户端名称"),

		field.Text("client_ip", "IP地址"),

		field.Text("script_path", "脚本路径"),

		field.Text("file_path", "文件路径"),

		field.Datetime("task_started_at", "开始时间").OnlyOnIndex(),

		field.Datetime("task_end_at", "完成时间").OnlyOnIndex(),

		field.Datetime("created_at", "创建时间").OnlyOnIndex(),

		field.Select("status", "状态").SetOptions(
			[]selectfield.Option{
				{Value: 1, Label: "未开始"},
				{Value: 2, Label: "已开始"},
				{Value: 3, Label: "已结束"},
			}),
	}
}

// 搜索
func (p *PhotoshopTask) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *PhotoshopTask) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		action.RunTask(),
		action.CloseTask(),
		action.OpenTask(),
		action.SyncTask(),
		actions.BatchDelete(),
		actions.BatchDisable(),
		actions.BatchEnable(),
		actions.Delete(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}
