package action

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type ClearTaskAction struct {
	actions.Action
}

// 清理本机任务，ClearTask() | ClearTask("清理本机任务")
func ClearTask(options ...interface{}) *ClearTaskAction {
	action := &ClearTaskAction{}

	// 文字
	action.Name = "清理本机任务"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *ClearTaskAction) Init(ctx *quark.Context) interface{} {

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	p.WithLoading = true

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	// 行为类型
	p.ActionType = "ajax"

	return p
}

// 执行行为句柄
func (p *ClearTaskAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	// 设置任务状态
	err := service.NewPhotoshopTaskService().Delete()
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	return ctx.JSON(200, message.Success("操作成功"))
}
