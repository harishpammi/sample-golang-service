package main

import (
	"fmt"
	"log"
	"sample-service/resource"

	"github.com/gocql/gocql"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Service started")

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "golang_users"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	var firstname string
	var lastname string
	var city string
	var state string

	if err := session.Query(`SELECT firstname,lastname,city,state FROM golang_users.users where id=? and deleted=false ALLOW FILTERING`,
		"users_0001").Consistency(gocql.One).Scan(&firstname, &lastname, &city, &state); err != nil {
		log.Fatal("Error ", err)
	}

	fmt.Println("Firstname", firstname)
	fmt.Println("Lastname", lastname)
	fmt.Println("City", city)
	fmt.Println("State", state)
	e := echo.New()
	e.GET("/users", resource.GetAllUsersEndPoint)

	e.Logger.Fatal(e.Start(":8070"))
}
