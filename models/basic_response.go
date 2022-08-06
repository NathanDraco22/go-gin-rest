package models

type BasicResponse struct {
	Ok        bool   `json:"ok"`
	Msg       string `json:"msg"`
	ExtraInfo string `json:"extraInfo"`
}
