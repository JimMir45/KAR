package domain

import "time"

type ActivityStatus string

const (
	ActivityStatusDraft     ActivityStatus = "draft"
	ActivityStatusOpen      ActivityStatus = "open"
	ActivityStatusClosed    ActivityStatus = "closed"
	ActivityStatusFinished  ActivityStatus = "finished"
	ActivityStatusCancelled ActivityStatus = "cancelled"
)

type Track struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:100" json:"name"`
	Address     string `gorm:"size:255" json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	LayoutImage string `gorm:"size:512" json:"layout_image"`
	Status      string `gorm:"size:20;default:active" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type Activity struct {
	ID                   uint           `gorm:"primaryKey" json:"id"`
	Title                string         `gorm:"size:200" json:"title"`
	TrackID              uint           `gorm:"index" json:"track_id"`
	ActivityDate         time.Time      `json:"activity_date"`
	Fee                  int64          `json:"fee"`
	MaxParticipants      int            `json:"max_participants"`
	CurrentParticipants  int            `gorm:"default:0" json:"current_participants"`
	RegistrationDeadline time.Time      `json:"registration_deadline"`
	CoverImage           string         `gorm:"size:512" json:"cover_image"`
	SponsorID            *uint          `json:"sponsor_id,omitempty"`
	Schedule             JSON           `gorm:"type:jsonb" json:"schedule"`
	Rules                JSON           `gorm:"type:jsonb" json:"rules"`
	Notices              JSON           `gorm:"type:jsonb" json:"notices"`
	FlagGuideImage       string         `gorm:"size:512" json:"flag_guide_image"`
	Status               ActivityStatus `gorm:"size:20;default:draft" json:"status"`
	CreatedBy            uint           `json:"created_by"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`

	Track   *Track   `gorm:"foreignKey:TrackID" json:"track,omitempty"`
	Sponsor *Sponsor `gorm:"foreignKey:SponsorID" json:"sponsor,omitempty"`
}

type RegistrationStatus string

const (
	RegStatusRegistered    RegistrationStatus = "registered"
	RegStatusCancelled     RegistrationStatus = "cancelled"
	RegStatusRefundPending RegistrationStatus = "refund_pending"
	RegStatusRefunded      RegistrationStatus = "refunded"
)

type Registration struct {
	ID              uint               `gorm:"primaryKey" json:"id"`
	ActivityID      uint               `gorm:"index" json:"activity_id"`
	UserID          uint               `gorm:"index" json:"user_id"`
	RegistrantName  string             `gorm:"size:50" json:"registrant_name"`
	RegistrantPhone string             `gorm:"size:20" json:"registrant_phone"`
	CarPlate        string             `gorm:"size:20" json:"car_plate"`
	CarModel        string             `gorm:"size:50" json:"car_model"`
	OrderID         *uint              `json:"order_id,omitempty"`
	Status          RegistrationStatus `gorm:"size:20;default:registered" json:"status"`
	Source          string             `gorm:"size:10;default:online" json:"source"`
	Remark          string             `gorm:"size:255" json:"remark"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`

	Activity *Activity `gorm:"foreignKey:ActivityID" json:"activity,omitempty"`
	User     *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type Sponsor struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"size:100" json:"name"`
	LogoURL string `gorm:"size:512" json:"logo_url"`
	LinkURL string `gorm:"size:512" json:"link_url"`
	Status  string `gorm:"size:20;default:active" json:"status"`
}
