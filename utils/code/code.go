package errmsg

type CodeType int

const (
	CodeSuccess CodeType = 10000 + iota
	CodeError
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeUploadFileFailed
	CodeUserNotLogin

	CodeNeedLogin
	CodeInvalidToken
)

var codeAndMsgMap = map[CodeType]string{
	CodeSuccess:          "Success",
	CodeInvalidParam:     "Invalid Param",
	CodeUserExist:        "User is exist",
	CodeUserNotExist:     "User is not exist",
	CodeUserNotLogin:     "User is not login",
	CodeInvalidPassword:  "Invalid username and password",
	CodeServerBusy:       "Server is busy",
	CodeError:            "Error",
	CodeUploadFileFailed: "Upload Filed Failed!",

	CodeNeedLogin:    "Need login",
	CodeInvalidToken: "Token is invalid",
}

func (c *CodeType) Msg() string {
	return codeAndMsgMap[*c]
}
