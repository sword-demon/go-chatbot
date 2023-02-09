package repository

type IOpenAI interface {
	DoChatGTP(question string) string
}