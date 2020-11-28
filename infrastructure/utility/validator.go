package utility

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-ozzo/ozzo-validation/v4"
)

func Validator(data io.ReadCloser, request interface{}, schema validation.MapRule) (interface{}, error) {
	err := json.NewDecoder(data).Decode(&request)

	if err != nil {
		return nil, fmt.Errorf("Invalid JSON body")
	}

	verr := validation.Validate(request, schema)

	if verr != nil {
		return nil, verr
	}

	return request, nil
}
