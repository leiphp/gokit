package response

import (
	"encoding/json"
	"net/http"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 格式化 []byte 输出 JSON 响应
func Success(w http.ResponseWriter, body []byte) {
	var obj interface{}
	err := json.Unmarshal(body, &obj)
	if err != nil {
		Error(w, 500, "返回数据解析失败: "+err.Error())
		return
	}
	writeJSON(w, http.StatusOK, Resp{
		Code: 0,
		Msg:  "success",
		Data: obj,
	})
}

// Error 返回错误信息
func Error(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, http.StatusOK, Resp{
		Code: code,
		Msg:  msg,
	})
}

func writeJSON(w http.ResponseWriter, status int, data Resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
