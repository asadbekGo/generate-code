// Branch ..
r.POST("/branch", h.CreateBranch)
r.GET("/branch/:branch_id", h.GetSingleBranch)
r.GET("/branch", h.GetBranchList)
r.PUT("/branch", h.UpdateBranch)
r.DELETE("/branch/:branch_id", h.DeleteBranch)

// User ..
r.POST("/user", h.CreateUser)
r.GET("/user/:user_id", h.GetSingleUser)
r.GET("/user", h.GetUserList)
r.PUT("/user", h.UpdateUser)
r.DELETE("/user/:user_id", h.DeleteUser)

// ChapterMachine ..
r.POST("/chapter-machine", h.CreateChapterMachine)
r.GET("/chapter-machine/:chapter_machine_id", h.GetSingleChapterMachine)
r.GET("/chapter-machine", h.GetChapterMachineList)
r.PUT("/chapter-machine", h.UpdateChapterMachine)
r.DELETE("/chapter-machine/:chapter_machine_id", h.DeleteChapterMachine)

// Machine ..
r.POST("/machine", h.CreateMachine)
r.GET("/machine/:machine_id", h.GetSingleMachine)
r.GET("/machine", h.GetMachineList)
r.PUT("/machine", h.UpdateMachine)
r.DELETE("/machine/:machine_id", h.DeleteMachine)

// Product ..
r.POST("/product", h.CreateProduct)
r.GET("/product/:product_id", h.GetSingleProduct)
r.GET("/product", h.GetProductList)
r.PUT("/product", h.UpdateProduct)
r.DELETE("/product/:product_id", h.DeleteProduct)

// ChapterStock ..
r.POST("/chapter-stock", h.CreateChapterStock)
r.GET("/chapter-stock/:chapter_stock_id", h.GetSingleChapterStock)
r.GET("/chapter-stock", h.GetChapterStockList)
r.PUT("/chapter-stock", h.UpdateChapterStock)
r.DELETE("/chapter-stock/:chapter_stock_id", h.DeleteChapterStock)

// StockProduct ..
r.POST("/stock-product", h.CreateStockProduct)
r.GET("/stock-product/:stock_product_id", h.GetSingleStockProduct)
r.GET("/stock-product", h.GetStockProductList)
r.PUT("/stock-product", h.UpdateStockProduct)
r.DELETE("/stock-product/:stock_product_id", h.DeleteStockProduct)

// WriteOff ..
r.POST("/write-off", h.CreateWriteOff)
r.GET("/write-off/:write_off_id", h.GetSingleWriteOff)
r.GET("/write-off", h.GetWriteOffList)
r.PUT("/write-off", h.UpdateWriteOff)
r.DELETE("/write-off/:write_off_id", h.DeleteWriteOff)

// Request ..
r.POST("/request", h.CreateRequest)
r.GET("/request/:request_id", h.GetSingleRequest)
r.GET("/request", h.GetRequestList)
r.PUT("/request", h.UpdateRequest)
r.DELETE("/request/:request_id", h.DeleteRequest)

// RequestProduct ..
r.POST("/request-product", h.CreateRequestProduct)
r.GET("/request-product/:request_product_id", h.GetSingleRequestProduct)
r.GET("/request-product", h.GetRequestProductList)
r.PUT("/request-product", h.UpdateRequestProduct)
r.DELETE("/request-product/:request_product_id", h.DeleteRequestProduct)

// TransferSend ..
r.POST("/transfer-send", h.CreateTransferSend)
r.GET("/transfer-send/:transfer_send_id", h.GetSingleTransferSend)
r.GET("/transfer-send", h.GetTransferSendList)
r.PUT("/transfer-send", h.UpdateTransferSend)
r.DELETE("/transfer-send/:transfer_send_id", h.DeleteTransferSend)

// TransferSendProduct ..
r.POST("/transfer-send-product", h.CreateTransferSendProduct)
r.GET("/transfer-send-product/:transfer_send_product_id", h.GetSingleTransferSendProduct)
r.GET("/transfer-send-product", h.GetTransferSendProductList)
r.PUT("/transfer-send-product", h.UpdateTransferSendProduct)
r.DELETE("/transfer-send-product/:transfer_send_product_id", h.DeleteTransferSendProduct)

// TransferReceive ..
r.POST("/transfer-receive", h.CreateTransferReceive)
r.GET("/transfer-receive/:transfer_receive_id", h.GetSingleTransferReceive)
r.GET("/transfer-receive", h.GetTransferReceiveList)
r.PUT("/transfer-receive", h.UpdateTransferReceive)
r.DELETE("/transfer-receive/:transfer_receive_id", h.DeleteTransferReceive)

// TransferReceiveProduct ..
r.POST("/transfer-receive-product", h.CreateTransferReceiveProduct)
r.GET("/transfer-receive-product/:transfer_receive_product_id", h.GetSingleTransferReceiveProduct)
r.GET("/transfer-receive-product", h.GetTransferReceiveProductList)
r.PUT("/transfer-receive-product", h.UpdateTransferReceiveProduct)
r.DELETE("/transfer-receive-product/:transfer_receive_product_id", h.DeleteTransferReceiveProduct)

