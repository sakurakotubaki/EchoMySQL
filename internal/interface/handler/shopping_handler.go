package handler

import (
	"net/http"
	"strconv"

	"myapi/internal/usecase"

	"github.com/labstack/echo/v4"
)

type ShoppingHandler struct {
	usecase usecase.ShoppingUsecase
}

// NewShoppingHandler creates a new shopping handler
func NewShoppingHandler(u usecase.ShoppingUsecase) *ShoppingHandler {
	return &ShoppingHandler{usecase: u}
}

func (h *ShoppingHandler) CreateItem(c echo.Context) error {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := h.usecase.CreateItem(req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Item created successfully"})
}

func (h *ShoppingHandler) GetAllItems(c echo.Context) error {
	items, err := h.usecase.GetAllItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ShoppingHandler) GetItem(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	item, err := h.usecase.GetItem(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
	}

	return c.JSON(http.StatusOK, item)
}

func (h *ShoppingHandler) UpdateItem(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var req struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = h.usecase.UpdateItem(uint(id), req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Item updated successfully"})
}

func (h *ShoppingHandler) DeleteItem(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	err = h.usecase.DeleteItem(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Item deleted successfully"})
}
