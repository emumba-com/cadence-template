package utils

import (
	"github.com/go-resty/resty/v2"
	"net/url"
	"orderServiceCadenceWorker/v1/env"
	"time"

	"go.uber.org/cadence"
	"go.uber.org/cadence/workflow"
)

func GetActivityOptions(duration time.Duration) workflow.ActivityOptions {
	initialIntervalSeconds := env.Env.RetryPolicyInitialInterval
	backOffCoefficient := env.Env.RetryPolicyBackoffCoefficient
	maxAttempts := env.Env.RetryPolicyMaxAttempts

	retryPolicy := &cadence.RetryPolicy{
		InitialInterval:    time.Second * time.Duration(initialIntervalSeconds),
		BackoffCoefficient: backOffCoefficient,
		MaximumInterval:    time.Second * 30,
		MaximumAttempts:    maxAttempts,
	}

	activityOptions := workflow.ActivityOptions{
		ScheduleToStartTimeout: duration * time.Minute,
		StartToCloseTimeout:    duration * time.Minute,
		//HeartbeatTimeout:       duration,
		RetryPolicy: retryPolicy,
	}

	return activityOptions
}

func GetActivityOptionsWithoutRetry(duration time.Duration) workflow.ActivityOptions {
	activityOptions := workflow.ActivityOptions{
		ScheduleToStartTimeout: duration * time.Minute,
		StartToCloseTimeout:    duration * time.Minute,
	}

	return activityOptions
}

func GetRestyClient() *resty.Client {
	return resty.New()
}

func GenerateURL(relativePath string, queryParam string, scheme string, host string) *url.URL {
	return &url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     relativePath,
		RawQuery: queryParam,
	}
}
