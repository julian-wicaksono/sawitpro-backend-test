// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	InsertUser(ctx context.Context, input UserData) (output UserId, err error)
	GetUserDataByPhoneNumber(ctx context.Context, input UserData) (userData UserData, userId UserId, err error)
	GetUserByPhoneNumber(ctx context.Context, input UserData) (userId UserId, err error)
	GetUserDataById(ctx context.Context, input UserId) (output UserData, err error)
	UpdatePhoneNumber(ctx context.Context, userData UserData, userId UserId) (err error)
	UpdateFullName(ctx context.Context, userData UserData, userId UserId) (err error)
	UpdateUserData(ctx context.Context, userData UserData, userId UserId) (err error)
	UpdateLoginCount(ctx context.Context, input UserId) (err error)
}
