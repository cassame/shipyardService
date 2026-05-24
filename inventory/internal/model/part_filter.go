package model

type PartsFilter struct {
	UUIDs                 []string
	Names                 []string
	Categories            []int32
	ManufacturerCountries []string
	Tags                  []string
}
