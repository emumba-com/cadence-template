# orderService Cadence Worker

**Register cadence domain**

`sudo docker run --network=host --rm ubercadence/cli:master --do orderServiceDomain domain register`

**Build a compiled binary executable** 

`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .`


**Building image**

`docker build -t <registry-path>:<tag> .`

**Pushing image**

`docker push <registry-path>:<tag>`


**ENV**

```
BUILD_ENV=dev

# cadence worker
CADENCE_SERVICE=<cadence-host:cadence-port> e.g localhost:7933
TASK_LIST_NAME=orderServiceTaskList
CADENCE_DOMAIN=orderServiceDomain
CADENCE_SERVICE_NAME=cadence-frontend
CLIENT_NAME=cadence-service-worker

RETRY_POLICY_INITIAL_INTERVAL_SECONDS=2
RETRY_POLICY_BACKOFF_COEFFICIENT=5
RETRY_POLICY_MAXIMUM_ATTEMPTS=3
LAST_HEALTH_PING_TIMEOUT=1


# internal services
USERS_SERVICE_ADDRESS=<asset-service-address> e.g http://localhost:8086
PAYMENT_SERVICE_ADDRESS=<superset-service-address> e.g http://localhost:9000
ORDERS_SERVICE_ADDRESS=<chatAgent-service-address> e.g  http://localhost:8084