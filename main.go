package main

import (
	"fmt"
	"github.com/robfig/cron"
	// "scheduler/scheduler"
)

func test() {
	fmt.Printf("asdf")
}

func main() {
	c := cron.New()
	c.AddFunc("@every 1s", func() { fmt.Println("Every hour on the half hour") })
	c.Start()
	for {
	}
}
