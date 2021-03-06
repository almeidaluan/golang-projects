package users

import (
	"bookstore_users_api/infraestructure/mysql/db_global"
	"bookstore_users_api/utils"
	"fmt"
	"strings"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id,first_name,last_name,email,date_created from users WHERE id = ? "
)

func (user *User) Get() *utils.RestError {

	db_global.Init()
	stmt, err := db_global.Client.Prepare(queryGetUser)

	if err != nil {
		return utils.InternalServerError(fmt.Sprintf("Error when trying to get user : %s "), err.Error())
	}
	defer stmt.Close()

	resultGetUser := stmt.QueryRow(user.Id)

	if err := resultGetUser.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return utils.NotFoundError(fmt.Sprintf("Error when trying to get user - NotFound : %d ", user.Id), err.Error())
		}
		return utils.InternalServerError(fmt.Sprintf("Error  when trying get user: %d  %s", user.Id, err.Error()), "Internal Server Error")
	}

	return nil
}

func (user *User) Save() *utils.RestError {
	db_global.Init()
	stmt, err := db_global.Client.Prepare(queryInsertUser)

	if err != nil {
		return utils.InternalServerError(fmt.Sprintf("Error when trying to save user : %s "), err.Error())
	}
	defer stmt.Close()

	resultInsert, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		return utils.InternalServerError(fmt.Sprintf("Error when trying to save user : %s "), err.Error())
	}
	userId, err := resultInsert.LastInsertId()

	if err != nil {
		return utils.InternalServerError(fmt.Sprintf("Error when trying to save user: %s"), err.Error())
	}

	user.Id = userId

	return nil
}

func BulkInsert() *utils.RestError {
	db_global.Init()

	data := []map[string]string{
		{"first_name": "Joao Pedro", "last_name": "Almeida", "email": "j@gmail.com", "date_created": "2020-05-05"},
		{"first_name": "Yasmin", "last_name": "Almeida", "email": "y@gmail.com", "date_created": "2020-05-05"},
	}

	sqlStr := "INSERT INTO users(first_name,last_name,email,date_created)"
	vals := []interface{}{}

	for _, row := range data {
		sqlStr += "(?,?,?,?)"
		vals = append(vals, row["first_name"], row["last_name"], row["email"], row["date_created"])
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]

	stmt, _ := db_global.Client.Prepare(sqlStr)

	defer stmt.Close()

	_, err := stmt.Exec(vals...)

	if err != nil {
		return utils.InternalServerError(fmt.Sprintf("Error when trying to save user"), err.Error())
	}
	return nil
}
