package cinful

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
)

//go:embed vendors.json
var b []byte

func init() {
	err := json.Unmarshal(b, &vendors)
	if err != nil {
		panic(err)
	}
}

var vendors []Vendor

type Vendor struct {
	Name     string `json:"name,omitempty"`
	Constant string `json:"constant,omitempty"`
	Env      any    `json:"env,omitempty"`
	PR       any    `json:"pr,omitempty"`
	Val      string `json:"val,omitempty"`
}

func (v *Vendor) String() string {
	vfmt := "Name:  %s\nConst: %s\nEnv:   %v\nVal:   %v"
	return fmt.Sprintf(vfmt, v.Name, v.Constant, v.Env, v.Val)
}

func PrintVendors() {
	for _, v := range vendors {
		fmt.Printf("%#v\n", v)
	}
}

func Info() *Vendor {
	// check vendor first, for more details
	for _, v := range vendors {
		switch vt := v.Env.(type) {
		// normally just a string
		case string:
			val := os.Getenv(vt)
			if val != "" {
				v.Val = val
				return &v
			}
		// a list of strings
		case []interface {}:
			for _, ev := range vt {
				val := os.Getenv(ev.(string))
				if val != "" {
					v.Env = ev.(string)
					v.Val = val
					return &v
				}
			}

		// an ENV var with a specific value
		case map[string]interface {}:
			for ek, ev := range vt {
				val := os.Getenv(ek)
				if val == ev.(string) {
					v.Env = ek
					v.Val, _ = ev.(string)
					return &v
				}
			}
			
		}
	}

	// check some common ENV vars
	for _, ce := range commonEnv {
		val := os.Getenv(ce)
		if val != "" {
			return &Vendor{
				Name: "Common Var",
				Constant: "COMMON",
				Env:      ce,
				Val:      val,
			}
		}
	}
	return nil
}

var commonEnv = []string{
	"CI",
	"CONTINUOUS_INTEGRATION",
	"BUILD_NUMBER",
	"RUN_ID",
}
