package action

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"gorm.io/gorm"
)

type SyncTaskAction struct {
	actions.Action
}

// 同步任务，SyncTask() | SyncTask("同步任务")
func SyncTask(options ...interface{}) *SyncTaskAction {
	action := &SyncTaskAction{}

	// 文字
	action.Name = "同步任务"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *SyncTaskAction) Init(ctx *quark.Context) interface{} {

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
func (p *SyncTaskAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	return ctx.JSON(200, message.Success("操作成功"))
}
