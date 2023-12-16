# FineTuningJobNEpochs

The number of epochs to train the model for. An epoch refers to one full cycle through the training dataset.
"auto" decides the optimal number of epochs based on the size of the dataset. If setting the number manually, we support any number between 1 and 50 epochs.


## Supported Types

### FineTuningJob1

```go
fineTuningJobNEpochs := shared.CreateFineTuningJobNEpochsFineTuningJob1(shared.FineTuningJob1{/* values here */})
```

### 

```go
fineTuningJobNEpochs := shared.CreateFineTuningJobNEpochsInteger(int64{/* values here */})
```

