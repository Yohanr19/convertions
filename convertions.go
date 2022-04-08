/*
 Takes value and show units convertions--
-Weight in Kilograms cnnverted to Lb
-Weight in Lb to Kilograms
-Temperature in Celsius to Fahrenheit
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
type Celsius float64
type Fahrenheit float64
type Meters float64
type Yards float64
type Foots float64
type Liters float64
type Gallons float64

//Define methods for type convertion

//Takes a value of type Kilograms(kg) and returns a corresponding value of type Pounds(lb)
func WKtoP(k Kilograms) Pounds { return Pounds(k * 2.20462262185) }

//Takes a value of type Pounds(lb) and returns a corresponding value of type Kilograms(kg)
func WPtoK(p Pounds) Kilograms { return Kilograms(p / 2.20462262185) }

// TCtoF converts a Celsius temperature to Fahrenheit.
func TCToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32) }

// TFtoC converts a Fahrenheit temperature to Celsius.
func TFToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5.0 / 9.0) }

// DMtoY converts a value of Meters(m) to Yards (yd)
func DMtoY(m Meters) Yards { return Yards(m / 0.9144) }

// DYtoM converts a value of Yards(yd) to Meters(m)
func DYtoM(y Yards) Meters { return Meters(y * 0.9144) }

// DMtoF converts a value of Meters(m)) to Foot(ft)
func DMtoF(m Meters) Foots { return Foots(m / 0.3048) }

// DFtoM converts a value of Foot(ft)) to Meters(m)
func DFtoM(f Foots) Meters { return Meters(f * 0.3048) }

// DYtoF converts a value of Yards(yd)) to Foot(ft)
func DYtoF(y Yards) Foots { return Foots(y * 3.0) }

// DFtoY converts a value of Foots(ft)) to Yards(yd)
func DFtoY(f Yards) Yards { return Yards(f / 3.0) }

// VLtoG converts a value of Liters(L)) to Gallons(gal(US))
func VLtoG(l Liters) Gallons { return Gallons(l * 0.2641720524) }

// VGtoL converts a value of Gallons(gal(US)) to Liters(L)
func VGtoL(g Gallons) Liters { return Liters(g / 0.2641720524) }

//Make each type implement Stringer interface with method String()
func (k Kilograms) String() string  { return fmt.Sprintf("%.2fkg", k) }
func (p Pounds) String() string     { return fmt.Sprintf("%.2flb", p) }
func (c Celsius) String() string    { return fmt.Sprintf("%.4f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.4f°F", f) }
func (m Meters) String() string     { return fmt.Sprintf("%.2fm", m) }
func (y Yards) String() string      { return fmt.Sprintf("%.4fyd", y) }
func (f Foots) String() string      { return fmt.Sprintf("%.4fft", f) }
func (l Liters) String() string     { return fmt.Sprintf("%.4fL", l) }
func (g Gallons) String() string    { return fmt.Sprintf("%.4fgal(US)", g) }