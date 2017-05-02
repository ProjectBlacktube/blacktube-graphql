package models

import (
  "encoding/json"
  "time"

  "github.com/markbates/pop"
  "github.com/markbates/validate"
  "github.com/markbates/validate/validators"
)

type User struct {
  ID        int       `json:"id" db:"id"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
  Name      string    `json:"name" db:"name"`
  Password  string    `json:"password" db:"password"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
  ju, _ := json.Marshal(u)
  return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
  ju, _ := json.Marshal(u)
  return string(ju)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
  return validate.Validate(
    &validators.StringIsPresent{Field: u.Name, Name: "Name"},
    &validators.StringIsPresent{Field: u.Password, Name: "Password"},
  ), nil
}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (u *User) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
  return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
  return validate.NewErrors(), nil
}
