// CashierRequest ..
v1.POST("/cashier-request", s.HandlerSupplier.CreateCashierRequest)
v1.GET("/cashier-request/:cashier_request_id", s.HandlerSupplier.GetSingleCashierRequest)
v1.GET("/cashier-request", s.HandlerSupplier.GetCashierRequestList)
v1.PUT("/cashier-request", s.HandlerSupplier.UpdateCashierRequest)
v1.DELETE("/cashier-request/:cashier_request_id", s.HandlerSupplier.DeleteCashierRequest)

// CashierRequestProduct ..
v1.POST("/cashier-request-product", s.HandlerSupplier.CreateCashierRequestProduct)
v1.GET("/cashier-request-product/:cashier_request_product_id", s.HandlerSupplier.GetSingleCashierRequestProduct)
v1.GET("/cashier-request-product", s.HandlerSupplier.GetCashierRequestProductList)
v1.PUT("/cashier-request-product", s.HandlerSupplier.UpdateCashierRequestProduct)
v1.DELETE("/cashier-request-product/:cashier_request_product_id", s.HandlerSupplier.DeleteCashierRequestProduct)

