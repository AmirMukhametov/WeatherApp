package terminal

import (
	"fmt"
)

func ReadCityInput() string {
	var city string
	fmt.Println("Введите город")
	fmt.Scan(&city)

	return city
}