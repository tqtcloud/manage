### 1. N/A

1. route definition

- Url: /api/user/login
- Method: POST
- Request: `LoginRequest`
- Response: `LoginResponse`

2. request definition



```golang
type LoginRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginResponse struct {
	Name string `json:"name"`
	Gender int64 `json:"gender"`
	Mobile string `json:"mobile"`
	Email string `json:"email"`
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
}
```

### 2. N/A

1. route definition

- Url: /api/user/register
- Method: POST
- Request: `RegisterRequest`
- Response: `RegisterResponse`

2. request definition



```golang
type RegisterRequest struct {
	Name string `json:"name"`
	Gender int64 `json:"gender"`
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	Email string `json:"email"`
}
```


3. response definition



```golang
type RegisterResponse struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Gender int64 `json:"gender"`
	Mobile string `json:"mobile"`
	Email string `json:"email"`
}
```

### 3. N/A

1. route definition

- Url: /api/user/userinfo
- Method: POST
- Request: `-`
- Response: `UserInfoResponse`

2. request definition



3. response definition



```golang
type UserInfoResponse struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Gender int64 `json:"gender"`
	Mobile string `json:"mobile"`
	Email string `json:"email"`
}
```

### 4. N/A

1. route definition

- Url: /api/user/userlist
- Method: POST
- Request: `UserListRequest`
- Response: `UserListResponse`

2. request definition



```golang
type UserListRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}
```


3. response definition



```golang
type UserListResponse struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Gender int64 `json:"gender"`
	Mobile string `json:"mobile"`
	Email string `json:"email"`
}
```

