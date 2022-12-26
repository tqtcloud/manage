### 1. "登录用户"

1. route definition

- Url: /user-api/v1/login
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

### 2. "自身用户信息查询"

1. route definition

- Url: /user-api/v1/userinfo
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

### 3. "分页用户"

1. route definition

- Url: /user-api/v1
- Method: GET
- Request: `UserListRequest`
- Response: `UserListResponse`

2. request definition



```golang
type UserListRequest struct {
	Page int64 `form:"page"`
	Limit int64 `form:"limit"`
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

### 4. "创建用户"

1. route definition

- Url: /user-api/v1/register
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

