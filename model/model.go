package model

import "time"

// 生产信息
type ProductionInfo struct {
	productionDate time.Time // 生产日期
	shelfLife      time.Time // 保质期
	batchNum       string    // 批次号
	manufacturer   string    // 生产厂家
}

// 基础产品信息
type BaseProductInfo struct {
	productName   string
	productBrand  string
	productSize   float64
	productWeight float64
	productNum    int
}

// 原料信息
type MaterialInfo struct {
	materialOrigin string // 原料产地
	supplier       string // 供应商
}

// 加工信息
type ProcessingInfo struct {
	processingLocation string // 加工地点
	processingMethod   string // 加工方法
}

// 物流与存储条件
type LogisticsInfo struct {
	storageConditions string  // 存储条件
	transportTemp     float64 // 运输温度
}

// 质检报告
type QualityReport struct {
	reportResult  string // 报告结果
	qualityAssess string // 质量评估
}
