package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Multiple DC MPP strings combined (PV Meter)",
		Sample: `power:
  type: calc
  add:
  - type: modbus
    model: sunspec
    value: 160:1:DCW # SunSpec Model 160 MPP string 1 DCW
    uri: 192.0.2.2:502
    id: 1
  - type: modbus
    model: sunspec
    value: 160:2:DCW # SunSpec Model 160 MPP string 2 DCW
    uri: 192.0.2.2:502
    id: 1`,
	}

	registry.Add(template)
}
