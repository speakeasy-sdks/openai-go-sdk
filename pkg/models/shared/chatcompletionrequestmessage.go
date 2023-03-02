package shared

type ChatCompletionRequestMessageRoleEnum string

const (
	ChatCompletionRequestMessageRoleEnumSystem    ChatCompletionRequestMessageRoleEnum = "system"
	ChatCompletionRequestMessageRoleEnumUser      ChatCompletionRequestMessageRoleEnum = "user"
	ChatCompletionRequestMessageRoleEnumAssistant ChatCompletionRequestMessageRoleEnum = "assistant"
)

type ChatCompletionRequestMessage struct {
	Content string                               `json:"content"`
	Name    *string                              `json:"name,omitempty"`
	Role    ChatCompletionRequestMessageRoleEnum `json:"role"`
}
