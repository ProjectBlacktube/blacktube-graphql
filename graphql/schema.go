package graphql

import (
	"github.com/graphql-go/graphql"
  "github.com/graphql-go/graphql/gqlerrors"
)

var userType *graphql.Object
var queryType *graphql.Object
// var userInputType *graphql.InputObject
var Schema graphql.Schema

func init() {
	userType = graphql.NewObject(graphql.ObjectConfig{
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

  // userInputType = graphql.NewInputObject(graphql.InputObjectConfig{
  //   Name: "User",
  //   Fields: graphql.Fields{
  //     "id": &graphql.Field{
  //       Type: graphql.Int,
  //     },
  //     "name": &graphql.Field{
  //       Type: graphql.String,
  //     },
  //     "password": &graphql.Field{
  //       Type: graphql.String,
  //     },
  //   },
  // })

	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "QueryUser",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: graphql.NewList(userType),
        Args: graphql.FieldConfigArgument{
          "id": &graphql.ArgumentConfig{
            Type: graphql.String,
          },
        },
        Description: "List of user(s)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          idQuery, isOK := p.Args["id"].(int)
          if isOK {
            return findUser(idQuery)
          }
          return allUsers()
				},
			},
		},
	})

  mutationType := graphql.NewObject(graphql.ObjectConfig{
    Name: "MutationUser",
    Fields: graphql.Fields{
      "createUser": &graphql.Field{
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
          
          user := User{
            Name: name,
            Password: pass,
          }
          return newUser(user)
        },
      },

      "updateUser": &graphql.Field{
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
          user, _ := findUser(idQuery)
          if !isOk {
           return nil, gqlerrors.Error{}
          }

          user.Name = name
          user.Password = pass
          return updateUser(user)
        },
      },
    },
  })

  Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
    Mutation: mutationType,
	})

}

