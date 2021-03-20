package transaction

import "backendEkost/user"

type GetKostTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
