package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-stimulate/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedCloudHashingStimulateServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterCloudHashingStimulateServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterCloudHashingStimulateHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
