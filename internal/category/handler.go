package category

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
	router.GET("/", h.Get)
}

func (h *Handler) Get(c *gin.Context) {

	userIDParam := c.Query("user_id")

	userID, err := uuid.Parse(userIDParam)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid user_id",
			},
		)
		return
	}

	categories, err := h.service.GetByUserID(userID)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		categories,
	)
}