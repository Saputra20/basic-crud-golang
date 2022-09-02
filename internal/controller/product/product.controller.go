package productctr

import (
	"basic-crud/internal/datastore"
	"basic-crud/internal/models"

	"gorm.io/gorm"
)

type ProductCtr struct{}

func (*ProductCtr) Products(dbConn *gorm.DB, req models.Pagination) (resp *[]datastore.Product, count int64, err error) {
	productds := datastore.Product{}
	products, count, err := productds.List(dbConn, &req)
	return products, count, err
}

func (*ProductCtr) ProductDetail(dbConn *gorm.DB, req models.HTTPAPIParamReq) (resp *datastore.Product, err error) {
	productds := datastore.Product{}
	product, err := productds.Get(dbConn, req.ID)
	return &product, err
}

func (*ProductCtr) ProductCreate(dbConn *gorm.DB, req models.HTTPAPIBodyReq) (err error) {
	productds := datastore.Product{}
	productds.Name = req.Name
	err = productds.Save(dbConn, productds)

	if err != nil {
		return err
	}

	return
}

func (*ProductCtr) ProductUpdate(dbConn *gorm.DB, id int, req models.HTTPAPIBodyReq) (err error) {
	productds := datastore.Product{}
	productds.Name = req.Name
	err = productds.Update(dbConn, id, productds)

	if err != nil {
		return err
	}

	return
}

func (*ProductCtr) ProductDelete(dbConn *gorm.DB, id int) (err error) {
	productds := datastore.Product{}
	err = productds.Delete(dbConn, id)

	if err != nil {
		return err
	}

	return
}
