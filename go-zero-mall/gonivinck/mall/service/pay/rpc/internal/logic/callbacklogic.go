package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/types/order"
	"mall/service/user/model"
	"mall/service/user/rpc/types/user"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	// todo: add your logic here and delete this line
	//查找用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
	if err != nil {
		return nil, err
	}

	//查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{Id: in.Oid})
	if err != nil {
		return nil, err
	}

	//查询支付是否存在
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Oid)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "支付不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	//支付金额与订单金额不同
	if in.Amount != res.Amount {
		return nil, status.Error(100, "支付金额与订单金额不符")
	}
	res.Source = in.Source
	res.Status = in.Status

	err = l.svcCtx.PayModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	//更新订单支付状态
	_, err = l.svcCtx.OrderRpc.Paid(l.ctx, &order.PaidRequest{Id: in.Oid})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &pay.CallbackResponse{}, nil
}
