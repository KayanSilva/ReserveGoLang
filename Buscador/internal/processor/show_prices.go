package processor

import (
	"buscador/internal/models"
	"fmt"
)

func ShowPricesAndAVG(priceChannel <-chan models.Price, done chan<- bool) {
	var totalPrices float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrices += price.Value
		countPrices++
		avgPrice := totalPrices / countPrices
		fmt.Printf("[%s] Preço recebido de %s |R$ %.2f| Preço médi até agora %.2f\n",
			price.Timestamp.Format("02-Jan-2006 15:04:05 "),
			price.StoreName,
			price.Value,
			avgPrice)
	}

	done <- true
}
