package moonshot

import "context"

type IUsers interface {
	Balance(ctx context.Context) (*UsersBalanceResponse, error)
}

type users struct {
	client *Client
}

// Users returns a new users controller
func (c *Client) Users() IUsers {
	return &users{client: c}
}

type UsersBalanceResponse struct {
	Code   int                       `json:"code"`
	Data   *UsersBalanceResponseData `json:"data"`
	Scode  string                    `json:"scode"`
	Status bool                      `json:"status"`
}
type UsersBalanceResponseData struct {
	// AvailableBalance including cash balance and voucher balance. When it is less than or equal to 0, the user cannot call the completions API
	AvailableBalance float64 `json:"available_balance"`
	// VoucherBalance will not be negative
	VoucherBalance float64 `json:"voucher_balance"`
	// CashBalance may be negative, which means that the user owes the cost. When it is negative, the AvailableBalance can be the amount of VoucherBalance
	CashBalance float64 `json:"cash_balance"`
}

// Balance returns the user's balance
func (u *users) Balance(ctx context.Context) (*UsersBalanceResponse, error) {
	const path = "/v1/users/me/balance"
	resp, err := u.client.HTTPClient().SetPath(path).Get(ctx)
	if err != nil {
		return nil, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	userBalance := new(UsersBalanceResponse)
	err = resp.Unmarshal(userBalance)
	if err != nil {
		return nil, err
	}
	return userBalance, nil
}
