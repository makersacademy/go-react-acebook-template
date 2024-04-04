# Backend Structure

The backend of this project can be found in the `api` directory.

It consists of a few parts. This page will go some way to explain the different
parts, and how they fit together.

![Diagram of the structure of the api](./diagrams/api_overview.png)

## Contents

- [main.go](#indexgo)
- [app.js](#appjs)
- [routes](#routes)
- [controllers](#controllers)
- [models](#models)
- [middleware](#middleware)
- [lib](#lib)
- [db](#db)

### main.go

This is the file we run when we want to start our server. It does two things:

1. Connect to the Postgresql database.
1. Creates our application
   1. Configures CORS settings.
   2. Sets up our api routes at `/users`, `/posts` and `/tokens`
1. Start our app, listening for incoming HTTP requests.

### routes

Our routes define the different request URLs that our api will respond to. They
also define which controllers will handle which requests. Each file in the
routes directory defines one route, with multiple endpoints.

For example, `posts.go` creates two routes:

- POST `<BACKEND_URL>/posts` to create a new post with `controllers.CreatePost`
- GET `<BACKEND_URL>/posts` to get all posts with `controllers.GetAllPosts`

### controllers

Controllers contain the functions that actually handle the requests that our app
recieves. Each controller is a function that takes one argument, the Gin
_context_. This struct has information about the request, and methods and fields
we can use to send a response.

The full list of features can be found
[here](https://pkg.go.dev/github.com/gin-gonic/gin#Context).

We can also add data to the `ctx` struct with _middleware_, such as the
`AuthenticationMiddleware` adding the `user_id` property, which can then be used
by controller functions.

We can respond to requests using the `ctx.JSON` method to send a JSON string.

### models

Our models are structs that handle interaction with the database. We create our
models using a library called Gorm. Each model is created with a
[_declaration_](https://gorm.io/docs/models.html), which defines the shape of
the data in the database, ie. what fields the entries in our database have.

These structs provide built-in methods, such as `.First`,
`.Last` and `.Find`. More info can be found
[here](https://gorm.io/docs/query.html).

To create entries in the database, we create a struct, and use the `.Save`
function:

```golang
	// testPost := models.Post{
	// 	Message: fmt.Sprintf("This is a test message created at %v!", time.Now()),
	// }
	// testPost.Save()
```


### middleware

Middleware are similar to our controller handler functions, but don't
necessarily end with a response. Instead they call a `next` method to pass on to
the next handler.

For example, the `AuthenticationMiddleware` makes sure that the request has been
sent with a vaild token. If it has, it passes on to the next handler, such as
the posts router. If it has an invalid token, it will send an error response
back to the client.


### auth

The lib directory simply contains functions which can be used from anywhere in
our app, for handling authentication.

### env

This is a small package for loading environment variables.
