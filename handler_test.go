package main

import (
	"context"
	"reflect"
	"testing"
	"userService/kitex_gen/user"
	"userService/util"
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

func getAuth(userName string) string {
	rdb := util.RedisInit()
	status := rdb.Get(context.Background(), userName)
	if status.Err() != nil {
		return ""
	}
	return status.Val()
}
func TestUserServiceImpl_Login(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user.LoginRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user.LoginResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"user not exist", args{
			ctx: context.Background(),
			req: &user.LoginRequest{
				UserName: "admi",
				UserPwd:  "123456",
			},
		}, &user.LoginResponse{
			Success: false,
			ErrMsg:  "user is not exist",
			Auth:    "",
		}, false},
		{"pwd error", args{
			ctx: context.Background(),
			req: &user.LoginRequest{
				UserName: "admin",
				UserPwd:  "1234",
			},
		}, &user.LoginResponse{
			Success: false,
			ErrMsg:  "password error",
			Auth:    "",
		}, false},
		{"success", args{
			ctx: context.Background(),
			req: &user.LoginRequest{
				UserName: "admin",
				UserPwd:  "123456",
			},
		}, &user.LoginResponse{
			Success: true,
			ErrMsg:  "",
			Auth:    "",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserServiceImpl{}
			gotResp, err := s.Login(tt.args.ctx, tt.args.req)
			tt.wantResp.Auth = getAuth(tt.args.req.GetUserName())
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Login() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
