package test

import (
	"fmt"
	"github.com/leiphp/gokit/pkg/sdk/weather"
	"testing"
)

func TestWeather(t *testing.T) {
	w := weather.NewClient()
	code, ok := weather.GetCityCode("广州")
	if !ok {
		panic("城市编码未找到")
	}
	info, err := w.Get(code)
	if err != nil {
		panic(err)
	}

	fmt.Printf("城市: %s\n当前温度: %s°C\n空气质量: %s\n",
		info.CityInfo.City, info.Data.Wendu, info.Data.Quality)

	if len(info.Data.Forecasts) > 0 {
		forecast := info.Data.Forecasts[0]
		fmt.Printf("今日天气: %s %s ~ %s（%s）\n",
			forecast.Date, forecast.Low, forecast.High, forecast.Type)
	}
}
