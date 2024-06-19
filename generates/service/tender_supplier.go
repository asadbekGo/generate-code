package supplier_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"warehouse/warehouse_go_storehouse_service/config"
	"warehouse/warehouse_go_storehouse_service/genproto/storehouse_supplier_service"
	"warehouse/warehouse_go_storehouse_service/grpc/client"
	"warehouse/warehouse_go_storehouse_service/models"
	"warehouse/warehouse_go_storehouse_service/pkg/logger"
	"warehouse/warehouse_go_storehouse_service/storage"
)

type TenderSupplierService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_supplier_service.UnimplementedTenderSupplierServiceServer
}

func NewTenderSupplierService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TenderSupplierService {
	return &TenderSupplierService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TenderSupplierService) CreateTenderSupplier(ctx context.Context, req *storehouse_supplier_service.CreateTenderSupplierRequest) (resp *storehouse_supplier_service.TenderSupplier, err error) {

	i.log.Info("---CreateTenderSupplier------>", logger.Any("req", req))

	pKey, err := i.strg.TenderSupplier().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTenderSupplier->TenderSupplier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.TenderSupplier().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyTenderSupplier->TenderSupplier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TenderSupplierService) GetByIDTenderSupplier(ctx context.Context, req *storehouse_supplier_service.TenderSupplierPrimaryKey) (resp *storehouse_supplier_service.TenderSupplier, err error) {

	i.log.Info("---GetByIDTenderSupplier------>", logger.Any("req", req))

	resp, err = i.strg.TenderSupplier().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDTenderSupplier->TenderSupplier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TenderSupplierService) GetListTenderSupplier(ctx context.Context, req *storehouse_supplier_service.GetListTenderSupplierRequest) (resp *storehouse_supplier_service.GetListTenderSupplierResponse, err error) {

	i.log.Info("---GetListTenderSupplier------>", logger.Any("req", req))

	resp, err = i.strg.TenderSupplier().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListTenderSupplier->TenderSupplier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TenderSupplierService) UpdateTenderSupplier(ctx context.Context, req *storehouse_supplier_service.UpdateTenderSupplierRequest) (resp *storehouse_supplier_service.TenderSupplier, err error) {

	i.log.Info("---UpdateTenderSupplier------>", logger.Any("req", req))

	rowsAffected, err := i.strg.TenderSupplier().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateTenderSupplier--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TenderSupplier().GetByPKey(ctx, &storehouse_supplier_service.TenderSupplierPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateTenderSupplier--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TenderSupplierService) UpdatePatchTenderSupplier(ctx context.Context, req *storehouse_supplier_service.UpdatePatchTenderSupplierRequest) (resp *storehouse_supplier_service.TenderSupplier, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.TenderSupplier().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.TenderSupplier().GetByPKey(ctx, &storehouse_supplier_service.TenderSupplierPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TenderSupplierService) DeleteTenderSupplier(ctx context.Context, req *storehouse_supplier_service.TenderSupplierPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteTenderSupplier------>", logger.Any("req", req))

	err = i.strg.TenderSupplier().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTenderSupplier->TenderSupplier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
