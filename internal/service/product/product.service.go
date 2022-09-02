package productsvc

import (
	"basic-crud/internal/datastore"

	"gorm.io/gorm"
)

func ProductDetail(dbConn *gorm.DB, id int) (resp *datastore.Product, err error) {
	return resp, err
}
