package main

type Account struct {
	Username string
	Password string
}

func (a *Account) Save() error {
	return nil
}
func (a *Account) Update() error {
	return nil
}
func (a *Account) Remove() error {
	return nil
}
//func LoadAccount(email string) (*Account, error) {
//	return &Account{Email: email}, nil
//}
