package service

import (
	"errors"
	"strings"

	data "github.com/lunaorg/luna-main-api/data/account"
	logger "github.com/lunaorg/luna-main-api/libs/log"
	"github.com/lunaorg/luna-main-api/types/viewmodel"
)

type AccountService struct {
	accountData *data.AccountData
}

func NewAccountService() (*AccountService, error) {
	accountData, err := data.NewAccountData()
	if err != nil {
		return nil, err
	}

	return &AccountService{
		accountData: accountData,
	}, nil
}

func (s *AccountService) Save(account types.SaveAccount) error {
	errs := account.Validate()
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	err := s.accountData.Save(account)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *AccountService) NewIncome(income types.AccountTransaction) error {
	errs := income.Validate()
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	err := s.accountData.NewIncome(income)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *AccountService) NewExpense(expense types.AccountTransaction) error {
	errs := expense.Validate()
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	err := s.accountData.NewExpense(expense)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
