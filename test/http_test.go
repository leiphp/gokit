package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/leiphp/gokit/pkg/core/httpclient"
	"github.com/leiphp/gokit/utils/response"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	// 创建一个模拟的 ResponseWriter
	w := httptest.NewRecorder()
	url := "http://127.0.0.1:8081/api/v1/transceiver/list"
	body, err := httpclient.Get(url)
	if err != nil {
		t.Fatal("请求失败:", err)
		response.Error(w, 500, "请求失败: "+err.Error())
		return
	}
	fmt.Println("body json:", string(body))
	// 尝试格式化 JSON 输出
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		t.Fatal("JSON格式化失败:", err)
	}
	fmt.Println("响应 JSON：", prettyJSON.String())
}
