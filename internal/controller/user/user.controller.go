package rolectr

import (
	"basic-crud/internal/datastore"
	"basic-crud/internal/models"
	usermdl "basic-crud/internal/models/user"

	"gorm.io/gorm"
)

type UserCtr struct{}

func (*UserCtr) Users(dbConn *gorm.DB, req models.Pagination) (resp *[]datastore.User, count int64, err error) {
	userds := datastore.User{}
	users, count, err := userds.List(dbConn, &req)
	return users, count, err
}

func (*UserCtr) UserDetail(dbConn *gorm.DB, req models.HTTPAPIParamReq) (resp *datastore.User, err error) {
	userds := datastore.User{}
	user, err := userds.Get(dbConn, req.ID)
	return &user, err
}

func (*UserCtr) UserCreate(dbConn *gorm.DB, req usermdl.UserRequest) (err error) {
	userds := datastore.User{}
	userds.Name = req.Name
	userds.RoleID = req.RoleID
	err = userds.Save(dbConn, userds)

	if err != nil {
		return err
	}

	return
}

func (*UserCtr) UserUpdate(dbConn *gorm.DB, id int, req usermdl.UserRequest) (err error) {
	userds := datastore.User{}
	userds.Name = req.Name
	userds.RoleID = req.RoleID
	err = userds.Update(dbConn, id, userds)

	if err != nil {
		return err
	}

	return
}

func (*UserCtr) UserDelete(dbConn *gorm.DB, id int) (err error) {
	userds := datastore.User{}
	err = userds.Delete(dbConn, id)

	if err != nil {
		return err
	}

	return
}
