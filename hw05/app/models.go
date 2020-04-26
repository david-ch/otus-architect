package main

type user struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type errorResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func fromError(e error) errorResp {
	return errorResp{Message: e.Error()}
}
