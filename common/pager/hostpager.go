package pager

import (
	"fmt"
	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"log"
)

type Pager struct {
	Req        ecs.DescribeInstancesRequest
	TotalCount int32
	PageSize   int32
	PageNumber int32
	RegionId   string
}

func NewPager(regionId string) *Pager {
	return &Pager{
		Req:        ecs.DescribeInstancesRequest{RegionId: tea.String(regionId)},
		TotalCount: 10,
		PageSize:   50,
		PageNumber: 0,
		RegionId:   regionId,
	}
}

func (ps *Pager) SetTotalCount(client *ecs.Client) {
	req := &ecs.DescribeInstancesRequest{
		PageSize:   tea.Int32(ps.PageSize),
		PageNumber: tea.Int32(1),
		RegionId:   tea.String(ps.RegionId),
	}
	resp, err := client.DescribeInstances(req)
	if err != nil {
		log.Println(err)
	}
	ps.TotalCount = *resp.Body.TotalCount
}

func (ps *Pager) CheckHasNext() bool {
	// 小于总量  页数加一
	fmt.Println(ps.PageNumber, ps.PageSize, ps.TotalCount)
	if ps.PageNumber*ps.PageSize <= ps.TotalCount {
		ps.PageNumber++
		return true
	} else {
		// 否则读取完毕，没有下一页了
		return false
	}
}

// HostsReq 分页请求主机数据
func (ps *Pager) HostsReq(client *ecs.Client) ([]ecs.DescribeInstancesResponseBodyInstancesInstance, error) {
	ps.Req.PageNumber = tea.Int32(ps.PageNumber)
	ps.Req.PageSize = tea.Int32(ps.PageSize)
	resp, err := client.DescribeInstances(&ps.Req)
	if err != nil {
		return nil, err
	}
	set := make([]ecs.DescribeInstancesResponseBodyInstancesInstance, 0)
	for _, v := range resp.Body.Instances.Instance {
		set = append(set, *v)
	}
	ps.TotalCount = *resp.Body.TotalCount
	return set, nil
}
