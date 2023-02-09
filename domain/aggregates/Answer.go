package aggregates

import "go-chatbot/domain/vo"

type Answer struct {
	Id string
	Object string
	Created int
	Model string
	Choices []vo.Choices
}
