package handlers

import (
	"net/http"
	"strconv"

	"github.com/Chintukr2004/go-ecommerce-backend/internal/models"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{
		Service: service,
	}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var p models.Product

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	err := h.Service.Create(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "product created"})
}

func (h *ProductHandler) GetAll(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	search := c.DefaultQuery("search", "")

	products, err := h.Service.GetAll(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch",
		})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	product, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var p models.Product

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	err := h.Service.Update(id, &p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product update",
	})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})

}
