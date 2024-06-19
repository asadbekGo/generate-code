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

type CashierRequestProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_supplier_service.UnimplementedCashierRequestProductServiceServer
}

func NewCashierRequestProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CashierRequestProductService {
	return &CashierRequestProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *CashierRequestProductService) CreateCashierRequestProduct(ctx context.Context, req *storehouse_supplier_service.CreateCashierRequestProductRequest) (resp *storehouse_supplier_service.CashierRequestProduct, err error) {

	i.log.Info("---CreateCashierRequestProduct------>", logger.Any("req", req))

	pKey, err := i.strg.CashierRequestProduct().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateCashierRequestProduct->CashierRequestProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.CashierRequestProduct().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyCashierRequestProduct->CashierRequestProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestProductService) GetByIDCashierRequestProduct(ctx context.Context, req *storehouse_supplier_service.CashierRequestProductPrimaryKey) (resp *storehouse_supplier_service.CashierRequestProduct, err error) {

	i.log.Info("---GetByIDCashierRequestProduct------>", logger.Any("req", req))

	resp, err = i.strg.CashierRequestProduct().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDCashierRequestProduct->CashierRequestProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestProductService) GetListCashierRequestProduct(ctx context.Context, req *storehouse_supplier_service.GetListCashierRequestProductRequest) (resp *storehouse_supplier_service.GetListCashierRequestProductResponse, err error) {

	i.log.Info("---GetListCashierRequestProduct------>", logger.Any("req", req))

	resp, err = i.strg.CashierRequestProduct().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListCashierRequestProduct->CashierRequestProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestProductService) UpdateCashierRequestProduct(ctx context.Context, req *storehouse_supplier_service.UpdateCashierRequestProductRequest) (resp *storehouse_supplier_service.CashierRequestProduct, err error) {

	i.log.Info("---UpdateCashierRequestProduct------>", logger.Any("req", req))

	rowsAffected, err := i.strg.CashierRequestProduct().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateCashierRequestProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CashierRequestProduct().GetByPKey(ctx, &storehouse_supplier_service.CashierRequestProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateCashierRequestProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CashierRequestProductService) UpdatePatchCashierRequestProduct(ctx context.Context, req *storehouse_supplier_service.UpdatePatchCashierRequestProductRequest) (resp *storehouse_supplier_service.CashierRequestProduct, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.CashierRequestProduct().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CashierRequestProduct().GetByPKey(ctx, &storehouse_supplier_service.CashierRequestProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CashierRequestProductService) DeleteCashierRequestProduct(ctx context.Context, req *storehouse_supplier_service.CashierRequestProductPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteCashierRequestProduct------>", logger.Any("req", req))

	err = i.strg.CashierRequestProduct().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteCashierRequestProduct->CashierRequestProduct->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
