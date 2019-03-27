// MetricConvInch can converts meter and inch
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Inch float64
type Meter float64

const (
	MeterPerInch Inch  = 0.3048
	InchPerMeter Meter = 3.2808
)

func (c Inch) String() string  { return fmt.Sprintf("%.2finch", c) }
func (m Meter) String() string { return fmt.Sprintf("%.2fmeter", m) }

func ItoM(c Inch) Meter { return Meter(MeterPerInch * c) }
func MtoI(m Meter) Inch { return Inch(InchPerMeter * m) }

func main() {
	if len(os.Args) < 2 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			convFunc(input.Text())
		}
	} else {
		for _, v := range os.Args[1:] {
			convFunc(v)
		}
	}
}

func convFunc(str string) {
	n, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	meter := Meter(n)
	inch := Inch(n)
	fmt.Printf("%s = %s, %s = %s\n", meter, MtoI(meter), inch, ItoM(inch))
}
