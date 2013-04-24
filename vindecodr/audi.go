package vindecodr

var carLocation = map[string]string {
	"TRU": "Audi - Hungary",
	"WAU": "Audi - Germany",
	"WA1": "Audi - Europe",
}

var carType = map[string]string {
	"TRU": "Audi - Hungary: Passenger Car",
	"WAU": "Audi - Germany: Passenger Car",
	"WA1": "Audi - Europe: SUV",
}

var series = map[string]string {
	"A": "A3 Avant Premium/A4 Premium/A5 Cab Premium/A6 Premium/R8 Coup",
	"B": "A4 Premium Quattro/A6 Premium Quattro/TT & TTS Cp Premium Quattro",
	"C": "A3 Avant Premium Quattro/A5 Premium Quattro/A5 Cab Premium Quattro/Audi Q5 Premium/Audi Q7 Premium",
	"D": "A4 Manual Premium Quattro/R8 Coup - Manual",
	"E": "A3 Avant - Manual Premium/A4 Premium+/A6 Premium+/A5 Manual Premium Quattro",
	"F": "A4 Premium+ Quattro/S4 Premium+ Quattro/A6 Premium+ Quattro/TT & TTS Cp Premium+ Quattro",
	"H": "A4 Manual Premium+ Quattro/S4 Manual Premium+ Quattro",
	"J": "A3 Avant Premium+/A5 Cab Premium+",
	"K": "A4 Prestige Quattro/S4 Prestige Quattro/A6 Prestige Quattro/S6 Prestige Quattro/TT & TTS Cp Prestige Quattro",
	"L": "A3 Avant Premium+ Quattro/A5 Premium+ Quattro/A5 Cab Premium+ Quattro/S5 Premium+ Quattro/A8 Sedan/Audi Q5 Premium+/Audi Q7 Premium+",
	"M": "A4 Manual Prestige Quattro/S4 Manual Prestige Quattro/A8L Sedan/Audi Q5 Premium+S-Line/Audi Q7 Premium+S-Line",
	"N": "A3 Avant-Manual Premium+",
	"R": "A5 Manual Premium+ Quattro/S5 Manual Premium+ Quattro",
	"S": "A4 Advant Premium+ Quattro/A6 Advant Premium+ Quattro/TT & TTS Rdstr Premium Quattro",
	"V": "S5 Prestige Quattro/S5 Cab Prestige Quattro/Audi Q5 Prestige/Audi Q7 Prestige",
	"W": "A4 Advant Premium+ Quattro/A6 Advant Premium+ Quattro/A5 Prestige Quattro S-Line/A5 Cab Prestige Quattro S-Line/Audi Q5 Prestige S-Line/Audi Q7 Prestige S-Line/TT Rdstr Premium+ Quattro",
	"3": "S5 Manual Prestige Quattro",
	"4": "A4 Avant Prestige Quattro/A6 Avant Prestige Quattro/A5 Manual Prestige Quattro S-Line/TT & TTS Rdstr Prestige Quattro",
}

var engine = map[string]string {
	"E": "4 cyl 2.0L 200hp (CBFA-Partial Zero Emissions Vehicle) A3",
	"F": "4 cyl 2.0L 211hp (CAEB) A4/A4q/A4 Avant q/A5q/A5 Cab/A5 Cab q",
	"G": "V6 3.0L 300hp (CCAA) A6 q/A6 Avant q or V6 3.0L 333hp (CCBA) S4/S5 Cab",
	"J": "4 cyl 2.0L TDI 140hp (CBEA) A3",
	"K": "V6 3.2L 265hp (CALA) A5 q/A6 CVT or V6 3.2L 270hp (CALB) Q5",
	"M": "V6 3.0L TDI 225hp (CATA) Q7",
	"N": "V10 5.2L 525hp (BUJ) R8 or V10 5.2L 435hp (BXA) S6",
	"U": "V8 4.2L 420hp (BYH) R8",
	"V": "V8 4.2L 350hp (BAR) Q7 or V8 4.2L 350hp (BVJ) A6q/A8/A8L or V8 4.2L 354hp (CAUA) S5 q",
	"Y": "VR6 3.6L 280hp (BHK) Q7",
	"1": "4 cyl 2.0L 265hp (CDMA) TTS Cpe/Rdstr",
}

var model = map[string]string {
	"4E": "A8/S8",
	"4F": "A6/S6",
	"4L": "Audi Q7",
	"42": "R8",
	"8F": "A5/S5 Cabriolet",
	"8H": "A4/S4 Cabriolet",
	"8J": "TT/TTS",
	"8K": "A4",
	"8P": "A3",
	"8R": "Audi Q5",
	"8T": "A5/S5",
}

var assemblyPlant = map[string]string {
	"A": "Ingolstadt",
	"D": "Bratislava",
	"K": "Karmann-Rheine",
	"N": "Neckarsulm",
	"1": "Gyor",
}

var modelYear = map[string]string {
	"J": "1988",
	"K": "1989",
	"L": "1990",
	"M": "1991",
	"N": "1992",
	"P": "1993",
	"R": "1994",
	"S": "1995",
	"T": "1996",
	"V": "1997",
	"W": "1998",
	"X": "1999",
	"Y": "2000",
	"1": "2001",
	"2": "2002",
	"3": "2003",
	"4": "2004",
	"5": "2005",
	"6": "2006",
	"7": "2007",
	"8": "2008",
	"9": "2009",
	"A": "2010",
}

// WAUAF78E08A042215

func Audi(vin VIN) (details VehicleDetails) {
	details.Location = carLocation[vin.WorldMfgCode + vin.Manufacturer]
	details.Make     = carType[vin.WorldMfgCode + vin.Manufacturer]
	details.Weight   = weightMap[string(vin.Attributes[0])]
	details.Model    = modelMap[string(vin.Attributes[1:3])]
	details.Engine   = engineMap[string(vin.Attributes[3])]
	details.Intro    = introMap[string(vin.Attributes[4])]
	details.Year     = modelYear[vin.ModelYear]
	details.City     = mfgMap[vin.MfgPlant]
	details.Serial   = vin.SerialNumber
	return
}
