package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type envFile struct {
	BuildEnv   string
	ServerPort string

	CadenceService           string
	TaskListName             string
	CadenceDomain            string
	ClientName               string
	CadenceServiceName       string
	CadenceWorkerServiceName string
}

func getEnvValue(key string, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}

	return defaultValue
}

var Env *envFile

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error loading .env file")
	}

	serverPort := getEnvValue("SERVER_PORT", "3001")

	buildEnv := getEnvValue("BUILD_ENV", "dev")

	cadenceService := getEnvValue("CADENCE_SERVICE", "localhost:7933")

	taskListName := getEnvValue("TASK_LIST_NAME", "orderServiceCadenceTaskList")

	cadenceDomain := getEnvValue("CADENCE_DOMAIN", "orderServiceCadenceDomain")

	clientName := getEnvValue("CLIENT_NAME", "cadence-service-worker")

	cadenceServiceName := getEnvValue("CADENCE_SERVICE_NAME", "cadence-frontend")

	cadenceWorkerServiceName := getEnvValue("CADENCE_WORKER_SERVICE_NAME", "orderServiceCadenceWorker")

	Env = &envFile{
		BuildEnv:                 buildEnv,
		ServerPort:               serverPort,
		CadenceService:           cadenceService,
		TaskListName:             taskListName,
		CadenceDomain:            cadenceDomain,
		ClientName:               clientName,
		CadenceServiceName:       cadenceServiceName,
		CadenceWorkerServiceName: cadenceWorkerServiceName,
	}
}
