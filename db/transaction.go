package db

import (
	"bank/models"

	"gorm.io/gorm"
)

func TransferMoney(formaccountid int64, toaccountid int64, amount int64) {
	type Account struct {
		Id         int64
		Owner      string
		Balance    int64
		Currency   string
		Created_at string
	}
	var user Account
	var user2 Account
	DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.First(&user, formaccountid).Error; err != nil {
			// return any error will rollback
			return err
		}

		if err := tx.First(&user2, toaccountid).Error; err != nil {
			// return any error will rollback
			return err
		}

		user.Balance = user.Balance - amount
		user2.Balance = user2.Balance + amount
		if err := tx.Model(&user).Updates(&models.Account{Balance: user.Balance}).Error; err != nil {
			return err
		}

		if err := tx.Model(&user2).Updates(&models.Account{Balance: user2.Balance}).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})
}

func UpdateAccount(accountid int64, amount int64) {
	var inp models.Account

	DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.First(&inp, accountid).Error; err != nil {
			// return any error will rollback
			return err
		}
		inp.Balance = inp.Balance + amount

		if err := tx.Model(&inp).Updates(&models.Account{Balance: inp.Balance}).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

}
