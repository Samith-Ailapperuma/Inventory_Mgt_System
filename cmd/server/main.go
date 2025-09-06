package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/config"
	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/repository"
	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World!")

	db := config.ConnectDB()
	defer db.Close()

	vendorRepo := repository.NewVendorRepository(db)
	salesRepo := repository.NewSalesRepository(db)
	vendorHandler := service.NewVendorHandler(vendorRepo)
	salesHandler := service.NewSalesHandler(salesRepo)

	router := mux.NewRouter()
	router.HandleFunc("/vendors", vendorHandler.GetVendors).Methods("GET")
	router.HandleFunc("/vendors", vendorHandler.CreateVendor).Methods("POST")
	router.HandleFunc("/allSales", salesHandler.GetAllSales).Methods("GET")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
