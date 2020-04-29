// Code generated by jsonschema. DO NOT EDIT.
package beerjson

import "encoding/json"
import "fmt"

// ID: https://raw.githubusercontent.com/beerjson/beerjson/master/json/timing.json

// The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step.
type TimingType struct {
	// Used to indicate when an addition is added based on a desired specific gravity. E.G. Add dry hop at when SG is 1.018.
	SpecificGravity *GravityType `json:"specific_gravity,omitempty"`
	// Used to indicate when an addition is added based on a desired specific gravity. eg Add brett when pH is 3.4.
	PH *AcidityType `json:"pH,omitempty"`
	// Used to indicate what step this ingredient timing addition is referencing. EG A value of 2 for add_to_fermentation would mean to add during the second fermentation step.
	Step *int32   `json:"step,omitempty"`
	Use  *UseType `json:"use,omitempty"`
	// What time during a process step is added, eg a value of 2 days for a dry hop addition would be added 2 days into the fermentation step.
	Time *TimeType `json:"time,omitempty"`
	// How long an ingredient addition remains, this was referred to as time in the BeerXML standard. E.G. A 40 minute hop boil additions means to boil for 40 minutes, and a 2 day duration for a dry hop means to remove it after 2 days.
	Duration *TimeType `json:"duration,omitempty"`
	// A continuous addition is spread out evenly and added during the entire process step, eg 60 minute IPA by dogfish head takes all ofthe hop additions and adds them throughout the entire boil.
	Continuous *bool `json:"continuous,omitempty"`
}

// Differentiates the specific process type when this ingredient addition is used.
type UseType string

func (s *UseType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	*s = UseType(v)

	switch *s {
	case UseType_AddToMash:
		return nil
	case UseType_AddToBoil:
		return nil
	case UseType_AddToFermentation:
		return nil
	case UseType_AddToPackage:
		return nil
	}
	return fmt.Errorf("UseType: value '%v' does not match any value", v)
}

const (
	UseType_AddToMash         UseType = "add_to_mash"
	UseType_AddToBoil         UseType = "add_to_boil"
	UseType_AddToFermentation UseType = "add_to_fermentation"
	UseType_AddToPackage      UseType = "add_to_package"
)
