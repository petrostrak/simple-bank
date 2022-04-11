#### Simple Bank
A backend application in GO that simulates simple bank functionality

##### To install golang migrate tool in debian
`https://github.com/golang-migrate/migrate/releases/tag/v4.15.1`

To create initial migrate files we run the following command
`migrate create -ext sql -dir db/migration -seq init_schema`