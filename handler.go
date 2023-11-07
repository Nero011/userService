package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	user "userService/kitex_gen/user"
	"userService/util"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: 并发环境测试
	db := util.MysqlInit()
	//sqlStr := "select user_name from user where	user_name = \"" + req.GetUserName() + "\""
	sqlStr := fmt.Sprintf(`select user_name from user where user_name = "%s"`, req.GetUserName())

	rows, err := db.Query(sqlStr)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	// 检测是否数据库中已存在用户
	exist := false
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		exist = true
		if err != nil {
			return nil, err
		}
	}
	if exist {
		resp = &user.RegisterResponse{
			Success: false,
			ErrMsg:  "user is exist",
		}
		return resp, nil
	}

	// 在数据库中插入新用户的数据
	sqlStr = fmt.Sprintf(`insert into user (user_name, user_pwd) values ("%s", "%s")`, req.GetUserName(), req.GetUserPwd())
	res, err := db.Exec(sqlStr)
	if err != nil {
		return nil, err
	}
	aff, _ := res.RowsAffected()
	println(aff)
	resp = &user.RegisterResponse{
		Success: true,
		ErrMsg:  "",
	}
	return resp, nil
}

// Login implements the UserServiceImpl interface.

func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: 并发环境的测试
	db := util.MysqlInit()
	sqlStr := fmt.Sprintf(`select user_name, user_pwd from user where user_name = "%s"`, req.GetUserName())
	row, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	exist := false
	for row.Next() { // 遍历row, 正常情况下row只有一行
		exist = true
		var name, pwd string
		err = row.Scan(&name, &pwd)
		if err != nil {
			return nil, err
		}
		if pwd != req.GetUserPwd() {
			resp = &user.LoginResponse{
				Success: false,
				ErrMsg:  "password error",
			}
			return resp, nil
		}
	}
	if exist == false {
		resp = &user.LoginResponse{
			Success: false,
			ErrMsg:  "user is not exist",
		}
		return resp, nil
	}

	// 生成响应
	rowString := req.GetUserName() + req.GetUserPwd() + time.Now().String()
	data := []byte(rowString)
	hash := sha256.Sum256(data)
	auth := hex.EncodeToString(hash[:])
	// 把auth写入redis已登陆用户缓存
	RedisDb := util.RedisInit()
	// TODO: 增加密钥过期时间
	sqlStr = fmt.Sprintf("set %s %s", req.GetUserName(), auth) // 暂时无过期时间
	status := RedisDb.Set(context.Background(), req.GetUserName(), auth, -1)

	if status.Err() != nil {
		resp = &user.LoginResponse{
			Success: false,
			ErrMsg:  "write redis err",
			Auth:    "",
		}
		return nil, err
	}
	resp = &user.LoginResponse{
		Success: true,
		ErrMsg:  "",
		Auth:    auth,
	}
	return
}
