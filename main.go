package main

import (
	"encuestify/data"
	"encuestify/routes"
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(fmt.Sprintf(":%s", data.Config.Port), routes.Router)
}
