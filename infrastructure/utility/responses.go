package utility

type ResponseError struct{
	Code			string	`json:"code"`
	Description	string	`json:"description"`
}

type Responses struct{
	Status		string		`json:"status"`
	Description	string		`json:"description"`
	Data		interface{}	`json:"data,omitempty"`
	Errors		[]ResponseError	`json:"error"`
}

func (r *Responses) ErrorResponse(desc string, data []ResponseError ){
	r.Status = "001"
	r.Description = desc
	r.Errors = data
}

func (r *Responses) SuccessResponse(status string, desc string, data interface{} ){
	r.Status = status
	r.Description = desc
	r.Data = data
}