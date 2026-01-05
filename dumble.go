package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocraft/dbr/v2"
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

	dumble, err := getDumble(sess, 3)
	fmt.Println("Fetched Dumble:", dumble, "Error:", err)

	out, _ := json.MarshalIndent(dumble, "", "  ")
	fmt.Println(string(out))

}
