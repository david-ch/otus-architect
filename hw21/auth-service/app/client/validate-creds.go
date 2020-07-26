package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/david-ch/otus-architect/auth-service/conf"
)

type validateCredsRq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type validateCredsRs struct {
	IsValid bool `json:"isValid"`
}

// ValidateCreds sends request to User Service. Returns true if creds are valid
func ValidateCreds(username string, password string) (bool, error) {
	rq := validateCredsRq{
		Username: username,
		Password: password,
	}

	jsonBytes, err := json.Marshal(rq)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", conf.UserSvcAddress()+"/validateCreds", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return false, err
	}

	req.Header.Add("X-Username", "tech")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	rs := validateCredsRs{}
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return false, err
	}

	return rs.IsValid, nil
}
