package domain

import "time"

type Role string

const (
	RoleUser       Role = "user"
	RoleCEO        Role = "ceo"
	RoleTechnician Role = "technician"
	RoleShareholder Role = "shareholder"
	RoleSupplier   Role = "supplier"
	RoleSuperAdmin Role = "super_admin"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OpenID    string    `gorm:"uniqueIndex;size:128" json:"-"`
	BOpenID   string    `gorm:"size:128" json:"-"`
	UnionID   string    `gorm:"size:128" json:"-"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Nickname  string    `gorm:"size:64" json:"nickname"`
	AvatarURL string    `gorm:"size:512" json:"avatar_url"`
	Role      Role      `gorm:"size:20;default:user" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	MemberProfile *MemberProfile `gorm:"foreignKey:UserID" json:"member_profile,omitempty"`
}

type MemberProfile struct {
	ID                   uint   `gorm:"primaryKey" json:"id"`
	UserID               uint   `gorm:"uniqueIndex" json:"user_id"`
	LevelID              uint   `json:"level_id"`
	Balance              int64  `gorm:"default:0" json:"balance"`
	TotalConsumed12M     int64  `gorm:"default:0" json:"total_consumed_12m"`
	IsShareholder        bool   `gorm:"default:false" json:"is_shareholder"`
	ShareholderDiscount  int    `gorm:"default:0" json:"shareholder_discount"`
	UpdatedAt            time.Time `json:"updated_at"`

	Level *MemberLevel `gorm:"foreignKey:LevelID" json:"level,omitempty"`
}

type MemberLevel struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Code       string `gorm:"uniqueIndex;size:10" json:"code"`
	Name       string `gorm:"size:20" json:"name"`
	Threshold  int64  `gorm:"default:0" json:"threshold"`
	Discount   int    `gorm:"default:100" json:"discount"`
	Privileges JSON   `gorm:"type:jsonb" json:"privileges"`
	SortOrder  int    `gorm:"default:0" json:"sort_order"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type BalanceLog struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"index" json:"user_id"`
	Type           string    `gorm:"size:20" json:"type"`
	Amount         int64     `json:"amount"`
	BalanceAfter   int64     `json:"balance_after"`
	RelatedOrderID *uint     `json:"related_order_id,omitempty"`
	Remark         string    `gorm:"size:255" json:"remark"`
	CreatedAt      time.Time `json:"created_at"`
}
