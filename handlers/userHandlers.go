package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-api/user"

	"gopkg.in/mgo.v2/bson"
)

func bodyUser(r *http.Request, u *user.User) error {

	if r.Body == nil {
		return errors.New("Request body is empty")
	}
	if u == nil {
		return errors.New("a user is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, u)
}
func userGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := user.GetAll()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"data": users, "success": true})
}
func userPostOne(w http.ResponseWriter, r *http.Request) {

	u := new(user.User)
	err := bodyUser(r, u)
	fmt.Println(err)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}

	u.ID = bson.NewObjectId()
	err = u.Save()
	if err != nil {
		if err == user.ErrorRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/users/"+u.ID.Hex())
	w.WriteHeader(http.StatusCreated)

}
