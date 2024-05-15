package main

type User struct {
	login string
	password string
	isAdmin bool
}

func (u User) getLogin() string{
	return u.login
}

func (u User) getPassword() string{
	return u.password
}

func (u User) getAdmin() bool {
	return u.isAdmin
}

func (u * User) setLogin(l string) {
	u.login = l
} 

func (u * User) setPassword(pass string) {
	u.password = pass
}

func (u * User) setAdminRoot(root bool) {
	u.isAdmin = root
}

