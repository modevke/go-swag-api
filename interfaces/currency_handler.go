package interfaces

import(
	"net/http"
)

type Currency struct{}

func(c *Currency) CreateCurrency(w http.ResponseWriter, r *http.Request){}
func(c *Currency) FetchCurrencyByID(w http.ResponseWriter, r *http.Request){}
func(c *Currency) FetchCurrencies(w http.ResponseWriter, r *http.Request){}
func(c *Currency) UpdateCurrency(w http.ResponseWriter, r *http.Request){}
func(c *Currency) SoftDeleteCurrency(w http.ResponseWriter, r *http.Request){}
func(c *Currency) HardDeleteCurrency(w http.ResponseWriter, r *http.Request){}