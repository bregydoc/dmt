package dmt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

type API struct {
	port   int
	host   string
	engine *gin.Engine
}

type ChannelStatus struct {
	Name         string    `json:"name"`
	LastActivity time.Time `json:"last_activity"`
}

func (api *API) run(dmt *Engine) error {
	// api.engine.Use(gin.BasicAuth(gin.Accounts{
	//
	// }))
	// TODO
	api.engine.GET("/status", func(c *gin.Context) {
		states := make([]ChannelStatus, 0)
		for _, ch := range dmt.channels {

			states = append(states, ChannelStatus{
				Name:         string(ch.Name()),
				LastActivity: time.Now(), // TODO...
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"channels": states,
		})
	})

	api.engine.POST("/emit", func(c *gin.Context) {
		task := Task{}
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		spew.Dump(task)

		if err := dmt.registerNewTask(task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok, your message has been delivered",
			"task":    task,
		})
	})

	return api.engine.Run(fmt.Sprintf(":%d", api.port))
}
