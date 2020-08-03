## Features

- Generate grapql with golang
- Support graphql connect with RESTapi
- Support graphql connect with mongodb

## Getting Started

##### clone this repository:
>git clone https://git.touchdevops.com/kanokp/graphql-gen.git

##### define schema: edit schema on schema.graphqls
```other
input PaginationInput {
    limit: Int
    offset: Int
}

input createInput  {
    name: String!
}

input updateInput  {
    id: String!
    name: String!
}

type Response  {
    id: String!
    name: String!
}

type ResponseList {
    total: Int
    response: [Response]!
}

type Pagination {
    total: Int
    nextPage: String
    previousPage: String
}

type Query {
    list(input: PaginationInput): ResponseList!
    read(input: ID!): Response!
}

type Mutation {
    create(input: createInput): Response!
    update(input: updateInput): Response!
    delete(input: String!): Response!
}
```

##### generate model:
> gqlgen generate

- graph/model/models_gen.go

```go
    package model
    
    type Pagination struct {
        Total        *int    `json:"total"`
        NextPage     *string `json:"nextPage"`
        PreviousPage *string `json:"previousPage"`
    }
```

- graph/schema.resolvers.go: implement resolvers

```go
    func (r *mutationResolver) Create(ctx context.Context, input *model.CreateInput) (*model.Response, error) {
        return r.Repo.Create(ctx, input)
    }
    
    func (r *mutationResolver) Update(ctx context.Context, input *model.UpdateInput) (*model.Response, error) {
        return r.Repo.Update(ctx, input)
    }
```
##### implement server
- choose plugin: REST API, mongoDB

```go

    conf := config.Get()
    repo := store.New(conf.MongoDBEndpoint, conf.MongoDBName, conf.MongoDBTableName)
    //repo := store.NewClient(client.New(), conf.HOST, conf.PATH)
    
    resolver := &graph.Resolver{Repo: repo}
    gqlConfig := generated.Config{Resolvers: resolver}
    schema := generated.NewExecutableSchema(gqlConfig)
    srv := handler.NewDefaultServer(schema)     

```

## Environment
[Environment file](graph/development/local.env)
```env
    MONGODB_ENDPOINT="mongodb://touch:secret@localhost:27017"     
    MONGODB_TABLE_NAME="company"                                  
    MONGODB_NAME="go_project"                                     
    PORT="8080"                                                   
    HOST="http://127.0.0.1:8081"                                  
    PATH_URL="company"               
```
            
## Running Project
> go run server.go

<img src="https://media.touchdevops.com/file/BjtDJCnlwX-Screen%20Shot%202563-07-30%20at%2017.10.12.png" width=500 alt="graphQL playground">


## sequence diagram
graphql with database
![](sequnce/database.png)

graphql with restful
![](sequnce/restful.png)