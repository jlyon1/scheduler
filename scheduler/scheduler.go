//Package scheduler provides functions and types for scheduling functions to
//run on a recurring basis
package scheduler

import (
  "time"
  "reflect"
)

type Scheduler struct{
  jobs []Job
}

type IdGen struct{
  existingIds []int
}

type Job struct{
  id int
  toRun interface{}
  args []string
  startTime time.Time
  endTime time.Time
  executed bool
}

//New returns a new scheduler instance for scheduling jobs
func New() (*Scheduler){
  return &Scheduler{}
}

//NewJob allows for the creation of a new job
func NewJob(function interface{}, args ...string)(*Job){
  j := Job{}
  j.toRun = function
  j.args = args
  j.id = -1
  return &j
}

func (s *Scheduler)AddJob(j *Job) (int){
  s.jobs = append(s.jobs, *j)
  id := len(s.jobs) -1
  return id
}

func (job *Job) Invoke() {
    v := reflect.ValueOf(job.toRun)
    rargs := make([]reflect.Value, len(job.args))
    for i, a := range job.args {
        rargs[i] = reflect.ValueOf(a)
    }
    v.Call(rargs)
}
