package model

type PageModel struct {
  Page int `json:"page,omitempty" v:"gt=0"`    // 页码
  Size int `json:"size,omitempty" v:"lte=100"` // 数量
}
