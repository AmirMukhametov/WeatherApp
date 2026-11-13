package parseweather

import (
	"fmt"
	"encoding/json"
)

// type User struct {
//     Name string `json:"name"`
//     Age  int    `json:"age"`
// }

func ParseWeather(res []byte) {
	var data string
    err := json.Unmarshal(res, &data) // ВАЖНО: передаем указатель &u
    if err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", data) 
}