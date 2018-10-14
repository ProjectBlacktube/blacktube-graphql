package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Video struct {
	ID          int       `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Duration    int       `json:"duration" db:"duration"`
	Key         string    `json:"key" db:"key"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Owner       int       `json:"owner" db:"owner"`
}

type VideoNested struct {
	ID          int       `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Duration    int       `json:"duration" db:"duration"`
	Key         string    `json:"key" db:"key"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Owner       User      `json:"owner"`
}

// String is not required by pop and may be deleted
func (v Video) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Videos is not required by pop and may be deleted
type Videos []Video

// String is not required by pop and may be deleted
func (v Videos) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (v *Video) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: v.Duration, Name: "Duration"},
		&validators.IntIsPresent{Field: v.Owner, Name: "Owner"},
		&validators.StringIsPresent{Field: v.Key, Name: "Key"},
		&validators.StringIsPresent{Field: v.Title, Name: "Title"},
		&validators.StringIsPresent{Field: v.Description, Name: "Description"},
	), nil
}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (v *Video) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (v *Video) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
