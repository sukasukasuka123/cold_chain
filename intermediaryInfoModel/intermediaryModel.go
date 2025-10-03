package intermediaryInfoModel

import "time"

// 入库信息（包括物流单号、冷库位置、温度、湿度、图片信息）
type InboundInfo struct {
	LogisticsNum string    `json:"logistics_num"`
	StoragePlace string    `json:"storage_place"`
	Temp         float64   `json:"temp"`
	Humidity     float64   `json:"humidity"`
	ImageInfo    []byte    `json:"image_info"`
	InboundDate  time.Time `json:"inbound_date"`
}
