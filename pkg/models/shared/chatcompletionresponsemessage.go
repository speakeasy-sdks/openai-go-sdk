package shared

type ChatCompletionResponseMessageRoleEnum string

const (
	ChatCompletionResponseMessageRoleEnumSystem    ChatCompletionResponseMessageRoleEnum = "system"
	ChatCompletionResponseMessageRoleEnumUser      ChatCompletionResponseMessageRoleEnum = "user"
	ChatCompletionResponseMessageRoleEnumAssistant ChatCompletionResponseMessageRoleEnum = "assistant"
)

type ChatCompletionResponseMessage struct {
	Content string                                `json:"content"`
	Role    ChatCompletionResponseMessageRoleEnum `json:"role"`
}
