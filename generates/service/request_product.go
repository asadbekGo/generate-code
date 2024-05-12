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

type RequestProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedRequestProductServiceServer
}

func NewRequestProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *RequestProductService {
	return &RequestProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *RequestProductService) CreateRequestProduct(ctx context.Context, req *storehouse_service.CreateRequestProductRequest) (resp *storehouse_service.RequestProduct, err error) {

	i.log.Info("---CreateRequestProduct------>", logger.Any("req", req))

	pKey, err := i.strg.RequestProduct().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateRequestProduct->RequestProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.RequestProduct().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyRequestProduct->RequestProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *RequestProductService) GetByIDRequestProduct(ctx context.Context, req *storehouse_service.RequestProductPrimaryKey) (resp *storehouse_service.RequestProduct, err error) {

	i.log.Info("---GetByIDRequestProduct------>", logger.Any("req", req))

	resp, err = i.strg.RequestProduct().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDRequestProduct->RequestProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *RequestProductService) GetListRequestProduct(ctx context.Context, req *storehouse_service.GetListRequestProductRequest) (resp *storehouse_service.GetListRequestProductResponse, err error) {

	i.log.Info("---GetListRequestProduct------>", logger.Any("req", req))

	resp, err = i.strg.RequestProduct().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListRequestProduct->RequestProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *RequestProductService) UpdateRequestProduct(ctx context.Context, req *storehouse_service.UpdateRequestProductRequest) (resp *storehouse_service.RequestProduct, err error) {

	i.log.Info("---UpdateRequestProduct------>", logger.Any("req", req))

	rowsAffected, err := i.strg.RequestProduct().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateRequestProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.RequestProduct().GetByPKey(ctx, &storehouse_service.RequestProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateRequestProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *RequestProductService) UpdatePatchRequestProduct(ctx context.Context, req *storehouse_service.UpdatePatchRequestProductRequest) (resp *storehouse_service.RequestProduct, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.RequestProduct().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.RequestProduct().GetByPKey(ctx, &storehouse_service.RequestProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *RequestProductService) DeleteRequestProduct(ctx context.Context, req *storehouse_service.RequestProductPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteRequestProduct------>", logger.Any("req", req))

	err = i.strg.RequestProduct().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteRequestProduct->RequestProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
