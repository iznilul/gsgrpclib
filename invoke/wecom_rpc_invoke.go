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
