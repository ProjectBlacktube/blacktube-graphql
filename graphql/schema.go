package graphql

import (
	"crypto/sha1"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/koneko096/blacktube-graphql/models"
	"log"
)

func BuildSchema(userManager *UserQueryManager, videoManager *VideoQueryManager) *graphql.Schema {
	userType := createUserType()

	videoType := createVideoType(userType)

	queryType := createQueryType(userType, userManager, videoType, videoManager)

	createUserMutation := createUserCreationMutation(userType, userManager)

	updateUserMutation := createUserUpdateMutation(userType, userManager)

	createVideoMutation := createVideoCreationMutation(videoType, videoManager)

	updateVideoMutation := createVideoUpdateMutation(videoType, videoManager)

	mutationType := createMutationType(createUserMutation, updateUserMutation, createVideoMutation, updateVideoMutation)

	schema, err := createSchema(queryType, mutationType)
	if err != nil {
		log.Panic(err)
	}

	return &schema
}

func createSchema(queryType *graphql.Object, mutationType *graphql.Object) (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}

func createVideoUpdateMutation(videoType *graphql.Object, videoManager *VideoQueryManager) *graphql.Field {
	return &graphql.Field{
		Type: videoType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"duration": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Description: "Update existing video",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			idQuery, isOk := p.Args["id"].(int)
			title, _ := p.Args["title"].(string)
			desc, _ := p.Args["description"].(string)
			dur, _ := p.Args["duration"].(int)

			v, _ := videoManager.findVideo(idQuery)
			video, _ := videoManager.fromNested(v)
			if !isOk {
				return nil, gqlerrors.Error{}
			}

			video.Title = title
			video.Duration = dur
			video.Description = desc

			return videoManager.updateVideo(video)
		},
	}
}

func createVideoCreationMutation(videoType *graphql.Object, videoManager *VideoQueryManager) *graphql.Field {
	return &graphql.Field{
		Type: videoType,
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"duration": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"ownerId": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Description: "Add new video",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			title, isOk := p.Args["title"].(string)
			if !isOk {
				return nil, gqlerrors.Error{}
			}

			desc, isOk := p.Args["description"].(string)
			if !isOk {
				return nil, gqlerrors.Error{}
			}

			dur, isOk := p.Args["duration"].(int)
			if !isOk {
				return nil, gqlerrors.Error{}
			}

			owner, isOk := p.Args["ownerId"].(int)
			if !isOk {
				return nil, gqlerrors.Error{}
			}

			key := fmt.Sprintf("%x", sha1.Sum([]byte(title)))

			video := models.Video{
				Title:       title,
				Key:         string(key),
				Description: desc,
				Duration:    dur,
				Owner:       owner,
			}

			return videoManager.newVideo(video)
		},
	}
}

func createUserUpdateMutation(userType *graphql.Object, userManager *UserQueryManager) *graphql.Field {
	return &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Description: "Update existing user",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			idQuery, isOk := p.Args["id"].(int)
			if !isOk {
				return nil, gqlerrors.Error{}
			}

			name, _ := p.Args["name"].(string)
			pass, _ := p.Args["password"].(string)

			user, _ := userManager.findUser(idQuery)
			if !isOk {
				return nil, gqlerrors.Error{}
			}

			user.Name = name
			user.Password = pass

			return userManager.updateUser(user)
		},
	}
}

func createUserCreationMutation(userType *graphql.Object, userManager *UserQueryManager) *graphql.Field {
	return &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Description: "Add new user",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			name, isOk := p.Args["name"].(string)
			if !isOk {
				return nil, gqlerrors.Error{}
			}

			pass, isOk := p.Args["password"].(string)
			if !isOk {
				return nil, gqlerrors.Error{}
			}

			user := models.User{
				Name:     name,
				Password: pass,
			}
			return userManager.newUser(user)
		},
	}
}

func createMutationType(
	createUserMutation *graphql.Field,
	updateUserMutation *graphql.Field,
	createVideoMutation *graphql.Field,
	updateVideoMutation *graphql.Field,
) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": createUserMutation,

			"updateUser": updateUserMutation,

			"createVideo": createVideoMutation,

			"updateVideo": updateVideoMutation,
		},
	})
}

func createQueryType(
	userType *graphql.Object,
	userManager *UserQueryManager,
	videoType *graphql.Object,
	videoManager *VideoQueryManager,
) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: graphql.NewList(userType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Description: "List of user(s)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(int)
					if isOK {
						return userManager.findUser(idQuery)
					}
					return userManager.allUsers()
				},
			},

			"video": &graphql.Field{
				Type: graphql.NewList(videoType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Description: "List of video(s)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(int)
					if isOK {
						return videoManager.findVideo(idQuery)
					}
					return videoManager.allVideos()
				},
			},
		},
	})
}

func createUserType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func createVideoType(userType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Video",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"key": &graphql.Field{
				Type: graphql.String,
			},
			"owner": &graphql.Field{
				Type: userType,
			},
			"duration": &graphql.Field{
				Type: graphql.Int,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}
