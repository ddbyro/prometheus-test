package data

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type myJSON struct {
	Array []string
}

func CreateJson() {
	currentTime := time.Now()
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()
	jsondat := &myJSON{Array: []string{strconv.Itoa(year), month, strconv.Itoa(day)}}
	encjson, _ := json.Marshal(jsondat)
	fmt.Println(string(encjson))
}
