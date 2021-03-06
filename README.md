# Scheduler [![Build Status](https://travis-ci.org/jlyon1/scheduler.svg?branch=master)](https://travis-ci.org/jlyon1/scheduler) [![GoDoc](https://godoc.org/github.com/jlyon1/scheduler/scheduler?status.svg)](https://godoc.org/github.com/jlyon1/scheduler/scheduler) [![Go Report Card](https://goreportcard.com/badge/github.com/jlyon1/scheduler)](https://goreportcard.com/report/github.com/jlyon1/scheduler) [![codecov](https://codecov.io/gh/jlyon1/scheduler/branch/master/graph/badge.svg)](https://codecov.io/gh/jlyon1/scheduler)
A execute tasks at a given time, on a recurring schedule, similar to cron, with the ability to create load and save state.

This is being created with the Rensselaer Polytechnic Institute Web Technologies Group's [Shuttle Tracker](https://github.com/wtg/shuttletracker) in mind, To allow tasks like enabling and disabling routes to be run on a scheduled basis.

### Scope:

- [X] Support Executing arbitrary functions using a custom type
- [X] Execute functions on a recurring basis
- [ ] Ensure they execute, even if the server was off when they were supposed to run
- [X] Allow those jobs to easily be removed (by some id)
- [ ] Persist Jobs over server reboot or application restart
- [X] Allow One Time scheduled Jobs
- [ ] Support Daylight Savings time

## Example

```
package main

import "fmt"
import "time"
import "github.com/jlyon1/scheduler"

func printSomething(s string) {
	fmt.Println(s)
}
func printSomethingElse() {
	fmt.Println("Something Else")
}

func main() {
	fmt.Println("Current Day:", time.Friday)
	s := scheduler.New()
	j := scheduler.NewJob(printSomething, "hey").EveryDay().At(time.Now().Add(10 * time.Second))

	fmt.Println("Job Time: ", j.GetExecTime())
	fmt.Println("Adding Job: ", s.AddJob(j))
	s.Run()

}

```
