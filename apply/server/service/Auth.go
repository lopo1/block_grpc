package service

import (
	pb "block_tool/user/proto"
	"block_tool/utils"
	"context"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *SearchService) AuthNonce(ctx context.Context, r *pb.AuthNonceRequest) (*pb.Response, error) {
	var msg string
	marshal := &pb.AuthNonceResponse{}
	marshal.Nonce = utils.GenValidateCode(6)
	//if ok := utils.CheckETHAddress(r.Address); !ok {
	//	return &pb.Response{Code: 402, Error: "address invalid"}, nil
	//}

	any, err := anypb.New(marshal)

	if err != nil {
		msg = err.Error()
	}
	return &pb.Response{Code: 200, Data: any, Error: msg}, nil
}
