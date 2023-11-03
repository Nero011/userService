package main

import (
	"context"
	"reflect"
	"testing"
	"userService/kitex_gen/user"
)

func TestUserServiceImpl_Register(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user.RegisterRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user.RegisterResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			"user is exist",
			args{
				ctx: context.Background(),
				req: &user.RegisterRequest{
					UserName: "admin",
					UserPwd:  "111",
				},
			},
			&user.RegisterResponse{
				Success: false,
				ErrMsg:  "user is exist",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserServiceImpl{}
			gotResp, err := s.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Register() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
