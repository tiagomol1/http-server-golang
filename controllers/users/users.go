package users

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	username string
	isAdmin  bool
	fullname string
}

func Index(w http.ResponseWriter, r *http.Request) {
	user := User{
		fullname: "Tiago Murilo Ochôa da Luz",
		isAdmin:  true,
		username: "tiago.murilo",
	}

	fmt.Println(user)
	userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(userJson)
}
