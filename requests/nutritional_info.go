package requests

type NutritionalInfo struct {
	Product                             string   `json:"Product"`
	BusinessPartener                    int      `json:"BusinessPartener"`
	Nutrient                            int      `json:"Nutrient"`
	NutrientContent                     *int     `json:"NutrientContent"`
	NutrientContentUnit                 *string  `json:"NutrientContentUnit"`
}
