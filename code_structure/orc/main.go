package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

var (
	dbClient *sql.DB
)

func init() {
	var err error
	dbClient, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s",
		"root", "root", "127.0.0.1:3306", "users_db"))
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

func main() {
	router := gin.Default()

	router.GET("/users/:id", handlerGetUser)
	router.POST("/users", handlerCreateUser)

	router.Run(":8080")
}

func handlerGetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user id"})
		return
	}

	user, err := GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func handlerCreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid json body"})
		return
	}

	if err := SaveUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func SaveUser(user *User) error {
	if user == nil {
		return errors.New("invalid user to save")
	}
	stmt, err := dbClient.Prepare("INSERT INTO users(email) VALUES(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Email)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = userId
	return nil
}

func GetUser(id int64) (*User, error) {
	stmt, err := dbClient.Prepare("SELECT id, email FROM users WHERE id=?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("user not found")
}
