package env

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type envFile struct {
	BuildEnv                      string
	CadenceService                string
	TaskListName                  string
	CadenceDomain                 string
	ClientName                    string
	CadenceServiceName            string
	LastHealthPingTimeout         int
	RetryPolicyInitialInterval    int
	RetryPolicyBackoffCoefficient float64
	RetryPolicyMaxAttempts        int32
	UserServiceAddress            string
	UserServiceHost               string
	UserServiceScheme             string
}

func getEnvValue(key string, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}

	return defaultValue
}

var Env *envFile

// nolint
func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error loading .env file")
	}

	buildEnv := getEnvValue("BUILD_ENV", "dev")

	lastHealthPingTimeout, err := strconv.Atoi(getEnvValue("LAST_HEALTH_PING_TIMEOUT", "1"))
	if err != nil {
		log.Println(err.Error())
	}

	retryPolicyInitialInterval, err := strconv.Atoi(getEnvValue("RETRY_POLICY_INITIAL_INTERVAL_SECONDS", "2"))
	if err != nil {
		log.Println(err.Error())
	}

	retryPolicyBackoffCoefficient, err := strconv.ParseFloat(getEnvValue("RETRY_POLICY_BACKOFF_COEFFICIENT", "5"), 64)
	if err != nil {
		log.Println(err.Error())
	}

	retryPolicyMaxAttempts, err := strconv.ParseInt(getEnvValue("RETRY_POLICY_MAXIMUM_ATTEMPTS", "3"), 10, 32)
	if err != nil {
		log.Println(err.Error())
	}

	cadenceService := getEnvValue("CADENCE_SERVICE", "localhost:7933")

	taskListName := getEnvValue("TASK_LIST_NAME", "orderServiceCadenceTaskList")

	cadenceDomain := getEnvValue("CADENCE_DOMAIN", "orderServiceCadenceDomain")

	clientName := getEnvValue("CLIENT_NAME", "cadence-service-worker")

	cadenceServiceName := getEnvValue("CADENCE_SERVICE_NAME", "cadence-frontend")

	userServiceAddress := getEnvValue("USER_SERVICE_ADDRESS", "http://localhost:8093")
	userServiceURL, err := url.Parse(userServiceAddress)
	if err != nil {
		log.Println(err.Error())
	}

	Env = &envFile{
		BuildEnv:                      buildEnv,
		LastHealthPingTimeout:         lastHealthPingTimeout,
		RetryPolicyInitialInterval:    retryPolicyInitialInterval,
		RetryPolicyBackoffCoefficient: retryPolicyBackoffCoefficient,
		RetryPolicyMaxAttempts:        int32(retryPolicyMaxAttempts),
		CadenceService:                cadenceService,
		TaskListName:                  taskListName,
		CadenceDomain:                 cadenceDomain,
		ClientName:                    clientName,
		CadenceServiceName:            cadenceServiceName,
		UserServiceAddress:            userServiceAddress,
		UserServiceHost:               userServiceURL.Host,
		UserServiceScheme:             userServiceURL.Scheme,
	}
}
