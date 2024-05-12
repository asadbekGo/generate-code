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

type TransferReceiveProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedTransferReceiveProductServiceServer
}

func NewTransferReceiveProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TransferReceiveProductService {
	return &TransferReceiveProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TransferReceiveProductService) CreateTransferReceiveProduct(ctx context.Context, req *storehouse_service.CreateTransferReceiveProductRequest) (resp *storehouse_service.TransferReceiveProduct, err error) {

	i.log.Info("---CreateTransferReceiveProduct------>", logger.Any("req", req))

	pKey, err := i.strg.TransferReceiveProduct().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTransferReceiveProduct->TransferReceiveProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.TransferReceiveProduct().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyTransferReceiveProduct->TransferReceiveProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferReceiveProductService) GetByIDTransferReceiveProduct(ctx context.Context, req *storehouse_service.TransferReceiveProductPrimaryKey) (resp *storehouse_service.TransferReceiveProduct, err error) {

	i.log.Info("---GetByIDTransferReceiveProduct------>", logger.Any("req", req))

	resp, err = i.strg.TransferReceiveProduct().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDTransferReceiveProduct->TransferReceiveProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferReceiveProductService) GetListTransferReceiveProduct(ctx context.Context, req *storehouse_service.GetListTransferReceiveProductRequest) (resp *storehouse_service.GetListTransferReceiveProductResponse, err error) {

	i.log.Info("---GetListTransferReceiveProduct------>", logger.Any("req", req))

	resp, err = i.strg.TransferReceiveProduct().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListTransferReceiveProduct->TransferReceiveProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferReceiveProductService) UpdateTransferReceiveProduct(ctx context.Context, req *storehouse_service.UpdateTransferReceiveProductRequest) (resp *storehouse_service.TransferReceiveProduct, err error) {

	i.log.Info("---UpdateTransferReceiveProduct------>", logger.Any("req", req))

	rowsAffected, err := i.strg.TransferReceiveProduct().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateTransferReceiveProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TransferReceiveProduct().GetByPKey(ctx, &storehouse_service.TransferReceiveProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateTransferReceiveProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TransferReceiveProductService) UpdatePatchTransferReceiveProduct(ctx context.Context, req *storehouse_service.UpdatePatchTransferReceiveProductRequest) (resp *storehouse_service.TransferReceiveProduct, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.TransferReceiveProduct().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TransferReceiveProduct().GetByPKey(ctx, &storehouse_service.TransferReceiveProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TransferReceiveProductService) DeleteTransferReceiveProduct(ctx context.Context, req *storehouse_service.TransferReceiveProductPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteTransferReceiveProduct------>", logger.Any("req", req))

	err = i.strg.TransferReceiveProduct().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTransferReceiveProduct->TransferReceiveProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
