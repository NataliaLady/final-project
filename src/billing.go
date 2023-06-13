package src

import (
	"fmt"
	"math"
	"os"

	"example.com/types/data"
)

func ReadBillingDataFromFile(fileName string) data.BillingData {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return data.BillingData{}
	}

	fileDataStr := string(fileData)

	var j float64
	var res uint8
	for i := len(fileDataStr) - 1; i >= 0; i-- {
		if fileData[i] != 49 {
			continue
		}

		res += uint8(math.Pow(2, j))
		j++
	}

	var billing data.BillingData
	billing.CreateCustomer = fileDataStr[0] == 49
	billing.Purchase = fileDataStr[1] == 49
	billing.Payout = fileDataStr[2] == 49
	billing.Recurring = fileDataStr[3] == 49
	billing.FraudControl = fileDataStr[4] == 49
	billing.CheckoutPage = fileDataStr[5] == 49

	fmt.Println("=================")
	fmt.Println("Состояние системы Billing:")
	fmt.Println("billing в десятичном формате:")
	fmt.Println(res)
	fmt.Println("=================")
	fmt.Println("Состояние системы Billing:")
	fmt.Println("Create Customer:", billing.CreateCustomer)
	fmt.Println("Purchase:", billing.Purchase)
	fmt.Println("Payout:", billing.Payout)
	fmt.Println("Recurring:", billing.Recurring)
	fmt.Println("Fraud Control:", billing.FraudControl)
	fmt.Println("Checkout Page:", billing.CheckoutPage)

	return billing
}
