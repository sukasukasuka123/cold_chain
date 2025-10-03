# cold_chain struct

--- **factory need struct**
``` go

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
```
--- **intermediary need struct**

```go
// 入库信息（包括物流单号、冷库位置、温度、湿度、图片信息）
type InboundInfo struct {
	LogisticsNum string    `json:"logistics_num"`
	StoragePlace string    `json:"storage_place"`
	Temp         float64   `json:"temp"`
	Humidity     float64   `json:"humidity"`
	ImageInfo    []byte    `json:"image_info"`
	InboundDate  time.Time `json:"inbound_date"`
}
```

--- **block chain need struct**

```go
// 交易
type Transaction struct {
	TxID string      `json:"tx_id"`
	Data interface{} `json:"data"`
}

// 区块头
type BlockHeader struct {
	Index     int       `json:"index"`
	Timestamp time.Time `json:"timestamp"`
	PrevHash  string    `json:"prev_hash"`
	Hash      string    `json:"hash"`
	Nonce     int       `json:"nonce"`
}

// 区块体
type BlockBody struct {
	Transactions []Transaction `json:"transactions"`
}

// 区块
type Block struct {
	Header BlockHeader `json:"header"`
	Body   BlockBody   `json:"body"`
}
```
