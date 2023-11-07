# ChatCompletionRequestMessage


## Supported Types

### SystemMessage

```go
chatCompletionRequestMessage := shared.CreateChatCompletionRequestMessageSystemMessage(shared.SystemMessage{/* values here */})
```

### UserMessage

```go
chatCompletionRequestMessage := shared.CreateChatCompletionRequestMessageUserMessage(shared.UserMessage{/* values here */})
```

### AssistantMessage

```go
chatCompletionRequestMessage := shared.CreateChatCompletionRequestMessageAssistantMessage(shared.AssistantMessage{/* values here */})
```

### ToolMessage

```go
chatCompletionRequestMessage := shared.CreateChatCompletionRequestMessageToolMessage(shared.ToolMessage{/* values here */})
```

### FunctionMessage

```go
chatCompletionRequestMessage := shared.CreateChatCompletionRequestMessageFunctionMessage(shared.FunctionMessage{/* values here */})
```

