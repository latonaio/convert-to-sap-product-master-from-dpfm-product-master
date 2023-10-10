package requests

type Calories struct {
	Product                    string   `json:"Product"`
	BusinessPartener           int      `json:"BusinessPartener"`
	CaloryUnitQuantitiy        int      `json:"CaloryUnitQuantitiy"`
	Calories                   *int     `json:"Calories"`
	CaloryUnit                 *string  `json:"CaloryUnit"`
}

