package response

const (
	ErrCodeSuccess  = 20001
	ErrInvalidToken = 30001
)

// message

var msg = map[int]string{
	ErrCodeSuccess:  "success",
	ErrInvalidToken: "Token invalid",
}
