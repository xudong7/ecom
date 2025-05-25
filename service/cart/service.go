package cart

import (
	"fmt"

	"github.com/xudong7/ecom/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product %d", item.ProductID)
		}
		productIds[i] = item.ProductID
	}
	return productIds, nil
}

func (h *Handler) createOrder(
	products []types.Product,
	items []types.CartItem,
	userID int,
) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	// check if products are available
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, nil
	}

	// calculate total price
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce quantity of products in db
	for _, item := range items {
		product := productMap[item.ProductID]
		// TODO: high concurrency problem
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}

	// create order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "123 Main St", // TODO: get address from request
	})
	if err != nil {
		return 0, 0, err
	}

	// create order items
	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}

func calculateTotalPrice(
	cartItems []types.CartItem,
	products map[int]types.Product,
) float64 {
	var totalPrice float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		totalPrice += float64(item.Quantity) * product.Price
	}

	return totalPrice
}

func checkIfCartIsInStock(
	cartItems []types.CartItem,
	products map[int]types.Product,
) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d not found", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("insufficient stock for product %d", item.ProductID)
		}
	}
	return nil
}
