FROM golang:1.17.7-alpine3.15 AS building-stage
WORKDIR /go_service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM ubuntu:18.04
ARG ENVIRONMENT=dev
RUN apt update && apt install -y wget vim curl unzip groff less && \
    VERSION=v4.24.5 && BINARY=yq_linux_amd64 && \
    wget https://github.com/mikefarah/yq/releases/download/${VERSION}/${BINARY} -O /usr/bin/yq &&\
    chmod +x /usr/bin/yq && \
    curl -LO "https://dl.k8s.io/release/v1.23.6/bin/linux/amd64/kubectl" && \
    chmod +x ./kubectl && \
    mv ./kubectl /usr/local/bin/kubectl && \
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64-2.0.30.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && ./aws/install -i /usr/local/aws-cli -b /usr/local/bin

COPY --from=building-stage /go_service/.aws /root/.aws
COPY --from=building-stage /go_service/orderServiceCadenceWorker /
COPY --from=building-stage /go_service/deploy /
RUN cd /root && pwd && ls -la /root/.aws && \
    aws eks update-kubeconfig --kubeconfig ~/.kube/config --region us-east-2 --name cdPaas-${ENVIRONMENT}-eks-cluster
CMD ["/orderServiceCadenceWorker"]
