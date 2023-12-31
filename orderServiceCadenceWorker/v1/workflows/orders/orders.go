package orders

import (
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"orderServiceCadenceWorker/v1/activities/communicationService"
	"orderServiceCadenceWorker/v1/activities/orderService"
	"orderServiceCadenceWorker/v1/activities/paymentService"
	"orderServiceCadenceWorker/v1/log"
	"orderServiceCadenceWorker/v1/models"
	"orderServiceCadenceWorker/v1/utils"
)

func ProcessAnOrder(ctx workflow.Context, orderDetails models.OrderDetails) error {
	var (
		logger       = log.GetLogger()
		ctxWithRetry = workflow.WithActivityOptions(ctx, utils.GetActivityOptions(utils.Duration5Minutes))
		err          error
	)

	logger.Info("Workflow ProcessAnOrder started", zap.String("OrderID", orderDetails.ID))

	defer func() {
		// this is a compensation activity that needs to be executed in case of any error
		if err == nil {
			return
		}

		logger.Info("Executing Compensation Activity UpdateOrderStatus", zap.String("OrderID", orderDetails.ID))

		err = workflow.ExecuteActivity(ctxWithRetry, orderService.UpdateOrderStatus, orderDetails.ID, utils.OrderStatusFailed).Get(ctx, nil)
		if err != nil {
			logger.Error("Error while executing UpdateOrderStatusActivity", zap.Error(err))
		}
	}()

	logger.Info("Executing Activity UpdateOrderStatus", zap.String("OrderID", orderDetails.ID))

	err = workflow.ExecuteActivity(ctxWithRetry, orderService.UpdateOrderStatus, orderDetails.ID, utils.OrderStatusProcessing).Get(ctx, nil)
	if err != nil {
		logger.Error("Error while executing UpdateOrderStatusActivity", zap.Error(err))

		return err
	}

	logger.Info("Executing Activity MakePayment", zap.String("OrderID", orderDetails.ID))

	err = workflow.ExecuteActivity(ctxWithRetry, paymentService.MakePayment, orderDetails.ID).Get(ctx, nil)
	if err != nil {
		logger.Error("Error while executing ProcessPaymentActivity", zap.Error(err))

		return err
	}

	logger.Info("Executing Activity ShipOrder", zap.String("OrderID", orderDetails.ID))

	err = workflow.ExecuteActivity(ctxWithRetry, orderService.ShipOrder, orderDetails.ID).Get(ctx, nil)
	if err != nil {
		logger.Error("Error while executing ShipOrderActivity", zap.Error(err))

		logger.Info("Executing Compensation Activity RefundPayment", zap.String("OrderID", orderDetails.ID))

		err = workflow.ExecuteActivity(ctxWithRetry, paymentService.RefundPayment, orderDetails.ID).Get(ctx, nil)
		if err != nil {
			logger.Error("Error while executing RefundPaymentActivity", zap.Error(err))
		}

		return err
	}

	logger.Info("Executing Activity SendEmail", zap.String("UserID", orderDetails.UserID))

	err = workflow.ExecuteActivity(ctxWithRetry, communicationService.SendEmail, orderDetails.UserID).Get(ctx, nil)
	if err != nil {
		logger.Error("Error while executing SendEmailActivity", zap.Error(err))

		return err
	}

	logger.Info("Executing Activity UpdateOrderStatus", zap.String("OrderID", orderDetails.ID))

	err = workflow.ExecuteActivity(ctxWithRetry, orderService.UpdateOrderStatus, orderDetails.ID, utils.OrderStatusCompleted).Get(ctx, nil)
	if err != nil {
		logger.Error("Error while executing UpdateOrderStatusActivity", zap.Error(err))

		return err
	}

	logger.Info("Workflow ProcessAnOrder completed", zap.String("OrderID", orderDetails.ID))

	return nil
}
