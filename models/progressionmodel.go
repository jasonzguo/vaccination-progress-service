package model

type ProgressionModel struct {
	ID                      string  `bson:"_id,omitempty" json:"id"`
	Country                 string  `bson:"country,omitempty" json:"country"`
	FullyVaccinedPercentage float64 `bson:"fullyVaccinedPercentage,omitempty" json:"fullyVaccinatedPercentage`
}
