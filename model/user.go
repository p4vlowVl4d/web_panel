package model

type UserInterface interface {
	getIdentifier() int
	getRole() RoleInterface
}

type RoleInterface interface {
	getName() string
}

type User struct {
	Id int
	Name, Email, Phone, Password  string
	Role  UserRole
}

func(u *User) getIdentifier() int {
	return u.Id
}

func (u *User) getRole() RoleInterface {
	return u.Role
}

type UserRole struct {
	Id 	  int
	Name  string
	Mode  int
}

func(r UserRole) getName() string {
	return r.Name
}