//Package scheduler provides functions and types for scheduling functions to
//run on a recurring basis, It supports running Jobs at a scheduled time based on current time
package scheduler

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Scheduler struct {
	Jobs  []Job `json: Jobs`
	MaxId int   `json: MaxId`
}

type Job struct {
	Id    int         `json:id`
	toRun interface{} `json:toRun`
	args  []string    `json:args`
	day   int         `json: day`
	time  time.Time   `json:time`
	once  bool        `json:once`
}

//RemoveJob allows you to remove a job by an id
func (s *Scheduler) RemoveJob(id int) bool {
	i := -1
	for idx, j := range s.Jobs {
		if j.Id == id {
			i = idx
		}
	}
	if i == -1 {
		return false
	} else {
		s.Jobs = append(s.Jobs[:i], s.Jobs[i+1:]...)
		return true
	}
}

//Run instructs the scheduler to start running Jobs on their timed basis
func (s *Scheduler) Run() {
	for {
		<-time.After(time.Second / 2)
		for idx, _ := range s.Jobs {
			if time.Now().After(s.Jobs[idx].time) {
				go s.Jobs[idx].Invoke()
				if !s.Jobs[idx].once {
					if s.Jobs[idx].day == 8 {
						s.Jobs[idx].time = s.Jobs[idx].time.AddDate(0, 0, 1)
					} else {
						s.Jobs[idx].time = s.Jobs[idx].time.AddDate(0, 0, 7)
					}
				}
			}
		}
	}
}

//New returns a new scheduler instance for scheduling Jobs
func New() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Export() string {
	val, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("error: ", err.Error())
	}
	return string(val)
}

//NewJob allows for the creation of a new job
func NewJob(function interface{}, args ...string) *Job {
	j := Job{}
	j.toRun = function
	j.args = args
	j.Id = -1
	return &j
}

//Every allows a job to be scheduled for every day
func (j *Job) Every(day int) *Job {
	j.day = day
	return j
}

//Every allows a job to be scheduled for every day, We use day 8 to represent every day
func (j *Job) EveryDay() *Job {
	j.day = 8
	return j
}

//GetDaySchedule returns the day of the week the job is to run
func (j *Job) GetDaySchedule() int {
	return j.day
}

//GetExecTime returns the time to execute the task
func (j *Job) GetExecTime() time.Time {
	return j.time
}

//At sets the time of day for the job to be run at, which is set to the current
//day if it is to be run every day, or the specified date
func (j *Job) At(t time.Time) *Job {
	j.time = time.Now()
	daysToAdd := 0
	duration := t.Sub(j.time)
	if int(t.Weekday()) != j.day && j.day != 8 {
		daysToAdd = (j.day - int(t.Weekday()))
		if int(t.Weekday()) > j.day {
			daysToAdd = 7 + daysToAdd
		}
	}
	j.time = j.time.Add(duration)
	j.time = j.time.AddDate(0, 0, daysToAdd)
	return j
}

//AddJob adds a job to the scheduler queue, Jobs contain all the information
//about when to run, and timing for the job
func (s *Scheduler) AddJob(j *Job) int {
	if j.Id != -1 {
		fmt.Println("Job Not added")
	}
	j.Id = s.MaxId
	s.MaxId += 1
	s.Jobs = append(s.Jobs, *j)
	return j.Id
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
