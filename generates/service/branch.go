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

type BranchService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedBranchServiceServer
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *BranchService {
	return &BranchService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *BranchService) CreateBranch(ctx context.Context, req *storehouse_service.CreateBranchRequest) (resp *storehouse_service.Branch, err error) {

	i.log.Info("---CreateBranch------>", logger.Any("req", req))

	pKey, err := i.strg.Branch().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Branch().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BranchService) GetByIDBranch(ctx context.Context, req *storehouse_service.BranchPrimaryKey) (resp *storehouse_service.Branch, err error) {

	i.log.Info("---GetByIDBranch------>", logger.Any("req", req))

	resp, err = i.strg.Branch().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BranchService) GetListBranch(ctx context.Context, req *storehouse_service.GetListBranchRequest) (resp *storehouse_service.GetListBranchResponse, err error) {

	i.log.Info("---GetListBranch------>", logger.Any("req", req))

	resp, err = i.strg.Branch().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BranchService) UpdateBranch(ctx context.Context, req *storehouse_service.UpdateBranchRequest) (resp *storehouse_service.Branch, err error) {

	i.log.Info("---UpdateBranch------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Branch().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateBranch--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Branch().GetByPKey(ctx, &storehouse_service.BranchPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateBranch--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *BranchService) UpdatePatchBranch(ctx context.Context, req *storehouse_service.UpdatePatchBranchRequest) (resp *storehouse_service.Branch, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Branch().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Branch().GetByPKey(ctx, &storehouse_service.BranchPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *BranchService) DeleteBranch(ctx context.Context, req *storehouse_service.BranchPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteBranch------>", logger.Any("req", req))

	err = i.strg.Branch().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
