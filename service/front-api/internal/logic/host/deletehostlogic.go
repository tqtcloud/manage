package host

import (
	"context"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"
	"strconv"

	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHostLogic {
	return &DeleteHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteHostLogic) DeleteHost(req *types.DeleteHostRequest) (resp *types.HostResponse, err error) {
	res, err := l.svcCtx.AliOperatorRpc.HostDelete(l.ctx, &operator.DeleteHostRequest{InstanceId: strconv.FormatInt(req.Id, 10)})
	if err != nil {
		l.Logger.Errorf("AliOperatorRpc HostDelete 删除错误 %s ", err)
		return nil, err
	}

	return &types.HostResponse{
		InstanceId:              res.InstanceId,
		Regionid:                res.Regionid,
		InstanceName:            res.InstanceName,
		ExpiredTime:             res.ExpiredTime,
		CreationTime:            res.CreationTime,
		KeypairName:             res.KeyPairName,
		Description:             res.Description,
		OsName:                  res.OsName,
		ImageId:                 res.ImageId,
		GpuAmount:               res.GpuAmount,
		Cpu:                     res.Cpu,
		Memory:                  res.Memory,
		OsType:                  res.OsType,
		InstanceChargeType:      res.InstanceChargeType,
		InternetMaxBandwidthOut: res.InternetMaxBandwidthOut,
		InternetMaxBandwidthIn:  res.InternetMaxBandwidthIn,
		Primaryip:               res.Primaryip,
		Publicip:                res.Publicip,
		EipAddresses:            res.EipAddresses,
		SecurityGroupId:         res.SecurityGroupId,
	}, nil
}
