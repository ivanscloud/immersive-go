package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}

var users []User

// ----------- CONTROLLER -----------------

// GET ALL USERS
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success GET All Users",
		"users":    users,
	})
}

// GET USER BY ID

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success Users Details",
		"users":    users[id],
	})
}

// DELETE USER BY ID

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users = append(users[:id], users[id+1:]...)
	return c.NoContent(http.StatusNoContent)
}

// UPDATE USER BY ID

func UpdateUserController(c echo.Context) error {
	newUser := new(User)
	if err := c.Bind(newUser); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if users[id].Name != "" {
		users[id].Name = newUser.Name
	}
	if users[id].Address != "" {
		users[id].Address = newUser.Address
	}
	if users[id].Password != "" {
		users[id].Password = newUser.Password
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success Update User",
		"user[id]": users[id],
	})
}

// CREATE NEW USER

func CreateUserController(c echo.Context) error {
	// Binding Data
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success Create User",
		"user":     user,
	})
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.Logger.Fatal(e.Start(":8000"))
}
