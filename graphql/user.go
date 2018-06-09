package graphql

import (
  "log"

  "github.com/gobuffalo/pop"
  "github.com/koneko096/blacktube-graphql/models"
)

type UserQueryManager struct {
  db *pop.Connection
}

func (manager *UserQueryManager) allUsers() (models.Users, error) {
  users := models.Users{}
  query := pop.Q(manager.db)

  err := query.All(&users)
  if err != nil {
    log.Panic(err)
  }

  return users, err
}

func (manager *UserQueryManager) findUser(id int) (models.User, error) {
  user := models.User{}
  err := manager.db.Find(&user, id)
  if err != nil {
    log.Panic(err)
  }

  return user, err
}

func (manager *UserQueryManager) newUser(user models.User) (models.User, error) {
  err := manager.db.Save(&user)
  if err != nil {
    log.Panic(err)
  }

  return user, err
}

func (manager *UserQueryManager) updateUser(user models.User) (models.User, error) {
  err := manager.db.Update(&user)
  if err != nil {
    log.Panic(err)
  }

  return user, err
}
