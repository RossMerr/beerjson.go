// Code generated by jsonschema. DO NOT EDIT.
package beerjson

import "encoding/json"
import "fmt"

// ID: https://raw.githubusercontent.com/beerjson/beerjson/master/json/hop.json

// HopAdditionType collects the attributes of each hop ingredient for use in a recipe hop bil.
type HopAdditionType struct {
	Year               *string             `json:"year,omitempty"`
	AlphaAcid          *PercentType        `json:"alpha_acid,omitempty"`
	BetaAcid           *PercentType        `json:"beta_acid,omitempty"`
	Name               *string             `json:"name,omitempty"`
	ProductId          *string             `json:"product_id,omitempty"`
	Origin             *string             `json:"origin,omitempty"`
	HopVarietyBaseForm *HopVarietyBaseForm `json:"form,omitempty"`
	Producer           *string             `json:"producer,omitempty"`
	// The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step.
	Timing TimingType            `json:"timing", validate:"required"`
	Amount HopAdditionTypeAmount `json:"amount", validate:"required,oneof"`
}

func (s *HopAdditionType) UnmarshalJSON(b []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil
	}

	hopAdditionTypeAmount := func() HopAdditionTypeAmount {
		raw, ok := m["amount"]
		if !ok {
			return nil
		}

		var volumeType VolumeType
		if err := json.Unmarshal(raw, &volumeType); err == nil {
			return &volumeType
		}

		var massType MassType
		if err := json.Unmarshal(raw, &massType); err == nil {
			return &massType
		}

		return nil
	}
	type Alias HopAdditionType
	aux := &struct {
		Amount HopAdditionTypeAmount `json:"amount", validate:"required,oneof"`
		*Alias
	}{
		Amount: hopAdditionTypeAmount(),
		Alias:  (*Alias)(s),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	s.Amount = aux.Amount

	return nil
}

// HopAdditionTypeAmount
type HopAdditionTypeAmount interface {
	HopAdditionTypeamount()
}
type HopInventoryType struct {
	Amount HopInventoryTypeAmount `json:"amount,omitempty", validate:"oneof"`
}

func (s *HopInventoryType) UnmarshalJSON(b []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil
	}

	hopInventoryTypeAmount := func() HopInventoryTypeAmount {
		raw, ok := m["amount"]
		if !ok {
			return nil
		}

		var volumeType VolumeType
		if err := json.Unmarshal(raw, &volumeType); err == nil {
			return &volumeType
		}

		var massType MassType
		if err := json.Unmarshal(raw, &massType); err == nil {
			return &massType
		}

		return nil
	}
	type Alias HopInventoryType
	aux := &struct {
		Amount HopInventoryTypeAmount `json:"amount,omitempty", validate:"oneof"`
		*Alias
	}{
		Amount: hopInventoryTypeAmount(),
		Alias:  (*Alias)(s),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	s.Amount = aux.Amount

	return nil
}

// HopInventoryTypeAmount
type HopInventoryTypeAmount interface {
	HopInventoryTypeamount()
}

// HopVarietyBase provides unique properties to identify individual records of a hop variety.
type HopVarietyBase struct {
	Name               string              `json:"name", validate:"required"`
	Producer           *string             `json:"producer,omitempty"`
	ProductId          *string             `json:"product_id,omitempty"`
	Origin             *string             `json:"origin,omitempty"`
	Year               *string             `json:"year,omitempty"`
	HopVarietyBaseForm *HopVarietyBaseForm `json:"form,omitempty"`
	AlphaAcid          PercentType         `json:"alpha_acid", validate:"required"`
	BetaAcid           *PercentType        `json:"beta_acid,omitempty"`
}

type HopVarietyBaseForm string

func (s *HopVarietyBaseForm) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	*s = HopVarietyBaseForm(v)

	switch *s {
	case HopVarietyBaseForm_Extract:
		return nil
	case HopVarietyBaseForm_Leaf:
		return nil
	case HopVarietyBaseForm_LeafWet:
		return nil
	case HopVarietyBaseForm_Pellet:
		return nil
	case HopVarietyBaseForm_Powder:
		return nil
	case HopVarietyBaseForm_Plug:
		return nil
	}
	return fmt.Errorf("HopVarietyBaseForm: value '%v' does not match any value", v)
}

const (
	HopVarietyBaseForm_Extract HopVarietyBaseForm = "extract"
	HopVarietyBaseForm_Leaf    HopVarietyBaseForm = "leaf"
	HopVarietyBaseForm_LeafWet HopVarietyBaseForm = "leaf (wet)"
	HopVarietyBaseForm_Pellet  HopVarietyBaseForm = "pellet"
	HopVarietyBaseForm_Powder  HopVarietyBaseForm = "powder"
	HopVarietyBaseForm_Plug    HopVarietyBaseForm = "plug"
)

// Used to differentiate which IBU formula is being used in a recipe. If formula is modified in any way, eg to support whirlpool/flameout additions etc etc, please use `Other` for transparency.
type IBUEstimateType struct {
	Method *IBUMethodType `json:"method,omitempty"`
}

type IBUMethodType string

func (s *IBUMethodType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	*s = IBUMethodType(v)

	switch *s {
	case IBUMethodType_Rager:
		return nil
	case IBUMethodType_Tinseth:
		return nil
	case IBUMethodType_Garetz:
		return nil
	case IBUMethodType_Other:
		return nil
	}
	return fmt.Errorf("IBUMethodType: value '%v' does not match any value", v)
}

const (
	IBUMethodType_Rager   IBUMethodType = "Rager"
	IBUMethodType_Tinseth IBUMethodType = "Tinseth"
	IBUMethodType_Garetz  IBUMethodType = "Garetz"
	IBUMethodType_Other   IBUMethodType = "Other"
)

// oil_content collects all information of a hop variety pertaining to oil content, polyphenols, and thiols. Each individual compound is expressed as a percent of the total oil measurement.
type OilContentType struct {
	BPinene *PercentType `json:"b_pinene,omitempty"`
	Pinene  *PercentType `json:"pinene,omitempty"`
	// The total amount of oil, including hydrocarbons, esters, and terpene alcohols in units of ml of oil per 100g of hop mass.
	TotalOilMlPer100g *float64     `json:"total_oil_ml_per_100g,omitempty"`
	Humulene          *PercentType `json:"humulene,omitempty"`
	Nerol             *PercentType `json:"nerol,omitempty"`
	Xanthohumol       *PercentType `json:"xanthohumol,omitempty"`
	Caryophyllene     *PercentType `json:"caryophyllene,omitempty"`
	Farnesene         *PercentType `json:"farnesene,omitempty"`
	Limonene          *PercentType `json:"limonene,omitempty"`
	Polyphenols       *PercentType `json:"polyphenols,omitempty"`
	Cohumulone        *PercentType `json:"cohumulone,omitempty"`
	Myrcene           *PercentType `json:"myrcene,omitempty"`
	Geraniol          *PercentType `json:"geraniol,omitempty"`
	Linalool          *PercentType `json:"linalool,omitempty"`
}

// VarietyInformation collects the attributes of a hop variety to store as record information.
type VarietyInformation struct {
	AlphaAcid              *PercentType            `json:"alpha_acid,omitempty"`
	VarietyInformationType *VarietyInformationType `json:"type,omitempty"`
	//  Defined as the percentage of hop alpha lost in 6 months of storage.
	PercentLost *PercentType `json:"percent_lost,omitempty"`
	Year        *string      `json:"year,omitempty"`
	Notes       *string      `json:"notes,omitempty"`
	// Oil Content information object.
	OilContent         *OilContentType     `json:"oil_content,omitempty"`
	Producer           *string             `json:"producer,omitempty"`
	ProductId          *string             `json:"product_id,omitempty"`
	HopVarietyBaseForm *HopVarietyBaseForm `json:"form,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Substitutes        *string             `json:"substitutes,omitempty"`
	Inventory          *HopInventoryType   `json:"inventory,omitempty"`
	Origin             *string             `json:"origin,omitempty"`
	BetaAcid           *PercentType        `json:"beta_acid,omitempty"`
}

type VarietyInformationType string

func (s *VarietyInformationType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	*s = VarietyInformationType(v)

	switch *s {
	case VarietyInformationType_Aroma:
		return nil
	case VarietyInformationType_Bittering:
		return nil
	case VarietyInformationType_Flavor:
		return nil
	case VarietyInformationType_AromaBittering:
		return nil
	case VarietyInformationType_BitteringFlavor:
		return nil
	case VarietyInformationType_AromaFlavor:
		return nil
	case VarietyInformationType_AromaBitteringFlavor:
		return nil
	}
	return fmt.Errorf("VarietyInformationType: value '%v' does not match any value", v)
}

const (
	VarietyInformationType_Aroma                VarietyInformationType = "aroma"
	VarietyInformationType_Bittering            VarietyInformationType = "bittering"
	VarietyInformationType_Flavor               VarietyInformationType = "flavor"
	VarietyInformationType_AromaBittering       VarietyInformationType = "aroma/bittering"
	VarietyInformationType_BitteringFlavor      VarietyInformationType = "bittering/flavor"
	VarietyInformationType_AromaFlavor          VarietyInformationType = "aroma/flavor"
	VarietyInformationType_AromaBitteringFlavor VarietyInformationType = "aroma/bittering/flavor"
)
