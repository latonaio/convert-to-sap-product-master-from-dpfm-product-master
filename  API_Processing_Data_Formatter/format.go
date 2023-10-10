package dpfm_api_input_reader

import (
	"convert-to-dpfm-product-master-from-sap-product-master/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.ProductMaster
	return &requests.General{
    
                Product:                            data.Product,
                ProductType:                        data.ProductType,
                BaseUnit:                           data.BaseUnit,
                ValidityStartDate:                  data.ValidityStartDate,
                ProductGroup:                       data.ProductGroup,
                GrossWeight:                        data.GrossWeight,
                NetWeight:                          data.NetWeight,
                WeightUnit:                         data.WeightUnit,
                InternalCapacityQuantity:           data.InternalCapacityQuantity,
                InternalCapacityQuantityUnit:       data.InternalCapacityQuantityUnit,
                SizeOrDimensionText:                data.SizeOrDimensionText,
                ProductStandardID:                  data.ProductStandardID,
	            IndustryStandardName:               data.IndustryStandardName,
                //ItemCategory:                       data.,
                //CountryOfOrigin:                    data.,
                CountryOfOriginLanguage:            data.CountryOfOriginLanguage,
                BarcodeType:                        data.BarcodeType,
                //ProductAccountAssignmentGroup:      data.,
                CreationDate:                       data.CreationDate,
                LastChangeDate:                     data.LastChangeDate,
                IsMarkedForDeletion:                data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToProductDescription() *requests.ProductDescription {
	data := sdc.ProductMaster
	return &requests.ProductDescription{

                Product:                data.Product,   
                Language:               data.Language,    
                ProductDescription:     data.ProductDescription,              

	}
}

func (sdc *SDC) ConvertToGeneralDoc() *requests.GeneralDoc {
	data := sdc.ProductMaster
	return &requests.GeneralDoc{

                Product:                   data.Product,     
                //DocType:                  data.,     
                //DocVersionID:             data.,          
                //DocID:                    data.,   
                //FileExtension:            data.,
                //FileName:                 data.,
                //FilePath:                 data.,
                //DocIssuerBusinessPartner: data.,

	}
}

func (sdc *SDC) ConvertToBusinessPartner() *requests.BusinessPartner {
	data := sdc.ProductMaster
	return &requests.BusinessPartner{

                Product:                   data.Product,  
                //BusinessPartner:           data.,          
                //ValidityEndDate:           data.,          
                ValidityStartDate:         data.ValidityStartDate,            
                //BusinessPartnerProduct:    data.,                 
                //IsMarkedForDeletion:       data.,              
	}
}

func (sdc *SDC) ConvertToProductDescriptionByBusinessPartner() *requests.ProductDescriptionByBusinessPartner {
	data := sdc.ProductMaster
	return &requests.ProductDescriptionByBusinessPartner{

                Product:             data.Product,  
                //BusinessPartner:     data.,          
                Language:            data.Language,   
                ProductDescription:  data.ProductDescription,             

	}
}

func (sdc *SDC) ConvertToBPPlant() *requests.BPPlant {
	data := sdc.ProductMaster
	return &requests.BPPlant{
                
       			Product:                                   		data.Product,
    			//BusinessPartner:                           		data.,
    			Plant:                                     		data.Plant,
    			//AvailabilityCheckType:                     		data.,
    			MRPType:                                   		data.MRPType,
    			MRPController:                             		data.MRPController,
    			ReorderThresholdQuantity:                  		data.ReorderThresholdQuantity,
    			PlanningTimeFence:                         		data.PlanningTimeFence,
    			MRPPlanningCalendar:                       		data.MRPPlanningCalendar,
    			SafetyStockQuantityInBaseUnit:             		data.SafetyStockQuantityInBaseUnit,
    			SafetyDuration:                            		data.SafetyDuration,
    			MaximumStockQuantityInBaseUnit:            		data.MaximumStockQuantityInBaseUnit,
    			MinimumDeliveryQuantityInBaseUnit:         		data.MinimumDeliveryQuantityInBaseUnit,
    			MinimumDeliveryLotSizeQuantityInBaseUnit:  		data.MinimumDeliveryLotSizeQuantityInBaseUnit,
    			StandardDeliveryLotSizeQuantityInBaseUnit: 		data.StandardDeliveryLotSizeQuantityInBaseUnit,
    			DeliveryLotSizeRoundingQuantityInBaseUnit: 		data.DeliveryLotSizeRoundingQuantityInBaseUnit,
    			MaximumDeliveryLotSizeQuantityInBaseUnit:  		data.MaximumDeliveryLotSizeQuantityInBaseUnit,
    			MaximumDeliveryQuantityInBaseUnit:         		data.MaximumDeliveryQuantityInBaseUnit,
    			DeliveryLotSizeIsFixed:                    		data.DeliveryLotSizeIsFixed,
    			StandardDeliveryDurationInDays:            		data.StandardDeliveryDurationInDays,
    			IsBatchManagementRequired:                 		data.IsBatchManagementRequired,
    			//BatchManagementPolicy:                     		data.,
    			//InventoryUnit:                             		data.,
    			ProfitCenter:                              		data.ProfitCenter,
    			IsMarkedForDeletion:                       		data.IsMarkedForDeletion,

	}
}

func (sdc *SDC) ConvertToBPPlantDoc() *requests.BPPlantDoc {
	data := sdc.ProductMaster
	return &requests.BPPlantDoc{     

  				Product:                        data.Product,
  				//BusinessPartner:                data.,
  				Plant:                          data.Plant,
  				//DocType:                        data.,
  				//DocVersionID:                   data.,
  				//DocID:                          data.,
  				FileExtension:                  data.FileExtension,
  				//FileName:                       data.,
  				FilePath:                       data.FilePath,
  				DocIssuerBusinessPartner:       data.DocIssuerBusinessPartner,

	}
}

func (sdc *SDC) ConvertToStorageLocation() *requests.StorageLocation {
	data := sdc.ProductMaster
	return &requests.StorageLocation{

                Product:                  data.Product,         
                //BusinessPartner:          data.,                 
                Plant:                    data.Plant,       
                StorageLocation:          data.StorageLocation,                 
                CreationDate:             data.CreationDate,
                //InventoryBlockStatus:     data.,
                //IsMarkedForDeletion:      data.,                     

	}
}

func (sdc *SDC) ConvertToMRPArea() *requests.MRPArea {
	data := sdc.ProductMaster
	return &requests.MRPArea{                
 
     			Product:                                   		data.Product,
    			//BusinessPartner:                           		data.,
    			Plant:                                     		data.Plant,
    			MRPArea:                                   		data.MRPArea,
    			StorageLocationForMRP:                     		data.StorageLocationForMRP,
    			MRPType:                                   		data.MRPType,
    			MRPController:                             		data.MRPController,
    			ReorderThresholdQuantity:                  		data.ReorderThresholdQuantity,
    			PlanningTimeFence:                         		data.PlanningTimeFence,
    			MRPPlanningCalendar:                       		data.MRPPlanningCalendar,
    			SafetyStockQuantityInBaseUnit:             		data.SafetyStockQuantityInBaseUnit,
    			SafetyDuration:                            		data.SafetyDuration,
    			MaximumStockQuantityInBaseUnit:            		data.MaximumStockQuantityInBaseUnit,
    			MinimumDeliveryQuantityInBaseUnit:         		data.MinimumDeliveryQuantityInBaseUnit,
    			MinimumDeliveryLotSizeQuantityInBaseUnit:  		data.MinimumDeliveryLotSizeQuantityInBaseUnit,
    			StandardDeliveryLotSizeQuantityInBaseUnit: 		data.StandardDeliveryLotSizeQuantityInBaseUnit,
    			DeliveryLotSizeRoundingQuantityInBaseUnit: 		data.DeliveryLotSizeRoundingQuantityInBaseUnit,
    			MaximumDeliveryLotSizeQuantityInBaseUnit:  		data.MaximumDeliveryLotSizeQuantityInBaseUnit,
    			MaximumDeliveryQuantityInBaseUnit:         		data.MaximumDeliveryQuantityInBaseUnit,
    			DeliveryLotSizeIsFixed:                    		data.DeliveryLotSizeIsFixed,
    			StandardDeliveryDurationInDays:            		data.StandardDeliveryDurationInDays,
    			IsMarkedForDeletion:                       		data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToWorkScheduling() *requests.WorkScheduling {
	data := sdc.ProductMaster
	return &requests.WorkScheduling{

                Product:                        data.Product,     
                //BusinessPartner:                data.,             
                Plant:                          data.Plant,   
                ProductionInvtryManagedLoc:     data.ProductionInvtryManagedLoc,                        
                ProductProcessingTime:          data.ProductProcessingTime,                   
                ProductionSupervisor:           data.ProductionSupervisor,                  
                ProductProductionQuantityUnit:  data.ProductProductionQuantityUnit,                           
                ProdnOrderIsBatchRequired:      data.ProdnOrderIsBatchRequired,                       
                MatlCompIsMarkedForBackflush:   data.MatlCompIsMarkedForBackflush,                          
                ProductionSchedulingProfile:    data.ProductionSchedulingProfile,                         
                //IsMarkedForDeletion:            data.,                 
 
	}
}

func (sdc *SDC) ConvertToTax() *requests.Tax {
	data := sdc.ProductMaster
	return &requests.Tax{

                Product:                   data.Product,   
                Country:                   data.Country,   
                ProductTaxCategory:        data.TaxCategory,       
                ProductTaxClassification:  data.TaxClassification,                    
 
	}
}

func (sdc *SDC) ConvertToAccounting() *requests.Accounting {
	data := sdc.ProductMaster
	return &requests.Accounting{

                Product:                  data.Product,     
                //BusinessPartner:          data.,             
                Plant:                    data.Plant,   
                ValuationClass:           data.ValuationClass,            
                //CostingPolicy:            data.,           
                PriceUnitQty:             data.PriceUnitQty,          
                StandardPrice:            data.StandardPrice,           
                MovingAveragePrice:       data.MovingAveragePrice,                
                PriceLastChangeDate:      data.PriceLastChangeDate,                 
                IsMarkedForDeletion:      data.IsMarkedForDeletion,                 
 
	}
}

func (sdc *SDC) ConvertToAllergen() *requests.Allergen {
	data := sdc.ProductMaster
	return &requests.Allergen{

                Product:                  data.Product,     
                //BusinessPartner:          data.,             
                Allergen:                 data.Allergen,   
                AllergenIsContained:      data.AllergenIsContained,            
 
	}
}

func (sdc *SDC) ConvertToNutritional() *requests.Nutritional {
	data := sdc.ProductMaster
	return &requests.Nutritional{

                Product:                  data.Product,     
                //BusinessPartner:          data.,             
                Nutrient:                 data.Nutrient,   
                NutrientContent:          data.NutrientContent, 
                NutrientContentUnit:      data.NutrientContentUnit,
 
	}
}

func (sdc *SDC) ConvertToCalories() *requests.Calories {
	data := sdc.ProductMaster
	return &requests.Calories{

                Product:                  data.Product,     
                //BusinessPartner:          data.,             
                CaloryUnitQuantitiy:      data.CaloryUnitQuantitiy,   
                Calories:                 data.Calories, 
                CaloryUnit:               data.CaloryUnit,
 
	}
}
