package main

import (
	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
	"orderServiceCadenceWorker/v1/activities/communicationService"
	"orderServiceCadenceWorker/v1/activities/orderService"
	"orderServiceCadenceWorker/v1/activities/paymentService"
	"orderServiceCadenceWorker/v1/cadenceClient"
	"orderServiceCadenceWorker/v1/env"
	"orderServiceCadenceWorker/v1/log"
	"orderServiceCadenceWorker/v1/workflows/orders"

	"os"
	"os/signal"
	"syscall"
)

var Domain = env.Env.CadenceDomain
var TaskListName = env.Env.TaskListName
var workerInstance worker.Worker

func main() {
	log.InitLogger(env.Env.BuildEnv, "./v1/logs/order-service-cadence-worker.log")

	logger := log.GetLogger()

	startWorker(logger, cadenceClient.BuildCadenceClient())

	osSignalsChan := make(chan os.Signal, 1)
	signal.Notify(osSignalsChan, syscall.SIGINT, syscall.SIGTERM)

	<-osSignalsChan // listen for shutdown signal

	workerInstance.Stop()
}

func startWorker(logger *zap.Logger, service workflowserviceclient.Interface) {
	workerOptions := worker.Options{
		Logger:       logger,
		MetricsScope: tally.NewTestScope(TaskListName, map[string]string{}),
	}

	workerInstance = worker.New(
		service,
		Domain,
		TaskListName,
		workerOptions,
	)

	registerProcessOrderWorkflowsAndActivities(workerInstance)

	if err := workerInstance.Start(); err != nil {
		panic("failed to start worker")
	}

	logger.Info("Started Worker.", zap.String("worker", TaskListName))
}

func registerProcessOrderWorkflowsAndActivities(workerInstance worker.Worker) {
	workerInstance.RegisterWorkflow(orders.ProcessAnOrder)

	workerInstance.RegisterActivity(orderService.UpdateOrderStatus)
	workerInstance.RegisterActivity(orderService.ShipOrder)

	workerInstance.RegisterActivity(paymentService.MakePayment)
	workerInstance.RegisterActivity(paymentService.RefundPayment)

	workerInstance.RegisterActivity(communicationService.SendEmail)
}
