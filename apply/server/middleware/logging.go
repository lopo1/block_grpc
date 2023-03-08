package middleware

import (
	"block_tool/user/server/global"
	"context"
	"google.golang.org/grpc"
	"time"
)

type Logging struct {
	Method string
	Req    interface{}
	Resp   interface{}
	Time   int64
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	//log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	elapsed := time.Since(start)
	defer func() {
		logging := Logging{
			Method: info.FullMethod,
			Req:    req,
			Resp:   resp,
			Time:   elapsed.Nanoseconds(),
		}
		global.GVA_LOG.Infof("[%s]  %v\n", global.GVA_CONFIG.System.Name, logging)
		//log.Logger.Info("[GRPC]:", zap.Any("logging", logging))
	}()
	//log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}
