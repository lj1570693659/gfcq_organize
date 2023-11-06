package user

import (
	"context"
	"database/sql"
	"github.com/lj1570693659/gfcq_product/internal/dao"
	"github.com/lj1570693659/gfcq_product/internal/library"
	"github.com/lj1570693659/gfcq_product/internal/model/do"
	"github.com/lj1570693659/gfcq_product/internal/model/entity"
	"github.com/lj1570693659/gfcq_product/internal/service"
	"time"

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

func (s *sUser) GetInfo(ctx context.Context, info entity.User) (entity.User, error) {
	var user entity.User
	query := dao.User.Ctx(ctx)
	if info.Id > 0 {
		query = query.Where(dao.User.Columns().Id, info.Id)
	}

	if len(info.WorkNumber) > 0 {
		query = query.Where(dao.User.Columns().WorkNumber, info.WorkNumber)
	}

	err := query.Scan(&user)
	if err == sql.ErrNoRows {
		return user, nil
	}
	return user, err
}

func (s *sUser) Create(ctx context.Context, info entity.User) error {
	data := do.User{
		UserName:   info.WorkNumber,
		WorkNumber: info.WorkNumber,
		Password:   library.Encrypt("123456"),
		EmployeeId: info.EmployeeId,
		CreateTime: time.Now().Format("2006-01-02"),
		UpdateTime: time.Now().Format("2006-01-02"),
	}
	_, err := dao.User.Ctx(ctx).Data(data).InsertAndGetId()
	return err
}

func (s *sUser) DeleteById(ctx context.Context, uid uint64) error {
	_, err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Delete()
	return err
}
