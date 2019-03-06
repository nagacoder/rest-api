package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"rest-api/user"

	"github.com/asdine/storm"
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
func userGetOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	u, err := user.GetOne(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"success": true, "data": u})
}

func userPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u := new(user.User)
	err := bodyUser(r, u)

	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	err = u.Save()
	if err != nil {
		if err == user.ErrorRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"success": true, "data": u})
}
func userPacthOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u, err := user.GetOne(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}

	err = bodyUser(r, u)

	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	err = u.Save()
	if err != nil {
		if err == user.ErrorRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"success": true, "data": u})
}
func userDeleteOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	err := user.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
