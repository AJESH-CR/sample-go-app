package dto

type Level struct {
	ID int `json:"level"`
	Approvers []int `json:"approvers"`
	Cond string `json:"cond"`
}
