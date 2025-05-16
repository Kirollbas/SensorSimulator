package dto

const (
	ModifierTypeConstantOffset      = "MODIFIER_TYPE_CONSTANT_OFFSET"
	ModifierTypeHysteresis          = "MODIFIER_TYPE_HYSTERESIS"
	ModifierTypeInertia             = "MODIFIER_TYPE_INERTIA"
	ModifierTypeNonLinearDependence = "MODIFIER_TYPE_NONLINEAR_DEPENDENCE"
	ModifierTypeProgressingOffset   = "MODIFIER_TYPE_PROGRESSING_OFFSET"
	ModifierTypeQuantization        = "MODIFIER_TYPE_QUANTIZATION"
	ModifierTypeRandomAddDash       = "MODIFIER_TYPE_RANDOM_ADD_DASH"
	ModifierTypeRandomFixedDash     = "MODIFIER_TYPE_RANDOM_FIXED_DASH"
	ModifierTypeWhiteNoise          = "MODIFIER_TYPE_WHITE_NOISE"
	ModifierTypeDependence          = "MODIFIER_TYPE_DEPENDENCE"
)

type Modifier struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type ConstantOffsetModifier struct {
	Offset float64 `json:"offset"`
}

type HysteresisModifier struct {
	Percentage int `json:"percentage"`
}

type InertitaModifier struct {
	Value  float64  `json:"value"`
	Period Duration `json:"period"`
}

type NonLinearDependenceModifier struct {
	Coefficient float64 `json:"coefficient"`
	Center      float64 `json:"center"`
}

type ProgressingOffsetModifier struct {
	Value    float64  `json:"value"`
	Interval Duration `json:"interval"`
}

type QuantizationModifier struct {
	Quant float64 `json:"quant"`
}

type RandomAddDashModifier struct {
	Generator       Prng     `json:"generator"`
	MinAddValue     float64  `json:"min_add_value"`
	MaxAddValue     float64  `json:"max_add_value"`
	MinDashDuration Duration `json:"min_dash_duration"`
	MaxDashDuration Duration `json:"max_dash_duration"`
	AvgPeriod       Duration `json:"avg_period"`
}

type RandomFixedDashModifier struct {
	Generator       Prng     `json:"generator"`
	Value           float64  `json:"value"`
	MinDashDuration Duration `json:"min_dash_duration"`
	MaxDashDuration Duration `json:"max_dash_duration"`
	AvgPeriod       Duration `json:"avg_period"`
}

type WhiteNoiseModifier struct {
	Generator Prng    `json:"generator"`
	MaxValue  float64 `json:"max_value"`
}

type DependenceModifier struct {
	SimulatorName string  `json:"simulator_name"`
	Center        float64 `json:"center"`
	Coefficient   float64 `json:"coefficient"`
}
