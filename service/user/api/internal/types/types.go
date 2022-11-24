// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Name         string `json:"name"`
	Gender       int64  `json:"gender"`
	Mobile       string `json:"mobile"`
	Email        string `json:"email"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Gender   int64  `json:"gender"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}

type UserInfoResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}

type UserListRequest struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type UserListResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}
