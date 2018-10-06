package graphql

import (
	"log"

	"github.com/gobuffalo/pop"
	"github.com/koneko096/blacktube-graphql/models"
)

type UserQueryManager struct {
	Db *pop.Connection
}

func (manager *UserQueryManager) allUsers() (models.Users, error) {
	users := models.Users{}
	query := pop.Q(manager.Db)

	err := query.All(&users)
	if err != nil {
		log.Panic(err)
	}

	return users, err
}

func (manager *UserQueryManager) findUser(id int) (models.User, error) {
	user := models.User{}
	err := manager.Db.Find(&user, id)
	if err != nil {
		log.Panic(err)
	}

	return user, err
}

func (manager *UserQueryManager) newUser(user models.User) (models.User, error) {
	err := manager.Db.Save(&user)
	if err != nil {
		log.Panic(err)
	}

	return user, err
}

func (manager *UserQueryManager) updateUser(user models.User) (models.User, error) {
	err := manager.Db.Update(&user)
	if err != nil {
		log.Panic(err)
	}

	return user, err
}
