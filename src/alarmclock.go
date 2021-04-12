package main

import (
	"fmt"
	"time"
)

func Remind(text string, pause time.Duration) {
	tickChannel := time.Tick(pause)
	for {
		select {
		case now := <-tickChannel:
			fmt.Printf(text, now.Format("15:04"))
		}
	}
}

func main() {
	go Remind("Klockan är %s: Dags att äta\n", 10*time.Second)
	go Remind("Klockan är %s: Dags att arbeta\n", 30*time.Second)
	go Remind("Klockan är %s: Dags att sova\n", 60*time.Second)
	select {}
}
