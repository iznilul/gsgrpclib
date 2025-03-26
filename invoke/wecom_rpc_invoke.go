package invoke

import (
	"context"
	"github.com/iznilul/gsgrpclib/client"
	"github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcGetUserInfo(ctx context.Context) (string, error) {
	code := ctx.Value("code")
	data, err := utils.ParseDataToAny(code)
	if err != nil {
		return "", err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserInfo", ao)
	if err != nil {
		return "", err
	}
	return vo.Msg, nil
}

func InvokeRpcSendTextMsg(ctx context.Context) error {
	requestUserID := ctx.Value("requestUserID")
	userID := ctx.Value("userID")
	msg := ctx.Value("msg")
	map1 := map[string]interface{}{
		"requestUserID": requestUserID,
		"userID":        userID,
		"message":       msg,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SendTextMsg", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSendWarnMsg(ctx context.Context) error {
	userID := ctx.Value("userID")
	msg := ctx.Value("msg")
	map1 := map[string]interface{}{
		"userID":  userID,
		"message": msg,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SendWarnMsg", ao)
	if err != nil {
		return err
	}
	return nil
}
