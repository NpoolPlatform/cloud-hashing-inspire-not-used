package userinvitationcode

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

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
	infos, err := db.Client().
		UserInvitationCode.
		Query().
		Where(
			userinvitationcode.And(
				userinvitationcode.InvitationCode(code),
			),
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

	info, err := db.Client().
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
	return nil, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetUserInvitationCodeByAppUserRequest) (*npool.GetUserInvitationCodeByAppUserResponse, error) {
	return nil, nil
}
