package serializer

const (
	OK                  = 200
	BadAuthError        = 401
	NotMatchError       = 403
	NotFoundError       = 404
	BadRequestError     = 402
	InternalServerError = 500
)

// Response 基础序列化器
type Response struct {
	Status uint        `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
