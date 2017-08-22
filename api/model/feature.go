package model

import "github.com/ryantomlinson/jaffa-api/db"

type FeatureFlag struct {
	Key 		string
	Version 	int
	On 		bool
}

// TODO: Change this to take in userid as a guid and add to select statement
func (f FeatureFlag) All() (features []FeatureFlag, err error) {
	_, err = db.GetDB().Select(&features, "SELECT * FROM FeatureFlags")
	return features, err
}