package routes

import (
	"echo-api/controllers"
	"echo-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UserGroup(g *echo.Group) {
	g.POST("", createUser)
	g.GET("/:id", findUser)
	g.GET("", listUsers)
	g.PUT("/:id", updateUser)
	g.DELETE("/:id", deleteUser)

}

func createUser(c echo.Context) error {
	var userReq models.User
	if err := c.Bind(&userReq); err != nil {
		return err
	}

	if err := controllers.CreateUser(&userReq); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": userReq})
}

func updateUser(c echo.Context) error {

	userIDString := c.Param("id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	var userReq models.User
	if err := c.Bind(&userReq); err != nil {
		return err
	}

	if err := controllers.UpdateUser(uint32(userID), &userReq); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "update user successfully."})
}

func deleteUser(c echo.Context) error {
	userIDString := c.Param("id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	err = controllers.DeleteUser(uint32(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "delete user successfully."})
}

func findUser(c echo.Context) error {
	userIDString := c.Param("id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	user, err := controllers.FindUser(uint32(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "find user", "data": user})
}

func listUsers(c echo.Context) error {
	perPageStr := c.QueryParam("perPage")
	pageStr := c.QueryParam("page")
	search := c.QueryParam("search")

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		perPage = 10
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	users, err := controllers.ListUser(search, uint32(page), uint32(perPage))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "list user ", "data": users})
}
