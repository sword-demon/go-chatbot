package main

import (
	"flag"
	"fmt"
	"go-chatbot/domain/repository/impl"
	"go-chatbot/infrastructure"
	"log"
)

func main() {
	infrastructure.InitConfig()
	key := infrastructure.AppConfig.OpenaiKey
	question := flag.String("question", "现在是几几年", "a question for ai to answer")
	flag.Parse()
	log.Printf("question=%s", *question)

	service := impl.NewOpenAIRepository()
	answer, err := service.DoChatGTP(key, *question)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(answer)
}
