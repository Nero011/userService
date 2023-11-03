package main

import (
	"context"
	"fmt"
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
