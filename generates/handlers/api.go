// Tender ..
v1.POST("/tender", s.HandlerSupplier.CreateTender)
v1.GET("/tender/:tender_id", s.HandlerSupplier.GetSingleTender)
v1.GET("/tender", s.HandlerSupplier.GetTenderList)
v1.PUT("/tender", s.HandlerSupplier.UpdateTender)
v1.DELETE("/tender/:tender_id", s.HandlerSupplier.DeleteTender)

