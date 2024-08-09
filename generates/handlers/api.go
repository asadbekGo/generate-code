// Client ..
v1.POST("/client", s.HandlerClient.CreateClient)
v1.GET("/client/:client_id", s.HandlerClient.GetSingleClient)
v1.GET("/client", s.HandlerClient.GetClientList)
v1.PUT("/client", s.HandlerClient.UpdateClient)
v1.DELETE("/client/:client_id", s.HandlerClient.DeleteClient)

// ClientContract ..
v1.POST("/client-contract", s.HandlerClient.CreateClientContract)
v1.GET("/client-contract/:client_contract_id", s.HandlerClient.GetSingleClientContract)
v1.GET("/client-contract", s.HandlerClient.GetClientContractList)
v1.PUT("/client-contract", s.HandlerClient.UpdateClientContract)
v1.DELETE("/client-contract/:client_contract_id", s.HandlerClient.DeleteClientContract)

