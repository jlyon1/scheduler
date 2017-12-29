package scheduler

import (
  "time"
  "reflect"
)

type Scheduler struct{
  jobs []Job

}

type Job struct{
  ToRun interface{}
  args []string
  startTime time.Time
  endTime time.Time
  executed bool
}

//New returns a new scheduler instance for scheduling jobs
func New() (*Scheduler){
  return &Scheduler{}
}

func NewJob(function interface{}, args ...string)(*Job){
  j := Job{}
  j.ToRun = function
  j.args = args
  return &j
}

func (job *Job) Invoke() {
    v := reflect.ValueOf(job.ToRun)
    rargs := make([]reflect.Value, len(job.args))
    for i, a := range job.args {
        rargs[i] = reflect.ValueOf(a)
    }
    v.Call(rargs)
}
