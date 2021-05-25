package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/lyondhill/workflow"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
		Logger:   garbageLogger{},
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "workflow", worker.Options{})

	w.RegisterWorkflow(workflow.ExecuteDAG)
	w.RegisterWorkflow(workflow.ExecuteNode)
	w.RegisterActivity(&workflow.Activities{})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

type garbageLogger struct {
}

func (garbageLogger) Debug(msg string, keyvals ...interface{}) {}
func (garbageLogger) Info(msg string, keyvals ...interface{})  {}
func (garbageLogger) Warn(msg string, keyvals ...interface{})  {}
func (garbageLogger) Error(msg string, keyvals ...interface{}) {}
