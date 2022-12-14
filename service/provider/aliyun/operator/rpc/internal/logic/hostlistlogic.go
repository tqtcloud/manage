package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tqtcloud/manage/common/xerr"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/model"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/internal/svc"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostListLogic {
	return &HostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HostListLogic) HostList(in *operator.GetListRequest) (*operator.GetListResponse, error) {
	list, err := l.svcCtx.HostsModel.FindAllPage(l.ctx, in)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PagerError), "Host  FindAllPage err:%v", err)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "其他错误 err:%v", err)
	}

	hostList := make([]*operator.DeleteHostResponse, 0)
	for _, item := range list {
		hostList = append(hostList, &operator.DeleteHostResponse{
			InstanceId:              item.InstanceId,
			Regionid:                item.Regionid,
			InstanceName:            item.InstanceName,
			ExpiredTime:             item.ExpiredTime,
			CreationTime:            item.CreationTime,
			KeyPairName:             item.KeyPairName,
			Description:             item.Description,
			OsName:                  item.OsName,
			ImageId:                 item.ImageId,
			GpuAmount:               item.GpuAmount,
			Cpu:                     item.Cpu,
			Memory:                  item.Memory,
			OsType:                  item.OsType,
			InstanceChargeType:      item.InstanceChargeType,
			InternetMaxBandwidthOut: item.InternetMaxBandwidthOut,
			InternetMaxBandwidthIn:  item.InternetMaxBandwidthIn,
			Primaryip:               item.Primaryip,
			Publicip:                item.Publicip,
			EipAddresses:            item.EipAddresses,
			SecurityGroupId:         item.SecurityGroupId,
		})
	}

	return &operator.GetListResponse{
		Data: hostList,
	}, nil
}
