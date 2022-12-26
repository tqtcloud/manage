### 1. "创建秘钥"

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

### 2. "删除秘钥"

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

### 3. "更新秘钥"

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

### 4. "分页查询秘钥"

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

### 5. "指定ID查询秘钥信息"

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

