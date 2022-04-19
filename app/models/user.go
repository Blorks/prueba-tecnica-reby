package models

type User struct {
	idUser  int64
	name    string
	email   string
	balance int64
}

func (user *User) Constructor(name string, email string, balance int64) {
	user.name = name
	user.email = email
	user.balance = balance
}

func (user *User) GetIdUser() int64 {
	return user.idUser
}

func (user *User) SetIdUser(idUser int64) {
	user.idUser = idUser
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) GetEmail() string {
	return user.email
}

func (user *User) SetEmail(email string) {
	user.email = email
}

func (user *User) GetBalance() int64 {
	return user.idUser
}

func (user *User) SetBalance(balance int64) {
	user.balance = balance
}

func (user *User) CheckUserBalance(balance int64) bool {
	return user.balance > balance
}
