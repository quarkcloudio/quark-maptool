package model

import (
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"gorm.io/gorm"
)

// 执行Photoshop任务
type PhotoshopTask struct {
	Id            int               `json:"id" gorm:"autoIncrement"`
	ClientName    string            `json:"client_name" gorm:"size:200;not null"`
	ClientIp      string            `json:"client_ip" gorm:"size:200;not null"`
	ScriptPath    string            `json:"script_path" gorm:"size:200;not null"`
	FilePath      string            `json:"file_path" gorm:"size:200;not null"`
	Status        int               `json:"status" gorm:"size:1;not null;default:1"`
	TaskStartedAt datetime.Datetime `json:"task_started_at"`
	TaskEndAt     datetime.Datetime `json:"task_end_at"`
	CreatedAt     datetime.Datetime `json:"created_at"`
	UpdatedAt     datetime.Datetime `json:"updated_at"`
	DeletedAt     gorm.DeletedAt    `json:"deleted_at"`
}
