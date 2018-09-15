package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	r := maxProfit(prices)

	fmt.Println(r)
}

func maxProfit(prices []int) int {
	profit := 0
	buy := 0
	sell := -1
	for i := 1; i < len(prices); i++ {
		// find the cheapest price
		if prices[i] < prices[buy] {
			buy = i
			continue
		}
		if prices[i]-prices[buy] > profit {
			sell = i
			profit = prices[i] - prices[buy]
		}

	}
	if sell == -1 {
		return 0
	}

	return profit
}
