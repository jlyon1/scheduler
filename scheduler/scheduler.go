//Package scheduler provides functions and types for scheduling functions to
//run on a recurring basis, It supports running jobs at a scheduled time based on current time
package scheduler

import (
  "time"
  "fmt"
  "reflect"
)


type Scheduler struct{
  jobs []Job
  maxId int
}

type Job struct{
  id int
  toRun interface{}
  args []string
  day int
  time time.Time
  executed bool
  once bool

}

//Run instructs the scheduler to start running jobs on their timed basis
func (s *Scheduler) Run(){
  for{
    <-time.After(time.Second/2)
    for idx,_ := range s.jobs{
      if(time.Now().After(s.jobs[idx].time) && !s.jobs[idx].executed){
        s.jobs[idx].Invoke()
        if(!s.jobs[idx].once){
          s.jobs[idx].time = s.jobs[idx].time.AddDate(0,0,1)
        }
      }
    }
  }
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

//Every allows a job to be scheduled for every day
func (j *Job)Every(day int) (*Job){
  j.day = day
  return j
}

//Every allows a job to be scheduled for every day, We use day 8 to represent every day
func (j *Job)EveryDay() (*Job){
  j.day = 8
  return j
}

//GetDaySchedule returns the day of the week the job is to run
func (j *Job)GetDaySchedule() (int){
  return j.day
}

//GetExecTime returns the time to execute the task
func (j *Job)GetExecTime() (time.Time){
  return j.time
}

//At sets the time of day for the job to be run at
func (j *Job)At(t time.Time) (*Job){
  j.time = t
  return j
}

//AddJob adds a job to the scheduler queue, jobs contain all the information
//about when to run, and timing for the job
func (s *Scheduler)AddJob(j *Job) (int){
  if(j.id != -1){
    fmt.Println("Job Not added")
  }
  j.id = s.maxId
  s.maxId += 1;
  s.jobs = append(s.jobs, *j)
  return j.id
}
//Invoke runs the function associated with a given job
func (job *Job) Invoke() {
    v := reflect.ValueOf(job.toRun)
    rargs := make([]reflect.Value, len(job.args))
    for i, a := range job.args {
        rargs[i] = reflect.ValueOf(a)
    }
    v.Call(rargs)
}
