package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/cron"
	"time"
)

type MyJob struct {
	Name string
}

func (j MyJob) Run() {
	fmt.Println(j.Name, time.Now())
}

func main() {
	langgo.Run(&cron.Instance{WithSeconds: true})
	wait := make(chan struct{})
	cron.Get().BindTaskAndSchedule("basic", "* * * * * *", MyJob{Name: "MyJob"})
	<-wait
}
