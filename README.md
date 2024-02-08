Warden
===
### Development Environment Setup
Start by installing Go, PostgreSQL, Redis and Mailhog on your system if they are not already available. While Mailhog is optional, the environment settings will need to be adjusted if you decide to use another mail delivery service during development. For example, using homebrew:
```
brew update && brew install go postgresql redis mailhog
```

Make sure the PostgreSQL, Redis and Mailhog services are running:
```
brew services start postgresql
brew services start redis
brew services start mailhog
```

Copy the provided `.env.example` file to `.env`. In order to connect to the database, you will need to replace `<YOUR_USERNAME>` with the name of the user you log into your system with. This is of course assuming that you used Homebrew to install PostgreSQL. Other than that, everything should work out of the box with the default values.

Next, install the module dependencies:
```
go mod download
```

Create the development database and run the data migrations:
```
createdb warden 
go run ./cmd/db upgrade
```

Start the server:
```
go run ./cmd/server
```
