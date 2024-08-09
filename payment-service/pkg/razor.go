package pkg

import (
	"errors"
	"fmt"
	"os"

	"github.com/razorpay/razorpay-go"
)

func Executerazorpay(order string, amount int) (string, error) {

	client := razorpay.NewClient(os.Getenv("RAZOR_KEY"), os.Getenv("RAZOR_SECRET"))
	fmt.Println("")
	fmt.Println(order)

	data := map[string]interface{}{
		"amount":   amount * 100,
		"currency": "INR",
		"receipt":  order,
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return "", errors.New("payment not initiated")
	}
	razorId, _ := body["id"].(string)
	return razorId, nil
}
