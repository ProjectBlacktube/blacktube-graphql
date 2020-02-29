package manager

import (
	"log"

	"github.com/ProjectBlacktube/blacktube-graphql/models"
	"github.com/gobuffalo/pop"
)

type UserQueryManager struct {
	Db *pop.Connection
}

func (manager *UserQueryManager) AllUsers() (models.Users, error) {
	users := models.Users{}
	query := pop.Q(manager.Db)

	err := query.All(&users)
	if err != nil {
		log.Panic(err)
	}

	return users, err
}

func (manager *UserQueryManager) FindUser(id string) (*models.User, error) {
	user := &models.User{}
	err := manager.Db.Find(&user, id)
	if err != nil {
		log.Panic(err)
	}

	return user, err
}

func (manager *UserQueryManager) NewUser(newUser *models.NewUser) (*models.User, error) {
	user := &models.User{
		Name:     newUser.Name,
		Password: newUser.Password,
	}
	err := manager.Db.Save(&user)
	if err != nil {
		log.Panic(err)
	}

	return user, err
}

func (manager *UserQueryManager) UpdateUser(user *models.User) (*models.User, error) {
	err := manager.Db.Update(&user)
	if err != nil {
		log.Panic(err)
	}

	return user, err
}

func (manager *UserQueryManager) DeleteUser(id string) (*models.User, error) {
	user, err := manager.FindUser(id)
	if err != nil {
		log.Panic(err)
	}

	err = manager.Db.Destroy(&user)
	if err != nil {
		log.Panic(err)
	}

	return user, err
}
