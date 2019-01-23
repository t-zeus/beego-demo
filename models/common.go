package models

// 定义 json 返回结构体
type CodeInfo struct {
	Code int    `json:"code"` // 对应 json 的 code
	Msg  string `json:"msg"`  // 对应 json 的 msg
}

func NewErrorInfo(info string) *CodeInfo {
	return &CodeInfo{-1, info}
}

func NewNormalInfo(info string) *CodeInfo {
	return &CodeInfo{0, info}
}
