package test

import (
	"fmt"
	"testing"
)
import "github.com/leiphp/gokit/pkg/sdk/ipinfo"

func TestIpInfo(t *testing.T) {
	ip, err := ipinfo.GetPublicIP()
	if err != nil {
		fmt.Errorf("获取公网 IP 失败: %v", err)
	}
	info, err := ipinfo.GetIPLocation(ip)
	if err != nil {
		fmt.Errorf("获取 IP 所在城市失败: %v", err)
	}
	fmt.Printf("当前 IP：%s\n所在城市：%s - %s - %s（ISP：%s）\n",
		info.Query, info.Country, info.RegionName, info.City, info.ISP)
}
