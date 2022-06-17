package userinvitationcode

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userinvitationcode" //nolint

	"github.com/AmirSoleimani/VoucherCodeGenerator/vcgen"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateUserInvitationCode(info *npool.UserInvitationCode) error {
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invlaid app id: %v", err)
	}
	return nil
}

func generateInvitationCode() (string, error) {
	vc := vcgen.New(&vcgen.Generator{
		Count:   1,
		Pattern: "###-###-####",
		Charset: "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM",
	})
	codes, err := vc.Run()
	if err != nil {
		return "", xerrors.Errorf("fail run invitation code generator: %v", err)
	}
	return (*codes)[0], nil
}

func codeExist(ctx context.Context, code string) (bool, error) {
	cli, err := db.Client()
	if err != nil {
		return false, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserInvitationCode.
		Query().
		Where(
			userinvitationcode.InvitationCode(code),
		).
		All(ctx)
	if err != nil {
		return false, xerrors.Errorf("fail query invitation code: %v", err)
	}
	if len(infos) == 0 {
		return false, nil
	}
	return true, nil
}

func dbRowToUserInvitationCode(row *ent.UserInvitationCode) *npool.UserInvitationCode {
	return &npool.UserInvitationCode{
		ID:             row.ID.String(),
		UserID:         row.UserID.String(),
		AppID:          row.AppID.String(),
		InvitationCode: row.InvitationCode,
		CreateAt:       row.CreateAt,
	}
}

func Create(ctx context.Context, in *npool.CreateUserInvitationCodeRequest) (*npool.CreateUserInvitationCodeResponse, error) {
	if err := validateUserInvitationCode(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	var code string
	var err error

	for {
		code, err = generateInvitationCode()
		if err != nil {
			return nil, xerrors.Errorf("fail generate invitation code: %v", err)
		}

		exist, err := codeExist(ctx, code)
		if err != nil {
			return nil, err
		}
		if exist {
			continue
		}
		break
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserInvitationCode.
		Create().
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetInvitationCode(code).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create invitation code: %v", err)
	}

	return &npool.CreateUserInvitationCodeResponse{
		Info: dbRowToUserInvitationCode(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserInvitationCodeRequest) (*npool.GetUserInvitationCodeResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserInvitationCode.
		Query().
		Where(
			userinvitationcode.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user invitation code: %v", err)
	}
	var info *npool.UserInvitationCode
	if len(infos) > 0 {
		info = dbRowToUserInvitationCode(infos[0])
	}

	return &npool.GetUserInvitationCodeResponse{
		Info: info,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetUserInvitationCodeByAppUserRequest) (*npool.GetUserInvitationCodeByAppUserResponse, error) {
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return nil, xerrors.Errorf("invlaid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserInvitationCode.
		Query().
		Where(
			userinvitationcode.And(
				userinvitationcode.AppID(uuid.MustParse(in.GetAppID())),
				userinvitationcode.UserID(uuid.MustParse(in.GetUserID())),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user invitation code: %v", err)
	}

	var info *npool.UserInvitationCode
	if len(infos) > 0 {
		info = dbRowToUserInvitationCode(infos[0])
	}

	return &npool.GetUserInvitationCodeByAppUserResponse{
		Info: info,
	}, nil
}

func GetByCode(ctx context.Context, in *npool.GetUserInvitationCodeByCodeRequest) (*npool.GetUserInvitationCodeByCodeResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserInvitationCode.
		Query().
		Where(
			userinvitationcode.InvitationCode(in.GetCode()),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user invitation code: %v", err)
	}

	var info *npool.UserInvitationCode
	if len(infos) > 0 {
		info = dbRowToUserInvitationCode(infos[0])
	}

	return &npool.GetUserInvitationCodeByCodeResponse{
		Info: info,
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetUserInvitationCodesRequest) (*npool.GetUserInvitationCodesResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserInvitationCode.
		Query().
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user invitation code: %v", err)
	}

	codes := []*npool.UserInvitationCode{}
	for _, info := range infos {
		codes = append(codes, dbRowToUserInvitationCode(info))
	}

	return &npool.GetUserInvitationCodesResponse{
		Infos: codes,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetUserInvitationCodesByAppRequest) (*npool.GetUserInvitationCodesByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserInvitationCode.
		Query().
		Where(
			userinvitationcode.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user invitation code: %v", err)
	}

	codes := []*npool.UserInvitationCode{}
	for _, info := range infos {
		codes = append(codes, dbRowToUserInvitationCode(info))
	}

	return &npool.GetUserInvitationCodesByAppResponse{
		Infos: codes,
	}, nil
}
