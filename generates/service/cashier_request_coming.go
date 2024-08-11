package client_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"warehouse/warehouse_go_storehouse_service/config"
	"warehouse/warehouse_go_storehouse_service/genproto/storehouse_client_service"
	"warehouse/warehouse_go_storehouse_service/grpc/client"
	"warehouse/warehouse_go_storehouse_service/models"
	"warehouse/warehouse_go_storehouse_service/pkg/logger"
	"warehouse/warehouse_go_storehouse_service/storage"
)

type CashierRequestComingService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_client_service.UnimplementedCashierRequestComingServiceServer
}

func NewCashierRequestComingService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CashierRequestComingService {
	return &CashierRequestComingService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *CashierRequestComingService) CreateCashierRequestComing(ctx context.Context, req *storehouse_client_service.CreateCashierRequestComingRequest) (resp *storehouse_client_service.CashierRequestComing, err error) {

	i.log.Info("---CreateCashierRequestComing------>", logger.Any("req", req))

	pKey, err := i.strg.CashierRequestComing().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateCashierRequestComing->CashierRequestComing->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.CashierRequestComing().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyCashierRequestComing->CashierRequestComing->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestComingService) GetByIDCashierRequestComing(ctx context.Context, req *storehouse_client_service.CashierRequestComingPrimaryKey) (resp *storehouse_client_service.CashierRequestComing, err error) {

	i.log.Info("---GetByIDCashierRequestComing------>", logger.Any("req", req))

	resp, err = i.strg.CashierRequestComing().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDCashierRequestComing->CashierRequestComing->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestComingService) GetListCashierRequestComing(ctx context.Context, req *storehouse_client_service.GetListCashierRequestComingRequest) (resp *storehouse_client_service.GetListCashierRequestComingResponse, err error) {

	i.log.Info("---GetListCashierRequestComing------>", logger.Any("req", req))

	resp, err = i.strg.CashierRequestComing().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListCashierRequestComing->CashierRequestComing->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestComingService) UpdateCashierRequestComing(ctx context.Context, req *storehouse_client_service.UpdateCashierRequestComingRequest) (resp *storehouse_client_service.CashierRequestComing, err error) {

	i.log.Info("---UpdateCashierRequestComing------>", logger.Any("req", req))

	rowsAffected, err := i.strg.CashierRequestComing().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateCashierRequestComing--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CashierRequestComing().GetByPKey(ctx, &storehouse_client_service.CashierRequestComingPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateCashierRequestComing--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CashierRequestComingService) UpdatePatchCashierRequestComing(ctx context.Context, req *storehouse_client_service.UpdatePatchCashierRequestComingRequest) (resp *storehouse_client_service.CashierRequestComing, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.CashierRequestComing().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CashierRequestComing().GetByPKey(ctx, &storehouse_client_service.CashierRequestComingPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CashierRequestComingService) DeleteCashierRequestComing(ctx context.Context, req *storehouse_client_service.CashierRequestComingPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteCashierRequestComing------>", logger.Any("req", req))

	err = i.strg.CashierRequestComing().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteCashierRequestComing->CashierRequestComing->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
