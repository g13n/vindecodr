package vindecodr

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/hoisie/mustache"
)

func init() {
	http.HandleFunc("/", displayTemplate)
	http.HandleFunc("/decoder", decodeVINNumber)
}

func displayTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintln(w, mustache.RenderFile(getTemplatePath("main.mustache"), map[string]interface{}{"vehicle": false}))
}

func getTemplatePath(template string) string {
	return path.Join(path.Join(os.Getenv("PWD"), "templates"), template)
}

type VehicleFunc func(VIN) VehicleDetails

var VehicleFuncMap = map[string]VehicleFunc {
	"HD": HarleyDavidson,
}

type VehicleDetails struct {
	Location string
	Make     string
	Weight   string
	Model    string
	Engine   string
	Intro    string
	Check    string
	Year     int
	City     string
	Serial   string
}

func decodeVINNumber(w http.ResponseWriter, r *http.Request) {
	vehicle := Vehicle{VIN: r.FormValue("vin"), Type: r.FormValue("type")}
	// lower(HTTP_X_REQUEST_WITH) == "xmlhttprequest"

	if len(vehicle.VIN) != 17 {
		fmt.Fprintln(w, mustache.RenderFile(getTemplatePath("main.mustache"), map[string]interface{}{"vehicle": Vehicle{VIN: vehicle.VIN, Type: vehicle.Type}, "error_vin": true}))
		return
	}

	vin, err := vehicle.Parse()
	if err != nil {
		fmt.Fprintln(w, mustache.RenderFile(getTemplatePath("main.mustache"), map[string]interface{}{"vehicle": Vehicle{VIN: vehicle.VIN, Type: ""}, "error_general": true}))
		return
	}

	if vehicleFunc, ok := VehicleFuncMap[vin.Manufacturer]; ok {
		fmt.Fprintln(w, mustache.RenderFile(getTemplatePath("main.mustache"), map[string]interface{}{"vehicle": vehicle}, map[string]interface{}{"details": vehicleFunc(vin)}))
	}
}
