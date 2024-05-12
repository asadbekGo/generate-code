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

type ChapterStockService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedChapterStockServiceServer
}

func NewChapterStockService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ChapterStockService {
	return &ChapterStockService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ChapterStockService) CreateChapterStock(ctx context.Context, req *storehouse_service.CreateChapterStockRequest) (resp *storehouse_service.ChapterStock, err error) {

	i.log.Info("---CreateChapterStock------>", logger.Any("req", req))

	pKey, err := i.strg.ChapterStock().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateChapterStock->ChapterStock->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.ChapterStock().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyChapterStock->ChapterStock->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ChapterStockService) GetByIDChapterStock(ctx context.Context, req *storehouse_service.ChapterStockPrimaryKey) (resp *storehouse_service.ChapterStock, err error) {

	i.log.Info("---GetByIDChapterStock------>", logger.Any("req", req))

	resp, err = i.strg.ChapterStock().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDChapterStock->ChapterStock->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ChapterStockService) GetListChapterStock(ctx context.Context, req *storehouse_service.GetListChapterStockRequest) (resp *storehouse_service.GetListChapterStockResponse, err error) {

	i.log.Info("---GetListChapterStock------>", logger.Any("req", req))

	resp, err = i.strg.ChapterStock().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListChapterStock->ChapterStock->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ChapterStockService) UpdateChapterStock(ctx context.Context, req *storehouse_service.UpdateChapterStockRequest) (resp *storehouse_service.ChapterStock, err error) {

	i.log.Info("---UpdateChapterStock------>", logger.Any("req", req))

	rowsAffected, err := i.strg.ChapterStock().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateChapterStock--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.ChapterStock().GetByPKey(ctx, &storehouse_service.ChapterStockPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateChapterStock--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ChapterStockService) UpdatePatchChapterStock(ctx context.Context, req *storehouse_service.UpdatePatchChapterStockRequest) (resp *storehouse_service.ChapterStock, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.ChapterStock().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.ChapterStock().GetByPKey(ctx, &storehouse_service.ChapterStockPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ChapterStockService) DeleteChapterStock(ctx context.Context, req *storehouse_service.ChapterStockPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteChapterStock------>", logger.Any("req", req))

	err = i.strg.ChapterStock().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteChapterStock->ChapterStock->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
