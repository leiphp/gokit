package ipinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// IPResponse 用于 IP 查询
type IPResponse struct {
	Query      string `json:"query"`
	Country    string `json:"country"`
	RegionName string `json:"regionName"`
	City       string `json:"city"`
	ISP        string `json:"isp"`
	Status     string `json:"status"`
	Message    string `json:"message"` // 失败时返回
}

// GetPublicIP 获取本机公网 IP
func GetPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}

// GetIPLocation 查询 IP 的位置信息
func GetIPLocation(ip string) (*IPResponse, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var info IPResponse
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	if info.Status != "success" {
		return nil, fmt.Errorf("IP 查询失败: %s", info.Message)
	}
	return &info, nil
}
