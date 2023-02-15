package main

import (
	"fmt"
	"os"
)

var balance int = 0
var flag []byte
var inventory = map[string]int{"apple": 0, "grape": 0, "banana": 0, "flag": 0}
var prices = map[string]int{"apple": 1, "grape": 2, "banana": 50, "flag": 1000000000000000000}

func mine_coins() {
	fmt.Println("How much do you want?")
	fmt.Print("> ")
	var coin int
	fmt.Scanf("%v", &coin)
	if coin > 100 {
		fmt.Println("you cannot mine more than 100 coins at a time!")
	} else {
		// * todo * implement sleep logic to prevent bruteforce
		//time.Sleep(1000 * time.Nanosecond)
		balance += coin
		fmt.Printf("%v coins added to balance.\n", coin)
	}
}

func list_items() {
	for k, v := range prices {
		fmt.Printf("%v: %v coin(s)\n", k, v)
	}
}

func buy_items() {
	fmt.Println("which item do you want to buy?")
	fmt.Print("> ")
	var s string
	fmt.Scanf("%v", &s)
	price, exists := prices[s]

	if !exists {
		fmt.Println("item doesn't exists.")
	} else if price > balance {
		fmt.Println("insufficient funds to buy that item.")
	} else {
		balance -= price
		inventory[s]++
		fmt.Println("item purchased successfully.")
	}
}

func view_status() {
	fmt.Println("--------------")
	fmt.Printf("your balance: %v\n", balance)
	fmt.Println("your inventory: ")
	for k, v := range inventory {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println("--------------")
}

func view_flag() {
	if inventory["flag"] > 0 {
		fmt.Println(string(flag))
	} else {
		fmt.Println("You have to buy flag first.")
	}
}

func menu() {
	fmt.Println("select a number: ")
	fmt.Println("--------------")
	fmt.Println("1: mine coins")
	fmt.Println("2: list items")
	fmt.Println("3: buy items")
	fmt.Println("4: view status")
	fmt.Println("5: view flag")
	fmt.Println("6: exit")
	fmt.Println("--------------")
	fmt.Print("> ")
}

func init() {
	dat, err := os.ReadFile("flag")

	if err != nil {
		panic(err)
	}

	flag = dat
}

func main() {
	fmt.Println("Hello Welcome to the hc-shop :)")
	for {
		menu()
		var a int
		fmt.Scanf("%v", &a)
		switch a {
		case 1:
			mine_coins()
		case 2:
			list_items()
		case 3:
			buy_items()
		case 4:
			view_status()
		case 5:
			view_flag()
		case 6:
			fmt.Println("Thank you for visiting :)")
			os.Exit(0)
		}
	}
}
