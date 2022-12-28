package host

import (
	"context"
	"github.com/tqtcloud/manage/service/provider/aliyun/operator/rpc/types/operator"
	"strconv"

	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIdHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIdHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdHostLogic {
	return &GetIdHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIdHostLogic) GetIdHost(req *types.GetHostIdRequest) (resp *types.HostResponse, err error) {
	res, err := l.svcCtx.AliOperatorRpc.HostGetId(l.ctx, &operator.GetIdHostRequest{InstanceId: strconv.FormatInt(req.Id, 10)})
	if err != nil {
		l.Logger.Errorf("AliOperatorRpc HostGetId 查询错误 %s ", err)
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
