package rolertr

import (
	rolectr "basic-crud/internal/controller/role"
	db "basic-crud/internal/db/postgres"
	"basic-crud/internal/models"
	rolemdl "basic-crud/internal/models/role"
	"basic-crud/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var roleCrt rolectr.RoleCtr

type RoleHandler struct{}

func (*RoleHandler) getRoles(c *gin.Context) {
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

	resp, count, err := roleCrt.Roles(dbConn, req)

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

func (*RoleHandler) getRole(c *gin.Context) {
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

	resp, err := roleCrt.RoleDetail(dbConn, req)

	if err != nil {
		resp = nil
	}

	c.JSON(200, models.HTTPResponse{
		Success: true,
		Data:    resp,
	})
}

func (*RoleHandler) createRole(c *gin.Context) {
	req := rolemdl.RoleRequest{}
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

	err = roleCrt.RoleCreate(dbConn, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPResponse{
			Success: false,
			Error:   "Failed create role",
		})
		c.Abort()
		return
	}

	c.JSON(200, models.HTTPResponse{
		Success: true,
	})
}

func (*RoleHandler) updateRole(c *gin.Context) {
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

	req := rolemdl.RoleRequest{}
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

	err = roleCrt.RoleUpdate(dbConn, param.ID, req)

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

func (*RoleHandler) deleteRole(c *gin.Context) {
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

	err = roleCrt.RoleDelete(dbConn, param.ID)

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
