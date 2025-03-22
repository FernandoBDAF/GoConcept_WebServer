// Package all binds all the routes into the specified app.
package all

import (
	"github.com/fernandobdaf/GoConcept_WebServer/app/domain/checkapp"
	// "github.com/fernandobdaf/GoConcept_WebServer/app/domain/homeapp"
	// "github.com/fernandobdaf/GoConcept_WebServer/app/domain/productapp"
	// "github.com/fernandobdaf/GoConcept_WebServer/app/domain/rawapp"
	// "github.com/fernandobdaf/GoConcept_WebServer/app/domain/tranapp"
	// "github.com/fernandobdaf/GoConcept_WebServer/app/domain/userapp"
	// "github.com/fernandobdaf/GoConcept_WebServer/app/domain/vproductapp"
	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/mux"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/domain/homebus"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/domain/homebus/stores/homedb"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/domain/productbus"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/domain/productbus/stores/productdb"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/domain/userbus"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/domain/userbus/stores/usercache"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/domain/userbus/stores/userdb"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/domain/vproductbus"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/domain/vproductbus/stores/vproductdb"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/sdk/delegate"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {

	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	// delegate := delegate.New(cfg.Log)
	// userBus := userbus.NewBusiness(cfg.Log, delegate, usercache.NewStore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB), time.Minute))
	// productBus := productbus.NewBusiness(cfg.Log, userBus, delegate, productdb.NewStore(cfg.Log, cfg.DB))
	// homeBus := homebus.NewBusiness(cfg.Log, userBus, delegate, homedb.NewStore(cfg.Log, cfg.DB))
	// vproductBus := vproductbus.NewBusiness(vproductdb.NewStore(cfg.Log, cfg.DB))

	checkapp.Routes(app, checkapp.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	// homeapp.Routes(app, homeapp.Config{
	// 	Log:        cfg.Log,
	// 	HomeBus:    homeBus,
	// 	AuthClient: cfg.AuthClient,
	// })

	// productapp.Routes(app, productapp.Config{
	// 	Log:        cfg.Log,
	// 	ProductBus: productBus,
	// 	AuthClient: cfg.AuthClient,
	// })

	// rawapp.Routes(app)

	// tranapp.Routes(app, tranapp.Config{
	// 	Log:        cfg.Log,
	// 	DB:         cfg.DB,
	// 	UserBus:    userBus,
	// 	ProductBus: productBus,
	// 	AuthClient: cfg.AuthClient,
	// })

	// userapp.Routes(app, userapp.Config{
	// 	Log:        cfg.Log,
	// 	UserBus:    userBus,
	// 	AuthClient: cfg.AuthClient,
	// })

	// vproductapp.Routes(app, vproductapp.Config{
	// 	Log:         cfg.Log,
	// 	UserBus:     userBus,
	// 	VProductBus: vproductBus,
	// 	AuthClient:  cfg.AuthClient,
	// })
}
