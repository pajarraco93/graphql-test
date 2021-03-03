package graphql

import "github.com/graphql-go/graphql"

var Group = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Group",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"genre": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var Album = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Album",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"composedBy": &graphql.Field{
				Type: Group,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var Song = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Song",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"appearsIn": &graphql.Field{
				Type: Album,
			},
		},
	},
)

var NewGroup = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "NewGroup",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"genre": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

var NewAlbum = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "NewAlbum",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"composedBy": &graphql.InputObjectFieldConfig{
				Type: NewGroup,
			},
			"year": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)

var NewSong = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "NewSong",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"appearsIn": &graphql.InputObjectFieldConfig{
				Type: NewAlbum,
			},
		},
	},
)

type Schema struct {
	libraryResolver Resolver
}

func (s Schema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"AllGroups": &graphql.Field{
				Type:        graphql.NewList(Group),
				Description: "List all groups",
				Resolve:     s.libraryResolver.AllGroups,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}

func (s Schema) Mutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"CreateGroup": &graphql.Field{
				Type:        Group,
				Description: "Create a new group",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: NewGroup,
					},
				},
				Resolve: s.libraryResolver.CreateGroup,
			},
			"CreateAlbum": &graphql.Field{
				Type:        Album,
				Description: "Create a new album",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: NewAlbum,
					},
				},
				Resolve: s.libraryResolver.CreateAlbum,
			},
			"CreateSong": &graphql.Field{
				Type:        Song,
				Description: "Create a new song",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: NewSong,
					},
				},
				Resolve: s.libraryResolver.CreateSong,
			},
		},
	}
	return graphql.NewObject(objectConfig)
}

func NewSchema(libraryResolver Resolver) Schema {
	return Schema{
		libraryResolver: libraryResolver,
	}
}
