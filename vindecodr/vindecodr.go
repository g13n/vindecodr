package vindecodr

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

func init() {
	http.HandleFunc("/", displayTemplate)
	http.HandleFunc("/decoder", decodeVINNumber)
}

func displayTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintln(w, vinFormTemplate)
}

const vinFormTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>VIN Decoder</title>
<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css" media="screen" />
<link rel="stylesheet" type="text/css" href="/css/main.css" media="screen" />
</head>
<body>
<header>
<h1>VIN Decoder</h1>
</header>

<form action="/decoder" class="form-horizontal" method="get">
<fieldset>
<legend>Enter VIN number and select the vehicle type</legend>
<div class="control-group">
<label class="control-label" for="vin">VIN Number:</label>
<div class="controls">
<input id="vin" name="vin" type="text" required>
</div>
</div>
<div class="control-group">
<label class="control-label" for="type">Vehicle:</label>
<div class="controls">
<select id="type" name="type">
<optgroup label="Motorcycle">
<option>Harley Davidson</option>
</optgroup>
</select>
</div>
</div>
<div class="form-actions">
<button class="btn-large btn-primary" type="submit">Decode</button>
</div>
</fieldset>
</form>

<footer>
<p class="muted text-center">Disclaimer: Use at your own risk.  No data is stored.</p>
</footer>
</body>
</html>
`

type Vehicle struct {
	VIN      string
	Type     string
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

var LocationMap = map[string]string {
	"1": "Domestic",
	"5": "International",
}

var MakeMap = map[string]string {
	"HD": "Harley Davidson",
}

var WeightMap = map[string]string {
	"1": "Heavyweight (901 cc and larger)",
	"4": "Lightweight (900 cc and smaller)",
	"8": "Sidecar",
}

var ModelMap = map[string]string {
	"CR": "Superlow (XL883L)",
	"CT": "1200 Custom (XL1200C)",
	"CZ": "Nightster (XL1200N)",
	"LA": "Sporster (XR1200)",
	"LC": "Forty-Eight (XL1200X)",
	"LD": "Sporster 1200 (XR1200X)",
	"LE": "Iron 883 (XL883N)",
	"LF": "Seventy-Two (XL1200V)",
	"LH": "Sporster (XL1200CP)",
	"GP": "Dyna Wide Glide (FXDWG)",
	"GV": "Dyna Super Glide Custom (FXDC)",
	"GX": "Dyna Street Bob (FXDB)",
	"GY": "Dyna Fat Bob (FXDF)",
	"GZ": "Dyna Switchback (FLD)",
	"BW": "Heritage Softail Classic (FLSTC)",
	"BX": "Fat Boy (FLSTF)",
	"JR": "Softail Slim (FLS)",
	"JD": "Softail Deluxe (FLSTN)",
	"JN": "Fat Boy Lo (FLSTFB)",
	"JP": "Blackline (FXS)",
	"FB": "Road King (FLHR)",
	"FC": "Ultra Classic Electra Glide (FLHTCU)",
	"FF": "Electra Glide Classic (FLHTC)",
	"FR": "Road King Classic (FLHRC)",
	"FL": "Touring FLHTCU Special Edition",
	"KB": "Street Glide (FLHX)",
	"KE": "Electra Glide Ultra Limited (FLHTK)",
	"KH": "Road Glide Custom (FLTRX)",
	"A" : "Tri Glide Ultra Classic (FLHTCUTG)",
	"HA": "V-Rod VRSCA",
	"HH": "Night Rod Special (VRSCDX)",
	"HP": "V-Rod Muscle (VRSCF)",
	"PY": "CVO Softail Convertible (FLSTSE3)",
	"PZ": "CVO Street Glide (FLHXSE)",
	"PR": "CVO Ultra Classic Electra Glide (FLHTCUSE7)",
}

var EngineMap = map[string]string {
	"A": "1130 Revolution (100 CV)",
	"B": "1450 Fuel Injected Counter Balanced",
	"C": "1550",
	"D": "1550 EFI",
	"E": "1690 EFI",
	"F": "1690 Balanced-EFI",
	"G": "1246 Revolution EFI",
	"H": "1250 ESPFI",
	"J": "1246",
	"M": "883 Evolution XL / 1690 (2008 and later)",
	"N": "1100 Evolution XL",
	"P": "1200 Evolution XL",
	"R": "1340 Evolution Fuel Injected",
	"S": "500 Single (Armstrong Military)",
	"V": "Twin Cam 88 Carburetor",
	"W": "Twin Cam 88 Fuel Injected",
	"Y": "Twin Cam 88 Counter Balanced - Carb",
	"Z": "1130 Revolution (115hp)",
	"1": "1450 EFI",
	"2": "883 ESPFI",
	"3": "1200 ESPFI",
	"4": "1584cc air cooled, fuel injected",
	"5": "1584 ESPFI",
	"6": "1200",
	"8": "1800 ESPFI",
	"9": "1800 ESPFI",
}

var IntroMap = map[string]string {
	"1": "Regular Introduction date / 49 State calibration",
	"2": "Mid Year Introduction date / 49 State calibration",
	"3": "Regular Introduction Date / California calibration",
	"4": "Cosmetic changes and/or special introductory date/ California calibration",
	"5": "Cosmetic changes and/or special introductory date / Californiacalibration",
	"6": "California/mid year/Claifornia calibration",
	"A": "Regular introduction date/CAN calibration",
	"B": "mid-year introduction date/CAN calibration",
	"C": "Regular introduction date/HDI calibration",
	"D": "Mid-year introduction date/HDI calibration",
	"E": "Regular introduction date/JPN calibration",
	"F": "Mid-Year introduction date/JPN calibration",
	"G": "Regular introduction date/AUS calibration",
	"H": "Regular introduction date/AUS calibration",
	"J": "Regular introduction date/BRZ calibration",
	"K": "Regular introduction date/BRZ calibration",
}

var VinCheckMap = map[string]string {
	"0": "Check digit",
	"1": "Check digit",
	"2": "Check digit",
	"3": "Check digit",
	"4": "Check digit",
	"5": "Check digit",
	"6": "Check digit",
	"7": "Check digit",
	"8": "Check digit",
	"9": "Check digit",
	"X": "Check digit",
}

var MfgMap = map[string]string {
	"A": "Tomahawk, WI",
	"B": "York, PA",
	"C": "Kansas City, MO",
	"D": "Manaus, Brazil",
	"E": "Buell East Troy",
	"K": "Kansas city",
}

var YearMap = map[string]int {
	"A": 2010,
	"B": 2011,
	"C": 2012,
	"D": 2013,
	"E": 2014,
}

func decodeVINNumber(w http.ResponseWriter, r *http.Request) {
	vin := r.FormValue("vin")
	v := Vehicle{VIN: vin, Type: r.FormValue("type")}
	// lower(HTTP_X_REQUEST_WITH) == "xmlhttprequest"

	re, err := regexp.Compile(`(\d)([A-Z]{2})(\d)([A-Z]{2})([A-Z0-9])([1-6A-K])([0-9X])(.)([ABCDEK])(.*)`)
	if err == nil {
		matches := re.FindStringSubmatch(vin)
		if matches != nil {
			if loc, ok := LocationMap[matches[1]]; ok {
				v.Location = loc
			}
			if mk, ok := MakeMap[matches[2]]; ok {
				v.Make = mk
			}
			if wt, ok := WeightMap[matches[3]]; ok {
				v.Weight = wt
			}
			if mod, ok := ModelMap[matches[4]]; ok {
				v.Model = mod
			}
			if eng, ok := EngineMap[matches[5]]; ok {
				v.Engine = eng
			}
			if intro, ok := IntroMap[matches[6]]; ok {
				v.Intro = intro
			}
			if chk, ok := VinCheckMap[matches[7]]; ok {
				v.Check = chk
			}
			v.Year, err = strconv.Atoi(matches[8])
			if err != nil {
				if year, ok := YearMap[matches[8]]; ok {
					v.Year = year
				}
			} else {
				v.Year += 2000
			}
			if city, ok := MfgMap[matches[9]]; ok {
				v.City = city
			}
			v.Serial    = matches[10]
		} else {
			err = errorTemplate.Execute(w, v)
			if err != nil {
				w.Header().Add("Content-Type", "text/html")
				fmt.Fprintln(w, errorTemplate)
			}
			return
		}
	}
	
	err = outputTemplate.Execute(w, v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var outputTemplate = template.Must(template.New("output").Parse(outputHTML))

const outputHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>VIN Decoded for: {{.VIN}} ({{.Type}})</title>
<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css" media="screen" />
<link rel="stylesheet" type="text/css" href="/css/main.css" media="screen" />
</head>
<body>
<header>
<h1>VIN Decoded for: {{.VIN}} ({{.Type}})</h1>
</header>

<form action="/decoder" class="form-horizontal" method="get">
<fieldset>
<legend>Enter VIN number and select the vehicle type</legend>
<div class="control-group">
<label class="control-label" for="vin">VIN Number:</label>
<div class="controls">
<input id="vin" name="vin" type="text" value="{{.VIN}}" required>
</div>
</div>
<div class="control-group">
<label class="control-label" for="type">Vehicle:</label>
<div class="controls">
<select id="type" name="type">
<optgroup label="Motorcycle">
<option>Harley Davidson</option>
</optgroup>
</select>
</div>
</div>
<div class="form-actions">
<button class="btn-primary" type="submit">Decode</button>
</div>
</fieldset>
</form>

<dl class="dl-horizontal">
<dt>Location</dt>
<dd>{{.Location}}</dd>
<dt>Make</dt>
<dd>{{.Make}}</dd>
<dt>Weight</dt>
<dd>{{.Weight}}</dd>
<dt>Model</dt>
<dd>{{.Model}}</dd>
<dt>Engine</dt>
<dd>{{.Engine}}</dd>
<dt>Special</dt>
<dd>{{.Intro}}</dd>
<dt>Special</dt>
<dd>{{.Check}}</dd>
<dt>Manufactured</dt>
<dd>{{.Year}}</dd>
<dt>Assembled</dt>
<dd>{{.City}}</dd>
<dt>Serial Number</dt>
<dd>{{.Serial}}</dd>
</dl>

<footer>
<p class="muted text-center">Disclaimer: Use at your own risk.  No data is stored.</p>
</footer>
</body>
</html>
`

var errorTemplate = template.Must(template.New("output").Parse(errorHTML))

const errorHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>VIN for: {{.VIN}} ({{.Type}})</title>
<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css" media="screen" />
<link rel="stylesheet" type="text/css" href="/css/main.css" media="screen" />
</head>
<body>
<header>
<h1>VIN for: {{.VIN}} ({{.Type}})</h1>
</header>

<p class="lead">
I'm unable to decode the VIN number provided. If you think the VIN is valid, please
submit a bug report here.
</p>

<footer>
<p class="muted text-center">Disclaimer: Use at your own risk.  No data is stored.</p>
</footer>
</body>
</html>
`
