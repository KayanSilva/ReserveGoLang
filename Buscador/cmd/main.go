package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/models"
	"buscador/internal/processor"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	priceChannel := make(chan models.Price, 4)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPricesAndAVG(priceChannel, done)

	<-done

	fmt.Printf("Total time: %s", time.Since(startTime).String())
}
