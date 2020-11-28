package main

import(
	rn "math/rand"
	"time"
	"os"
	"log"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/modevke/go-swag-api/interfaces"
)

// STRUCTURE LOGS
func init(){
	rn.Seed(time.Now().UnixNano())
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main(){

	var err error
	defer func(){
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// LOAD ENVIRONMENT VARIABLES
	if logerr := godotenv.Load(); logerr != nil {
		err = fmt.Errorf("env file is missing")
		return
	}

	// LOAD GLOBAL ENV VARIABLES
	port := os.Getenv("API_PORT")
	host := os.Getenv("API_HOST")

	// TODO VALIDATE ENV VARIABLES

	
	// LOAD ROUTES
	routes := interfaces.Routing()
	app := http.Server{
		Handler: routes,
		Addr: host+":"+port,

		ReadTimeout:       1 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	if serr := app.ListenAndServe(); serr != http.ErrServerClosed{
		err = fmt.Errorf("Unable to start server: %v", serr)
		return
	}


}