package request

type UploadPet struct {
	ASCII       string `json:"ascii" binding:"omitempty,required"`
	Description string `json:"description" binding:"omitempty,required"`
}
