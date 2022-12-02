package entity

// WxUser undefined
type WxUser struct {
  Common
  Openid    string `gorm:"openid" json:"openid" mapstructure:"openid"`          // 小程序用户openid
  Nickname  string `gorm:"nickname" json:"nickname" mapstructure:"nickname"`    // 用户昵称
  Avatarurl string `gorm:"avatarurl" json:"avatarurl" mapstructure:"avatarurl"` // 用户头像
  Gender    int8   `gorm:"gender" json:"gender" mapstructure:"gender"`          // 性别   0 男  1  女  2 人妖
  Country   string `gorm:"country" json:"country" mapstructure:"country"`       // 所在国家
  Province  string `gorm:"province" json:"province" mapstructure:"province"`    // 省份
  City      string `gorm:"city" json:"city" mapstructure:"city"`                // 城市
  Language  string `gorm:"language" json:"language" mapstructure:"language"`
  Mobile    string `gorm:"mobile" json:"mobile" mapstructure:"mobile"` // 手机号码
}

// TableName 表名称
func (*WxUser) TableName() string {
  return "wx_user"
}
