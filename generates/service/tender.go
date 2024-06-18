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

type TenderService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	storehouse_service.UnimplementedTenderServiceServer
}

func NewTenderService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TenderService {
	return &TenderService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TenderService) CreateTender(ctx context.Context, req *storehouse_service.CreateTenderRequest) (resp *storehouse_service.Tender, err error) {

	i.log.Info("---CreateTender------>", logger.Any("req", req))

	pKey, err := i.strg.Tender().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTender->Tender->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Tender().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyTender->Tender->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TenderService) GetByIDTender(ctx context.Context, req *storehouse_service.TenderPrimaryKey) (resp *storehouse_service.Tender, err error) {

	i.log.Info("---GetByIDTender------>", logger.Any("req", req))

	resp, err = i.strg.Tender().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetByIDTender->Tender->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TenderService) GetListTender(ctx context.Context, req *storehouse_service.GetListTenderRequest) (resp *storehouse_service.GetListTenderResponse, err error) {

	i.log.Info("---GetListTender------>", logger.Any("req", req))

	resp, err = i.strg.Tender().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetListTender->Tender->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TenderService) UpdateTender(ctx context.Context, req *storehouse_service.UpdateTenderRequest) (resp *storehouse_service.Tender, err error) {

	i.log.Info("---UpdateTender------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Tender().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateTender--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Tender().GetByPKey(ctx, &storehouse_service.TenderPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateTender--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TenderService) UpdatePatchTender(ctx context.Context, req *storehouse_service.UpdatePatchTenderRequest) (resp *storehouse_service.Tender, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))
	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Tender().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Tender().GetByPKey(ctx, &storehouse_service.TenderPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdatePatchOrder--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TenderService) DeleteTender(ctx context.Context, req *storehouse_service.TenderPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteTender------>", logger.Any("req", req))

	err = i.strg.Tender().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTender->Tender->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
