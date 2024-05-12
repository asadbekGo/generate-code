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

type ProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedProductServiceServer
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ProductService {
	return &ProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ProductService) CreateProduct(ctx context.Context, req *storehouse_service.CreateProductRequest) (resp *storehouse_service.Product, err error) {

	i.log.Info("---CreateProduct------>", logger.Any("req", req))

	pKey, err := i.strg.Product().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateProduct->Product->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Product().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyProduct->Product->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ProductService) GetByIDProduct(ctx context.Context, req *storehouse_service.ProductPrimaryKey) (resp *storehouse_service.Product, err error) {

	i.log.Info("---GetByIDProduct------>", logger.Any("req", req))

	resp, err = i.strg.Product().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDProduct->Product->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ProductService) GetListProduct(ctx context.Context, req *storehouse_service.GetListProductRequest) (resp *storehouse_service.GetListProductResponse, err error) {

	i.log.Info("---GetListProduct------>", logger.Any("req", req))

	resp, err = i.strg.Product().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListProduct->Product->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ProductService) UpdateProduct(ctx context.Context, req *storehouse_service.UpdateProductRequest) (resp *storehouse_service.Product, err error) {

	i.log.Info("---UpdateProduct------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Product().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Product().GetByPKey(ctx, &storehouse_service.ProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ProductService) UpdatePatchProduct(ctx context.Context, req *storehouse_service.UpdatePatchProductRequest) (resp *storehouse_service.Product, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Product().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Product().GetByPKey(ctx, &storehouse_service.ProductPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ProductService) DeleteProduct(ctx context.Context, req *storehouse_service.ProductPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteProduct------>", logger.Any("req", req))

	err = i.strg.Product().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteProduct->Product->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
