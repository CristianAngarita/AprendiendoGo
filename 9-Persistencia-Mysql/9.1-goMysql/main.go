package main

import (
	"fmt"
	"goMysql/db"
	"goMysql/models"
)

func main() {
	db.Connect()
	//fmt.Print(db.ExistTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	//db.TruncateTable("users")
	//user := models.CreateUser("Cristian2", "Contraseña2", "Cristian2@gmail.com")
	//fmt.Println(user)
	/* users := models.ListUsers()
	fmt.Println(users) */

	/* 	user := models.GetUser(2)
	   	fmt.Println(user)
	   	 	user.Username = "Fredys"
	   	   	user.Password = "fredy123"
	   	   	user.Email = "Fredy@gmail.com"
	   	   	user.Save()

	   	user.Delete() */
	db.TruncateTable("users")
	fmt.Println(models.ListUsers())
	db.Close()
	//db.Ping()
}
