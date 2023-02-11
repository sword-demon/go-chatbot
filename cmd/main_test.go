package main

import (
	"fmt"
	"go-chatbot/common"
	"go-chatbot/domain/repository/impl"
	"go-chatbot/infrastructure"
	"log"
	"testing"
)

func TestAsk(t *testing.T) {
	infrastructure.InitConfig()
	question := "孙悟空的师傅是谁"
	repository := impl.NewOpenAIRepository()
	gtp, err := repository.DoChatGTP(question)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	fmt.Println(gtp)
}

func TestGetAppConfig(t *testing.T) {
	infrastructure.InitConfig()
	fmt.Println(common.AppConfig.OpenaiKey)
}