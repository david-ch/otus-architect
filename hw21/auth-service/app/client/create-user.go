package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/david-ch/otus-architect/auth-service/conf"
)

type createUserRq struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

// CreateUser sends request to User Service
func CreateUser(username string, firstName string, lastName string, password string) error {
	rq := createUserRq{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	}

	jsonBytes, err := json.Marshal(rq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", conf.UserSvcAddress()+"/profile", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	req.Header.Add("X-Username", "tech")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
