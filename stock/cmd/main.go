package main

import (
	"fmt"
	"stock/internal/models"
	"stock/internal/services"
)

func main() {
	fmt.Println("Stock system is running...")

	itens := []models.Item{
		{Id: 1, Name: "Laptop", Quantity: 10, Price: 1500.00},
		{Id: 2, Name: "Mouse", Quantity: -50, Price: 25.00},
		{Id: 3, Name: "Keyboard", Quantity: 30, Price: 45.00},
		{Id: 4, Name: "Monitor", Quantity: 20, Price: 300.00},
	}
	stock := services.NewStockService()
	for _, item := range itens {
		err := stock.AddItem(item, "admin")
		if err != nil {
			fmt.Println("Error adding item:", err)
		} else {
			fmt.Println("Item added:", item.Info())
		}
	}

	for _, item := range stock.GetItems() {
		fmt.Printf("Id %d | Product: %s | Quantity: %d | Price: %.2f \n",
			item.Id, item.Name, item.Quantity, item.Price)
	}

	for _, log := range stock.GetLogs() {
		fmt.Printf("[%s] Action: %s | User: %s | Item ID: %d | Quantity: %d | Reason: %s\n",
			log.Timestamp.Local().Format("02-Jan-2006 15:04:05"), log.Action, log.User, log.ItemID, log.Quantity, log.Reason)
	}

	fmt.Printf("Total cost of all items in stock: R$ %.2f\n", stock.CalculateTotalCost())

	itemsSearch, err := services.FindBy(itens, func(item models.Item) bool {
		return item.Name == "Laptop"
	})

	if err != nil {
		fmt.Println("Error getting items by name:", err)
	}

	fmt.Println("Items found: ", itemsSearch)

	alura := services.Supplier{
		CNPJ:    "12.345.678/0001-95",
		Contact: "Alura Suporte",
		City:    "São Paulo",
	}
	fmt.Println("Supplier Info:", alura.GetInfo())

	if alura.CheckAvailability(5, 10) {
		fmt.Println("Supplier has enough stock available.")
	} else {
		fmt.Println("Supplier does not have enough stock available.")
	}

	fmt.Println("Stock system finished running.")
}
