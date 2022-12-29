package routes

import (
	"github.com/gorilla/mux"
	cart "github.com/sanjana-jadhav-searce/shopping-catlog/pkg/controllers/Carts"
	category "github.com/sanjana-jadhav-searce/shopping-catlog/pkg/controllers/Category"
	inventory "github.com/sanjana-jadhav-searce/shopping-catlog/pkg/controllers/Inventory"
	products "github.com/sanjana-jadhav-searce/shopping-catlog/pkg/controllers/Products"
)

var RegisterShoppingCart = func(router *mux.Router) {
	router.HandleFunc("/create", products.CreateProduct).Methods("POST")
	router.HandleFunc("/products", products.GetProducts).Methods("GET")
	router.HandleFunc("/product/{name}", products.GetProductByName).Methods("GET")
	router.HandleFunc("/updateproduct", products.UpdateProduct).Methods("PUT")
	router.HandleFunc("/deleteproduct", products.DeleteProduct).Methods("DELETE")

	router.HandleFunc("/createinventory", inventory.CreateInventoryProduct).Methods("POST")
	router.HandleFunc("/productinventory/{name}", inventory.GetProductByNameInInventory).Methods("GET")
	router.HandleFunc("/productsinventory", inventory.GetProductsInventory).Methods("GET")
	router.HandleFunc("/updateproductinventory", inventory.UpdateProductInventory).Methods("PUT")
	router.HandleFunc("/deleteproductinventory", inventory.DeleteProductInventory).Methods("DELETE")

	router.HandleFunc("/createcategory", category.CreateCategory).Methods("POST")
	router.HandleFunc("/namecategory/{name}", category.GetCategoryByName).Methods("GET")
	router.HandleFunc("/categories", category.GetCategories).Methods("GET")
	router.HandleFunc("/deletecategory", category.DeleteCategory).Methods("DELETE")

	router.HandleFunc("/createcart", cart.CreateCart).Methods("POST")
	router.HandleFunc("/additemstocart", cart.AddItemsToCart).Methods("POST")
	router.HandleFunc("/additemtocart", cart.AddItemToCart).Methods("POST")
	router.HandleFunc("/getcart", cart.GetCart).Methods("GET")
	router.HandleFunc("/deletecart", cart.DeleteCart).Methods("DELETE")

}
