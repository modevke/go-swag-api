package interfaces

import(
	"log"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/v4"

	"github.com/modevke/go-swag-api/domain/entity"
	"github.com/modevke/go-swag-api/infrastructure/utility"
)

type CountryHandler struct{}


// CREATE COUNTRY
var createCountrySchema = validation.Map(
	validation.Key("name", validation.Required, validation.Length(1, 100)),
	validation.Key("iso", validation.Required, validation.Length(1, 100)),
	validation.Key("iso3", validation.Required, validation.Length(1, 100)),
	validation.Key("numcode", validation.Required, validation.Length(1, 100)),
	validation.Key("phonecode", validation.Required, validation.Length(1, 100)),
	validation.Key("currency", validation.Required, validation.Length(1, 100)),
)

func(c *CountryHandler) CreateCountry(res http.ResponseWriter, req *http.Request){

	var request entity.Country
	dt, dterr := utility.Validator(req.Body, request, createCountrySchema)
	
	if dterr != nil {
		log.Printf("Something went wrong: %v\n", dterr)
	} else {
		log.Printf("Data is ok: %v\n", dt)
	}

}


func(c *CountryHandler) FetchCountryByID(res http.ResponseWriter, req *http.Request){}

func(c *CountryHandler) FetchCountryByISO(res http.ResponseWriter, req *http.Request){}
func(c *CountryHandler) FetchCountries(res http.ResponseWriter, req *http.Request){}
func(c *CountryHandler) UpdateCountry(res http.ResponseWriter, req *http.Request){}
func(c *CountryHandler) SoftDeleteCountry(res http.ResponseWriter, req *http.Request){}
func(c *CountryHandler) HardDeleteCountry(res http.ResponseWriter, req *http.Request){}