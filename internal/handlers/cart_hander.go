package handlers

import (
	"net/http"

	"github.com/Chintukr2004/go-ecommerce-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	Service *services.CartService
}

type AddCartRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func NewCartHandler(service *services.CartService) *CartHandler {
	return &CartHandler{
		Service: service,
	}
}

func (h *CartHandler) Add(c *gin.Context) {
	userID := c.GetInt("user_id")

	var req AddCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.Service.Add(userID, req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "added to cart "})
}

func (h *CartHandler) Get(c *gin.Context) {
	userID := c.GetInt("user_id")

	items, err := h.Service.GetByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetched",
		})
		return
	}

	c.JSON(http.StatusOK, items)
}
