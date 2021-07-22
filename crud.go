package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func main() {
	e := echo.New()
	e.GET("/", GetUsers)
	e.GET("/:id", GetOneUsers)
	e.POST("/", PostUser)
	e.PUT("/:id", UpdateUser)
	e.DELETE("/:id", DeleteUser)

	users = []User{
		{Id: 1, Name: "Farras", Email: "farras@gmail.com", Password: "12345"},
		{Id: 2, Name: "Muti", Email: "MutMut@gmail.com", Password: "hgdfhfd"},
		{Id: 3, Name: "Wildan", Email: "Wawew@gmail.com", Password: "uewewewe"},
	}
	e.Logger.Fatal(e.Start(":8080"))
}

func PostUser(c echo.Context) error {
	createUser := User{}
	c.Bind(&createUser)
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    createUser,
	})
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetOneUsers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	id -= 1
	if id == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users[id],
	})
}

func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	id -= 1
	if id == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	update := User{}
	c.Bind(&update)
	users[id] = update
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users[id],
	})
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			if i == len(users)-1 {
				users = users[:len(users)-1]
				return c.JSON(http.StatusOK, map[string]interface{}{
					"messages": "success get all users",
					"users":    users,
				})
			}
			users = users[i+1:]
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get all users",
				"users":    users,
			})
		}

	}
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "invalid id",
	})
}
