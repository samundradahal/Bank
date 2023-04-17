package models

import "time"

type Account struct {
	Id        int64
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt time.Time
}

type Entries struct {
	Id        int64
	AccountId int64
	Amount    int64
	CreatedAt time.Time
}

type Transfers struct {
	From_Accountid int64
	To_Accountid   int64
	Amount         int64
	CreatedAt      time.Time
}

type Users struct {
	Username          string
	HashedPassword    string
	Email             string
	FullName          string
	PasswordChangedAt time.Time
	CreatedAt         time.Time
}
   