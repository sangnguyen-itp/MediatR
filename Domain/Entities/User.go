package Entities

import "mediatR/Domain/Common"

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       int64  `json:"dob"`
	Common.BaseEntity
	Common.AuditableBaseEntity
}
