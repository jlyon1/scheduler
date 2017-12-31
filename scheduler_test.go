package scheduler_test

import "testing"
import "fmt"
import "time"
import "github.com/jlyon1/scheduler"
import "os"
import "io"
import "bytes"

var s *scheduler.Scheduler

func printSomething(s string) {
	fmt.Printf(s)
}
func printSomethingElse() {
	s.RemoveJob(0)
}

func TestRemoveJob(t *testing.T){
	old := os.Stdout
	r, w, _ := os.Pipe()

	fmt.Println("Taking over stdout to test job removal")

	os.Stdout = w
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()


	s = scheduler.New()
	j := scheduler.NewJob(printSomething, "-").EveryDay().At(time.Now())
	s.AddJob(j)
	go s.Run()
	if (!s.RemoveJob(0)){
		t.Errorf("Could not remove job 0")
	}
	<-time.After(time.Second * 5)

	os.Stdout = old
	//Close pipes
	r.Close()
	w.Close()

	val := <-outC
	if val != "" {
		t.Errorf("We should have '' got %s", val)
	}
	fmt.Println(val)
}

func TestRun(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()

	fmt.Println("Taking over stdout to test prints")

	os.Stdout = w
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	s = scheduler.New()

	j := scheduler.NewJob(printSomething, "-").EveryDay().At(time.Now())
	j2 := scheduler.NewJob(printSomething, "-").EveryDay().At(time.Now().Add(time.Second))
	s.AddJob(j)
	s.AddJob(j2)

	go s.Run()
	//After five seconds we should see two and only two prints
	<-time.After(time.Second * 5)

	os.Stdout = old
	//Close pipes
	r.Close()
	w.Close()

	val := <-outC
	if val != "--" {
		t.Errorf("We should have -- got %s", val)
	}
	fmt.Println(val)
}
