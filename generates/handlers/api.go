// Coming ..
v1.POST("/coming", s.HandlerClient.CreateComing)
v1.GET("/coming/:coming_id", s.HandlerClient.GetSingleComing)
v1.GET("/coming", s.HandlerClient.GetComingList)
v1.PUT("/coming", s.HandlerClient.UpdateComing)
v1.DELETE("/coming/:coming_id", s.HandlerClient.DeleteComing)

