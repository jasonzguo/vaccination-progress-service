package model

type ProgressionModel struct {
	ID                      string  `bson:"_id,omitempty" json:"id"`
	Country                 string  `bson:"country,omitempty" json:"country"`
	FullyVaccinedPercentage float64 `bson:"fullyVaccinedPercentage,omitempty" json:"fullyVaccinatedPercentage"`
}

type MetaData struct {
	Count  int64  `bson:"count,omitempty" json:"count"`
	LastId string `bson:"lastId,omitempty" json:"lastId"`
}

type PaginatedProgressionModel struct {
	Data []ProgressionModel `bson:"data,omitempty" json:"data"`
	Meta MetaData           `bson:"meta,omitempty" json:"meta"`
}
