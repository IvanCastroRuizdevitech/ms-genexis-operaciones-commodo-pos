package entities

type User struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Identification string `json:"identification"`
	Status         string `json:"status"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	ProfileID      int64  `json:"profile_id"`
	Tag            string `json:"tag"`
}
