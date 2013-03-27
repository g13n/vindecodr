package vindecodr

import (
	"errors"
	"regexp"
	"strconv"
)

type Vehicle struct {
	VIN  string
	Type string
}

// Structure (17 character):
// http://forums.audiworld.com/vindecoder.php#vin_intro
type VIN struct {
	WorldMfgCode string	// char[1]
	Manufacturer string	// char[2]
	Attributes   string	// char[5]
	CheckDigit   string	// char[1]
	ModelYear    int
	MfgPlant     string	// char[1]
	SerialNumber string	// char[6]
}

var WorldMfgCodeMap = map[string]string {
	"1": "Domestic",
	"5": "International",
}

var yearMap = map[string]int {
	"A": 2010,
	"B": 2011,
	"C": 2012,
	"D": 2013,
	"E": 2014,
}

var CheckDigitError = errors.New("Check-digit mismatch")
var VINError = errors.New("Invalid VIN")

// Parse the given Vehicle structure and return as VIN
func (v Vehicle) Parse() (VIN, error) {
	var chkErr error

	if len(v.VIN) != 17 {
		return VIN{}, VINError
	}

	chkErr = nil
	if !isValidVIN(v.VIN) {
		chkErr = CheckDigitError
	}

	re, err := regexp.Compile(`(\d)([A-Z]{2})([A-Z0-9]{5})(\d)([A-Z0-9])([A-Z])(\d{6})`)
	if err != nil {
		return VIN{}, err
	}

	match := re.FindStringSubmatch(v.VIN)
	if match != nil {
		v := VIN{}
		v.WorldMfgCode = match[1]
		v.Manufacturer = match[2]
		v.Attributes   = match[3]
		v.CheckDigit   = match[4]
		if year, err := strconv.Atoi(match[5]); err != nil {
			if y, ok := yearMap[match[5]]; ok {
				v.ModelYear = y
			} else {
				return VIN{}, errors.New("Invalid year")
			}
		} else {
			v.ModelYear = year + 2000
		}
		v.MfgPlant     = match[6]
		v.SerialNumber = match[7]
		return v, chkErr
	}
	return VIN{}, errors.New("Match Error")
}

func (v VIN) Stringer() string {
	var s string

	s = v.WorldMfgCode + ", " + v.Manufacturer + ", " +
		v.Attributes + ", " + v.CheckDigit + ", " +
		strconv.Itoa(v.ModelYear) + ", " + v.MfgPlant + ", " +
		v.SerialNumber
	return s
}

// http://en.wikipedia.org/wiki/Vehicle_Identification_Number#Transliterating_the_numbers
var checkDigitMap = map[string]int {
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	// I: nil
	"J": 1,
	"K": 2,
	"L": 3,
	"M": 4,
	"N": 5,
	// O: nil
	"P": 7,
	// Q: nil
	"R": 9,
	"S": 2,
	"T": 3,
	"U": 4,
	"V": 5,
	"W": 6,
	"X": 7,
	"Y": 8,
	"Z": 9,
}

// http://en.wikipedia.org/wiki/Vehicle_Identification_Number#Weights_used_in_calculation
var weightTable = map[int]int {
	 1: 8,
	 2: 7,
	 3: 6,
	 4: 5,
	 5: 4,
	 6: 3,
	 7: 2,
	 8: 10,
	 9: 0,
	10: 9,
	11: 8,
	12: 7,
	13: 6,
	14: 5,
	15: 4,
	16: 3,
	17: 2,
}

// Return true if VIN is valid, otherwise return false.
//
// This logic is based on the following ``check digit'' algorithm.
// http://en.wikipedia.org/wiki/Vehicle_Identification_Number#Check_digit_calculation
func isValidVIN(vin string) bool {
	for i := 0; i < 10; i++ {
		checkDigitMap[strconv.Itoa(i)] = i;
	}

	i   := 1
	sum := 0
	for _, c := range vin {
		var val int
		var ok bool

		if val, ok = checkDigitMap[string(c)]; !ok {
			return false
		}
		sum += val * weightTable[i]
		i++
	}

	var chkDigit string

	reminder := sum % 11
	if reminder == 10 {
		chkDigit = "X"
	} else {
		chkDigit = strconv.Itoa(reminder)
	}

	if chkDigit != vin[8:9] {
		return false
	}

	return true
}
