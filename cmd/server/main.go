package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/config"
	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/repository"
	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World!")

	db := config.ConnectDB()
	defer db.Close()

	corsConfig := config.LoadCorsConfig()

	vendorRepo := repository.NewVendorRepository(db)
	salesRepo := repository.NewSalesRepository(db)
	itemRepo := repository.NewItemRepository(db)
	vendorHandler := service.NewVendorHandler(vendorRepo)
	salesHandler := service.NewSalesHandler(salesRepo)
	itemHandler := service.NewItemHandler(itemRepo)

	router := mux.NewRouter()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins(corsConfig.AllowedOrigins),
		handlers.AllowedMethods(corsConfig.AllowedMethods),
		handlers.AllowedHeaders(corsConfig.AllowedHeaders),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/vendors", vendorHandler.GetVendors).Methods("GET")
	router.HandleFunc("/vendors", vendorHandler.CreateVendor).Methods("POST")
	router.HandleFunc("/allSales", salesHandler.GetAllSales).Methods("GET")
	router.HandleFunc("/itemsOfSale", salesHandler.GetAllItemsInSale).Methods("GET")
	router.HandleFunc("/addItemToSale", salesHandler.AddItemToSale).Methods("POST")
	router.HandleFunc("/createSale", salesHandler.CreateSale).Methods("POST")
	router.HandleFunc("/getAllItems", itemHandler.GetAllItems).Methods("GET")
	router.HandleFunc("/addItem", itemHandler.AddNewItem).Methods("POST")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}
