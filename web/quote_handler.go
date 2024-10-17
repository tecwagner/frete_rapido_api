package web

import (
	"bytes"
	"io"
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

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	output, err := h.UseCase.Execute(ctx, dto)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}
