package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello, Dumble!")

	// Open Connecrion to Database
	conn, err := dbr.Open("sqlite3", "./dumble.db", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	sess := conn.NewSession(nil)
	fmt.Println("Database connection established.")

	/*
		dumble, err := getDumble(sess, 3)
		fmt.Println("Fetched Dumble:", dumble, "Error:", err)

		out, _ := json.MarshalIndent(dumble, "", "  ")
		fmt.Println(string(out))
	*/

	// Start Echo Server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the Dumble API!")
	})

	e.GET("/dumble/:id", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}
		dumble, err := getDumble(sess, id)
		return c.JSON(http.StatusOK, map[string]interface{}{"data": dumble, "error": err})
	})

	e.GET("/dumbles", func(c echo.Context) error {
		dumbles, err := getAllDumbles(sess)
		return c.JSON(http.StatusOK, map[string]interface{}{"data": dumbles, "error": err})
	})

	// Add Likes by id
	e.GET("/dumble/:id/like", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}
		_, err = sess.Update("dumbles").Set("likes", dbr.Expr("likes + 1")).Where("dumble_id = ?", id).Exec()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to add like"})
		}
		dumble, err := getDumble(sess, id)
		return c.JSON(http.StatusOK, map[string]interface{}{"data": dumble, "error": err})
	})

	e.Logger.Fatal(e.Start(":8080"))

	fmt.Println("Goodbye")

}
