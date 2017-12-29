# Scheduler
A execute tasks at a given time, on a recurring schedule, similar to cron, with the ability to create load and save state.

This is being created with applications to the Rensselaer Polytechnic Institute Web Technologies Group's [Shuttle Tracker](https://github.com/wtg/shuttletracker) in mind, To allow tasks like enabling and disabling routes to be run on a scheduled basis.

### Scope:

- [X] Support Executing arbitrary functions using a custom type
- [X] Execute functions on a recurring basis
- [ ] Ensure they execute, even if the server was off when they were supposed to run
- [ ] Allow those jobs to easily be removed (by some id)
- [ ] Persist Jobs over server reboot or application restart
- [X] Allow One Time scheduled Jobs
- [ ] Support Daylight Savings time

## Example

```
func EnableRoute(){}
s := Scheduler.New()   
s.Run()             
id := s.Execute(EnableRoute).Every("SUN").At("1:00:00")
s.Remove(id)          
state := s.GetState()
s.SetState(s)         
```
