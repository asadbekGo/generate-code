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

type ChapterMachineService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedChapterMachineServiceServer
}

func NewChapterMachineService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ChapterMachineService {
	return &ChapterMachineService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ChapterMachineService) CreateChapterMachine(ctx context.Context, req *storehouse_service.CreateChapterMachineRequest) (resp *storehouse_service.ChapterMachine, err error) {

	i.log.Info("---CreateChapterMachine------>", logger.Any("req", req))

	pKey, err := i.strg.ChapterMachine().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateChapterMachine->ChapterMachine->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.ChapterMachine().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyChapterMachine->ChapterMachine->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ChapterMachineService) GetByIDChapterMachine(ctx context.Context, req *storehouse_service.ChapterMachinePrimaryKey) (resp *storehouse_service.ChapterMachine, err error) {

	i.log.Info("---GetByIDChapterMachine------>", logger.Any("req", req))

	resp, err = i.strg.ChapterMachine().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDChapterMachine->ChapterMachine->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ChapterMachineService) GetListChapterMachine(ctx context.Context, req *storehouse_service.GetListChapterMachineRequest) (resp *storehouse_service.GetListChapterMachineResponse, err error) {

	i.log.Info("---GetListChapterMachine------>", logger.Any("req", req))

	resp, err = i.strg.ChapterMachine().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListChapterMachine->ChapterMachine->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ChapterMachineService) UpdateChapterMachine(ctx context.Context, req *storehouse_service.UpdateChapterMachineRequest) (resp *storehouse_service.ChapterMachine, err error) {

	i.log.Info("---UpdateChapterMachine------>", logger.Any("req", req))

	rowsAffected, err := i.strg.ChapterMachine().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateChapterMachine--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.ChapterMachine().GetByPKey(ctx, &storehouse_service.ChapterMachinePrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateChapterMachine--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ChapterMachineService) UpdatePatchChapterMachine(ctx context.Context, req *storehouse_service.UpdatePatchChapterMachineRequest) (resp *storehouse_service.ChapterMachine, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.ChapterMachine().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.ChapterMachine().GetByPKey(ctx, &storehouse_service.ChapterMachinePrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ChapterMachineService) DeleteChapterMachine(ctx context.Context, req *storehouse_service.ChapterMachinePrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteChapterMachine------>", logger.Any("req", req))

	err = i.strg.ChapterMachine().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteChapterMachine->ChapterMachine->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
