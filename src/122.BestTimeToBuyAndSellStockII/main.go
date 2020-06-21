package main

import "fmt"

func main() {
	//prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{1, 2, 3, 4, 5}
	profit := maxProfit(prices)
	fmt.Print("profit:", profit)
}

func maxProfit(prices []int) int {
	buy := 0
	profits := 0
	for i := 1; i < len(prices); i++ {
		if prices[buy] > prices[i] {
			buy = i
			continue
		}
		// tomorrow's price is lower
		if i < len(prices)-1 && prices[i] <= prices[i+1] {
			continue
		}
		profits += prices[i] - prices[buy]
		buy = i
	}
	return profits
}
