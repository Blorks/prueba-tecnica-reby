package models

type User struct {
	IdUser  int `gorm:"primaryKey"`
	Name    string
	Email   string
	Balance int
}

func (user *User) Constructor(IdUser int, name string, email string, balance int) {
	user.Name = name
	user.Email = email
	user.Balance = balance
}

func (user *User) CheckUserBalance(balance int) bool {
	return user.Balance > balance
}
