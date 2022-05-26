package api_model

import (
	"encoding/json"
	"log"
)

type APIResponse[T any] struct {
	Result  *T     `json:"result"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (p APIResponse[T]) ToByteList() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		log.Panic(err)
	}
	return bytes
}

func (p APIResponse[T]) ToMap() map[string]any {
	res := make(map[string]any)
	byteList := p.ToByteList()
	err := json.Unmarshal(byteList, &res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func NewSuccessAPIRes[T any](data *T) APIResponse[T] {
	return APIResponse[T]{Result: data, Code: 0, Message: "Result available"}
}

func NewErrorAPIRes(msg string, code ErrorCode) APIResponse[any] {
	return APIResponse[any]{Result: nil, Message: msg, Code: int(code)}
}
