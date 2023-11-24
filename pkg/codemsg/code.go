// package codemsg contain project self-define code and msgs
package codemsg

const (
	StatusOK = iota
	SelfDeinfeStatsu1
	SelfDeinfeStatsu2
)

var msgs = map[int]string{
	StatusOK:          "ok",
	SelfDeinfeStatsu1: "self deinfe statsu 1",
	SelfDeinfeStatsu2: "self deinfe statsu 2",
}

func Msg(code int) string {
	return msgs[code]
}
