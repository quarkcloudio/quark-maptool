package action

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
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
	// 需要递归遍历的目录路径
	dirPath := "./web/map"

	// 获取主机名
	hostname, err := os.Hostname()
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	clientIp, err := utils.ClientIp()
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	execPath, _ := os.Executable()
	execDir := filepath.Dir(execPath)

	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", err)
			return ctx.JSON(200, message.Error(err.Error()))
		}
		if info.IsDir() && len(strings.Split(path, "\\")) == 4 {
			getExecDir := strings.ReplaceAll(execDir+path, "\\", "\\\\")
			scriptFilePath := "./web/app/storage/scripts/" + strings.ReplaceAll(path, "\\", "_") + ".jsx"
			p.MakeScript("./web/app/script_templates/changecolor.jsx", scriptFilePath, "127.0.0.1:3000", "地图调色", "调整", getExecDir)
			service.NewPhotoshopTaskService().Insert(
				model.PhotoshopTask{
					ClientIp:   clientIp,
					ClientName: hostname,
					FilePath:   path,
					ScriptPath: scriptFilePath,
					Status:     1,
				},
			)
		}
		return nil
	})

	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("操作成功"))
}

func (p *SyncTaskAction) MakeScript(inputFilePath string, outputFilePath string, tplApiUrl string, tplActionSetName string, tplActionName string, tplInputFolderPath string) error {

	// 读取文件内容
	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		return err
	}

	// 转换为字符串
	text := string(content)

	// API_URL
	modifiedText := strings.ReplaceAll(text, "API_URL", tplApiUrl)

	// ACTION_SET_NAME
	modifiedText = strings.ReplaceAll(modifiedText, "ACTION_SET_NAME", tplActionSetName)

	// API_URL
	modifiedText = strings.ReplaceAll(modifiedText, "ACTION_NAME", tplActionName)

	// INPUT_FOLDER_PATH
	modifiedText = strings.ReplaceAll(modifiedText, "INPUT_FOLDER_PATH", tplInputFolderPath)

	// 将修改后的内容写入新文件
	err = os.WriteFile(outputFilePath, []byte(modifiedText), 0644)
	if err != nil {
		return err
	}

	return nil
}
