package types

type SaveAccount struct {
	AccountName string `json:"account_name"`
}

func (t *SaveAccount) Validate() (e []string) {
	if t.AccountName == "" {
		e = append(e, "empty `account_name`")
	}

	return e
}

type AccountTransaction struct {
	FromAccountID string `json:"from_account_id"`
	ToAccountID   string `json:"to_account_id"`

	Title  string  `json:"title"`
	Note   string  `json:"note"`
	Amount float64 `json:"amount"`

	Type string `json:"type"`
}

func (t *AccountTransaction) Validate() (e []string) {
	if t.Title == "" {
		e = append(e, "invalid `title`")
	}
	if t.FromAccountID == "" {
		e = append(e, "invalid `from_account_id`")
	}
	if t.Amount <= 0 {
		e = append(e, "invalid `amount`")
	}

	return e
}
