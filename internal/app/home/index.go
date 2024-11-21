package home

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

// 结构体
type Index struct{}

// 首页
func (p *Index) Index(ctx *quark.Context) error {
	return ctx.Render(200, "index.html", map[string]interface{}{
		"content": "Hello, world!",
	})
}

// 任务完成
func (p *Index) TaskDone(ctx *quark.Context) error {
	taskPath := ctx.QueryParam("taskPath")
	execPath, _ := os.Executable()
	execDir := filepath.Dir(execPath)
	getTaskPath := strings.ReplaceAll(taskPath, execDir+"\\", "")
	service.NewPhotoshopTaskService().UpdateByFilePath(
		getTaskPath,
		model.PhotoshopTask{
			TaskEndAt: datetime.Now(),
			Status:    3,
		},
	)
	return ctx.JSONOk(taskPath)
}
