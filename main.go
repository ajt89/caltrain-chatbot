package main

import (
	"fmt"
	"github.com/ajt89/caltrain-chatbot/caltrain"
)

func main() {
	data := caltrain.GetRealTime()

	fmt.Printf("Header timestamp: %d\n", data.RealTime.Header.Timestamp)
}
