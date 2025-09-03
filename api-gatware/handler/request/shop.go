package request

type ShopReq struct {
	Page int64 `gorm:"page" json:"page"`
	Size int64 `gorm:"size" json:"size"`
}
