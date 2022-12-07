package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)
func NewAliCloudClient(ak, sk, region string) *AliCloudClient {
	return &AliCloudClient{
		Region:       region,
		AccessKey:    ak,
		AccessSecret: sk,
	}
}

type AliCloudClient struct {
	AccessKey    string
	AccessSecret string
	Region       string

	ecsConn   *ecs.Client
}

//EcsClient 创建 EcsClient 客户端
func (c *AliCloudClient) EcsClient() (*ecs.Client, error) {
	if c.ecsConn != nil {
		return c.ecsConn, nil
	}

	client, err := ecs.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("ecs.aliyuncs.com"),
		RegionId:        tea.String(c.Region),
	})
	if err != nil {
		return nil, err
	}

	c.ecsConn = client
	return client, nil
}