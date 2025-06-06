package ipinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
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

var ipProviders = []string{
	"https://api.myip.la",
	"https://api.ip.sb/ip",
	"http://checkip.amazonaws.com",
}

// GetPublicIP 获取本机公网 IP
func GetPublicIP() (string, error) {
	for _, api := range ipProviders {
		resp, err := http.Get(api)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		ipBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		ip := strings.TrimSpace(string(ipBytes))
		if net.ParseIP(ip) != nil {
			return ip, nil
		}
	}
	return "", fmt.Errorf("所有公网 IP 提供者请求失败")
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
