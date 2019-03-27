# Carlyboard

App backend for a simple idea board 

By Carly Ilg (https://github.com/cilg017) and Nate Vaughan (https://github.com/natevaughan/)

## Dependencies

 - Go 
 - Mysql

## Setup Instructions

Add the carlyboard mysql database by running `mysql -u<YOUR_USER> -p < database.sql`, replacing `<YOUR_USER>` with an appropriate mysql user and entering your password when prompted.

Copy appconfig.yml.sample to appconfig.yml:
`cp appconfig.yml.sample appconfig.yml`

Edit `appconfig.yml` with your mysql connection string, replacing "root" with your mysql user and adding a password after the colon.

Run the app:

`go run carlyboard.go db.go appconfig.go`
