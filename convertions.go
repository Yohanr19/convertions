/*
 Takes value and show units convertions--
-Weight in Kilograms cnnverted to Lb
-Weight in Lb to Kilograms
-Temperature in Celcius to Fahrenheit
-Temperature in Fahrenheit to Celcius
-Distance in Meters to Yards
-Distance in Yards to Meters
-Distance in Meters to Foot
-Distance in Foot to Meters
-Distance in Foot to Yards
-Distance in Yards to Foot
-Volume in Liters to Gallons
-Volume to Gallons in Liters
*/
package convertions

import (
	"fmt"
)

// Define a type for each unit , each underlying type will be float64

type Kilograms float64
type Pounds float64
type Celcius float64
type Fahrenheit float64
type Meters float64
type Yards float64
type Foots float64
type Liters float64
type Gallons float64

//Define methods for type convertion

//Takes a value of type Kilograms(kg) and returns a corresponding value of type Pounds(lb)
func WKtoP(k Kilograms) Pounds {
	return Pounds(k * 2.20462262185)
}

//Takes a value of type Pounds(lb) and returns a corresponding value of type Kilograms(kg)
func WPtoK(p Pounds) Kilograms {
	return Kilograms(p / 2.20462262185)
}

//Make each type implement Stringer interface with method String()

func (k Kilograms) String() string {
	return fmt.Sprintf("%.2fkg", k)
}
func (p Pounds) String() string {
	return fmt.Sprintf("%.2flb", p)
}
