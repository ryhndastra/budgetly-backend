package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/", h.Create)
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid user_id",
			},
		)
		return
	}

	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid category_id",
			},
		)
		return
	}

	transaction := &Transaction{
		UserID:     userID,
		CategoryID: categoryID,
		Title:      req.Title,
		Amount:     req.Amount,
		Note:       req.Note,
		Type:       req.Type,
	}

	result, err := h.service.Create(transaction)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		result,
	)
}