package server

import ()

type RentDetailRes struct {
	Status      string       `json:"introduceStatus"`
	DisplayNo   int          `json:"displayNo"`
	IsEdit      string       `json:"isEdit"`
	ServerProcs []ServerProc `json:"serveProcList"`
	OrderMap    Order        `json:"orderMap"`
}

type ServerProc struct {
	AliasName   string `json:"aliasName"`
	IntroduceId string `json:"introduceId"`
	DisplayNo   int    `json:"displayNo"`
	ServeCd     string `json:"serveCd"`
	Time        string `json:"time"`
	Title       string `json:"title"`
}

type Order struct {
	IntroduceType   string `json:"introduceType"`
	IntroduceId     string `json:"introduceId"`
	IntroduceStatus string `json:"introduceStatus"`
	EntName         string `json:"entName"`
	Phone           string `json:"phone"`
	ApplicantName   string `json:"applicantName"`
	RequirementType string `json:"requirementType"`
}

type RentDetailLocal struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
	From    string `json:"from"`
	To      string `json:"to"`
}
