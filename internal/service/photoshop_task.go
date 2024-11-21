package service

import (
	"os"

	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
)

type PhotoshopTaskService struct{}

func NewPhotoshopTaskService() *PhotoshopTaskService {
	return &PhotoshopTaskService{}
}

// 插入数据
func (p *PhotoshopTaskService) Insert(data model.PhotoshopTask) error {
	err := db.Client.Create(&data).Error
	return err
}

// 更改状态
func (p *PhotoshopTaskService) UpdateByFilePath(taskFilePath string, data model.PhotoshopTask) error {
	clientIp, err := utils.ClientIp()
	if err != nil {
		return err
	}

	// 获取主机名
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	err = db.Client.
		Model(model.PhotoshopTask{}).
		Where("client_name", hostname).
		Where("client_ip", clientIp).
		Where("file_path = ?", taskFilePath).
		Updates(&data).Error

	return err
}

// 清理本机任务
func (p *PhotoshopTaskService) Delete() (err error) {
	clientIp, err := utils.ClientIp()
	if err != nil {
		return err
	}

	// 获取主机名
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	err = db.Client.
		Where("client_name", hostname).
		Where("client_ip", clientIp).
		Delete(&model.PhotoshopTask{}).Error

	return err
}

// 获取列表
func (p *PhotoshopTaskService) GetList() (list []model.PhotoshopTask, err error) {
	clientIp, err := utils.ClientIp()
	if err != nil {
		return nil, err
	}

	// 获取主机名
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	err = db.Client.
		Where("client_name", hostname).
		Where("client_ip", clientIp).
		Where("status", 1).
		Find(&list).Error

	return list, err
}

// 获取单条数据
func (p *PhotoshopTaskService) GetOneByStatus(status int) (info model.PhotoshopTask, err error) {
	clientIp, err := utils.ClientIp()
	if err != nil {
		return info, err
	}

	// 获取主机名
	hostname, err := os.Hostname()
	if err != nil {
		return info, err
	}

	err = db.Client.
		Where("client_name", hostname).
		Where("client_ip", clientIp).
		Where("status", status).
		First(&info).Error

	return info, err
}

// 获取单条数据
func (p *PhotoshopTaskService) GetOneFilePath(path string) (info model.PhotoshopTask) {
	db.Client.
		Where("file_path = ?", path).
		First(&info)

	return info
}
