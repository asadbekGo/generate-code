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

type TransferSendService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedTransferSendServiceServer
}

func NewTransferSendService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TransferSendService {
	return &TransferSendService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TransferSendService) CreateTransferSend(ctx context.Context, req *storehouse_service.CreateTransferSendRequest) (resp *storehouse_service.TransferSend, err error) {

	i.log.Info("---CreateTransferSend------>", logger.Any("req", req))

	pKey, err := i.strg.TransferSend().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTransferSend->TransferSend->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.TransferSend().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyTransferSend->TransferSend->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferSendService) GetByIDTransferSend(ctx context.Context, req *storehouse_service.TransferSendPrimaryKey) (resp *storehouse_service.TransferSend, err error) {

	i.log.Info("---GetByIDTransferSend------>", logger.Any("req", req))

	resp, err = i.strg.TransferSend().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDTransferSend->TransferSend->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferSendService) GetListTransferSend(ctx context.Context, req *storehouse_service.GetListTransferSendRequest) (resp *storehouse_service.GetListTransferSendResponse, err error) {

	i.log.Info("---GetListTransferSend------>", logger.Any("req", req))

	resp, err = i.strg.TransferSend().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListTransferSend->TransferSend->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TransferSendService) UpdateTransferSend(ctx context.Context, req *storehouse_service.UpdateTransferSendRequest) (resp *storehouse_service.TransferSend, err error) {

	i.log.Info("---UpdateTransferSend------>", logger.Any("req", req))

	rowsAffected, err := i.strg.TransferSend().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateTransferSend--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TransferSend().GetByPKey(ctx, &storehouse_service.TransferSendPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateTransferSend--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TransferSendService) UpdatePatchTransferSend(ctx context.Context, req *storehouse_service.UpdatePatchTransferSendRequest) (resp *storehouse_service.TransferSend, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.TransferSend().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TransferSend().GetByPKey(ctx, &storehouse_service.TransferSendPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TransferSendService) DeleteTransferSend(ctx context.Context, req *storehouse_service.TransferSendPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteTransferSend------>", logger.Any("req", req))

	err = i.strg.TransferSend().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTransferSend->TransferSend->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
