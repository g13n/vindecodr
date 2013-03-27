package vindecodr

var makeMap = map[string]string {
	"HD": "Harley Davidson",
}

var weightMap = map[string]string {
	"1": "Heavyweight (901 cc and larger)",
	"4": "Lightweight (900 cc and smaller)",
	"8": "Sidecar",
}

var modelMap = map[string]string {
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

var engineMap = map[string]string {
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

var introMap = map[string]string {
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

var mfgMap = map[string]string {
	"A": "Tomahawk, WI",
	"B": "York, PA",
	"C": "Kansas City, MO",
	"D": "Manaus, Brazil",
	"E": "Buell East Troy",
	"K": "Kansas city",
}

func HarleyDavidson(vin VIN) (details VehicleDetails) {
	details.Location = WorldMfgCodeMap[vin.WorldMfgCode]
	details.Make     = makeMap[vin.Manufacturer]
	details.Weight   = weightMap[string(vin.Attributes[0])]
	details.Model    = modelMap[string(vin.Attributes[1:3])]
	details.Engine   = engineMap[string(vin.Attributes[3])]
	details.Intro    = introMap[string(vin.Attributes[4])]
	details.Year     = vin.ModelYear;
	details.City     = mfgMap[vin.MfgPlant]
	details.Serial   = vin.SerialNumber
	return
}
