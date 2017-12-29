package main

import "fmt"
import "github.com/jlyon1/scheduler/scheduler"

func printSomething() {
	fmt.Println("Something")
}

func main() {
	fmt.Printf("asdf")
	s := scheduler.New()
	j := scheduler.NewJob(printSomething)
	s.AddJob(j)
}
