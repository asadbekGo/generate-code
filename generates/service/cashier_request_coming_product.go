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

type CashierRequestComingProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_client_service.UnimplementedCashierRequestComingProductServiceServer
}

func NewCashierRequestComingProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CashierRequestComingProductService {
	return &CashierRequestComingProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *CashierRequestComingProductService) CreateCashierRequestComingProduct(ctx context.Context, req *storehouse_client_service.CreateCashierRequestComingProductRequest) (resp *storehouse_client_service.CashierRequestComingProduct, err error) {

	i.log.Info("---CreateCashierRequestComingProduct------>", logger.Any("req", req))

	pKey, err := i.strg.CashierRequestComingProduct().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateCashierRequestComingProduct->CashierRequestComingProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.CashierRequestComingProduct().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyCashierRequestComingProduct->CashierRequestComingProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestComingProductService) GetByIDCashierRequestComingProduct(ctx context.Context, req *storehouse_client_service.CashierRequestComingProductPrimaryKey) (resp *storehouse_client_service.CashierRequestComingProduct, err error) {

	i.log.Info("---GetByIDCashierRequestComingProduct------>", logger.Any("req", req))

	resp, err = i.strg.CashierRequestComingProduct().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDCashierRequestComingProduct->CashierRequestComingProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestComingProductService) GetListCashierRequestComingProduct(ctx context.Context, req *storehouse_client_service.GetListCashierRequestComingProductRequest) (resp *storehouse_client_service.GetListCashierRequestComingProductResponse, err error) {

	i.log.Info("---GetListCashierRequestComingProduct------>", logger.Any("req", req))

	resp, err = i.strg.CashierRequestComingProduct().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListCashierRequestComingProduct->CashierRequestComingProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestComingProductService) UpdateCashierRequestComingProduct(ctx context.Context, req *storehouse_client_service.UpdateCashierRequestComingProductRequest) (resp *storehouse_client_service.CashierRequestComingProduct, err error) {

	i.log.Info("---UpdateCashierRequestComingProduct------>", logger.Any("req", req))

	rowsAffected, err := i.strg.CashierRequestComingProduct().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateCashierRequestComingProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CashierRequestComingProduct().GetByPKey(ctx, &storehouse_client_service.CashierRequestComingProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateCashierRequestComingProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CashierRequestComingProductService) UpdatePatchCashierRequestComingProduct(ctx context.Context, req *storehouse_client_service.UpdatePatchCashierRequestComingProductRequest) (resp *storehouse_client_service.CashierRequestComingProduct, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.CashierRequestComingProduct().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CashierRequestComingProduct().GetByPKey(ctx, &storehouse_client_service.CashierRequestComingProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CashierRequestComingProductService) DeleteCashierRequestComingProduct(ctx context.Context, req *storehouse_client_service.CashierRequestComingProductPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteCashierRequestComingProduct------>", logger.Any("req", req))

	err = i.strg.CashierRequestComingProduct().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteCashierRequestComingProduct->CashierRequestComingProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
