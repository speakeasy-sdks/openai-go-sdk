package shared

type CreateClassificationResponseSelectedExamples struct {
	Document *int64  `json:"document,omitempty"`
	Label    *string `json:"label,omitempty"`
	Text     *string `json:"text,omitempty"`
}

type CreateClassificationResponse struct {
	Completion       *string                                        `json:"completion,omitempty"`
	Label            *string                                        `json:"label,omitempty"`
	Model            *string                                        `json:"model,omitempty"`
	Object           *string                                        `json:"object,omitempty"`
	SearchModel      *string                                        `json:"search_model,omitempty"`
	SelectedExamples []CreateClassificationResponseSelectedExamples `json:"selected_examples,omitempty"`
}
