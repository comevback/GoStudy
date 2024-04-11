package main

import (
	"fmt"
)

func profitCal() {
	var taxRate float64
	var revenue float64
	var expenses float64
	var EBT float64
	var profit float64
	var ratio float64

	getUserInput2("Enter the revenue: ", &revenue)

	fmt.Println("Profit Calculator:")
	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue) //如果你传递 revenue 而不是 &revenue 给 fmt.Scan，函数将接收 revenue 的一个副本，任何在函数内对该副本的修改都不会影响原始的 revenue 变量。这就是为什么你需要传递 revenue 的地址（即 &revenue），这样 fmt.Scan 就可以直接修改 revenue 变量的值了。
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	EBT = revenue - expenses
	profit = calProfit(EBT, taxRate)
	ratio = (profit / revenue)

	fmt.Println("The profit before tax is: ", EBT)
	fmt.Println("The profit after tax is: ", profit)
	fmt.Println("The profit ratio is: ", ratio)
}

func calProfit(EBT float64, taxRate float64) float64 {
	var profit float64 = EBT * (1 - (taxRate / 100))
	return profit
}

func getUserInput2(noti string, input *float64) { // *float64 is a pointer to a float64
	fmt.Print(noti)
	fmt.Scan(input)
}
