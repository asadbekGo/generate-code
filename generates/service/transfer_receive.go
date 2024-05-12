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

type TransferReceiveService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedTransferReceiveServiceServer
}

func NewTransferReceiveService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TransferReceiveService {
	return &TransferReceiveService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TransferReceiveService) CreateTransferReceive(ctx context.Context, req *storehouse_service.CreateTransferReceiveRequest) (resp *storehouse_service.TransferReceive, err error) {

	i.log.Info("---CreateTransferReceive------>", logger.Any("req", req))

	pKey, err := i.strg.TransferReceive().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTransferReceive->TransferReceive->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.TransferReceive().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyTransferReceive->TransferReceive->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferReceiveService) GetByIDTransferReceive(ctx context.Context, req *storehouse_service.TransferReceivePrimaryKey) (resp *storehouse_service.TransferReceive, err error) {

	i.log.Info("---GetByIDTransferReceive------>", logger.Any("req", req))

	resp, err = i.strg.TransferReceive().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDTransferReceive->TransferReceive->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferReceiveService) GetListTransferReceive(ctx context.Context, req *storehouse_service.GetListTransferReceiveRequest) (resp *storehouse_service.GetListTransferReceiveResponse, err error) {

	i.log.Info("---GetListTransferReceive------>", logger.Any("req", req))

	resp, err = i.strg.TransferReceive().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListTransferReceive->TransferReceive->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferReceiveService) UpdateTransferReceive(ctx context.Context, req *storehouse_service.UpdateTransferReceiveRequest) (resp *storehouse_service.TransferReceive, err error) {

	i.log.Info("---UpdateTransferReceive------>", logger.Any("req", req))

	rowsAffected, err := i.strg.TransferReceive().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateTransferReceive--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TransferReceive().GetByPKey(ctx, &storehouse_service.TransferReceivePrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateTransferReceive--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TransferReceiveService) UpdatePatchTransferReceive(ctx context.Context, req *storehouse_service.UpdatePatchTransferReceiveRequest) (resp *storehouse_service.TransferReceive, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.TransferReceive().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TransferReceive().GetByPKey(ctx, &storehouse_service.TransferReceivePrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TransferReceiveService) DeleteTransferReceive(ctx context.Context, req *storehouse_service.TransferReceivePrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteTransferReceive------>", logger.Any("req", req))

	err = i.strg.TransferReceive().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTransferReceive->TransferReceive->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
