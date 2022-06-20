package apis

var Department = map[string]string{

	"销售部":     "Sales",
	"供应链管理中心": "SupplyChainManagement",
	"研究与合作中心": "ResearchCollaborationCenter",
	"财务管理中心":  "FinancialManagementCenter",
	"人力资源中心":  "HumanResourceCenter",
	"产能利用中心":  "CapacityUtilizationCenter",

	"Sales":                       "销售部",
	"SupplyChainManagement":       "供应链管理中心",
	"FinancialManagementCenter":   "财务管理中心",
	"ResearchCollaborationCenter": "研究与合作中心",
	"HumanResourceCenter":         "人力资源中心",
	"CapacityUtilizationCenter":   "产能利用中心",
}

var Sales = map[string]string{
	"到账金额": "元",
	"销售收入": "元",

	"现金流贡献额": "元",
	"部门费用额":  "元",
	"部门费用率":  "%",
	"人员离职率":  "%",
}

var SupplyChainManagement = map[string]string{
	"供货准时率":   "%",
	"采购成本变动率": "%",
	"库存变动率":   "%",

	"现金流贡献额": "元",
	"部门费用额":  "元",
	"部门费用率":  "%",
	"人员离职率":  "%",
}

var ResearchCollaborationCenter = map[string]string{
	"项目数量": "个",
	"项目收入": "元",
	"售后服务": "次",
	"售后收入": "元",
	"投诉":   "次",

	"现金流贡献额": "元",
	"部门费用额":  "元",
	"部门费用率":  "%",
	"人员离职率":  "%",
}

var FinancialManagementCenter = map[string]string{
	"现金流贡献额": "元",
	"部门费用额":  "元",
	"部门费用率":  "%",
	"人员离职率":  "%",

	"财务成本": "元",
	"生成成本": "元",
	"管理费用": "元",
	"现金流":  "元",
}

var HumanResourceCenter = map[string]string{
	"现金流贡献额": "元",
	"部门费用额":  "元",
	"部门费用率":  "%",
	"人员离职率":  "%",

	"在职人员": "名",
	"薪资总额": "元",
	"离职率":  "%",
}

var CapacityUtilizationCenter = map[string]string{
	"现金流贡献额": "元",
	"部门费用额":  "元",
	"部门费用率":  "%",
	"人员离职率":  "%",

	"总产能":     "B",
	"平均每日生产量": "B",
	"产能利用率":   "%",
	"设备数量":    "台",
	"直接生产人员":  "名",
}

var Indexes = map[string]map[string]string{
	"Sales":                       Sales,
	"SupplyChainManagement":       SupplyChainManagement,
	"ResearchCollaborationCenter": ResearchCollaborationCenter,
	"FinancialManagementCenter":   FinancialManagementCenter,
	"HumanResourceCenter":         HumanResourceCenter,
	"CapacityUtilizationCenter":   CapacityUtilizationCenter,
}

var CommonIndexes = []string{
	"现金流贡献额",
	"部门费用额",
	"部门费用率",
	"人员离职率",
}

var Order = map[string][]string{
	"Sales":                       append(CommonIndexes, []string{"到账金额", "销售收入"}...),
	"SupplyChainManagement":       append(CommonIndexes, []string{"供货准时率", "采购成本变动率", "库存变动率"}...),
	"ResearchCollaborationCenter": append(CommonIndexes, []string{"项目数量", "项目收入", "售后服务", "售后收入", "投诉"}...),
	"FinancialManagementCenter":   append(CommonIndexes, []string{"财务成本", "生成成本", "管理费用", "现金流"}...),
	"HumanResourceCenter":         append(CommonIndexes, []string{"在职人员", "薪资总额", "离职率"}...),
	"CapacityUtilizationCenter":   append(CommonIndexes, []string{"总产能", "平均每日生产量", "产能利用率", "设备数量", "直接生产人员"}...),
}

