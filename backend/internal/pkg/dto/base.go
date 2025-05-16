package dto

const (
	BaseTypeBezier   = "BASE_TYPE_BEZIER"
	BaseTypeConstant = "BASE_TYPE_CONSTANT"
	BaseTypeLinear   = "BASE_TYPE_LINEAR"
	BaseTypeSinewave = "BASE_TYPE_SINEWAVE"
)

type Base struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type ConstantBase struct {
	Value float64 `json:"value"`
}

type CommonBase struct {
	Generator Prng     `json:"generator"`
	MinValue  float64  `json:"min_value"`
	MaxValue  float64  `json:"max_value"`
	MinPeriod Duration `json:"min_period"`
	MaxPeriod Duration `json:"max_period"`
}
