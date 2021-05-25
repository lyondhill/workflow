package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/lyondhill/workflow"
	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
	"gopkg.in/yaml.v3"
)

func main() {
	var dagFile string
	flag.StringVar(&dagFile, "dagFile", "simpleDAG.json", "dagFile specify the json file for the dsl workflow.")
	flag.Parse()

	data, err := ioutil.ReadFile(dagFile)
	if err != nil {
		log.Fatalln("failed to load dsl config file", err)
	}
	dag := &workflow.DAG{}
	if err := yaml.Unmarshal(data, dag); err != nil {
		log.Fatalln("failed to unmarshal dsl config", err)
	}
	fmt.Printf("%+v\n", dag)

	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowID := uuid.New()
	workflowOptions := client.StartWorkflowOptions{
		ID:        "workflow_" + workflowID,
		TaskQueue: "workflow",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflow.ExecuteDAG, dag)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

}
