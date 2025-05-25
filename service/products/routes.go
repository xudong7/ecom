package products

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xudong7/ecom/types"
	"github.com/xudong7/ecom/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// valid JSON payload
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"invalid payload: %v", err,
		))
		return
	}

	// check if the product already exists
	_, err := h.store.GetProductByName(payload.Name)
	// sql.ErrNoRows means the product does not exist but no error
	if err != nil && err != sql.ErrNoRows {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if err == nil { // product already exists
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"product with name %s already exists", payload.Name,
		))
		return
	}

	// if not, create the product
	product := &types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	}
	if err := h.store.CreateProduct(*product); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, product)
}
