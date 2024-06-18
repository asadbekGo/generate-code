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

type SupplierService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedSupplierServiceServer
}

func NewSupplierService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SupplierService {
	return &SupplierService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *SupplierService) CreateSupplier(ctx context.Context, req *storehouse_service.CreateSupplierRequest) (resp *storehouse_service.Supplier, err error) {

	i.log.Info("---CreateSupplier------>", logger.Any("req", req))

	pKey, err := i.strg.Supplier().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateSupplier->Supplier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Supplier().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeySupplier->Supplier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SupplierService) GetByIDSupplier(ctx context.Context, req *storehouse_service.SupplierPrimaryKey) (resp *storehouse_service.Supplier, err error) {

	i.log.Info("---GetByIDSupplier------>", logger.Any("req", req))

	resp, err = i.strg.Supplier().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDSupplier->Supplier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SupplierService) GetListSupplier(ctx context.Context, req *storehouse_service.GetListSupplierRequest) (resp *storehouse_service.GetListSupplierResponse, err error) {

	i.log.Info("---GetListSupplier------>", logger.Any("req", req))

	resp, err = i.strg.Supplier().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListSupplier->Supplier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SupplierService) UpdateSupplier(ctx context.Context, req *storehouse_service.UpdateSupplierRequest) (resp *storehouse_service.Supplier, err error) {

	i.log.Info("---UpdateSupplier------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Supplier().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateSupplier--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Supplier().GetByPKey(ctx, &storehouse_service.SupplierPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateSupplier--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *SupplierService) UpdatePatchSupplier(ctx context.Context, req *storehouse_service.UpdatePatchSupplierRequest) (resp *storehouse_service.Supplier, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Supplier().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Supplier().GetByPKey(ctx, &storehouse_service.SupplierPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *SupplierService) DeleteSupplier(ctx context.Context, req *storehouse_service.SupplierPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteSupplier------>", logger.Any("req", req))

	err = i.strg.Supplier().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteSupplier->Supplier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
