package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-product-master-from-sap-product-master/DPFM_API_Input_Reader"
	"strconv"
)

func OutputFormatter(
	sdc *dpfm_api_input_reader.SDC,
	//psdc *dpfm_api_processing_formatter.ProcessingFormatterSDC,
	osdc *Output,
) error {
	//general := ConvertToGeneral(*sdc, *psdc)
	general := ConvertToGeneral(*sdc)

	//osdc.DataConcatenation = DataConcatenation{
	//	General: general,
	//}

	osdc.General = *general

	osdc.ServiceLabel = "SAPProductMasterCreates"
	osdc.APISchema = "SAPProductMasterCreates"
	osdc.APIProcessingResult = getBoolPtr(true)

	return nil
}

func float32ToString(f *float32) *string {
	if f != nil {
		str := strconv.FormatFloat(float64(*f), 'f', -1, 32)
		return &str
	}
	return nil
}

func ConvertToGeneral(
	sdc dpfm_api_input_reader.SDC,
	// psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) *General {
//	product := ""
//	productType := ""
	industrySector := "M"
//	baseUnit := ""
//	validityStartDate := ""
//	productGroup := ""
	division := "01"
//	grossWeight := ""
//	netWeight := ""
//	weightUnit := ""
//	sizeOrDimensionText := ""
//	productStandardID := ""
//	industryStandardName := ""
//	itemCategoryGroup := ""
//	countryOfOrigin := ""
//	creationDate := ""
//	lastChangeDate := ""
//	isMarkedForDeletion := ""

	general := &General{
		Product:                   sdc.Message.General.Product,
		ProductType:               sdc.Message.General.ProductType,
		IndustrySector:            &industrySector,
		BaseUnit:                  sdc.Message.General.BaseUnit,
		ValidityStartDate:         sdc.Message.General.ValidityStartDate,
		ProductGroup:              sdc.Message.General.ProductGroup,
		Division:                  &division,
		GrossWeight:               &float32ToString(sdc.Message.General.GrossWeight),
		NetWeight:                 &float32ToString(sdc.Message.General.NetWeight),
		WeightUnit:                &float32ToString(sdc.Message.General.WeightUnit),
		SizeOrDimensionText:       sdc.Message.General.SizeOrDimensionText,
		ProductStandardID:         sdc.Message.General.ProductStandardID,
		IndustryStandardName:      sdc.Message.General.IndustryStandardName,
		ItemCategoryGroup:         sdc.Message.General.ItemCategoryGroup,
		CountryOfOrigin:           sdc.Message.General.CountryOfOrigin,
		CreationDate:              sdc.Message.General.CreationDate,
		LastChangeDate:            sdc.Message.General.LastChangeDate,
		IsMarkedForDeletion:       sdc.Message.General.IsMarkedForDeletion,
	}

	return general
}

func getBoolPtr(b bool) *bool {
	return &b
}
