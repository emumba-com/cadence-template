package cadenceClient

import (
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
	"orderServiceCadenceWorker/v1/env"
)

var HostPort = env.Env.CadenceService
var Domain = env.Env.CadenceDomain
var ClientName = env.Env.ClientName
var CadenceService = env.Env.CadenceServiceName

func BuildCadenceClient() workflowserviceclient.Interface {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(ClientName))
	if err != nil {
		panic("Failed to setup TChannel")
	}

	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: ClientName,
		Outbounds: yarpc.Outbounds{
			CadenceService: {Unary: ch.NewSingleOutbound(HostPort)},
		},
	})
	if err = dispatcher.Start(); err != nil {
		panic("Failed to start dispatcher")
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(CadenceService))
}

func GetNewCadenceClient() (client.Client, error) {
	service := BuildCadenceClient()

	return client.NewClient(service, Domain, &client.Options{}), nil
}
