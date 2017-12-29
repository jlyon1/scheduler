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
	j := scheduler.NewJob(printSomething, "heyo").EveryDay().At(time.Now().Add(10 * time.Second))

	fmt.Println("Job Time: ", j.GetExecTime())
	fmt.Println("Adding Job: ", s.AddJob(j))
	s.Run()

}
