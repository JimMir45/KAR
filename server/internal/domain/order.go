package domain

import "time"

type OrderType string

const (
	OrderTypeRegistration OrderType = "registration"
	OrderTypeShop         OrderType = "shop"
	OrderTypeRecharge     OrderType = "recharge"
	OrderTypeService      OrderType = "service"
)

type PaymentStatus string

const (
	PaymentStatusUnpaid          PaymentStatus = "unpaid"
	PaymentStatusPaid            PaymentStatus = "paid"
	PaymentStatusRefundPending   PaymentStatus = "refund_pending"
	PaymentStatusRefunded        PaymentStatus = "refunded"
	PaymentStatusPartialRefunded PaymentStatus = "partial_refunded"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusWaiting   OrderStatus = "waiting"
	OrderStatusWorking   OrderStatus = "working"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID                   uint          `gorm:"primaryKey" json:"id"`
	OrderNo              string        `gorm:"uniqueIndex;size:32" json:"order_no"`
	UserID               uint          `gorm:"index" json:"user_id"`
	Type                 OrderType     `gorm:"size:20" json:"type"`
	TotalAmount          int64         `json:"total_amount"`
	DiscountAmount       int64         `gorm:"default:0" json:"discount_amount"`
	PayAmount            int64         `json:"pay_amount"`
	DepositAmount        int64         `gorm:"default:0" json:"deposit_amount"`
	PaymentMethod        string        `gorm:"size:20" json:"payment_method"`
	PaymentStatus        PaymentStatus `gorm:"size:20;default:unpaid" json:"payment_status"`
	DepositStatus        string        `gorm:"size:20;default:none" json:"deposit_status"`
	Status               OrderStatus   `gorm:"size:20;default:pending" json:"status"`
	ActivityID           *uint         `json:"activity_id,omitempty"`
	AppointmentDate      *time.Time    `json:"appointment_date,omitempty"`
	Remark               string        `gorm:"size:255" json:"remark"`
	WechatTransactionID  string        `gorm:"size:64" json:"wechat_transaction_id"`
	CreatedAt            time.Time     `json:"created_at"`
	PaidAt               *time.Time    `json:"paid_at,omitempty"`
	UpdatedAt            time.Time     `json:"updated_at"`

	User  *User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Items []OrderItem `gorm:"foreignKey:OrderID" json:"items,omitempty"`
}

type OrderItem struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	OrderID      uint      `gorm:"index" json:"order_id"`
	ProductID    *uint     `json:"product_id,omitempty"`
	ProductName  string    `gorm:"size:200" json:"product_name"`
	ProductPrice int64     `json:"product_price"`
	Quantity     int       `gorm:"default:1" json:"quantity"`
	Subtotal     int64     `json:"subtotal"`
	CreatedAt    time.Time `json:"created_at"`
}
