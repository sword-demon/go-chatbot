package main

import (
	"fmt"
	"go-chatbot/domain/repository/impl"
	"go-chatbot/infrastructure"
	"log"
	"testing"
)

func TestAsk(t *testing.T) {
	infrastructure.InitConfig()
	question := "孙悟空的师傅是谁"
	key := infrastructure.AppConfig.OpenaiKey
	repository := impl.NewOpenAIRepository()
	gtp, err := repository.DoChatGTP(key, question)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	fmt.Println(gtp)
}

func TestGetAppConfig(t *testing.T) {
	infrastructure.InitConfig()
	fmt.Println(infrastructure.AppConfig.OpenaiKey)
}