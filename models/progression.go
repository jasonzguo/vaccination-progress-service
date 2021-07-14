package model

type ProgressionModel struct {
	ID                      string  `bson:"_id,omitempty" json:"id"`
	Country                 string  `bson:"country,omitempty" json:"country"`
	FullyVaccinedPercentage float64 `bson:"fullyVaccinedPercentage,omitempty" json:"fullyVaccinatedPercentage"`
}

type MetaData struct {
	Count  int64  `bson:"count,omitempty"`
	LastId string `bson:"lastId,omitempty"`
}

type PaginatedProgressionModel struct {
	Data []ProgressionModel `bson:"data,omitempty"`
	Meta MetaData           `bson:"meta,omitempty"`
}
