package action

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
	"gorm.io/gorm"
)

type CloseTaskAction struct {
	actions.Action
}

// 关闭任务，RunTask() | RunTask("关闭任务")
func CloseTask(options ...interface{}) *CloseTaskAction {
	action := &CloseTaskAction{}

	// 文字
	action.Name = "关闭任务"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *CloseTaskAction) Init(ctx *quark.Context) interface{} {

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	p.WithLoading = true

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	p.Type = "primary"

	// 行为类型
	p.ActionType = "ajax"

	return p
}

// 执行行为句柄
func (p *CloseTaskAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	// 设置任务状态
	utils.SetConfig("TASK_STATUS", "0")
	return ctx.JSON(200, message.Success("操作成功"))
}
