type CashierRequestRepoI interface {
	Create(ctx context.Context, req *storehouse_supplier_service.CreateCashierRequestRequest) (resp *storehouse_supplier_service.CashierRequestPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *storehouse_supplier_service.CashierRequestPrimaryKey) (resp *storehouse_supplier_service.CashierRequest, err error)
	GetAll(ctx context.Context, req *storehouse_supplier_service.GetListCashierRequestRequest) (resp *storehouse_supplier_service.GetListCashierRequestResponse, err error)
	Update(ctx context.Context, req *storehouse_supplier_service.UpdateCashierRequestRequest) (rowsAffected int64, err error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *storehouse_supplier_service.CashierRequestPrimaryKey) error
}
type CashierRequestProductRepoI interface {
	Create(ctx context.Context, req *storehouse_supplier_service.CreateCashierRequestProductRequest) (resp *storehouse_supplier_service.CashierRequestProductPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *storehouse_supplier_service.CashierRequestProductPrimaryKey) (resp *storehouse_supplier_service.CashierRequestProduct, err error)
	GetAll(ctx context.Context, req *storehouse_supplier_service.GetListCashierRequestProductRequest) (resp *storehouse_supplier_service.GetListCashierRequestProductResponse, err error)
	Update(ctx context.Context, req *storehouse_supplier_service.UpdateCashierRequestProductRequest) (rowsAffected int64, err error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *storehouse_supplier_service.CashierRequestProductPrimaryKey) error
}
