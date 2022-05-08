package src

// 配置文件
type Config struct {
	// 主机地址
	Host string
	// 实例UUID
	UUID string
	// 实例TOKEN
	Token string
}

// 解析后的数据
type EncryptoData struct {
	EncData      string
	EncDek       string
	BusinessName string
	KekVersion   int
}

type kek struct {
	ID       int64  `json:"id"`
	UUID     string `json:"uuid"`
	Version  int    `json:"version"`
	Content  string `json:"content"`
	Business int64  `json:"business"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
	Status   int    `json:"status"`
	Note     string `json:"note"`
	// 不在数据库中储存
	Algorithm string `json:"algorithm"`
}

// 业务结构体
type business struct {
	ID             int64  `json:"id"`
	Name           string `json:"name" form:"name" query:"name"`
	IsAutoRotate   int    `json:"isAutoRotate" form:"isAutoRotate" query:"isAutoRotate"`
	NextRotateDate string `json:"nextRotateDate"`
	IsStrict       int    `json:"isStrict" form:"isStrict" query:"isStrict"`
	UUID           string `json:"uuid"`
	Kek            int64  `json:"kek"`
	CreateAt       string `json:"createAt"`
	UpdateAt       string `json:"updateAt"`
	Status         int    `json:"status"`
	Note           string `json:"note" form:"note" query:"note"`
	Algorithm      string `json:"algorithm" form:"algorithm" query:"algorithm"`
	Keks           []kek  `json:"keks"`
}
