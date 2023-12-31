package orders

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/cadence/client"
	"net/http"
	"orderService/clients/cadenceClient"
	"orderService/env"
	"orderService/log"
	"orderService/models"
	"orderService/utils"
	"time"
)

type Server struct {
	Router        *gin.Engine
	RouterGroup   *gin.RouterGroup
	CadenceClient cadenceClient.CadStore
}

func (server *Server) PlaceAnOrder(ctx *gin.Context) {
	logger := log.GetLogger()
	logger.Info("PlaceAnOrder endpoint called")

	var orderDetails models.OrderDetails
	if err := ctx.ShouldBindJSON(&orderDetails); err != nil {
		logger.Error(err.Error())
		utils.BuildResponse(ctx, http.StatusBadRequest, utils.ERROR, err.Error(), nil)

		return
	}

	workflowOptions := client.StartWorkflowOptions{
		//ID:                              wID.String(),
		TaskList:                        env.Env.TaskListName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	err := server.CadenceClient.TriggerProcessOrderWorkflow(
		ctx,
		workflowOptions,
		orderDetails)
	if err != nil {
		logger.Error(err.Error())
		utils.BuildResponse(ctx, http.StatusBadRequest, utils.ERROR, "Couldn't Trigger PlaceAnOrder Workflow", nil)

		return
	}

	utils.BuildResponse(ctx, http.StatusCreated, utils.SUCCESS, "", "order placed")
	logger.Info("PlaceAnOrder endpoint returned successfully")
}
