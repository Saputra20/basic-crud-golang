package rolectr

import (
	"basic-crud/internal/datastore"
	"basic-crud/internal/models"
	rolemdl "basic-crud/internal/models/role"

	"gorm.io/gorm"
)

type RoleCtr struct{}

func (*RoleCtr) Roles(dbConn *gorm.DB, req models.Pagination) (resp *[]datastore.Role, count int64, err error) {
	roleds := datastore.Role{}
	roles, count, err := roleds.List(dbConn, &req)
	return roles, count, err
}

func (*RoleCtr) RoleDetail(dbConn *gorm.DB, req models.HTTPAPIParamReq) (resp *datastore.Role, err error) {
	roleds := datastore.Role{}
	role, err := roleds.Get(dbConn, req.ID)
	return &role, err
}

func (*RoleCtr) RoleCreate(dbConn *gorm.DB, req rolemdl.RoleRequest) (err error) {
	roleds := datastore.Role{}
	roleds.Name = req.Name
	err = roleds.Save(dbConn, roleds)

	if err != nil {
		return err
	}

	return
}

func (*RoleCtr) RoleUpdate(dbConn *gorm.DB, id int, req rolemdl.RoleRequest) (err error) {
	roleds := datastore.Role{}
	roleds.Name = req.Name
	err = roleds.Update(dbConn, id, roleds)

	if err != nil {
		return err
	}

	return
}

func (*RoleCtr) RoleDelete(dbConn *gorm.DB, id int) (err error) {
	roleds := datastore.Role{}
	err = roleds.Delete(dbConn, id)

	if err != nil {
		return err
	}

	return
}
