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

type RequestService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedRequestServiceServer
}

func NewRequestService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *RequestService {
	return &RequestService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *RequestService) CreateRequest(ctx context.Context, req *storehouse_service.CreateRequestRequest) (resp *storehouse_service.Request, err error) {

	i.log.Info("---CreateRequest------>", logger.Any("req", req))

	pKey, err := i.strg.Request().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateRequest->Request->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Request().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyRequest->Request->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *RequestService) GetByIDRequest(ctx context.Context, req *storehouse_service.RequestPrimaryKey) (resp *storehouse_service.Request, err error) {

	i.log.Info("---GetByIDRequest------>", logger.Any("req", req))

	resp, err = i.strg.Request().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDRequest->Request->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *RequestService) GetListRequest(ctx context.Context, req *storehouse_service.GetListRequestRequest) (resp *storehouse_service.GetListRequestResponse, err error) {

	i.log.Info("---GetListRequest------>", logger.Any("req", req))

	resp, err = i.strg.Request().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListRequest->Request->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *RequestService) UpdateRequest(ctx context.Context, req *storehouse_service.UpdateRequestRequest) (resp *storehouse_service.Request, err error) {

	i.log.Info("---UpdateRequest------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Request().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateRequest--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Request().GetByPKey(ctx, &storehouse_service.RequestPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateRequest--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *RequestService) UpdatePatchRequest(ctx context.Context, req *storehouse_service.UpdatePatchRequestRequest) (resp *storehouse_service.Request, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Request().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Request().GetByPKey(ctx, &storehouse_service.RequestPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *RequestService) DeleteRequest(ctx context.Context, req *storehouse_service.RequestPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteRequest------>", logger.Any("req", req))

	err = i.strg.Request().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteRequest->Request->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
