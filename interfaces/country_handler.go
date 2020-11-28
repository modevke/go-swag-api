package interfaces

import(
	"log"
	"net/http"
	"encoding/json"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"

	"github.com/modevke/go-swag-api/domain/entity"
	"github.com/modevke/go-swag-api/infrastructure/utility"
)


// SAMPLE DATA
var sample []entity.Country = []entity.Country{
	entity.Country{
		Base: entity.Base{
			ID: "123",
		},
		Name: "Kenya",
		Iso: "KE",
		Iso3: "KEN",
		Numcode: "404",
		Phonecode: "254",
		Currency: "KES",
	},
	entity.Country{
		Base: entity.Base{
			ID: "456",
		},
		Name: "Uganda",
		Iso: "UG",
		Iso3: "UGN",
		Numcode: "406",
		Phonecode: "256",
		Currency: "UGX",
	},
}


type CountryHandler struct{}



func(c *CountryHandler) CreateCountry(res http.ResponseWriter, req *http.Request){

	var request entity.Country

	am := utility.Validator(req.Body, &request, 
		validation.Field(&request.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&request.Iso, validation.Required, validation.Length(1, 2)),
		validation.Field(&request.Iso3, validation.Required, validation.Length(1, 3)),
		validation.Field(&request.Numcode, validation.Required, validation.Length(1, 100)),
		validation.Field(&request.Phonecode, validation.Required, validation.Length(1, 100)),
		validation.Field(&request.Currency, validation.Required, validation.Length(1, 3)),
	)

	if am != nil {
		log.Printf("Something went wrong: %v\n", am)
		return
	}

	log.Println(request.)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(request)


	// var response utility.Responses

	// dt, dterr := utility.Validator(req.Body, request, createCountrySchema)
	
	// if dterr != nil {
	// 	log.Printf("Something went wrong: %v\n", dterr)
	// 	// response.ErrorResponse('Data validation error', )
	// 	return
	// } 

	// jsonString, _ := json.Marshal(dt)
	// if err := json.Unmarshal(jsonString, &request); err != nil {
	// 	log.Printf("Failed to unmarshal")
	// 	return
	// }
	// res.Header().Set("Content-Type", "application/json")
	// res.WriteHeader(http.StatusOK)
	// json.NewEncoder(res).Encode(request)

}


func(c *CountryHandler) FetchCountryByID(res http.ResponseWriter, req *http.Request){


	vars := mux.Vars(req)
	log.Printf("Data passed: %v\n", vars["id"])

	var data entity.Country
	for _, val := range sample{
		if val.Base.ID == vars["id"] {
			data = val
		}
	}
	
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)


}

func(c *CountryHandler) FetchCountryByISO(res http.ResponseWriter, req *http.Request){}
func(c *CountryHandler) FetchCountries(res http.ResponseWriter, req *http.Request){}
func(c *CountryHandler) UpdateCountry(res http.ResponseWriter, req *http.Request){}
func(c *CountryHandler) SoftDeleteCountry(res http.ResponseWriter, req *http.Request){}
func(c *CountryHandler) HardDeleteCountry(res http.ResponseWriter, req *http.Request){}