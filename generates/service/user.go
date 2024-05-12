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

type UserService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *UserService {
	return &UserService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *UserService) CreateUser(ctx context.Context, req *storehouse_service.CreateUserRequest) (resp *storehouse_service.User, err error) {

	i.log.Info("---CreateUser------>", logger.Any("req", req))

	pKey, err := i.strg.User().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.User().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *UserService) GetByIDUser(ctx context.Context, req *storehouse_service.UserPrimaryKey) (resp *storehouse_service.User, err error) {

	i.log.Info("---GetByIDUser------>", logger.Any("req", req))

	resp, err = i.strg.User().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *UserService) GetListUser(ctx context.Context, req *storehouse_service.GetListUserRequest) (resp *storehouse_service.GetListUserResponse, err error) {

	i.log.Info("---GetListUser------>", logger.Any("req", req))

	resp, err = i.strg.User().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *UserService) UpdateUser(ctx context.Context, req *storehouse_service.UpdateUserRequest) (resp *storehouse_service.User, err error) {

	i.log.Info("---UpdateUser------>", logger.Any("req", req))

	rowsAffected, err := i.strg.User().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.User().GetByPKey(ctx, &storehouse_service.UserPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateUser--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *UserService) UpdatePatchUser(ctx context.Context, req *storehouse_service.UpdatePatchUserRequest) (resp *storehouse_service.User, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.User().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.User().GetByPKey(ctx, &storehouse_service.UserPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *UserService) DeleteUser(ctx context.Context, req *storehouse_service.UserPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteUser------>", logger.Any("req", req))

	err = i.strg.User().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}