package main

import (
	"fmt"
	"time"
)

func main() {

	var currentTime time.Time = time.Now()
	fmt.Println(currentTime.Local())



	formatter :=  "2006-01-02 15:04:05"
	value := "2022-12-31 23:59:59"


	t, err := time.Parse(formatter, value)
		if err != nil{
			fmt.Println(err)
		}else{
			fmt.Println(t.Local())
		}


}
