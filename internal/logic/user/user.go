package user

import (
	"context"
	"github.com/lj1570693659/gfcq_product/internal/dao"
	"github.com/lj1570693659/gfcq_product/internal/model/do"
	"github.com/lj1570693659/gfcq_product/internal/service"

	"github.com/lj1570693659/gfcq_protoc/pbentity"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(&sUser{})
}

func (s *sUser) GetById(ctx context.Context, uid uint64) (*pbentity.User, error) {
	var user *pbentity.User
	err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Scan(&user)
	return user, err
}

func (s *sUser) DeleteById(ctx context.Context, uid uint64) error {
	_, err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Delete()
	return err
}
