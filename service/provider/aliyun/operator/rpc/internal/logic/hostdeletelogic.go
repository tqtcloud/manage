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

type HostDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHostDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostDeleteLogic {
	return &HostDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HostDeleteLogic) HostDelete(in *operator.DeleteHostRequest) (*operator.DeleteHostResponse, error) {
	// 查询需要删除的 ID 数据
	resp, err := l.svcCtx.HostsModel.FindOne(l.ctx, in.InstanceId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.InstanceNoExistError), "operator 查询 Host ID err:%v,Task:%+v", err, in.InstanceId)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "其他错误 err:%v", err)
	}

	err = l.svcCtx.HostsModel.Delete(l.ctx, in.InstanceId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "其他错误 err:%v", err)
	}

	return &operator.DeleteHostResponse{
		InstanceId:              resp.InstanceId,
		Regionid:                resp.Regionid,
		InstanceName:            resp.InstanceName,
		ExpiredTime:             resp.ExpiredTime,
		CreationTime:            resp.CreationTime,
		KeyPairName:             resp.KeyPairName,
		Description:             resp.Description,
		OsName:                  resp.OsName,
		ImageId:                 resp.ImageId,
		GpuAmount:               resp.GpuAmount,
		Cpu:                     resp.Cpu,
		Memory:                  resp.Memory,
		OsType:                  resp.OsType,
		InstanceChargeType:      resp.InstanceChargeType,
		InternetMaxBandwidthOut: resp.InternetMaxBandwidthOut,
		InternetMaxBandwidthIn:  resp.InternetMaxBandwidthIn,
		Primaryip:               resp.Primaryip,
		Publicip:                resp.Publicip,
		EipAddresses:            resp.EipAddresses,
		SecurityGroupId:         resp.SecurityGroupId,
	}, nil

}
