package models

import "goMysql/db"

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}
type Users []User

const UserSchema string = `CREATE TABLE users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(127) NOT NULL,
	email VARCHAR(64),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

//Crear usuario e inseratar a db

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.insert()
	return user
}

func (user *User) insert() {
	sql := "INSERT INTO users SET username = ?, password = ?, email = ?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}

func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id = ?"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}

func (user *User) update() {
	sql := "UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id = ?"
	db.Exec(sql, user.Id)
}
