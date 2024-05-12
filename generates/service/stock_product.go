package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"warehouse/warehouse_go_storehouse_service/config"
	"warehouse/warehouse_go_storehouse_service/genproto/storehouse_service"
	"warehouse/warehouse_go_storehouse_service/grpc/client"
	"warehouse/warehouse_go_storehouse_service/models"
	"warehouse/warehouse_go_storehouse_service/pkg/logger"
	"warehouse/warehouse_go_storehouse_service/storage"
)

type StockProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedStockProductServiceServer
}

func NewStockProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *StockProductService {
	return &StockProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *StockProductService) CreateStockProduct(ctx context.Context, req *storehouse_service.CreateStockProductRequest) (resp *storehouse_service.StockProduct, err error) {

	i.log.Info("---CreateStockProduct------>", logger.Any("req", req))

	pKey, err := i.strg.StockProduct().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateStockProduct->StockProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.StockProduct().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyStockProduct->StockProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *StockProductService) GetByIDStockProduct(ctx context.Context, req *storehouse_service.StockProductPrimaryKey) (resp *storehouse_service.StockProduct, err error) {

	i.log.Info("---GetByIDStockProduct------>", logger.Any("req", req))

	resp, err = i.strg.StockProduct().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDStockProduct->StockProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *StockProductService) GetListStockProduct(ctx context.Context, req *storehouse_service.GetListStockProductRequest) (resp *storehouse_service.GetListStockProductResponse, err error) {

	i.log.Info("---GetListStockProduct------>", logger.Any("req", req))

	resp, err = i.strg.StockProduct().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListStockProduct->StockProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *StockProductService) UpdateStockProduct(ctx context.Context, req *storehouse_service.UpdateStockProductRequest) (resp *storehouse_service.StockProduct, err error) {

	i.log.Info("---UpdateStockProduct------>", logger.Any("req", req))

	rowsAffected, err := i.strg.StockProduct().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateStockProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.StockProduct().GetByPKey(ctx, &storehouse_service.StockProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateStockProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *StockProductService) UpdatePatchStockProduct(ctx context.Context, req *storehouse_service.UpdatePatchStockProductRequest) (resp *storehouse_service.StockProduct, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.StockProduct().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.StockProduct().GetByPKey(ctx, &storehouse_service.StockProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *StockProductService) DeleteStockProduct(ctx context.Context, req *storehouse_service.StockProductPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteStockProduct------>", logger.Any("req", req))

	err = i.strg.StockProduct().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteStockProduct->StockProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
