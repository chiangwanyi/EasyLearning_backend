package serializer

const (
	// OK 正常
	OK = 200
	// BadAuthError 无权限
	BadAuthError = 401
	// NotMatchError 信息匹配错误
	NotMatchError = 403
	// NotFoundError 查找为空
	NotFoundError = 404
	// BadRequestError 请求数据内容或格式错误
	BadRequestError = 402
	// InternalServerError 内部错误
	InternalServerError = 500
)

// Response 基础序列化器
type Response struct {
	Status uint        `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
