package requests

type Allergen struct {
	Product                    string   `json:"Product"`
	BusinessPartener           *int     `json:"BusinessPartener"`
	Allergen                   *string  `json:"Allergen"`
	AllergenIsContained        *bool    `json:"AllergenIsContained"`

}
