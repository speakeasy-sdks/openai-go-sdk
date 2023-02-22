package shared

type CreateModerationResponseResultsCategories struct {
	Hate            bool `json:"hate"`
	HateThreatening bool `json:"hate/threatening"`
	SelfHarm        bool `json:"self-harm"`
	Sexual          bool `json:"sexual"`
	SexualMinors    bool `json:"sexual/minors"`
	Violence        bool `json:"violence"`
	ViolenceGraphic bool `json:"violence/graphic"`
}

type CreateModerationResponseResultsCategoryScores struct {
	Hate            float64 `json:"hate"`
	HateThreatening float64 `json:"hate/threatening"`
	SelfHarm        float64 `json:"self-harm"`
	Sexual          float64 `json:"sexual"`
	SexualMinors    float64 `json:"sexual/minors"`
	Violence        float64 `json:"violence"`
	ViolenceGraphic float64 `json:"violence/graphic"`
}

type CreateModerationResponseResults struct {
	Categories     CreateModerationResponseResultsCategories     `json:"categories"`
	CategoryScores CreateModerationResponseResultsCategoryScores `json:"category_scores"`
	Flagged        bool                                          `json:"flagged"`
}

type CreateModerationResponse struct {
	ID      string                            `json:"id"`
	Model   string                            `json:"model"`
	Results []CreateModerationResponseResults `json:"results"`
}
