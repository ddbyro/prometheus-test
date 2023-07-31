package data

import (
	"fmt"
	"time"
)

func Get_date() {
	currentTime := time.Now()
	fmt.Println("Year   :", currentTime.Year())
	fmt.Println("Month  :", currentTime.Month())
	fmt.Println("Day    :", currentTime.Day())
}
