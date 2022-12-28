package host

import (
	"context"
	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListHostLogic {
	return &GetListHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListHostLogic) GetListHost(req *types.GetHostListRequest) (itemp []*types.HostResponse, err error) {
	item, err := l.svcCtx.AliOperatorRpc.HostList(l.ctx, &operator.GetListRequest{
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		l.Logger.Errorf("AliOperatorRpc HostList 分页查询错误 %s ", err)
		return nil, err
	}

	hostList := make([]*types.HostResponse, 0)

	for _, item := range item.Data {
		hostList = append(hostList, &types.HostResponse{
			InstanceId:              item.InstanceId,
			Regionid:                item.Regionid,
			InstanceName:            item.InstanceName,
			ExpiredTime:             item.ExpiredTime,
			CreationTime:            item.CreationTime,
			KeypairName:             item.KeyPairName,
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

	return hostList, nil
}
