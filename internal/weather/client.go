package client

import (
	"fmt"
	"net/http"
	"io"
	"weatherapp/internal/config"
	"weatherapp/internal/display/terminal"
	"weatherapp/internal/display/parseweather"
)

func Fetch() {
	city := terminal.ReadCityInput()
	apiKey := config.GetApiKey()
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city +"&appid=" + apiKey + "&units=metric&lang=ru")
	if err != nil {
		fmt.Println("Fetch error", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error ReadAll", err)
		return
	}
	parseweather.ParseWeather(body)
}