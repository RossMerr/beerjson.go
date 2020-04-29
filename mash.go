// Code generated by jsonschema. DO NOT EDIT.
package beerjson

// ID: https://raw.githubusercontent.com/beerjson/beerjson/master/json/mash.json

// This defines the procedure for performing unique mashing processes.
type MashProcedureType struct {
	Name string `json:"name", validate:"required"`
	// Initial grain temperature prior to the start of the mash.
	GrainTemperature TemperatureType `json:"grain_temperature", validate:"required"`
	Notes            *string         `json:"notes,omitempty"`
	MashSteps        []MashStepType  `json:"mash_steps", validate:"required"`
}
