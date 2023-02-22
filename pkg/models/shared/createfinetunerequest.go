package shared

type CreateFineTuneRequest struct {
	BatchSize                    *int64    `json:"batch_size,omitempty"`
	ClassificationBetas          []float64 `json:"classification_betas,omitempty"`
	ClassificationNClasses       *int64    `json:"classification_n_classes,omitempty"`
	ClassificationPositiveClass  *string   `json:"classification_positive_class,omitempty"`
	ComputeClassificationMetrics *bool     `json:"compute_classification_metrics,omitempty"`
	LearningRateMultiplier       *float64  `json:"learning_rate_multiplier,omitempty"`
	Model                        *string   `json:"model,omitempty"`
	NEpochs                      *int64    `json:"n_epochs,omitempty"`
	PromptLossWeight             *float64  `json:"prompt_loss_weight,omitempty"`
	Suffix                       *string   `json:"suffix,omitempty"`
	TrainingFile                 string    `json:"training_file"`
	ValidationFile               *string   `json:"validation_file,omitempty"`
}
