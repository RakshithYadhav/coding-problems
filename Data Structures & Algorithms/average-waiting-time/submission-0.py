class Solution:
    def averageWaitingTime(self, customers: List[List[int]]) -> float:
        waitTime = 0
        startTime = 0
        totalWaitTime = 0
        for customer in customers:
            arrivalTime, orderTime = customer[0], customer[1]
            if startTime == 0 or arrivalTime > startTime:
                startTime = arrivalTime
            waitTime =  (startTime + orderTime) - arrivalTime
            startTime += orderTime
            print(waitTime)
            totalWaitTime += waitTime

        print(totalWaitTime)
        print(customers)
        return totalWaitTime / len(customers)
        

        