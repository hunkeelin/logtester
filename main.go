package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type JsonLog struct {
	Msg  string `json:"msg"`
	Time string `json:"time"`
}

func main() {
	for {
		// Creating time object
		t := time.Now()
		obj := JsonLog{
			Msg:  "this is klin msg logtester",
			Time: t.Format("2006-01-02T15:04:05.999Z"),
		}

		time.Sleep(1 * time.Second)
		fmt.Println("klinlogtester: printing json")
		b, err := json.Marshal(obj)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))

	}
}

// "time":"2022-03-18T16:39:41.938Z"
