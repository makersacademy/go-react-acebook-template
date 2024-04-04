## Acebook Go/React Template

In this project, you are task with working on an existing application. A
significant part of the challenge will be to familiarise yourself with the
codebase you've inherited, as you work to **improve and extend** it.

### Structure

This repo contains two applications:

- A frontend React App
- A backend api server

These two applications will communicate through HTTP requests, and need to be
run separately.

### Documentation

[More documentation of the codebase and its architecture can be found here.](./DOCUMENTATION.md)
It's recommended you all read the suggested docs _after making sure the whole
setup below worked for everyone_. Then work together on a diagram describing how
the application works.

### Card wall

https://trello.com/invite/b/YyHEoIyD/ATTIe230926302b5d8daad7ead41ef861018888CCDA1/xenon-acebook-meta

### Quickstart

### Install Node.js

If you haven't already, make sure you have node and NVM installed.

1. Install Node Version Manager (NVM)
   ```
   brew install nvm
   ```
   Then follow the instructions to update your `~/.bash_profile`.
2. Open a new terminal
3. Install the latest version of [Node.js](https://nodejs.org/en/), (`20.5.0` at
   time of writing).
   ```
   nvm install 20
   ```

### Install Go

Follow the instructions here: https://go.dev/doc/install

### Set up your project

1. Have one team member fork this repository
2. Rename the fork to `acebook-<team name>`
3. Every team member clone the fork to their local machine
4. Install dependencies for both the `frontend` and `api` applications:
   ```
   cd frontend
   npm install
   cd ../api
   go get .
   ```
5. Install an ESLint plugin for your editor, for example
   [ESLint for VSCode](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)
6. Start Postgresql

   ```
   brew services start postgresql
   ```
7. Create your databases:
   ```
   createdb acebook
   createdb acebook_test
   ```

### Setting up environment variables.

We need to create two `.env` files, one in the frontend and one in the api.

#### Frontend

Create a file `frontend/.env` with the following contents:

```
VITE_BACKEND_URL="http://localhost:8082"
```

#### Backend

Create a file `api/.env` with the following contents:

```
POSTGRES_URL="postgresql://localhost:5432/acebook"
JWT_SECRET="secret"
```

For an explanation of these environment variables, see the documentation.

### How to run the server and use the app

1. Start the server application (in the `api` directory) in dev mode:

```
; cd api
; go run main.go
```

2. Start the front end application (in the `frontend` directory)

In a new terminal session...

```
; cd frontend
; npm run dev
```

You should now be able to open your browser and go to the
`http://localhost:5173/signup` to create a new user.

Then, after signing up, you should be able to log in by going to
`http://localhost:5173/login`.

After logging in, you won't see much but you can create posts using PostMan and
they should then show up in the browser if you refresh the page.
