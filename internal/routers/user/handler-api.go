package userrtr

import (
	userctr "basic-crud/internal/controller/user"
	db "basic-crud/internal/db/postgres"
	"basic-crud/internal/models"
	usermdl "basic-crud/internal/models/user"
	"basic-crud/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var userCtr userctr.UserCtr

type UserHandler struct{}

func (*UserHandler) getUsers(c *gin.Context) {
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

	resp, count, err := userCtr.Users(dbConn, req)

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

func (*UserHandler) getUser(c *gin.Context) {
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

	resp, err := userCtr.UserDetail(dbConn, req)

	if err != nil {
		resp = nil
	}

	c.JSON(200, models.HTTPResponse{
		Success: true,
		Data:    resp,
	})
}

func (*UserHandler) createUser(c *gin.Context) {
	req := usermdl.UserRequest{}
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

	err = userCtr.UserCreate(dbConn, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Failed create user",
		})
		c.Abort()
		return
	}

	c.JSON(200, models.HTTPResponse{
		Success: true,
	})
}

func (*UserHandler) updateUser(c *gin.Context) {
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

	req := usermdl.UserRequest{}
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

	err = userCtr.UserUpdate(dbConn, param.ID, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Failed update user",
		})
		c.Abort()
		return
	}

	c.JSON(200, models.HTTPResponse{
		Success: true,
	})
}

func (*UserHandler) deleteUser(c *gin.Context) {
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

	err = userCtr.UserDelete(dbConn, param.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Failed delete role",
		})
		c.Abort()
		return
	}

	c.JSON(200, models.HTTPResponse{
		Success: true,
	})
}
