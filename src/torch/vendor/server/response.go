package server

import ()

// RentDetailRes 租赁明细
type RentDetailRes struct {
	Status      string       `json:"introduceStatus"`
	DisplayNo   int          `json:"displayNo"`
	IsEdit      string       `json:"isEdit"`
	ServerProcs []ServerProc `json:"serveProcList"`
	OrderMap    Order        `json:"orderMap"`
}

// ServerProc 请求用户信息
type ServerProc struct {
	AliasName   string `json:"aliasName"`
	IntroduceId string `json:"introduceId"`
	DisplayNo   int    `json:"displayNo"`
	ServeCd     string `json:"serveCd"`
	Time        string `json:"time"`
	Title       string `json:"title"`
}

// Order 节点操作记录
type Order struct {
	IntroduceType   string `json:"introduceType"`
	IntroduceId     string `json:"introduceId"`
	IntroduceStatus string `json:"introduceStatus"`
	EntName         string `json:"entName"`
	Phone           string `json:"phone"`
	ApplicantName   string `json:"applicantName"`
	RequirementType string `json:"requirementType"`
}


// RentDetailLocal 租赁信息本地化
type RentDetailLocal struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
	From    string `json:"from"`
	To      string `json:"to"`
}

// StatisticDuration 时长统计
type StatisticDuration struct{
	Duration int `json:"duration"`
	Count int `json:"count"`
}