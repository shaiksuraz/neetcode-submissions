func maxProfit(prices []int) int {
	var profit int
	var buyingPrice = 101
	for i:=0; i< len(prices); i++ {
		buyingPrice = min(buyingPrice, prices[i])
		sellingPrice := prices[i]
		profit = max(profit,sellingPrice - buyingPrice)
	}
	return profit
}
