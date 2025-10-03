package factoryInfoModel

import "time"

// 生产信息
type ProductionInfo struct {
	ProductionDate time.Time `json:"production_date"`
	ShelfLife      time.Time `json:"shelf_life"`
	BatchNum       string    `json:"batch_num"`
	Manufacturer   string    `json:"manufacturer"`
}

// 基础产品信息
type BaseProductInfo struct {
	ProductName   string  `json:"product_name"`
	ProductBrand  string  `json:"product_brand"`
	ProductSize   float64 `json:"product_size"`
	ProductWeight float64 `json:"product_weight"`
	ProductNum    int     `json:"product_num"`
}

// 原料信息
type MaterialInfo struct {
	MaterialOrigin string `json:"material_origin"`
	Supplier       string `json:"supplier"`
}

// 加工信息
type ProcessingInfo struct {
	ProcessingLocation string `json:"processing_location"`
	ProcessingMethod   string `json:"processing_method"`
}

// 物流与运输存储条件
type LogisticsInfo struct {
	StorageConditions string  `json:"storage_conditions"`
	TransportTemp     float64 `json:"transport_temp"`
}

// 质检报告
type QualityReport struct {
	ReportResult  string `json:"report_result"`
	QualityAssess string `json:"quality_assess"`
}

// 工厂出厂时产品全量信息（一个交易的数据载体）
type AllProductInfoInFactory struct {
	ProductionInfo  ProductionInfo  `json:"production_info"`
	BaseProductInfo BaseProductInfo `json:"base_product_info"`
	MaterialInfo    MaterialInfo    `json:"material_info"`
	ProcessingInfo  ProcessingInfo  `json:"processing_info"`
	LogisticsInfo   LogisticsInfo   `json:"logistics_info"`
	QualityReport   QualityReport   `json:"quality_report"`
}
