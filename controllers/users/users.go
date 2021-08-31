package users

import (
	"encoding/json"
	"net/http"
)

type userStruct struct {
	username string
	fullname string
}

func Index(w http.ResponseWriter, r *http.Request) {

	user := userStruct{"tiago.murilo", "Tiago Murilo Ochoa da Luz"}

	usersJson, err := json.Marshal(user.fullname)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(usersJson)
}
