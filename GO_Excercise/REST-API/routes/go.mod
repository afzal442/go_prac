module REST-API/routes

go 1.16

require (
	REST-API/handlers v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.4
)

replace REST-API/handlers => ../handlers
