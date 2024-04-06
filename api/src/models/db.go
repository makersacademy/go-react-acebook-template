package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	// provides the PostgreSQL database driver for GORM, allowing communication
	// with a PostgreSQL database.
	"gorm.io/gorm"
	// the GORM package, that abstracts away the low-level SQL operations and
	// provides a higher-level API for working with databases. GORM allows for
	// its own internal structs to be embedded into the code structs.
	// gorm.Model for example contains fields that store info about creation
	// updating or deletion of the struct containing it.
	"os"
)

var Database *gorm.DB

// declares a global variable Database of type *gorm.DB, which will hold
// the connection to the database whenever the application is running.
// It's accessible from other files within the models package so that
// it can be referred to and used by the methods within the Post and User
// models to manipulate the Database. User and Post models define the structs
// and the outline of the Database tables.

func OpenDatabaseConnection() {
	// function responsible for establishing a connection to the PostgreSQL database.
	connection_string := os.Getenv("POSTGRES_URL")
	// This code retrieves the PostgreSQL connection string from the environment variable
	// POSTGRES_URL = "PROTOCOL(postgres) :// HOST_IP (localhost and port) / DATABASE_NAME"

	// If we were using a cloud database like the Render Database service for PostgreSQL,
	// the connection string would typically contain additional information such as the
	// username, password, and SSL/TLS configuration.
	// The above line would look like this instead:
	// connection_string := os.Getenv("POSTGRES_URL_EXTERNAL")
	// and it would refer to a slightly different environment variable:
	// POSTGRES_URL_EXTERNAL = "PROTOCOL (postgres) :// USERNAME : PASSWORD @ HOSTNAME . frankfurt-postgres.render.com / DATABASE NAME"

	fmt.Println(connection_string)
	// prints the connection string on the backend terminal for debugging purposes

	var err error
	// err variable will capture any errors that might occur in the following line.
	Database, err = gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	// Attempts to open a connection to the PostgreSQL database using the gorm.Open() function.
	// It passes the PostgreSQL connection string obtained earlier and an
	// empty configuration object (&gorm.Config{}).
	// The configuration object indicates that the database connection is being established with
	// DEFAULT configurations, which can still be customized later on.
	// The Database variable is reassigned with the newly created DB Connection
	// and it's now ready to be used by the User and Post models.

	if err != nil {
		panic(err)
	}
	// if an error occurs during the creation of a DB connection the program panics and
	// terminates at this stage, returning the error message.

	fmt.Println("Successfully connected to database")
	// if no error occurs the above message is printed to the terminal
}

func AutoMigrateModels() {
	Database.AutoMigrate(&User{})
	Database.AutoMigrate(&Post{})
	Database.AutoMigrate(&Comment{})
	// This is the 4th function that runs when main is launched (1st the environment variables
	// are loaded, 2nd the app is setup and 3rd the database connection is established).
	// This function is responsible for automatically migrating (creating or updating)
	// the database tables based on the defined models.
	// This means that inspects the struct's fields and generates the corresponding table structure
	// in the database. Each field in the struct typically corresponds to a column in the database table
	// and GORM maps the Go data types to appropriate database column types.
	// If the table already exists, modifications might be performed if the stuct architecture has changed.
	// It passes pointers to instances of the User and Post structs, which represent
	// the models for user and post data.
}
