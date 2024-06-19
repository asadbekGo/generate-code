package supplier_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"warehouse/warehouse_go_storehouse_service/config"
	"warehouse/warehouse_go_storehouse_service/genproto/storehouse_supplier_service"
	"warehouse/warehouse_go_storehouse_service/grpc/client"
	"warehouse/warehouse_go_storehouse_service/models"
	"warehouse/warehouse_go_storehouse_service/pkg/logger"
	"warehouse/warehouse_go_storehouse_service/storage"
)

type CashierRequestService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_supplier_service.UnimplementedCashierRequestServiceServer
}

func NewCashierRequestService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CashierRequestService {
	return &CashierRequestService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *CashierRequestService) CreateCashierRequest(ctx context.Context, req *storehouse_supplier_service.CreateCashierRequestRequest) (resp *storehouse_supplier_service.CashierRequest, err error) {

	i.log.Info("---CreateCashierRequest------>", logger.Any("req", req))

	pKey, err := i.strg.CashierRequest().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateCashierRequest->CashierRequest->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.CashierRequest().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyCashierRequest->CashierRequest->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestService) GetByIDCashierRequest(ctx context.Context, req *storehouse_supplier_service.CashierRequestPrimaryKey) (resp *storehouse_supplier_service.CashierRequest, err error) {

	i.log.Info("---GetByIDCashierRequest------>", logger.Any("req", req))

	resp, err = i.strg.CashierRequest().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDCashierRequest->CashierRequest->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestService) GetListCashierRequest(ctx context.Context, req *storehouse_supplier_service.GetListCashierRequestRequest) (resp *storehouse_supplier_service.GetListCashierRequestResponse, err error) {

	i.log.Info("---GetListCashierRequest------>", logger.Any("req", req))

	resp, err = i.strg.CashierRequest().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListCashierRequest->CashierRequest->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestService) UpdateCashierRequest(ctx context.Context, req *storehouse_supplier_service.UpdateCashierRequestRequest) (resp *storehouse_supplier_service.CashierRequest, err error) {

	i.log.Info("---UpdateCashierRequest------>", logger.Any("req", req))

	rowsAffected, err := i.strg.CashierRequest().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateCashierRequest--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CashierRequest().GetByPKey(ctx, &storehouse_supplier_service.CashierRequestPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateCashierRequest--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CashierRequestService) UpdatePatchCashierRequest(ctx context.Context, req *storehouse_supplier_service.UpdatePatchCashierRequestRequest) (resp *storehouse_supplier_service.CashierRequest, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.CashierRequest().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CashierRequest().GetByPKey(ctx, &storehouse_supplier_service.CashierRequestPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CashierRequestService) DeleteCashierRequest(ctx context.Context, req *storehouse_supplier_service.CashierRequestPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteCashierRequest------>", logger.Any("req", req))

	err = i.strg.CashierRequest().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteCashierRequest->CashierRequest->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
