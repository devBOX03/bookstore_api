package users

import (
	"net/http"
	"strconv"

	"github.com/devBOX03/bookstore_user_api/domain/users"
	"github.com/devBOX03/bookstore_user_api/services"
	"github.com/devBOX03/bookstore_user_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserId(c *gin.Context) (int64, *errors.RestError) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("User id should be number")
	}
	return userId, nil

}

func Create(c *gin.Context) {
	var user users.User
	if jsonErr := c.ShouldBindJSON(&user); jsonErr != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, serviceErr := services.CreateUser(user)
	if serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, getUserErr := getUserId(c)
	if getUserErr != nil {
		c.JSON(getUserErr.Status, getUserErr)
		return
	}
	user, serviceErr := services.GetUser(userId)
	if serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userId, getUserErr := getUserId(c)
	if getUserErr != nil {
		c.JSON(getUserErr.Status, getUserErr)
		return
	}

	var user users.User
	if jsonErr := c.ShouldBindJSON(&user); jsonErr != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
	}
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch
	result, serviceErr := services.UpdateUser(isPartial, user)
	if serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, getUserErr := getUserId(c)
	if getUserErr != nil {
		c.JSON(getUserErr.Status, getUserErr)
		return
	}
	if serviceErr := services.DeleteUser(userId); serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		err := errors.NewBadRequestError("status is missing")
		c.JSON(err.Status, err)
		return
	}

	result, serviceErr := services.SearchUser(status)
	if serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
