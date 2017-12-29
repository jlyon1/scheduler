package main

import "fmt"
import "time"
import "github.com/jlyon1/scheduler/scheduler"

func printSomething(s string) {
	fmt.Println(s)
}
func printSomethingElse() {
	fmt.Println("Something Else")
}

func main() {
	fmt.Println("Current Day:", time.Friday)
	s := scheduler.New()
	fmt.Println("Current Time: ", time.Now())
	j := scheduler.NewJob(printSomething, "Scheduled Print").EveryDay().At(time.Now().Add(2 * time.Second))
	fmt.Println("Job Time: ", j.GetExecTime())
	// const longForm = "Jan 2 2006 15:04:05 (MST)"
	// t, _ := time.Parse(longForm, "Dec 30 2017 13:30:34 (EST)")
	fmt.Println("Adding Job: ", s.AddJob(j))
	s.Run()

}
