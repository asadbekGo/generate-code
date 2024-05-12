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

type TransferSendProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedTransferSendProductServiceServer
}

func NewTransferSendProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TransferSendProductService {
	return &TransferSendProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TransferSendProductService) CreateTransferSendProduct(ctx context.Context, req *storehouse_service.CreateTransferSendProductRequest) (resp *storehouse_service.TransferSendProduct, err error) {

	i.log.Info("---CreateTransferSendProduct------>", logger.Any("req", req))

	pKey, err := i.strg.TransferSendProduct().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTransferSendProduct->TransferSendProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.TransferSendProduct().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyTransferSendProduct->TransferSendProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferSendProductService) GetByIDTransferSendProduct(ctx context.Context, req *storehouse_service.TransferSendProductPrimaryKey) (resp *storehouse_service.TransferSendProduct, err error) {

	i.log.Info("---GetByIDTransferSendProduct------>", logger.Any("req", req))

	resp, err = i.strg.TransferSendProduct().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDTransferSendProduct->TransferSendProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferSendProductService) GetListTransferSendProduct(ctx context.Context, req *storehouse_service.GetListTransferSendProductRequest) (resp *storehouse_service.GetListTransferSendProductResponse, err error) {

	i.log.Info("---GetListTransferSendProduct------>", logger.Any("req", req))

	resp, err = i.strg.TransferSendProduct().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListTransferSendProduct->TransferSendProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferSendProductService) UpdateTransferSendProduct(ctx context.Context, req *storehouse_service.UpdateTransferSendProductRequest) (resp *storehouse_service.TransferSendProduct, err error) {

	i.log.Info("---UpdateTransferSendProduct------>", logger.Any("req", req))

	rowsAffected, err := i.strg.TransferSendProduct().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateTransferSendProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TransferSendProduct().GetByPKey(ctx, &storehouse_service.TransferSendProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateTransferSendProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TransferSendProductService) UpdatePatchTransferSendProduct(ctx context.Context, req *storehouse_service.UpdatePatchTransferSendProductRequest) (resp *storehouse_service.TransferSendProduct, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.TransferSendProduct().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TransferSendProduct().GetByPKey(ctx, &storehouse_service.TransferSendProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TransferSendProductService) DeleteTransferSendProduct(ctx context.Context, req *storehouse_service.TransferSendProductPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteTransferSendProduct------>", logger.Any("req", req))

	err = i.strg.TransferSendProduct().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTransferSendProduct->TransferSendProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
