### 1. "创建任务"

1. route definition

- Url: /api/task/v1/create
- Method: POST
- Request: `CreateRequest`
- Response: `CreateResponse`

2. request definition



```golang
type CreateRequest struct {
	TaskName string `json:"taskname"`
	Vendor string `json:"vendor"`
	TaskType string `json:"tasktype"`
	SecretId int64 `json:"secret_id"`
	Region string `json:"vendor"`
	UserId int64 `json:"userid"`
}
```


3. response definition



```golang
type CreateResponse struct {
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

### 2. "删除任务"

1. route definition

- Url: /api/task/v1/delete/:id
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

### 3. "分页查询任务"

1. route definition

- Url: /api/task/v1/:page/:limit
- Method: GET
- Request: `GetListRequest`
- Response: `GetListResponse`

2. request definition



```golang
type GetListRequest struct {
	Page int64 `path:"page"`
	Limit int64 `path:"limit"`
}
```


3. response definition



```golang
type GetListResponse struct {
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

### 4. "指定ID查询任务信息"

1. route definition

- Url: /api/task/v1/:id
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

