package logic

import (
	"context"
	"fmt"
	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg/errors"
	"github.com/tqtcloud/manage/common/pager"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/provider/aliyun/client"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/model"
	"strings"
	"time"

	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostSyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHostSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostSyncLogic {
	return &HostSyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HostSyncLogic) HostSync(in *operator.CreateHostRequest) (*operator.GetListResponse, error) {
	startFuncTime := time.Now() // 获取当前时间
	hosts := make([]*operator.DeleteHostResponse, 0)

	aliClient := client.NewAliCloudClient(in.AccessKeyId, in.AccessKeySecret, in.Region)
	ecsClient, err := aliClient.EcsClient()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ProviderNewClientError), "阿里云 ECS client 创建错误 err:%v,user:%+v", err, in.AccessKeyId)
	}
	pag := pager.NewPager(in.Region)
	ines := make([]ecs.DescribeInstancesResponseBodyInstancesInstance, 0)
	pag.SetTotalCount(ecsClient)
	for pag.CheckHasNext() {
		l.Infof("第 %d 页数据 ,总共数据 %d 条", pag.PageNumber, pag.TotalCount)
		set, err := pag.HostsReq(ecsClient)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PagerError), "分页主机错误 err:%v,user:%+v", err, pag.Req)
		}
		var totalSucceed int64 = 0
		var totalFailed int64 = 0
		for _, v := range set {

			ines = append(ines, v)
			//l.Infof("id为：%s 实例名：%s \n", *v.InstanceId, *v.InstanceName)
			sqData, err := l.svcCtx.HostsModel.FindOne(l.ctx, *v.InstanceId)
			//if err == model.ErrNotFound {
			//	//return nil, errorx.NewDefaultError("InstanceId不存在,即将同步该条主机信息")
			//	l.Logger.Infof("InstanceId: %s 主机名：%s 不存在,即将同步该条主机信息", *v.InstanceId, *v.InstanceName)
			//}
			// 如果没有数据则创建任务
			if err == model.ErrNotFound {
				l.Logger.Infof("InstanceId: %s 主机名：%s 不存在,即将同步该条主机信息", *v.InstanceId, *v.InstanceName)

				newInstanceHost := model.Hosts{
					InstanceId:              tea.StringValue(v.InstanceId),
					Regionid:                tea.StringValue(v.RegionId),
					InstanceName:            tea.StringValue(v.InstanceName),
					ExpiredTime:             tea.StringValue(v.ExpiredTime),
					CreationTime:            tea.StringValue(v.CreationTime),
					KeyPairName:             tea.StringValue(v.KeyPairName),
					Description:             tea.StringValue(v.Description),
					OsName:                  tea.StringValue(v.OSName),
					ImageId:                 tea.StringValue(v.ImageId),
					GpuAmount:               int64(tea.Int32Value(v.GPUAmount)),
					Cpu:                     int64(tea.Int32Value(v.Cpu)),
					Memory:                  int64(tea.Int32Value(v.Memory)),
					OsType:                  tea.StringValue(v.OSType),
					InstanceChargeType:      tea.StringValue(v.InstanceChargeType),
					InternetMaxBandwidthOut: int64(tea.Int32Value(v.InternetMaxBandwidthOut)),
					InternetMaxBandwidthIn:  int64(tea.Int32Value(v.InternetMaxBandwidthIn)),
					Primaryip:               fmt.Sprintf(strings.Join(l.parsePrivateIp(&v), ",")),
					Publicip:                fmt.Sprintf(strings.Join(tea.StringSliceValue(v.PublicIpAddress.IpAddress), ",")),
					EipAddresses:            tea.StringValue(v.EipAddress.IpAddress),
					SecurityGroupId:         fmt.Sprintf(strings.Join(tea.StringSliceValue(v.SecurityGroupIds.SecurityGroupId), ",")),
				}
				l.Logger.Infof("实例id为：%s 实例名：%s, 地域: %s,", newInstanceHost.InstanceId, newInstanceHost.InstanceName, newInstanceHost.Regionid)
				_, err := l.svcCtx.HostsModel.Insert(context.Background(), &newInstanceHost)
				//time.Sleep(5 * time.Millisecond)
				if err != nil {
					l.Logger.Errorf("实例创建错误: %s", err)
					totalFailed++
					continue
				}
				// 添加到 model 切片中
				hosts = append(hosts, &operator.DeleteHostResponse{
					InstanceId:              newInstanceHost.InstanceId,
					Regionid:                newInstanceHost.Regionid,
					InstanceName:            newInstanceHost.InstanceName,
					ExpiredTime:             newInstanceHost.ExpiredTime,
					CreationTime:            newInstanceHost.CreationTime,
					KeyPairName:             newInstanceHost.KeyPairName,
					Description:             newInstanceHost.Description,
					OsName:                  newInstanceHost.OsName,
					ImageId:                 newInstanceHost.ImageId,
					GpuAmount:               newInstanceHost.GpuAmount,
					Cpu:                     newInstanceHost.Cpu,
					Memory:                  newInstanceHost.Memory,
					OsType:                  newInstanceHost.OsType,
					InstanceType:            newInstanceHost.InstanceType,
					InstanceChargeType:      newInstanceHost.InstanceChargeType,
					InternetMaxBandwidthOut: newInstanceHost.InternetMaxBandwidthOut,
					InternetMaxBandwidthIn:  newInstanceHost.InternetMaxBandwidthIn,
					Primaryip:               newInstanceHost.Primaryip,
					Publicip:                newInstanceHost.Publicip,
					EipAddresses:            newInstanceHost.EipAddresses,
					SecurityGroupId:         newInstanceHost.SecurityGroupId,
				})
				totalSucceed++
				continue
			}

			if tea.StringValue(v.InstanceId) == sqData.InstanceId {
				newInstanceHost := model.Hosts{
					InstanceId:              tea.StringValue(v.InstanceId),
					Regionid:                tea.StringValue(v.RegionId),
					InstanceName:            tea.StringValue(v.InstanceName),
					ExpiredTime:             tea.StringValue(v.ExpiredTime),
					CreationTime:            tea.StringValue(v.CreationTime),
					KeyPairName:             tea.StringValue(v.KeyPairName),
					Description:             tea.StringValue(v.Description),
					OsName:                  tea.StringValue(v.OSName),
					ImageId:                 tea.StringValue(v.ImageId),
					GpuAmount:               int64(tea.Int32Value(v.GPUAmount)),
					Cpu:                     int64(tea.Int32Value(v.Cpu)),
					Memory:                  int64(tea.Int32Value(v.Memory)),
					OsType:                  tea.StringValue(v.OSType),
					InstanceChargeType:      tea.StringValue(v.InstanceChargeType),
					InternetMaxBandwidthOut: int64(tea.Int32Value(v.InternetMaxBandwidthOut)),
					InternetMaxBandwidthIn:  int64(tea.Int32Value(v.InternetMaxBandwidthIn)),
					Primaryip:               fmt.Sprintf(strings.Join(l.parsePrivateIp(&v), ",")),
					Publicip:                fmt.Sprintf(strings.Join(tea.StringSliceValue(v.PublicIpAddress.IpAddress), ",")),
					EipAddresses:            tea.StringValue(v.EipAddress.IpAddress),
					SecurityGroupId:         fmt.Sprintf(strings.Join(tea.StringSliceValue(v.SecurityGroupIds.SecurityGroupId), ",")),
				}
				l.Logger.Infof("实例存在,执行更新 实例id为：%s 实例名：%s, 地域: %s,", newInstanceHost.InstanceId, newInstanceHost.InstanceName, newInstanceHost.Regionid)

				err := l.svcCtx.HostsModel.Update(context.Background(), &newInstanceHost)
				//time.Sleep(5 * time.Millisecond)

				if err != nil {
					l.Logger.Errorf("实例同步更新错误: %s", err)
					totalFailed++
					continue
				}
				hosts = append(hosts, &operator.DeleteHostResponse{
					InstanceId:              newInstanceHost.InstanceId,
					Regionid:                newInstanceHost.Regionid,
					InstanceName:            newInstanceHost.InstanceName,
					ExpiredTime:             newInstanceHost.ExpiredTime,
					CreationTime:            newInstanceHost.CreationTime,
					KeyPairName:             newInstanceHost.KeyPairName,
					Description:             newInstanceHost.Description,
					OsName:                  newInstanceHost.OsName,
					ImageId:                 newInstanceHost.ImageId,
					GpuAmount:               newInstanceHost.GpuAmount,
					Cpu:                     newInstanceHost.Cpu,
					Memory:                  newInstanceHost.Memory,
					OsType:                  newInstanceHost.OsType,
					InstanceType:            newInstanceHost.InstanceType,
					InstanceChargeType:      newInstanceHost.InstanceChargeType,
					InternetMaxBandwidthOut: newInstanceHost.InternetMaxBandwidthOut,
					InternetMaxBandwidthIn:  newInstanceHost.InternetMaxBandwidthIn,
					Primaryip:               newInstanceHost.Primaryip,
					Publicip:                newInstanceHost.Publicip,
					EipAddresses:            newInstanceHost.EipAddresses,
					SecurityGroupId:         newInstanceHost.SecurityGroupId,
				})
				totalSucceed++
			}
		}
		l.Infof("本次同步ECS数据共计 %d 条：", len(ines))
	}
	l.Logger.Infof("主机同步完成,同步耗时: %0.2f 秒", time.Since(startFuncTime).Seconds())

	return &operator.GetListResponse{
		Data: hosts,
	}, nil
}

func (l *HostSyncLogic) parsePrivateIp(ins *ecs.DescribeInstancesResponseBodyInstancesInstance) []string {
	ips := []string{}

	if ins.NetworkInterfaces == nil {
		return ips
	}

	// 优先通过网卡查询主私网IP地址
	for _, nc := range ins.NetworkInterfaces.NetworkInterface {
		for _, ip := range nc.PrivateIpSets.PrivateIpSet {
			if tea.BoolValue(ip.Primary) {
				ips = append(ips, tea.StringValue(ip.PrivateIpAddress))
			}
		}
	}
	if len(ips) > 0 {
		return ips
	}

	// 查询InnerIpAddress属性
	if len(ins.InnerIpAddress.IpAddress) > 0 {
		return tea.StringSliceValue(ins.InnerIpAddress.IpAddress)
	}

	// 通过专有网络VPC属性查询内网Ip
	return tea.StringSliceValue(ins.VpcAttributes.PrivateIpAddress.IpAddress)
}
