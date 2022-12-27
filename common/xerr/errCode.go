package xerr

//成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const (
	// 服务器一般错误（默认错误）
	ServerCommonError uint32 = 100001
	// 请求参数错误
	ReuqestParamError uint32 = 100002
	// token 过期错误
	TokenExpireError uint32 = 100003
	// token 生成错误
	TokenGenerateError uint32 = 100004
	// 数据库错误
	DbError uint32 = 100005
	// 数据库更新影响行数为0错误
	DbUpdateAffectedZeroError uint32 = 100006
)

// 用户模块错误码
const (
	// 用户或密码错误
	UsernamePwdError uint32 = 200001
	// 用户已存在,请联系管理员
	UsernameExistError uint32 = 200002
	// 数据库id递增插入错误,请联系管理员
	UserDbInsertError uint32 = 200003
	// 用户不存在
	UsernameNoExistError uint32 = 200004
)

// 秘钥模块错误码
const (
	// 秘钥已存在,请联系管理员
	SecretIDExistError uint32 = 300001
	// 数据库id递增插入错误,请联系管理员
	SecretDbInsertError uint32 = 300002
	// 秘钥不存在
	SecretIDNoExistError uint32 = 300003
	// 秘钥更新失败
	SecretUpdateError uint32 = 300004
)

// 任务模块错误码
const (
	// 任务已存在,请重新输入
	TaskIDExistError uint32 = 400001
	// 任务回调失败
	TaskCallbackError uint32 = 400002
	// 任务id递增插入错误,请联系管理员
	TaskDbInsertError uint32 = 400003
	// 任务不存在
	TaskIDNoExistError uint32 = 400004
	// 任务更新失败
	TaskUpdateError uint32 = 400005
	TaskVendorError uint32 = 400006
	TaskTypeError   uint32 = 400007
)
