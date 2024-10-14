package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	findmetric "github.com/tecwagner/frete_rapido_api/internal/useCase/find_metric"
)

type WebMetricsHandler struct {
	UseCase findmetric.FindMentricUseCase
}

func NewWebMetricsHandler(useCase findmetric.FindMentricUseCase) *WebMetricsHandler {
	return &WebMetricsHandler{
		UseCase: useCase,
	}
}

func (h *WebMetricsHandler) GetMetrics(c *gin.Context) {
	lastQuotesParam := c.Query("last_quotes")
	var lastQuotes *int

	if lastQuotesParam != "" {
		lastQuotesValue, err := strconv.Atoi(lastQuotesParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "last_quotes must be an integer"})
			return
		}
		lastQuotes = &lastQuotesValue
	}

	metrics, err := h.UseCase.Execute(c.Request.Context(), lastQuotes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, metrics)
}
