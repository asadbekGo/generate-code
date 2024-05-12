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

type WriteOffService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedWriteOffServiceServer
}

func NewWriteOffService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *WriteOffService {
	return &WriteOffService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *WriteOffService) CreateWriteOff(ctx context.Context, req *storehouse_service.CreateWriteOffRequest) (resp *storehouse_service.WriteOff, err error) {

	i.log.Info("---CreateWriteOff------>", logger.Any("req", req))

	pKey, err := i.strg.WriteOff().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateWriteOff->WriteOff->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.WriteOff().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyWriteOff->WriteOff->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *WriteOffService) GetByIDWriteOff(ctx context.Context, req *storehouse_service.WriteOffPrimaryKey) (resp *storehouse_service.WriteOff, err error) {

	i.log.Info("---GetByIDWriteOff------>", logger.Any("req", req))

	resp, err = i.strg.WriteOff().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDWriteOff->WriteOff->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *WriteOffService) GetListWriteOff(ctx context.Context, req *storehouse_service.GetListWriteOffRequest) (resp *storehouse_service.GetListWriteOffResponse, err error) {

	i.log.Info("---GetListWriteOff------>", logger.Any("req", req))

	resp, err = i.strg.WriteOff().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListWriteOff->WriteOff->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *WriteOffService) UpdateWriteOff(ctx context.Context, req *storehouse_service.UpdateWriteOffRequest) (resp *storehouse_service.WriteOff, err error) {

	i.log.Info("---UpdateWriteOff------>", logger.Any("req", req))

	rowsAffected, err := i.strg.WriteOff().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateWriteOff--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.WriteOff().GetByPKey(ctx, &storehouse_service.WriteOffPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateWriteOff--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *WriteOffService) UpdatePatchWriteOff(ctx context.Context, req *storehouse_service.UpdatePatchWriteOffRequest) (resp *storehouse_service.WriteOff, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.WriteOff().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.WriteOff().GetByPKey(ctx, &storehouse_service.WriteOffPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *WriteOffService) DeleteWriteOff(ctx context.Context, req *storehouse_service.WriteOffPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteWriteOff------>", logger.Any("req", req))

	err = i.strg.WriteOff().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteWriteOff->WriteOff->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
