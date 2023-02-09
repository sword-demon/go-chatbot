package impl

import (
	"encoding/json"
	"fmt"
	"go-chatbot/common"
	"go-chatbot/domain/aggregates"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type OpenAIRepository struct {
}

const (
	openAiKey = ""
)

func (o *OpenAIRepository) DoChatGTP(question string) string {
	client := &http.Client{}
	body := fmt.Sprintf(`{"model": "text-davinci-003", "prompt": "%s", "temperature": 0, "max_tokens": 1024}`, question)
	request, _ := http.NewRequest("POST", "https://api.openai.com/v1/completions", strings.NewReader(body))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+openAiKey)
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("post data error:%v\\n", err)
		return "请求失败"
	}
	defer resp.Body.Close()
	jsonStr, err := ParseResponse2String(resp)
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}
	var answer aggregates.Answer
	err = json.Unmarshal([]byte(jsonStr), &answer)
	if err != nil {
		return err.Error()
	}
	answers := common.NewStringBuilder("")
	for _, choice := range answer.Choices {
		answers.Append(choice.Text)
	}
	return answers.ToString()
}

func ParseResponse2String(response *http.Response) (string, error) {
	body, err := ioutil.ReadAll(response.Body)
	return string(body), err
}

func NewOpenAIRepository() *OpenAIRepository {
	return &OpenAIRepository{}
}
