package graphql

import (
  "log"

  "github.com/icalF/blacktube-graphql/models"
  "github.com/markbates/pop"
)

type Users models.Users
type User models.User

func allUsers() (Users, error) {
  db, err := pop.Connect("development")
  if err != nil {
    log.Panic(err)
  }

  users := Users{}
  query := pop.Q(db)
  err = query.All(&users)
  if err != nil {
    log.Panic(err)
  }

	return users, err
}

func findUser(id int) (User, error) {
  db, err := pop.Connect("development")
  if err != nil {
    log.Panic(err)
  }

  user := User{}
  err = db.Find(&user, id)
  if err != nil {
    log.Panic(err)
  }

  return user, err
}

func newUser(user User) (User, error) {
  db, err := pop.Connect("development")
  if err != nil {
    log.Panic(err)
  }

  err = db.Save(&user)
  if err != nil {
    log.Panic(err)
  }

  return user, err
}

func updateUser(user User) (User, error) {
  db, err := pop.Connect("development")
  if err != nil {
    log.Panic(err)
  }

  err = db.Update(&user)
  if err != nil {
    log.Panic(err)
  }

  return user, err
}

