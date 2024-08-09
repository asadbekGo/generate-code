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

type ClientService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_client_service.UnimplementedClientServiceServer
}

func NewClientService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ClientService {
	return &ClientService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ClientService) CreateClient(ctx context.Context, req *storehouse_client_service.CreateClientRequest) (resp *storehouse_client_service.Client, err error) {

	i.log.Info("---CreateClient------>", logger.Any("req", req))

	pKey, err := i.strg.Client().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateClient->Client->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Client().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyClient->Client->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ClientService) GetByIDClient(ctx context.Context, req *storehouse_client_service.ClientPrimaryKey) (resp *storehouse_client_service.Client, err error) {

	i.log.Info("---GetByIDClient------>", logger.Any("req", req))

	resp, err = i.strg.Client().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDClient->Client->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ClientService) GetListClient(ctx context.Context, req *storehouse_client_service.GetListClientRequest) (resp *storehouse_client_service.GetListClientResponse, err error) {

	i.log.Info("---GetListClient------>", logger.Any("req", req))

	resp, err = i.strg.Client().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListClient->Client->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ClientService) UpdateClient(ctx context.Context, req *storehouse_client_service.UpdateClientRequest) (resp *storehouse_client_service.Client, err error) {

	i.log.Info("---UpdateClient------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Client().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateClient--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Client().GetByPKey(ctx, &storehouse_client_service.ClientPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateClient--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ClientService) UpdatePatchClient(ctx context.Context, req *storehouse_client_service.UpdatePatchClientRequest) (resp *storehouse_client_service.Client, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Client().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Client().GetByPKey(ctx, &storehouse_client_service.ClientPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ClientService) DeleteClient(ctx context.Context, req *storehouse_client_service.ClientPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteClient------>", logger.Any("req", req))

	err = i.strg.Client().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteClient->Client->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
