package controller

import (
	"bank/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	user := models.Account{Owner: "Samundra", Balance: 0, Currency: "USD"}

	result := DB.Create(&user)

	require.NoError(t, result.Error)
	require.NotEmpty(t, user)

}
