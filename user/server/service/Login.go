package service

import (
	pb "block_tool/user/proto"
	"block_tool/user/server/global"
	"block_tool/user/server/model"
	"block_tool/user/server/model/request"
	"block_tool/utils"
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/anypb"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

func (s *SearchService) Login(ctx context.Context, r *pb.LoginRequest) (*pb.Response, error) {
	var login LoginMethod
	var err error
	//login.Sign = r.Sign
	getValue := reflect.ValueOf(login)
	method := GetLoginFunc(r.Method)
	// 或者使用方法名称
	m := getValue.MethodByName(method)
	if !m.IsValid() {
		return &pb.Response{Code: 401, Error: err.Error()}, nil
	}

	marshal := &pb.LoginAny{}
	tt := m.Call([]reflect.Value{reflect.ValueOf(marshal), reflect.ValueOf(r)})
	//data := tt[0].Interface().(map[string]interface{})
	if !tt[0].IsNil() {
		err = tt[0].Interface().(error)
		return &pb.Response{Code: 402, Error: err.Error()}, nil
	}

	any, err := anypb.New(marshal)

	var msg string
	if err != nil {
		msg = err.Error()
	}
	return &pb.Response{Code: 200, Data: any, Error: msg}, nil
}

type LoginMethod struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	Sign        string
	NickName    string
	AuthorityId uint
}

func GetLoginFunc(method string) string {
	if method == "" {
		return "Block"
	}
	return utils.FirstUpper(method)
}

func (l LoginMethod) Block(r *pb.LoginAny, req *pb.LoginRequest) (err error) {
	var user model.User
	if ok := utils.CheckETHAddress(req.Address); !ok {
		err = errors.New("invalid address")
		return
	}
	ok, err := utils.CheckSign(req.Code, req.Sign, req.Address)
	if err != nil {
		return
	}
	if !ok {
		return errors.New("invalid sign")
	}
	req.Address = strings.ToLower(req.Address)
	err = global.GVA_DB.Where("address = ?", req.Address).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if user.Nonce == req.Code || req.Code == "" {
		err = errors.New("nonce invalid")
		return
	}
	if user.Id == 0 {
		user.Address = req.Address
		user.Created = time.Now().Unix()
		user.Uuid = utils.Uuid()
	}

	user.Nonce = req.Code
	user.UpData = time.Now().Unix()
	err = global.GVA_DB.Save(&user).Error
	if err != nil {
		return
	}
	token, err := TokenNext(request.BaseClaims{
		UUID:     uuid.MustParse(user.Uuid),
		ID:       user.Id,
		Username: user.Address,
		NickName: user.Username,
	})
	if err != nil {
		return
	}
	r.Id = user.Id
	r.Username = user.Username
	r.Address = req.Address
	r.Token = token
	return
}
func (l LoginMethod) Phone(r *pb.LoginAny) (err error) {
	r.Id = 2
	r.Username = "test"
	r.Address = "0x1111111"
	r.Token = "tsttttttttt"
	return
}
func (l LoginMethod) Email(r *pb.LoginAny) (err error) {
	r.Id = 2
	r.Username = "test"
	r.Address = "0x1111111"
	r.Token = "tsttttttttt"
	return
}
