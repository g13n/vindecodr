package vindecodr

import (
	"fmt"
	"net/http"

	"github.com/hoisie/mustache"
)

func init() {
	http.HandleFunc("/", displayTemplate)
	http.HandleFunc("/decoder", decodeVINNumber)
}

func displayTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintln(w, mustache.RenderFile(getTemplatePath("application.mustache"), map[string]interface{}{"vehicle": false}))
}

func getTemplatePath(template string) string {
	return "./templates/" + template
}

type VehicleFunc func(VIN) VehicleDetails

var VehicleFuncMap = map[string]VehicleFunc {
	"HD": HarleyDavidson,
	"AU": Audi,
	"A1": Audi,
}

type VehicleDetails struct {
	Location string
	Make     string
	Weight   string
	Model    string
	Engine   string
	Intro    string
	Check    string
	Year     string
	City     string
	Serial   string
}

func decodeVINNumber(w http.ResponseWriter, r *http.Request) {
	vehicle := Vehicle{VIN: r.FormValue("vin"), Type: r.FormValue("type")}
	// lower(HTTP_X_REQUEST_WITH) == "xmlhttprequest"

	vin, err := vehicle.Parse()
	if err != CheckDigitError && err != nil {
		if err == VINError {
			fmt.Fprintln(w, mustache.RenderFile(getTemplatePath("application.mustache"), map[string]interface{}{"vehicle": Vehicle{VIN: vehicle.VIN, Type: ""}, "error": true, "error_vin": true}))
		} else {
			fmt.Fprintln(w, mustache.RenderFile(getTemplatePath("application.mustache"), map[string]interface{}{"vehicle": Vehicle{VIN: vehicle.VIN, Type: ""}, "error": true, "error_general": true}))
		}
		return
	}

	if vehicleFunc, ok := VehicleFuncMap[vin.Manufacturer]; ok {
		fmt.Fprintln(w, mustache.RenderFile(getTemplatePath("application.mustache"), map[string]interface{}{"vehicle": vehicle}, map[string]interface{}{"details": vehicleFunc(vin), "error_check": err}))
	}
}
