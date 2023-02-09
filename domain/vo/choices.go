package vo

// Choices 选择模型
type Choices struct {
	Text string `json:"text"`
	Index int `json:"index"`
	Logprobs string `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}