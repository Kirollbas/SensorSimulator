package dto

const (
	PRNGTypePCG     = "PRNG_TYPE_PCG"
	PRNGTypeXoshiro = "PRNG_TYPE_XOSHIRO"
)

type Prng struct {
	Type string `json:"type"`
	Seed int64  `json:"seed"`
}
