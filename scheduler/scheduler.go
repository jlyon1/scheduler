package scheduler

import (
  "time"
  "reflect"
)

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
