package main

import (
	"fmt"
	"go-chatbot/domain/repository/impl"
	"testing"
)

func TestAsk(t *testing.T) {
	question := "使用 java 写一个冒泡排序"
	repository := impl.NewOpenAIRepository()
	gtp := repository.DoChatGTP(question)
	fmt.Println(gtp)
}
