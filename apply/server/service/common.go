package service

import (
	pb "block_tool/user/proto"
	"block_tool/user/server/global"
	"block_tool/user/server/model/request"
	"block_tool/utils"
)

type SearchService struct {
	*pb.UnimplementedSearchServiceServer
}

// TokenNext 登录以后签发jwt
func TokenNext(req request.BaseClaims) (token string, err error) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(req)
	return j.CreateToken(claims)

	//if !global.GVA_CONFIG.System.UseMultipoint {
	//	response.OkWithDetailed(systemRes.LoginResponse{
	//		User:      user,
	//		Token:     token,
	//		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	//	}, "登录成功", c)
	//	return
	//}
	//
	//if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
	//	if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
	//		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
	//		response.FailWithMessage("设置登录状态失败", c)
	//		return
	//	}
	//	response.OkWithDetailed(systemRes.LoginResponse{
	//		User:      user,
	//		Token:     token,
	//		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	//	}, "登录成功", c)
	//} else if err != nil {
	//	global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
	//	response.FailWithMessage("设置登录状态失败", c)
	//} else {
	//	var blackJWT system.JwtBlacklist
	//	blackJWT.Jwt = jwtStr
	//	if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
	//		response.FailWithMessage("jwt作废失败", c)
	//		return
	//	}
	//	if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
	//		response.FailWithMessage("设置登录状态失败", c)
	//		return
	//	}
	//	response.OkWithDetailed(systemRes.LoginResponse{
	//		User:      user,
	//		Token:     token,
	//		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	//	}, "登录成功", c)
	//}
}
