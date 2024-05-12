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

type MachineService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedMachineServiceServer
}

func NewMachineService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *MachineService {
	return &MachineService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *MachineService) CreateMachine(ctx context.Context, req *storehouse_service.CreateMachineRequest) (resp *storehouse_service.Machine, err error) {

	i.log.Info("---CreateMachine------>", logger.Any("req", req))

	pKey, err := i.strg.Machine().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateMachine->Machine->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Machine().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyMachine->Machine->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *MachineService) GetByIDMachine(ctx context.Context, req *storehouse_service.MachinePrimaryKey) (resp *storehouse_service.Machine, err error) {

	i.log.Info("---GetByIDMachine------>", logger.Any("req", req))

	resp, err = i.strg.Machine().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDMachine->Machine->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *MachineService) GetListMachine(ctx context.Context, req *storehouse_service.GetListMachineRequest) (resp *storehouse_service.GetListMachineResponse, err error) {

	i.log.Info("---GetListMachine------>", logger.Any("req", req))

	resp, err = i.strg.Machine().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListMachine->Machine->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *MachineService) UpdateMachine(ctx context.Context, req *storehouse_service.UpdateMachineRequest) (resp *storehouse_service.Machine, err error) {

	i.log.Info("---UpdateMachine------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Machine().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateMachine--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Machine().GetByPKey(ctx, &storehouse_service.MachinePrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateMachine--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *MachineService) UpdatePatchMachine(ctx context.Context, req *storehouse_service.UpdatePatchMachineRequest) (resp *storehouse_service.Machine, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Machine().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Machine().GetByPKey(ctx, &storehouse_service.MachinePrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *MachineService) DeleteMachine(ctx context.Context, req *storehouse_service.MachinePrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteMachine------>", logger.Any("req", req))

	err = i.strg.Machine().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteMachine->Machine->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
