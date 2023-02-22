package shared

type CreateAnswerResponseSelectedDocuments struct {
	Document *int64  `json:"document,omitempty"`
	Text     *string `json:"text,omitempty"`
}

type CreateAnswerResponse struct {
	Answers           []string                                `json:"answers,omitempty"`
	Completion        *string                                 `json:"completion,omitempty"`
	Model             *string                                 `json:"model,omitempty"`
	Object            *string                                 `json:"object,omitempty"`
	SearchModel       *string                                 `json:"search_model,omitempty"`
	SelectedDocuments []CreateAnswerResponseSelectedDocuments `json:"selected_documents,omitempty"`
}
