//go:generate mockgen -source=clientStore.go -source=workflowRunner.go -destination=mocks/mock_cadence.go
package cadenceClient

import (
	"context"
	"go.uber.org/cadence/client"
	"orderService/models"
)

type WorkflowRunner interface {
	TriggerProcessOrderWorkflow(ctx context.Context,
		workFlowOptions client.StartWorkflowOptions,
		orderDetails models.OrderDetails) error
}

var _ WorkflowRunner = (*Workflows)(nil)
