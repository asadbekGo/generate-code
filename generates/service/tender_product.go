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

type TenderProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedTenderProductServiceServer
}

func NewTenderProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TenderProductService {
	return &TenderProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TenderProductService) CreateTenderProduct(ctx context.Context, req *storehouse_service.CreateTenderProductRequest) (resp *storehouse_service.TenderProduct, err error) {

	i.log.Info("---CreateTenderProduct------>", logger.Any("req", req))

	pKey, err := i.strg.TenderProduct().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTenderProduct->TenderProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.TenderProduct().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyTenderProduct->TenderProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TenderProductService) GetByIDTenderProduct(ctx context.Context, req *storehouse_service.TenderProductPrimaryKey) (resp *storehouse_service.TenderProduct, err error) {

	i.log.Info("---GetByIDTenderProduct------>", logger.Any("req", req))

	resp, err = i.strg.TenderProduct().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDTenderProduct->TenderProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TenderProductService) GetListTenderProduct(ctx context.Context, req *storehouse_service.GetListTenderProductRequest) (resp *storehouse_service.GetListTenderProductResponse, err error) {

	i.log.Info("---GetListTenderProduct------>", logger.Any("req", req))

	resp, err = i.strg.TenderProduct().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListTenderProduct->TenderProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TenderProductService) UpdateTenderProduct(ctx context.Context, req *storehouse_service.UpdateTenderProductRequest) (resp *storehouse_service.TenderProduct, err error) {

	i.log.Info("---UpdateTenderProduct------>", logger.Any("req", req))

	rowsAffected, err := i.strg.TenderProduct().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateTenderProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TenderProduct().GetByPKey(ctx, &storehouse_service.TenderProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateTenderProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TenderProductService) UpdatePatchTenderProduct(ctx context.Context, req *storehouse_service.UpdatePatchTenderProductRequest) (resp *storehouse_service.TenderProduct, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.TenderProduct().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TenderProduct().GetByPKey(ctx, &storehouse_service.TenderProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TenderProductService) DeleteTenderProduct(ctx context.Context, req *storehouse_service.TenderProductPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteTenderProduct------>", logger.Any("req", req))

	err = i.strg.TenderProduct().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTenderProduct->TenderProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
