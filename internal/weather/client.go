package client

import (
	"fmt"
	"net/http"
	"io"
	"weatherapp/internal/config"
)

func Fetch() {
	apiKey := config.GetApiKey()
	resp, err := http.Get("http://api.openweathermap.org/geo/1.0/direct?q=Moscow&limit=5&appid=" + apiKey)
	if err != nil {
		fmt.Println("Fetch error", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error ReadAll", err)
	}
	fmt.Println(string(body))
}