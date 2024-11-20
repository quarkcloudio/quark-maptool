package utils

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/quarkcloudio/quark-go/v3/service"
)

// 获取文件路径
func GetFilePath(id interface{}) string {
	if id == nil {
		return ""
	}

	return service.NewFileService().GetPath(id)
}

// 获取多文件路径
func GetFilePaths(id interface{}) []string {
	if id == nil {
		return nil
	}

	return service.NewFileService().GetPaths(id)
}

// 获取图片路径
func GetPicturePath(id interface{}) string {
	if id == nil {
		return ""
	}

	return service.NewPictureService().GetPath(id)
}

// 获取多图片路径
func GetPicturePaths(id interface{}) []string {
	if id == nil {
		return nil
	}

	return service.NewPictureService().GetPaths(id)
}

// 获取配置
func GetConfig(key string) string {
	return service.NewConfigService().GetValue(key)
}

// 获取域名
func GetDomain() string {

	http := ""
	domain := service.NewConfigService().GetValue("WEB_SITE_DOMAIN")
	ssl := service.NewConfigService().GetValue("SSL_OPEN")
	if domain != "" {
		if ssl == "1" {
			http = "https://"
		} else {
			http = "http://"
		}
	}

	return http + domain
}

// 内容中的地址替换
func ReplaceContentSrc(content string) string {

	reg := regexp.MustCompile(`src="(/[^"]*)"`)

	return reg.ReplaceAllStringFunc(content, func(src string) string {
		return "src= \"" + GetDomain() + src[strings.Index(src, "\"")+1:] + "\""
	})
}

// 正则验证
// expr 正则表达式
// content 要验证的内容
func CheckRegex(expr, content string) bool {

	r, err := regexp.Compile(expr)
	if err != nil {
		return false
	}

	return r.MatchString(content)
}

func ClientIp() (ip string, err error) {
	err = nil
	ip = ""
	// 获取网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return ip, err
	}

	for _, iface := range interfaces {
		// 跳过未启用或环回接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口地址
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Printf("Error getting addresses for interface %s: %v\n", iface.Name, err)
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && ipNet.IP.To4() != nil { // 只获取 IPv4 地址
				ip = ipNet.IP.String()
			}
		}
	}
	return ip, err
}
