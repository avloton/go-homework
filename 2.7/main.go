/*
Создать SQLite Базу данных с таблицей users. Атрибуты таблицы: id, email, password.
Заполнить таблицу случайными данными.

Написать модуль на ЯП Go, принимающий на вход пару логин-пароль, возвращаюший ответ в зависимости от входных данных:
-пользователь не найден
-пароль неверный
-пароль верный
*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	_"modernc.org/sqlite"
)

type User struct {
	id int
	email string
	password string
}

func getUser(db *sql.DB, email string) (*User, error) {
	row := db.QueryRow("select id, email, password from users where email = ?", email)
	u := User{}
	err := row.Scan(&u.id, &u.email, &u.password)
	if err != nil {
		return &User{}, err
	}
	return &u, nil
}

func main() {
	db, err := sql.Open("sqlite", "./mydb.db")
	if err != nil {
		log.Fatal("Failed to open database")
	}
	defer db.Close()

	//Задаем логин и пароль
	email := "user1@mail.ru"
	password := "123"

	//Проверяем логин и пароль
	user, err := getUser(db, email)
	if err != nil {
		fmt.Println("Пользователь не найден")
		return
	}
	if user.password != password {
		fmt.Println("Пароль неверный")
	} else {
		fmt.Println("Пароль верный")
	}
}