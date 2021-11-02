package routes

import (
	"GolangAPI/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// route untuk user dan login
	app.Post("/api/registeruser", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/getuser", controllers.GetData)
	app.Put("/api/updateuser", controllers.UpdateData)
	app.Delete("/api/updateuser", controllers.Deletedata)

	//route untuk produk
	app.Post("/api/registerproduk", controllers.RegisterProduk)
	app.Get("/api/getproduk", controllers.GetDataProduk)
	app.Put("/api/updateproduk", controllers.UpdateDataProduk)
	app.Delete("/api/updateproduk", controllers.DeletedataProduk)
}
