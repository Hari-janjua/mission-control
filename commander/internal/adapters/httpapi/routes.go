package httpapi

import (
	"mission-control/commander/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, missionSvc *service.MissionService, authSvc *service.AuthService) {
	r.POST("/missions", func(c *gin.Context) {
		var payload struct {
			Command string `json:"command"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := missionSvc.CreateMission(payload.Command)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"mission_id": id})
	})

	r.GET("/missions/:id", func(c *gin.Context) {
		id := c.Param("id")
		m, err := missionSvc.GetMission(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	})

	r.POST("/auth/token", func(c *gin.Context) {
		var req struct {
			SoldierID string `json:"soldier_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, err := authSvc.GenerateToken(req.SoldierID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})
}
