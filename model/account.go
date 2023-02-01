package model

type Account struct {
	Id            int    `json:"id"`
	Bank          string `json:"bank" validate:"required"`
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	CheckInterval int    `json:"check_interval,omitempty" validate:"required"`
	ActiveTime    int    `json:"active_time,omitempty" validate:"required"`
	AutoLogout    bool   `json:"auto_logout,omitempty" validate:"required"`
}

type CreateAccount struct {
	Bank          string `json:"bank" validate:"required"`
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	CheckInterval int    `json:"check_interval,omitempty" validate:"required"`
	ActiveTime    int    `json:"active_time,omitempty" validate:"required"`
	AutoLogout    bool   `json:"auto_logout,omitempty" validate:"required"`
}

type UpdateAccount struct {
	Id            int    `json:"id"`
	Bank          string `json:"bank"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	AccountNumber string `json:"account_number"`
	CheckInterval int    `json:"check_interval"`
	ActiveTime    int    `json:"active_time"`
	AutoLogout    bool   `json:"auto_logout,omitempty"`
}

type AccountResponse struct {
	Id            int    `json:"id"`
	Bank          string `json:"bank"`
	Username      string `json:"username"`
	AccountNumber string `json:"account_number"`
	CheckInterval int    `json:"check_interval"`
	ActiveTime    int    `json:"active_time"`
	AutoLogout    bool   `json:"auto_logout"`
}
