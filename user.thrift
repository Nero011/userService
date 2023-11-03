namespace go userService

struct RegisterRequest{
    1: string userName
    2: string userPwd
}

struct RegisterResponse{
    1: bool success
    2: string errMsg
}

service UserService{
    RegisterResponse Register(1:RegisterRequest req)
}