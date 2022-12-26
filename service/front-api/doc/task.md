### 1. "创建任务"

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

### 2. "删除任务"

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

### 3. "分页查询任务"

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

### 4. "指定ID查询任务信息"

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

