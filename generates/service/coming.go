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

type ComingService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_client_service.UnimplementedComingServiceServer
}

func NewComingService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ComingService {
	return &ComingService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ComingService) CreateComing(ctx context.Context, req *storehouse_client_service.CreateComingRequest) (resp *storehouse_client_service.Coming, err error) {

	i.log.Info("---CreateComing------>", logger.Any("req", req))

	pKey, err := i.strg.Coming().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateComing->Coming->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Coming().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyComing->Coming->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ComingService) GetByIDComing(ctx context.Context, req *storehouse_client_service.ComingPrimaryKey) (resp *storehouse_client_service.Coming, err error) {

	i.log.Info("---GetByIDComing------>", logger.Any("req", req))

	resp, err = i.strg.Coming().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDComing->Coming->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ComingService) GetListComing(ctx context.Context, req *storehouse_client_service.GetListComingRequest) (resp *storehouse_client_service.GetListComingResponse, err error) {

	i.log.Info("---GetListComing------>", logger.Any("req", req))

	resp, err = i.strg.Coming().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListComing->Coming->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ComingService) UpdateComing(ctx context.Context, req *storehouse_client_service.UpdateComingRequest) (resp *storehouse_client_service.Coming, err error) {

	i.log.Info("---UpdateComing------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Coming().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateComing--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Coming().GetByPKey(ctx, &storehouse_client_service.ComingPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateComing--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ComingService) UpdatePatchComing(ctx context.Context, req *storehouse_client_service.UpdatePatchComingRequest) (resp *storehouse_client_service.Coming, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Coming().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Coming().GetByPKey(ctx, &storehouse_client_service.ComingPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ComingService) DeleteComing(ctx context.Context, req *storehouse_client_service.ComingPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteComing------>", logger.Any("req", req))

	err = i.strg.Coming().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteComing->Coming->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
