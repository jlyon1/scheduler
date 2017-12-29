package main

import "fmt"
import "time"
import "github.com/jlyon1/scheduler/scheduler"

var s *scheduler.Scheduler

func printSomething(s string) {
	fmt.Println(s)
}
func printSomethingElse() {
	s.RemoveJob(0)
}

func main() {
	fmt.Println("Current Day:", time.Friday)
	s = scheduler.New()
	fmt.Println("Current Time: ", time.Now())
	j := scheduler.NewJob(printSomething, "Scheduled Print").EveryDay().At(time.Now().Add(6 * time.Second))
	j2 := scheduler.NewJob(printSomethingElse).EveryDay().At(time.Now().Add(10 * time.Second))

	fmt.Println("Job Time: ", j.GetExecTime())
	fmt.Println("Adding Job: ", s.AddJob(j))
	fmt.Println("Adding Job: ", s.AddJob(j2))
	s.Run()

}
