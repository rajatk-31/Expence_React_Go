package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	. "sample/models"

	"gopkg.in/mgo.v2/bson"
)

func RegisterUser(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var user User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		RespondWithError(res, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	// if err := dao.Insert(user); err != nil {
	// 	respondWithError(res, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	fmt.Println(user)
	RespondWithJson(res, http.StatusCreated, user)

}
