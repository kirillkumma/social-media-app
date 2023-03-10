package entity

import "fmt"

type User struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func (u User) Validate() error {
	if len(u.Name) < 4 {
		return fmt.Errorf("min length of name is 4 charachters")
	}
	return nil
}

func NewUser(name string) (User, error) {
	u := User{
		Name: name,
	}

	err := u.Validate()
	if err != nil {
		return User{}, err
	}

	return u, nil
}
