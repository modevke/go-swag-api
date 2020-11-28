package interfaces

import(
	"net/http"
	"github.com/gorilla/mux"
)

func Routing() http.Handler{
	r := mux.NewRouter()

	// ROUTER GROUPS
	api := r.PathPrefix("/api/v1").Subrouter()


	// Country APIs
	countryRoutes := api.PathPrefix("/country").Subrouter()
	country := CountryHandler{}
	
	countryRoutes.HandleFunc("/create", country.CreateCountry).Methods("POST")
	countryRoutes.HandleFunc("/fetch-id/{id}", country.FetchCountryByID).Methods("GET")
	countryRoutes.HandleFunc("/fetch-iso/{iso}", country.FetchCountryByISO).Methods("GET")
	countryRoutes.HandleFunc("/fetch-all", country.FetchCountries).Methods("GET")
	countryRoutes.HandleFunc("/update", country.UpdateCountry).Methods("PUT")
	countryRoutes.HandleFunc("/delete-soft", country.SoftDeleteCountry).Methods("DELETE")
	countryRoutes.HandleFunc("/delete-hard", country.HardDeleteCountry).Methods("DELETE")


	// Currency APIs
	currencyRoutes := api.PathPrefix("/currency").Subrouter()
	currency := Currency{}

	currencyRoutes.HandleFunc("/create", currency.CreateCurrency).Methods("POST")
	currencyRoutes.HandleFunc("/fetch-id/{id}", currency.FetchCurrencyByID).Methods("GET")
	currencyRoutes.HandleFunc("/fetch-all", currency.FetchCurrencies).Methods("GET")
	currencyRoutes.HandleFunc("/update", currency.UpdateCurrency).Methods("PUT")
	currencyRoutes.HandleFunc("/delete-soft", currency.SoftDeleteCurrency).Methods("DELETE")
	currencyRoutes.HandleFunc("/delete-hard", currency.HardDeleteCurrency).Methods("DELETE")

	return r
}