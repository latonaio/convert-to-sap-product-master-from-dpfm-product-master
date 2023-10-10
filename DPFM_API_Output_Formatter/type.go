package dpfm_api_output_formatter

type Output struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	APIStatusCode     int    `json:"api_status_code"`
	RuntimeSessionID  string `json:"runtime_session_id"`
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
	APIType           string `json:"api_type"`
	//DataConcatenation   DataConcatenation `json:"DataConcatenation"`
	Header              Header   `json:"ProductionOrderConfirmation"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type Message struct {
	General             []*General	`json:"General"`
}

type General struct {
	Product             	*string `json:"Product"`
	ProductType         	*string `json:"ProductType"`
	IndustrySector      	*string `json:"IndustrySector"`
	BaseUnit            	*string `json:"BaseUnit"`
	ValidityStartDate   	*string `json:"ValidityStartDate"`
	ProductGroup        	*string `json:"ProductGroup"`
	Division            	*string `json:"Division"`
	GrossWeight         	*string `json:"GrossWeight"`
	NetWeight           	*string `json:"NetWeight"`
	WeightUnit          	*string `json:"WeightUnit"`
	SizeOrDimensionText 	*string `json:"SizeOrDimensionText"`
	ProductStandardID   	*string `json:"ProductStandardID"`
	IndustryStandardName    *string	`json:"IndustryStandardName"`
	ItemCategoryGroup		*string `json:"ItemCategoryGroup"`
	CountryOfOrigin			*string `json:"CountryOfOrigin"`
	CreationDate        	*string `json:"CreationDate"`
	LastChangeDate      	*string `json:"LastChangeDate"`
	IsMarkedForDeletion 	*bool   `json:"IsMarkedForDeletion"`
}
