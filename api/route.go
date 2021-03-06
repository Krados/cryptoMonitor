package api

import (
	"cryptoMonitor/cache"
	"cryptoMonitor/lib"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start() {
	r := gin.Default()
	r.GET("/status/:symbol", func(c *gin.Context) {
		var resp lib.DirectionPrediction
		valByte, err := cache.Get().Get([]byte(c.Param("symbol")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		err = json.Unmarshal(valByte, &resp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, resp)
	})
	r.Run()
}
