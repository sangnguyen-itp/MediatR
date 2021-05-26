package Common

type BaseEntity struct {
	ID string `json:"id" gorm:"primaryKey"`
}