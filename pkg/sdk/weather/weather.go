package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	BaseURL string
	Client  *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: "http://t.weather.sojson.com/api/weather/city/",
		Client:  &http.Client{Timeout: 5 * time.Second},
	}
}

type WeatherResponse struct {
	CityInfo struct {
		City string `json:"city"`
	} `json:"cityInfo"`
	Data struct {
		Shidu     string `json:"shidu"`   // 湿度
		Quality   string `json:"quality"` // 空气质量
		Wendu     string `json:"wendu"`   // 当前温度
		Forecasts []struct {
			Date string `json:"ymd"`
			High string `json:"high"`
			Low  string `json:"low"`
			Type string `json:"type"`
			Week string `json:"week"`
		} `json:"forecast"`
	} `json:"data"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (c *Client) Get(cityCode string) (*WeatherResponse, error) {
	resp, err := c.Client.Get(c.BaseURL + cityCode)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result WeatherResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.Status != 200 {
		return nil, fmt.Errorf("天气服务错误: %s", result.Message)
	}

	return &result, nil
}
