package graphql

import (
  "log"

  "github.com/markbates/pop"
  "github.com/icalF/blacktube-graphql/models"
)

func allUsers() (models.Users, error) {
  db, err := pop.Connect("development")
  if err != nil {
    log.Panic(err)
  }

  users := models.Users{}
  query := pop.Q(db)
  err = query.All(&users)
  if err != nil {
    log.Panic(err)
  }

  return users, err
}

func findUser(id int) (models.User, error) {
  db, err := pop.Connect("development")
  if err != nil {
    log.Panic(err)
  }

  user := models.User{}
  err = db.Find(&user, id)
  if err != nil {
    log.Panic(err)
  }

  return user, err
}

func newUser(user models.User) (models.User, error) {
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

func updateUser(user models.User) (models.User, error) {
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
