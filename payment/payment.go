package payment

import (
	"github.com/viniciuswebdev/golang-unit-tests/database"
	"github.com/viniciuswebdev/golang-unit-tests/entity"
)

type PaymentService struct {
	attemptHistoryRepository database.AttemptHistory
	gateway                  Gateway
}

func NewPaymentService(
	attemptHistoryRepository database.AttemptHistory,
	gateway Gateway,
) *PaymentService {
	return &PaymentService{
		attemptHistoryRepository: attemptHistoryRepository,
		gateway:                  gateway,
	}
}

func (c *PaymentService) IsAuthorized(user entity.User, creditCard entity.CreditCard) (bool, error) {
	totalAttempts, err := c.attemptHistoryRepository.CountFailures(user)
	if err != nil {
		return false, err
	}

	if totalAttempts > 5 {
		return false, nil
	}

	isAuthorized, err := c.gateway.IsAuthorized(user, creditCard)
	if err != nil {
		return false, err
	}

	if !isAuthorized {
		err := c.attemptHistoryRepository.IncrementFailure(user)
		if err != nil {
			return false, err
		}
	}

	return isAuthorized, nil
}
