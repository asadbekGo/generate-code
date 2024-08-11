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

type ChapterStockUserService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_client_service.UnimplementedChapterStockUserServiceServer
}

func NewChapterStockUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ChapterStockUserService {
	return &ChapterStockUserService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ChapterStockUserService) CreateChapterStockUser(ctx context.Context, req *storehouse_client_service.CreateChapterStockUserRequest) (resp *storehouse_client_service.ChapterStockUser, err error) {

	i.log.Info("---CreateChapterStockUser------>", logger.Any("req", req))

	pKey, err := i.strg.ChapterStockUser().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateChapterStockUser->ChapterStockUser->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.ChapterStockUser().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyChapterStockUser->ChapterStockUser->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ChapterStockUserService) GetByIDChapterStockUser(ctx context.Context, req *storehouse_client_service.ChapterStockUserPrimaryKey) (resp *storehouse_client_service.ChapterStockUser, err error) {

	i.log.Info("---GetByIDChapterStockUser------>", logger.Any("req", req))

	resp, err = i.strg.ChapterStockUser().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDChapterStockUser->ChapterStockUser->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ChapterStockUserService) GetListChapterStockUser(ctx context.Context, req *storehouse_client_service.GetListChapterStockUserRequest) (resp *storehouse_client_service.GetListChapterStockUserResponse, err error) {

	i.log.Info("---GetListChapterStockUser------>", logger.Any("req", req))

	resp, err = i.strg.ChapterStockUser().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListChapterStockUser->ChapterStockUser->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ChapterStockUserService) UpdateChapterStockUser(ctx context.Context, req *storehouse_client_service.UpdateChapterStockUserRequest) (resp *storehouse_client_service.ChapterStockUser, err error) {

	i.log.Info("---UpdateChapterStockUser------>", logger.Any("req", req))

	rowsAffected, err := i.strg.ChapterStockUser().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateChapterStockUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.ChapterStockUser().GetByPKey(ctx, &storehouse_client_service.ChapterStockUserPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateChapterStockUser--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ChapterStockUserService) UpdatePatchChapterStockUser(ctx context.Context, req *storehouse_client_service.UpdatePatchChapterStockUserRequest) (resp *storehouse_client_service.ChapterStockUser, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.ChapterStockUser().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.ChapterStockUser().GetByPKey(ctx, &storehouse_client_service.ChapterStockUserPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ChapterStockUserService) DeleteChapterStockUser(ctx context.Context, req *storehouse_client_service.ChapterStockUserPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteChapterStockUser------>", logger.Any("req", req))

	err = i.strg.ChapterStockUser().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteChapterStockUser->ChapterStockUser->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
