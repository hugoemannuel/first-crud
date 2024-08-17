package service

import (
	"first-crud/connection"
	"first-crud/dto"
	"log"
)

func GetAll() []dto.User {
	db := connection.Connect()
	rows, err := db.Query("SELECT id, name, email, password FROM user")
	if err != nil {
		log.Fatal(err)
	}
	users := make([]dto.User, 0)
	for rows.Next() {
		var user dto.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	defer db.Close()

	return users
}

func Create(name string, password string, email string) (dto.User, error) {
	db := connection.Connect()
	result, errQuery := db.Exec("INSERT INTO user (name, password, email) VALUES (?, ?, ?)", name, password, email)
	if errQuery != nil {
		return dto.User{}, errQuery
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return dto.User{}, err
	}
	user := FindOne(int(userID))

	defer db.Close()

	return user, nil
}

func FindOne(id int) dto.User {
	db := connection.Connect()
	rows, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	users := dto.User{}
	for rows.Next() {
		var user dto.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
			log.Fatal(err)
		}
		users = user
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	defer db.Close()

	return users
}

func Exclude(id int) error {
	db := connection.Connect()
	_, err := db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}
