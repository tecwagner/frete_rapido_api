package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	createquote "github.com/tecwagner/frete_rapido_api/internal/useCase/create_quote"
)

type WebQuoteHandler struct {
	UseCase createquote.CreateQuoteUseCase
}

func NewWebQuoteHandler(useCase createquote.CreateQuoteUseCase) *WebQuoteHandler {
	return &WebQuoteHandler{
		UseCase: useCase,
	}
}

func (h *WebQuoteHandler) CreateQuote(c *gin.Context) {
	var dto createquote.CreateQuoteInputDTO

	// Bind JSON input to DTO using Gin
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Execute the use case
	ctx := c.Request.Context()
	output, err := h.UseCase.Execute(ctx, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quote"})
		return
	}

	// Return the output in JSON format
	c.JSON(http.StatusCreated, output)
}
