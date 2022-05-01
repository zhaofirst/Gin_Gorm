package handlers

import (
	"demo/dbfile"
	"demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUser is handler for create user.
func CreateUser(dc *dbfile.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		err := ctx.Bind(&user)
		if err != nil {
			ctx.String(http.StatusBadRequest, "bind user error: %v", err)
			return
		}
		err = dc.CreateUser(user)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "create user error: %v", err)
			return
		}
		ctx.Status(http.StatusCreated)
		ctx.String(http.StatusCreated, "Successfully created")
	}
}

// ReadUser is handler for read user.
func ReadUser(dc *dbfile.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam) // idparam is a int value
		if err != nil {
			ctx.String(http.StatusBadRequest, "parse param error: %v", err)
			return
		}

		user, err := dc.ReadUser(id)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "read user error: %v", err)
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

// UpdateUser is handler for update user.
func UpdateUser(dc *dbfile.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		err := ctx.Bind(&user)
		if err != nil {
			ctx.String(http.StatusBadRequest, "bind user error: %v", err)
			return
		}

		err = dc.UpdateUser(user)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "update user error: %v", err)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// DeleteUser is handler for delete user.
func DeleteUser(dc *dbfile.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.String(http.StatusBadRequest, "parse param error: %v", err)
			return
		}

		err = dc.DeleteUser(id)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "delete user error: %v", err)
			return
		}

		ctx.Status(http.StatusOK)
		ctx.String(http.StatusOK, "Successfully deleted")
	}
}
