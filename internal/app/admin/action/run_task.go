package action

import (
	"os/exec"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
	"gorm.io/gorm"
)

type RunTaskAction struct {
	actions.Action
}

// 同步任务，RunTask() | RunTask("同步任务")
func RunTask(options ...interface{}) *RunTaskAction {
	action := &RunTaskAction{}

	// 文字
	action.Name = "启动服务"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *RunTaskAction) Init(ctx *quark.Context) interface{} {

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
func (p *RunTaskAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	go p.task()

	return ctx.JSON(200, message.Success("操作成功"))
}

func (p *RunTaskAction) task() {

	// 设置任务状态
	utils.SetConfig("TASK_STATUS", "1")

	// 创建调度器
	s := gocron.NewScheduler(time.Local)

	// 每 10 秒执行一次任务
	s.Every(5).Seconds().Do(func() {
		taskStatus := utils.GetConfig("TASK_STATUS")
		if taskStatus == "1" {
			hasDoingTask, _ := service.NewPhotoshopTaskService().GetOneByStatus(2)
			if hasDoingTask.Id == 0 {
				taskInfo, _ := service.NewPhotoshopTaskService().GetOneByStatus(1)
				p.doTask(taskInfo)
			}
		}
	})

	// 启动调度器
	s.StartAsync()
}

func (p *RunTaskAction) doTask(task model.PhotoshopTask) {
	// Photoshop 安装路径（根据实际情况调整）
	photoshopPath := `photoshop.exe`

	// 构造命令
	cmd := exec.Command(photoshopPath, task.ScriptPath)

	// 执行命令
	cmd.CombinedOutput()

	service.NewPhotoshopTaskService().UpdateByFilePath(
		task.FilePath,
		model.PhotoshopTask{
			TaskStartedAt: datetime.Now(),
			Status:        2,
		},
	)
}
