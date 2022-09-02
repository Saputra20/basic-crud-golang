package productrtr

import (
	productctr "basic-crud/internal/controller/product"
	db "basic-crud/internal/db/postgres"
	"basic-crud/internal/models"
	"basic-crud/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var productCrt productctr.ProductCtr

type ProductHandler struct{}

func (*ProductHandler) getProducts(c *gin.Context) {
	req := utils.GeneratePaginationFromRequest(c)

	dbConn, err := db.PGConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Error DB connection",
		})
		c.Abort()
		return
	}

	resp, count, err := productCrt.Products(dbConn, req)

	if err != nil {
		resp = nil
	}

	meta := utils.GeneratePaginateFromResponse(req, count)

	c.JSON(200, models.HTTPResponse{
		Success: true,
		Data:    resp,
		Meta:    meta,
	})
}

func (*ProductHandler) getProduct(c *gin.Context) {
	req := models.HTTPAPIParamReq{}
	err := c.ShouldBindUri(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   err.Error(),
		})
		c.Abort()
		return
	}

	dbConn, err := db.PGConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Error DB connection",
		})
		c.Abort()
		return
	}

	resp, err := productCrt.ProductDetail(dbConn, req)

	if err != nil {
		resp = nil
	}

	c.JSON(200, models.HTTPResponse{
		Success: true,
		Data:    resp,
	})
}

func (*ProductHandler) createProduct(c *gin.Context) {
	req := models.HTTPAPIBodyReq{}
	err := c.ShouldBindBodyWith(&req, binding.JSON)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   err.Error(),
		})
		c.Abort()
		return
	}

	dbConn, err := db.PGConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Error DB connection",
		})
		c.Abort()
		return
	}

	err = productCrt.ProductCreate(dbConn, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Failed create product",
		})
		c.Abort()
		return
	}

	c.JSON(200, models.HTTPResponse{
		Success: true,
	})
}

func (*ProductHandler) updateProduct(c *gin.Context) {
	param := models.HTTPAPIParamReq{}
	err := c.ShouldBindUri(&param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   err.Error(),
		})
		c.Abort()
		return
	}

	req := models.HTTPAPIBodyReq{}
	err = c.ShouldBindBodyWith(&req, binding.JSON)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   err.Error(),
		})
		c.Abort()
		return
	}

	dbConn, err := db.PGConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Error DB connection",
		})
		c.Abort()
		return
	}

	err = productCrt.ProductUpdate(dbConn, param.ID, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Failed update product",
		})
		c.Abort()
		return
	}

	c.JSON(200, models.HTTPResponse{
		Success: true,
	})
}

func (*ProductHandler) deleteProduct(c *gin.Context) {
	param := models.HTTPAPIParamReq{}
	err := c.ShouldBindUri(&param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   err.Error(),
		})
		c.Abort()
		return
	}

	dbConn, err := db.PGConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Error DB connection",
		})
		c.Abort()
		return
	}

	err = productCrt.ProductDelete(dbConn, param.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Failed delete product",
		})
		c.Abort()
		return
	}

	c.JSON(200, models.HTTPResponse{
		Success: true,
	})
}
