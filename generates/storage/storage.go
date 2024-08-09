type ClientRepoI interface {
	Create(ctx context.Context, req *storehouse_client_service.CreateClientRequest) (resp *storehouse_client_service.ClientPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *storehouse_client_service.ClientPrimaryKey) (resp *storehouse_client_service.Client, err error)
	GetAll(ctx context.Context, req *storehouse_client_service.GetListClientRequest) (resp *storehouse_client_service.GetListClientResponse, err error)
	Update(ctx context.Context, req *storehouse_client_service.UpdateClientRequest) (rowsAffected int64, err error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *storehouse_client_service.ClientPrimaryKey) error
}
type ClientContractRepoI interface {
	Create(ctx context.Context, req *storehouse_client_service.CreateClientContractRequest) (resp *storehouse_client_service.ClientContractPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *storehouse_client_service.ClientContractPrimaryKey) (resp *storehouse_client_service.ClientContract, err error)
	GetAll(ctx context.Context, req *storehouse_client_service.GetListClientContractRequest) (resp *storehouse_client_service.GetListClientContractResponse, err error)
	Update(ctx context.Context, req *storehouse_client_service.UpdateClientContractRequest) (rowsAffected int64, err error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *storehouse_client_service.ClientContractPrimaryKey) error
}
