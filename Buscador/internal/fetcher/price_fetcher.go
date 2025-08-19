package fetcher

import (
	"buscador/internal/models"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func FetchPrices(priceChannel chan<- models.Price) {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		priceChannel <- FetchriceFromSite1("AAPL")
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchriceFromSite2("SSNHZ")
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchriceFromSite3("GOOG")
	}()

	go func() {
		defer wg.Done()
		FetchAndSendMultiplePrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

func FetchriceFromSite1(symbol string) models.Price {
	time.Sleep(1 * time.Second)
	fmt.Printf("Fetching price for %s from Site 1...\n", symbol)
	return models.Price{
		StoreName: "Site 1",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchriceFromSite2(symbol string) models.Price {
	time.Sleep(3 * time.Second)
	fmt.Printf("Fetching price for %s from Site 2...\n", symbol)
	return models.Price{
		StoreName: "Site 2",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchriceFromSite3(symbol string) models.Price {
	time.Sleep(2 * time.Second)
	fmt.Printf("Fetching price for %s from Site 3...\n", symbol)
	return models.Price{
		StoreName: "Site 3",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchAndSendMultiplePrices(priceChannel chan<- models.Price) {
	prices := []models.Price{
		{StoreName: "Site 4", Value: rand.Float64() * 100, Timestamp: time.Now()},
		{StoreName: "Site 5", Value: rand.Float64() * 100, Timestamp: time.Now()},
		{StoreName: "Site 6", Value: rand.Float64() * 100, Timestamp: time.Now()},
	}

	for _, price := range prices {
		priceChannel <- price
	}
}
