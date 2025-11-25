package entities

type AssignTagRequest struct {
	Tag            string `json:"tag" binding:"required"`
	Identification string `json:"identification" binding:"required"`
}

type AssignTagResult struct {
	ID int64 `json:"id"`
}
