package handlers

import (
	"net/http"

	"github.com/Chintukr2004/go-ecommerce-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{
		Service: service,
	}
}

func (h *OrderHandler) Checkout(c *gin.Context) {
	userID := c.GetInt("user_id")

	err := h.Service.Checkout(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order placed"})

}

func (h *OrderHandler) Get(c *gin.Context) {
	userID := c.GetInt("user_id")

	orders, err := h.Service.GetByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
