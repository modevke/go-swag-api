package utility

type Responses struct{
	Status		string		`json:"status"`
	Description	string		`json:"description"`
	Data		interface{}	`json:"data,omitempty"`
	Errors		interface{}	`json:errors,omitempty`
}