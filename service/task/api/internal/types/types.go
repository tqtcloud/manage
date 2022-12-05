// Code generated by goctl. DO NOT EDIT.
package types

type CreateRequest struct {
	TaskName string `json:"taskname"`
	Vendor   string `json:"vendor"`
	TaskType string `json:"tasktype"`
	SecretId int64  `json:"secret_id"`
	Region   string `json:"vendor"`
	UserId   int64  `json:"userid"`
}

type CreateResponse struct {
	Id         int64  `json:"id"`
	TaskName   string `json:"taskname"`
	Vendor     string `json:"vendor"`
	TaskType   string `json:"tasktype"`
	SecretId   int64  `json:"secret_id"`
	Region     string `json:"secret_desc"`
	TaskUser   string `json:"taskuser"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Start_At   string `json:"start_at"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type DeleteRequest struct {
	Id int64 `path:"id"`
}

type DeleteResponse struct {
	Id           int64  `json:"id"`
	TaskName     string `json:"taskname"`
	Vendor       string `json:"vendor"`
	TaskType     string `json:"tasktype"`
	SecretId     int64  `json:"secret_id"`
	Region       string `json:"secret_desc"`
	TaskUser     string `json:"taskuser"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	Start_At     string `json:"start_at"`
	End_At       string `json:"end_at"`
	TotalSucceed int64  `json:"total_succeed"`
	TotalFailed  int64  `json:"total_failed"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

type GetListRequest struct {
	Page  int64 `path:"page"`
	Limit int64 `path:"limit"`
}

type GetListResponse struct {
	Id           int64  `json:"id"`
	TaskName     string `json:"taskname"`
	Vendor       string `json:"vendor"`
	TaskType     string `json:"tasktype"`
	SecretId     int64  `json:"secret_id"`
	Region       string `json:"secret_desc"`
	TaskUser     string `json:"taskuser"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	Start_At     string `json:"start_at"`
	End_At       string `json:"end_at"`
	TotalSucceed int64  `json:"total_succeed"`
	TotalFailed  int64  `json:"total_failed"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

type GetIdRequest struct {
	Id int64 `path:"id"`
}