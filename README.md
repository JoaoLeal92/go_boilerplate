# go_boilerplate
Backend Boilerplate using go, containing routes for user creation and authentication.
This project assumes you have a postgresql instance running on docker.

# Run project locally

 - Clone repo
 - Create a `Makefile` using `Makefile.example` as a model, filling in local credentials
 - Create a `config.toml` file using `config.example.toml` as a model
 
 ## Creating database
 
 Run the following command
 
 ```bash
 $ make createdb
 ```
