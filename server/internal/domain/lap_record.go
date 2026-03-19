package domain

import "time"

type LapRecord struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	ActivityID        uint      `gorm:"index" json:"activity_id"`
	TrackID           uint      `gorm:"index" json:"track_id"`
	UserID            uint      `gorm:"index" json:"user_id"`
	CarModel          string    `gorm:"size:50" json:"car_model"`
	Horsepower        int       `json:"horsepower"`
	Powertrain        string    `gorm:"size:50" json:"powertrain"`
	ModificationLevel int       `gorm:"default:0" json:"modification_level"`
	TireModel         string    `gorm:"size:50" json:"tire_model"`
	LapTimeMs         int64     `json:"lap_time_ms"`
	Temperature       int       `json:"temperature"`
	RecordedAt        time.Time `json:"recorded_at"`
	RecordedBy        uint      `json:"recorded_by"`
	CreatedAt         time.Time `json:"created_at"`

	Track    *Track    `gorm:"foreignKey:TrackID" json:"track,omitempty"`
	User     *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Activity *Activity `gorm:"foreignKey:ActivityID" json:"activity,omitempty"`
}
