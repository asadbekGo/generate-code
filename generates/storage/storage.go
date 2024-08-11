type ComingRepoI interface {
	Create(ctx context.Context, req *storehouse_client_service.CreateComingRequest) (resp *storehouse_client_service.ComingPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *storehouse_client_service.ComingPrimaryKey) (resp *storehouse_client_service.Coming, err error)
	GetAll(ctx context.Context, req *storehouse_client_service.GetListComingRequest) (resp *storehouse_client_service.GetListComingResponse, err error)
	Update(ctx context.Context, req *storehouse_client_service.UpdateComingRequest) (rowsAffected int64, err error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *storehouse_client_service.ComingPrimaryKey) error
}
