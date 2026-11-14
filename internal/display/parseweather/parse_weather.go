package parseweather

import (
	"fmt"
	"encoding/json"
)

type WeatherResponse struct {
	Weather []struct {
		Main string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`

	Main struct {
		Temp float64 `json:"temp"`
		Humidity int `json:"humidity"`
	} `json:"main"`

	Name string `json:"name"`
}


func ParseWeather(res []byte) {
	var data WeatherResponse
    err := json.Unmarshal(res, &data) 
    if err != nil {
        fmt.Println("Ошибка парсинга:", err)
		return
    }

	fmt.Println("Город", data.Name)
	fmt.Println("Температура", data.Main.Temp)
	fmt.Println("Влажность", data.Main.Humidity)
	fmt.Println(data.Weather[0].Description)
}