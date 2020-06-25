package handlers

import (
	"fmt"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// DefaultHandler -- Handle
type DefaultHandler struct{}

// Category -- Model
type Category struct {
	Name             string `json:"name"`
	TotalRestaurants int    `json:"totalRestaurants"`
}

// Restaurant -- Model
type Restaurant struct {
	Name     string   `json:"name"`
	Category Category `json:"category"`
}

// Index @ DefaultHandler
func (h *DefaultHandler) Index(c *routing.Context) error {
	fmt.Println("Hit 'DefaultHandler@index'!")

	category := Category{"Burgers", 10}
	restaurant := Restaurant{"Mc Donalds", category}
	c.Write(restaurant)
	return nil
}
