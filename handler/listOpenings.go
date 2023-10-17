package handler

import (
	"net/http"

	"github.com/Guilherme-daCosta/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningsHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSucess(c, "list-openings", openings)
}
