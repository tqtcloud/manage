### 1. "创建秘钥"

1. route definition

- Url: /api/secret/v1/create
- Method: POST
- Request: `CreateRequest`
- Response: `CreateResponse`

2. request definition



```golang
type CreateRequest struct {
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```


3. response definition



```golang
type CreateResponse struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```

### 2. "删除秘钥"

1. route definition

- Url: /api/secret/v1/delete/:id
- Method: DELETE
- Request: `DeleteRequest`
- Response: `DeleteResponse`

2. request definition



```golang
type DeleteRequest struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type DeleteResponse struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```

### 3. "更新秘钥"

1. route definition

- Url: /api/secret/v1/update
- Method: PUT
- Request: `UpdateRequest`
- Response: `UpdateResponse`

2. request definition



```golang
type UpdateRequest struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```


3. response definition



```golang
type UpdateResponse struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```

### 4. "分页查询秘钥"

1. route definition

- Url: /api/secret/v1/getlist
- Method: POST
- Request: `GetListRequest`
- Response: `GetListResponse`

2. request definition



```golang
type GetListRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}
```


3. response definition



```golang
type GetListResponse struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```

### 5. "指定ID查询秘钥信息"

1. route definition

- Url: /api/secret/v1/:id
- Method: GET
- Request: `GetIdRequest`
- Response: `GetListResponse`

2. request definition



```golang
type GetIdRequest struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type GetListResponse struct {
	Id int64 `json:"id"`
	Vendor string `json:"vendor"`
	AccessKeyId string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
}
```

