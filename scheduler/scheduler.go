package scheduler

import (
  "time"
  "reflect"
)

/*
Scope:
-Support Executing arbitrary functions on a recurring basis
  -Ensure they execute, even if the server was off when they were supposed to run
-Allow those jobs to easily be removed (by some id)
-Persist Jobs over server reboot
-Allow One Time scheduled Jobs
-Support Daylight Savings time
*/


type Job struct{
  ToRun interface{}
  args []string
  startTime time.Time
  endTime time.Time
  executed bool
}

func (job *Job) Invoke() {
    v := reflect.ValueOf(job.ToRun)
    rargs := make([]reflect.Value, len(job.args))
    for i, a := range job.args {
        rargs[i] = reflect.ValueOf(a)
    }
    v.Call(rargs)
}
