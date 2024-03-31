// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type UserData struct {
	FullName    string
	PhoneNumber string
	Password    string
}

type UserId struct {
	UserID int64
}
