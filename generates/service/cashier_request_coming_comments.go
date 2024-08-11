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

type CashierRequestComingCommentsService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_client_service.UnimplementedCashierRequestComingCommentsServiceServer
}

func NewCashierRequestComingCommentsService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CashierRequestComingCommentsService {
	return &CashierRequestComingCommentsService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *CashierRequestComingCommentsService) CreateCashierRequestComingComments(ctx context.Context, req *storehouse_client_service.CreateCashierRequestComingCommentsRequest) (resp *storehouse_client_service.CashierRequestComingComments, err error) {

	i.log.Info("---CreateCashierRequestComingComments------>", logger.Any("req", req))

	pKey, err := i.strg.CashierRequestComingComments().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateCashierRequestComingComments->CashierRequestComingComments->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.CashierRequestComingComments().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyCashierRequestComingComments->CashierRequestComingComments->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestComingCommentsService) GetByIDCashierRequestComingComments(ctx context.Context, req *storehouse_client_service.CashierRequestComingCommentsPrimaryKey) (resp *storehouse_client_service.CashierRequestComingComments, err error) {

	i.log.Info("---GetByIDCashierRequestComingComments------>", logger.Any("req", req))

	resp, err = i.strg.CashierRequestComingComments().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDCashierRequestComingComments->CashierRequestComingComments->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestComingCommentsService) GetListCashierRequestComingComments(ctx context.Context, req *storehouse_client_service.GetListCashierRequestComingCommentsRequest) (resp *storehouse_client_service.GetListCashierRequestComingCommentsResponse, err error) {

	i.log.Info("---GetListCashierRequestComingComments------>", logger.Any("req", req))

	resp, err = i.strg.CashierRequestComingComments().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListCashierRequestComingComments->CashierRequestComingComments->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CashierRequestComingCommentsService) UpdateCashierRequestComingComments(ctx context.Context, req *storehouse_client_service.UpdateCashierRequestComingCommentsRequest) (resp *storehouse_client_service.CashierRequestComingComments, err error) {

	i.log.Info("---UpdateCashierRequestComingComments------>", logger.Any("req", req))

	rowsAffected, err := i.strg.CashierRequestComingComments().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateCashierRequestComingComments--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CashierRequestComingComments().GetByPKey(ctx, &storehouse_client_service.CashierRequestComingCommentsPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateCashierRequestComingComments--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CashierRequestComingCommentsService) UpdatePatchCashierRequestComingComments(ctx context.Context, req *storehouse_client_service.UpdatePatchCashierRequestComingCommentsRequest) (resp *storehouse_client_service.CashierRequestComingComments, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.CashierRequestComingComments().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CashierRequestComingComments().GetByPKey(ctx, &storehouse_client_service.CashierRequestComingCommentsPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CashierRequestComingCommentsService) DeleteCashierRequestComingComments(ctx context.Context, req *storehouse_client_service.CashierRequestComingCommentsPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteCashierRequestComingComments------>", logger.Any("req", req))

	err = i.strg.CashierRequestComingComments().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteCashierRequestComingComments->CashierRequestComingComments->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
