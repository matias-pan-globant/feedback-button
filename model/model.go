package model

type Count struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	Neutral int  `json:"neutral"`
}