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

### 5. "创建秘钥"

1. route definition

- Url: /secret-api/v1/create
- Method: POST
- Request: `CreateSecretRequest`
- Response: `CreateSecretResponse`

2. request definition



```golang
type CreateSecretRequest struct {
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```


3. response definition



```golang
type CreateSecretResponse struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```

### 6. "删除秘钥"

1. route definition

- Url: /secret-api/v1/delete/:id
- Method: DELETE
- Request: `DeleteSecretRequest`
- Response: `DeleteSecretResponse`

2. request definition



```golang
type DeleteSecretRequest struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type DeleteSecretResponse struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```

### 7. "更新秘钥"

1. route definition

- Url: /secret-api/v1/update
- Method: PUT
- Request: `UpdateSecretRequest`
- Response: `UpdateSecretResponse`

2. request definition



```golang
type UpdateSecretRequest struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```


3. response definition



```golang
type UpdateSecretResponse struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```

### 8. "分页查询秘钥"

1. route definition

- Url: /secret-api/v1
- Method: GET
- Request: `GetSecretListRequest`
- Response: `GetSecretListResponse`

2. request definition



```golang
type GetSecretListRequest struct {
	Page int64 `form:"page"`
	Limit int64 `form:"limit"`
}
```


3. response definition



```golang
type GetSecretListResponse struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```

### 9. "指定ID查询秘钥信息"

1. route definition

- Url: /secret-api/v1/:id
- Method: GET
- Request: `GetSecretIdRequest`
- Response: `GetSecretListResponse`

2. request definition



```golang
type GetSecretIdRequest struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type GetSecretListResponse struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```

### 10. "创建任务"

1. route definition

- Url: /task-api/v1/create
- Method: POST
- Request: `CreateTaskRequest`
- Response: `CreateTaskResponse`

2. request definition



```golang
type CreateTaskRequest struct {
	TaskName string `json:"taskname"`
	Vendor int64 `json:"vendor"`
	TaskType int64 `json:"tasktype"`
	SecretId int64 `json:"secret_id"`
	Region string `json:"region"`
	UserId int64 `json:"userid"`
}
```


3. response definition



```golang
type CreateTaskResponse struct {
	Id int64 `json:"id"`
	TaskName string `json:"taskname"`
	Vendor string `json:"vendor"`
	TaskType string `json:"tasktype"`
	SecretId int64 `json:"secret_id"`
	Region string `json:"secret_desc"`
	TaskUser string `json:"taskuser"`
	Status string `json:"status"`
	Message string `json:"message"`
	Start_At string `json:"start_at"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
```

### 11. "删除任务"

1. route definition

- Url: /task-api/v1/delete/:id
- Method: DELETE
- Request: `DeleteTaskRequest`
- Response: `DeleteTaskResponse`

2. request definition



```golang
type DeleteTaskRequest struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type DeleteTaskResponse struct {
	Id int64 `json:"id"`
	TaskName string `json:"taskname"`
	Vendor string `json:"vendor"`
	TaskType string `json:"tasktype"`
	SecretId int64 `json:"secret_id"`
	Region string `json:"secret_desc"`
	TaskUser string `json:"taskuser"`
	Status string `json:"status"`
	Message string `json:"message"`
	Start_At string `json:"start_at"`
	End_At string `json:"end_at"`
	TotalSucceed int64 `json:"total_succeed"`
	TotalFailed int64 `json:"total_failed"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
```

### 12. "分页查询任务"

1. route definition

- Url: /task-api/v1
- Method: GET
- Request: `GetTaskListRequest`
- Response: `GetTaskListResponse`

2. request definition



```golang
type GetTaskListRequest struct {
	Page int64 `form:"page"`
	Limit int64 `form:"limit"`
}
```


3. response definition



```golang
type GetTaskListResponse struct {
	Id int64 `json:"id"`
	TaskName string `json:"taskname"`
	Vendor string `json:"vendor"`
	TaskType string `json:"tasktype"`
	SecretId int64 `json:"secret_id"`
	Region string `json:"secret_desc"`
	TaskUser string `json:"taskuser"`
	Status string `json:"status"`
	Message string `json:"message"`
	Start_At string `json:"start_at"`
	End_At string `json:"end_at"`
	TotalSucceed int64 `json:"total_succeed"`
	TotalFailed int64 `json:"total_failed"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
```

### 13. "指定ID查询任务信息"

1. route definition

- Url: /task-api/v1/:id
- Method: GET
- Request: `GetTaskIdRequest`
- Response: `GetTaskListResponse`

2. request definition



```golang
type GetTaskIdRequest struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type GetTaskListResponse struct {
	Id int64 `json:"id"`
	TaskName string `json:"taskname"`
	Vendor string `json:"vendor"`
	TaskType string `json:"tasktype"`
	SecretId int64 `json:"secret_id"`
	Region string `json:"secret_desc"`
	TaskUser string `json:"taskuser"`
	Status string `json:"status"`
	Message string `json:"message"`
	Start_At string `json:"start_at"`
	End_At string `json:"end_at"`
	TotalSucceed int64 `json:"total_succeed"`
	TotalFailed int64 `json:"total_failed"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
```

