package entity

type Currency struct {
	Base
	Name   string	`json:"name"`
	Code   string	`json:"code"`
	Symbol string	`json:"symbol"`
	Status string	`json:"status"`
}