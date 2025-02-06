package main

func maxProfit(prices []int) int {
	maxProfit := 0
	windowStart := 0

	for windowEnd := 1; windowEnd < len(prices); windowEnd++ { // start evaluation on second day because buy/sell same day = 0
		if prices[windowEnd] <= prices[windowStart] { // if descending move windowStart right
			windowStart = windowEnd
		} else {
			if maxProfit < prices[windowEnd]-prices[windowStart] { // if ascending move windowEnd right
				maxProfit = prices[windowEnd] - prices[windowStart]
			}
		}
	}

	return maxProfit
}

func main() {
	println("121. Best Time to Buy and Sell Stock")

	// Input: prices = [10,1,5,6,7,1]
	// Output: 6
	println(maxProfit([]int{10, 1, 5, 6, 7, 1}))

	// Input: prices = [10,8,7,5,2]
	// Output: 0
	println(maxProfit([]int{10, 8, 7, 5, 2}))

	// 1,2,4,2,5,7,2,4,9,0,9
	// 9
	println(maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9}))

}
