package main

import (
	"context"
	userservice "userService/kitex_gen/userService"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *userservice.RegisterRequest) (resp *userservice.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}
