package domain

import (
	"errors"
)

func (user *User) Save() error {
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

func (user *User) Get() error {
	stmt, err := dbClient.Prepare("SELECT id, email FROM users WHERE id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return err
		}
		return nil
	}
	return errors.New("user not found")
}
