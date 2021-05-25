package workflow

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

type Node struct {
	Next        *Node      `json:"next,omitempty"`
	Condition   *Condition `json:"condition,omitempty"`
	Action      *Action    `json:"action,omitempty"`
	Alternative *Node      `json:"alternative,omitempty"`
}

type Condition struct {
	Activity string   `json:"activity"`
	Args     []string `json:"args"`
}

type Action struct {
	Activity string `json:"activity"`
}

type DAG struct {
	Head *Node `json:"head"`
}

// ExecuteDAG Workflow.
func ExecuteDAG(ctx workflow.Context, dag *DAG) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)

	err := workflow.ExecuteChildWorkflow(ctx, ExecuteNode, dag.Head).Get(ctx, nil)
	if err != nil {
		logger.Error("Get greeting failed.", "Error", err)
		return err
	}

	return nil
}

// Node workflow will be created as a child workflow for every node in a dag
// it will attempt to:
// 1 -> execute an activity for the action if the node contains an action
// 2 -> check to see if a conditional statement is present and select the alternative node if the condition is false
// 3 -> execute the next node if the condition was not present or the result was false.
func ExecuteNode(ctx workflow.Context, node *Node) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)

	if node.Action != nil {
		var result string
		err := workflow.ExecuteActivity(ctx, node.Action.Activity).Get(ctx, &result)
		if err != nil {
			logger.Error("Get greeting failed.", "Error", err)
			return err
		}
		fmt.Printf("RESULTS FROM ACTIVITY %s: %s\n", node.Action.Activity, result)

	}

	if node.Condition != nil && node.Alternative != nil {
		var primary bool
		err := workflow.ExecuteActivity(ctx, node.Condition.Activity, node.Condition.Args).Get(ctx, &primary)
		if err != nil {
			logger.Error("Get greeting failed.", "Error", err)
			return err
		}
		fmt.Printf("RESULTS FROM CONDITION %s: %t\n", node.Condition.Activity, primary)
		if !primary {
			fmt.Printf("EXECUTING CHILD WORKFLOW: %s\n", node.Alternative.Action.Activity)
			err := workflow.ExecuteChildWorkflow(ctx, ExecuteNode, node.Alternative).Get(ctx, nil)
			if err != nil {
				logger.Error("Get greeting failed.", "Error", err)
				return err
			}
			return nil
		}
	}

	if node.Next != nil {
		fmt.Printf("EXECUTING CHILD WORKFLOW: %s\n", node.Next.Action.Activity)
		err := workflow.ExecuteChildWorkflow(ctx, ExecuteNode, node.Next).Get(ctx, nil)
		if err != nil {
			logger.Error("Get greeting failed.", "Error", err)
			return err
		}
	}
	return nil
}
