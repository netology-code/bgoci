package main

import (
	"fmt"
	"github.com/netology-code/bgoci/pkg/card"
)

func main() {
	svc := card.NewService("Netology Bank")
	svcVal := *svc

	visa := (*svc).IssueCard("Visa", "RUB")
	master := svcVal.IssueCard("MasterCard", "USD")

	fmt.Println(visa.Name)
	fmt.Println(visa.Owner.Name)

	fmt.Println(visa)
	fmt.Println(master)
}
