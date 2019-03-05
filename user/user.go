package user

import (
	"errors"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID      bson.ObjectId `json:"id" storm:"id"`
	Name    string        `json:"name"`
	Email   string        `json:"email"`
	Address string        `json:"address"`
}

const (
	dbPath = "users.db"
)

var (
	ErrorRecordInvalid = errors.New("Invalid Record")
)

func GetAll() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetOne(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}
	return db.DeleteStruct(u)
}
func (u *User) Save() error {
	if err := u.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(u)

}

// validate make sure that record contains valid data
func (u *User) validate() error {
	if u.Name == "" {
		return ErrorRecordInvalid
	}
	return nil

}
