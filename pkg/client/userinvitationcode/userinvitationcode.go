//nolint:nolintlint,dupl
package userinvitationcode

import (
	"context"
	"time"

	"google.golang.org/grpc"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second
var IsTest = true

type handler func(context.Context, npool.CloudHashingInspireClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var conn *grpc.ClientConn
	var err error

	conn, err = grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewCloudHashingInspireClient(conn)

	return handler(_ctx, cli)
}

func GetUserInvitationCodeOnly(ctx context.Context, conds *npool.Conds) (*npool.UserInvitationCode, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.GetUserInvitationCodeByCode(ctx, &npool.GetUserInvitationCodeByCodeRequest{
			Code: conds.GetInvitationCode().GetValue(),
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.UserInvitationCode), nil
}
