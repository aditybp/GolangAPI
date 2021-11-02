package controllers

import (
	"GolangAPI/models"
	"GolangAPI/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const SecretKey = "secret"
func Register(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Nama: data["nama"],
		Email: data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return ctx.JSON(user)
}

func Login(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0{
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "user tidak ada",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": "password salah",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour*24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message" : "Tidak dapat diakses",
		})
	}

	cookie := fiber.Cookie{
		Name: "jet",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "succes",
	})

	return ctx.JSON(token)
}

func GetData(ctx *fiber.Ctx) error {

	var users []models.User
	database.DB.Find(&users)

	if len(users) == 0 {
		return ctx.Status(404).JSON(fiber.Map{"Status" : "error", "data" : nil})
	}

	return ctx.JSON(users)
}

func UpdateData(ctx *fiber.Ctx) error {

	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	type updatedata struct {
		Nama string	`json:"nama panjang"`
		Email string `json:"email@email.com"`
	}

	var users models.User

	database.DB.Find("id = ?", data["id"])

	if users.Id == 0{
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "user tidak ada",
		})
	}

	var updateData updatedata
	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	users.Nama = updateData.Nama
	users.Email = updateData.Email

	database.DB.Save(&users)

	return ctx.JSON(users)
}

func Deletedata(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}
	var users models.User

	database.DB.Find("id = ?", data["id"])

	if users.Id == 0{
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "user tidak ada",
		})
	}

	err := database.DB.Delete(&users, "id = ?", data["id"]).Error

	if err != nil{
		return ctx.JSON(fiber.Map{
			"message" : "gagal delete",
		})
	}

	return ctx.JSON(fiber.Map{
		"message" : "telah ter-delete",
	})
}

// function untuk tabel produk

func RegisterProduk(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	produk := models.Produk{
		Nama_Barang	: data["nama_barang"],
		Kode_Barang: data["kode_barang"],
		Harga_Barang: data["harga_barang"],
	}

	database.DB.Create(&produk)

	return ctx.JSON(produk)
}

func GetDataProduk(ctx *fiber.Ctx) error {

	var produk []models.Produk
	database.DB.Find(&produk)

	if len(produk) == 0 {
		return ctx.Status(404).JSON(fiber.Map{"Status" : "error", "data" : nil})
	}

	return ctx.JSON(produk)
}

func UpdateDataProduk(ctx *fiber.Ctx) error {

	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	type updateproduk struct {
		Nama_Barang string	`json:"nama barang"`
		Kode_Barang string `json:"A011111"`
		Harga_Barang string `json:"40000"`
	}

	var produk models.Produk

	database.DB.Find("id = ?", data["id"])

	if produk.Id == 0{
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "user tidak ada",
		})
	}

	var updateDataProduk updateproduk
	if err := ctx.BodyParser(&updateDataProduk); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	produk.Nama_Barang = updateDataProduk.Nama_Barang
	produk.Kode_Barang = updateDataProduk.Kode_Barang
	produk.Harga_Barang = updateDataProduk.Harga_Barang

	database.DB.Save(&produk)

	return ctx.JSON(produk)
}

func DeletedataProduk(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}
	var produk models.Produk

	database.DB.Find("id = ?", data["id"])

	if produk.Id == 0{
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "produk tidak ada",
		})
	}

	err := database.DB.Delete(&produk, "id = ?", data["id"]).Error

	if err != nil{
		return ctx.JSON(fiber.Map{
			"message" : "gagal delete",
		})
	}

	return ctx.JSON(fiber.Map{
		"message" : "telah ter-delete",
	})
}
