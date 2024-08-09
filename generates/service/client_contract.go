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

type ClientContractService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_client_service.UnimplementedClientContractServiceServer
}

func NewClientContractService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ClientContractService {
	return &ClientContractService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ClientContractService) CreateClientContract(ctx context.Context, req *storehouse_client_service.CreateClientContractRequest) (resp *storehouse_client_service.ClientContract, err error) {

	i.log.Info("---CreateClientContract------>", logger.Any("req", req))

	pKey, err := i.strg.ClientContract().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateClientContract->ClientContract->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.ClientContract().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyClientContract->ClientContract->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ClientContractService) GetByIDClientContract(ctx context.Context, req *storehouse_client_service.ClientContractPrimaryKey) (resp *storehouse_client_service.ClientContract, err error) {

	i.log.Info("---GetByIDClientContract------>", logger.Any("req", req))

	resp, err = i.strg.ClientContract().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDClientContract->ClientContract->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ClientContractService) GetListClientContract(ctx context.Context, req *storehouse_client_service.GetListClientContractRequest) (resp *storehouse_client_service.GetListClientContractResponse, err error) {

	i.log.Info("---GetListClientContract------>", logger.Any("req", req))

	resp, err = i.strg.ClientContract().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListClientContract->ClientContract->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ClientContractService) UpdateClientContract(ctx context.Context, req *storehouse_client_service.UpdateClientContractRequest) (resp *storehouse_client_service.ClientContract, err error) {

	i.log.Info("---UpdateClientContract------>", logger.Any("req", req))

	rowsAffected, err := i.strg.ClientContract().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateClientContract--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.ClientContract().GetByPKey(ctx, &storehouse_client_service.ClientContractPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateClientContract--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ClientContractService) UpdatePatchClientContract(ctx context.Context, req *storehouse_client_service.UpdatePatchClientContractRequest) (resp *storehouse_client_service.ClientContract, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.ClientContract().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.ClientContract().GetByPKey(ctx, &storehouse_client_service.ClientContractPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ClientContractService) DeleteClientContract(ctx context.Context, req *storehouse_client_service.ClientContractPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteClientContract------>", logger.Any("req", req))

	err = i.strg.ClientContract().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteClientContract->ClientContract->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
