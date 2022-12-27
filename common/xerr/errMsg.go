package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	// 全局错误
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[ReuqestParamError] = "参数错误"
	message[TokenExpireError] = "token失效，请重新登陆"
	message[TokenGenerateError] = "生成token失败"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[DbUpdateAffectedZeroError] = "更新数据影响行数为0"

	// 用户模块错误
	message[UsernamePwdError] = "用户或密码错误"
	message[UsernameExistError] = "用户已存在,请联系管理员"
	message[UserDbInsertError] = "数据库id递增插入错误,请联系管理员"
	message[UsernameNoExistError] = "用户不存在"

	// 秘钥模块错误
	message[SecretIDExistError] = "秘钥已存在,请联系管理员"
	message[SecretDbInsertError] = "数据库id递增插入错误,请联系管理员"
	message[SecretIDNoExistError] = "秘钥不存在"
	message[SecretUpdateError] = "秘钥更新失败"

	// 任务模块错误
	message[TaskIDExistError] = "任务已存在,请重新输入"
	message[TaskCallbackError] = "任务回调失败"
	message[TaskDbInsertError] = "任务id递增插入错误,请联系管理员"
	message[TaskIDNoExistError] = "任务不存在"
	message[TaskUpdateError] = "任务更新失败"
	message[TaskVendorError] = "未定义的云厂商"
	message[TaskTypeError] = "未定义的类型"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
