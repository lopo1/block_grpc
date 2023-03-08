package main

import (
	"block_tool/log"
	pb "block_tool/user/proto"
	"block_tool/user/server/core"
	"block_tool/user/server/global"
	"block_tool/user/server/initialize"
	"block_tool/user/server/middleware"
	"block_tool/user/server/service"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
)

//func (s *SearchService) mustEmbedUnimplementedSearchServiceServer() {
//
//}

const PORT = "9001"

func init() {
	core.Viper()
	log.ZapInterceptor()
	//model.Setup()
	global.GVA_DB = initialize.Gorm()
}
func main() {
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			middleware.RecoveryInterceptor,
			middleware.LoggingInterceptor,
		),
	}
	server := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(server, &service.SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		//log.Fatalf("net.Listen err: %v", err)
		global.GVA_LOG.Fatal("net.Listen err:", err)
	}

	server.Serve(lis)
}
