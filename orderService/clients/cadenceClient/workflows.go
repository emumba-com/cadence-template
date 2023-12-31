package cadenceClient

import (
	"context"
	"go.uber.org/cadence/client"
	"orderService/env"
	"orderService/log"

	"orderService/models"
)

var (
	ProcessOrderWorkflow = env.Env.CadenceWorkerServiceName + "/v1/workflows/orders.ProcessAnOrder"
)

func (wr *Workflows) TriggerProcessOrderWorkflow(ctx context.Context,
	workFlowOptions client.StartWorkflowOptions,
	orderDetails models.OrderDetails) error {
	logger := log.GetLogger()
	logger.Info("TriggerProcessOrderWorkflow endpoint called")

	workflowRun, err := wr.client.ExecuteWorkflow(
		ctx,
		workFlowOptions,
		ProcessOrderWorkflow,
		orderDetails)
	if err != nil {
		return err
	}

	var emptyInterface interface{}
	err = workflowRun.Get(ctx, &emptyInterface)

	return err
}
