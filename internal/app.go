package internal

import (
	"application/internal/domain/user"
	"application/internal/infra"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
var App = fx.New(
	fx.Provide(infra.NewServer),
	fx.Provide(infra.NewDatabase),
	fx.Invoke(func(db *gorm.DB) {}),
	fx.Invoke(func(server *infra.Server) {}),
	user.UsersModule,
)
