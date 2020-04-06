package payment

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/viniciuswebdev/golang-unit-tests/entity"
)

type Gateway interface {
	IsAuthorized(user entity.User, creditCard entity.CreditCard) (bool, error)
	Pay(creditCard entity.CreditCard, amount int) error
}

type MyPayment struct{}

func NewMyPayment() *MyPayment {
	return &MyPayment{}
}

func (c *MyPayment) IsAuthorized(user entity.User, creditCard entity.CreditCard) (bool, error) {

	message := map[string]interface{}{
		"credit_card": creditCard.Number,
		"expiration":  creditCard.Expiration,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(
		"https://mypayment.just.example/authorization",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return false, err
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	if result["result"] == "authorized" {
		return true, nil
	}

	return false, nil
}

func (c *MyPayment) Pay(creditCard entity.CreditCard, amount int) error {

	message := map[string]interface{}{
		"credit_card": creditCard.Number,
		"expiration":  creditCard.Expiration,
		"amount":      strconv.Itoa(amount),
	}

	data, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = http.Post(
		"https://mypayment.just.example/payment",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	return nil
}
