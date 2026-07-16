class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        # You have a list of prices where price[i] is the price of meta stock
        # [537,564,600,511..]
        # You can choose to buy meta stock on any single day [not multiple days]
        # You can choose to sell it in the future.
        # Find out what is the best combination of buy day and sell day that you will have the highest profit.

        # constraints
        # the max length of prices is just 100 a brute force could easily satisfy.

        # what exactly profit means = buy low and sell high.
        # now if you want the max profit = buy max(low) and sell max(high) or buy the lowest and sell the highest.

        # 1. Brute Force 
        # Go through the prices array. and consider that you buy on every ith day and check the max profit you can get.


        maxProfit = 0
        if len(prices) == 1:
            return maxProfit

        for i in range(len(prices)):
            for j in range(i+1, len(prices)):
                profit = prices[j] - prices[i] # buy - sell
                maxProfit = max(profit, maxProfit)
        return maxProfit

        # Now lets check if this passes edges cases.
        # what are the possible inputs.
        # Prices[], Price itself.
        # Prices[] = 1, N elements.
        # Prices[i] can be 0 so you can have variable elements or duplicate.
        # the code handles variable and duplicate.
        # one edge case if len(price) = 1 which means even if you buy he cant sell which means profit is 0

        