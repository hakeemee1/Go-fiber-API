package controllers

import (

	"go-fiber-test/database"
	m "go-fiber-test/models"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, This is API")
}

// FactCalc ฟังก์ชั่นนี้ทำหารของตัวเลขที่ส่งเข้ามา
func FactCalc(c *fiber.Ctx) error {
	// ตรวจสอบค่าที่ส่งเข้ามาในฟังก์ชั่น
	num := c.Params("num")

	// ประกาศตัวแปร factCal ใช้ในการคำนวณ
	factCal := 1

	// แปลงตัวเลขที่ส่งเข้ามาเป็นชนิดของตัวเลข
	numInt, err := strconv.Atoi(num)

	// ตรวจสอบว่ามีข้อผิดพลาดหรือไม่
	if err != nil {
		// ถ้ามีข้อผิดพลาดส่งกลับเป็น JSON ที่มีข้อผิดพลาด
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "invalid number",
		})
	}

	// ใช้วนลูปเพื่อคำนวณหารของตัวเลข
	for i := 1; i <= numInt; i++ {
		factCal *= i
	}

	// สร้างข้อความผลลัพธ์
	result := num + "! = " + strconv.Itoa(factCal)

	// ส่งข้อความผลลัพธ์กลับเป็นสตริง
	return c.SendString(result)
}

func AsciiCalc(c *fiber.Ctx) error {
	taxID := c.Query("tax_id")

	if taxID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "tax_id is required",
		})
	}

	asciiValues := make([]int, 0, len(taxID))
	for _, char := range taxID {
		asciiValues = append(asciiValues, int(char))
	}

	return c.JSON(fiber.Map{
		"ascii_values": asciiValues,
	})
}

func Register(c *fiber.Ctx) error {

	user := new(m.Register)
	if err := c.BodyParser((user)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	validate := validator.New()
	validate.RegisterValidation("username_validate", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(fl.Field().String())
	})
	validate.RegisterValidation("web_validate", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[a-z0-9\.]+$`).MatchString(fl.Field().String())
	})
	err := validate.Struct(user)
	if err != nil {
		fieldErrors := make(map[string]string)

		for _, e := range err.(validator.ValidationErrors) {
			if e.Field() == "Email" && e.Tag() == "email" {
				fieldErrors[strings.ToLower(e.Field())] = "Invalid email"
			} else if e.Field() == "Username" && e.Tag() == "username_validate" {
				fieldErrors[strings.ToLower(e.Field())] = "ใช้อักษรภาษาอังกฤษ (a-z), (A-Z), ตัวเลข (0-9) และเครื่องหมาย (_), (-) เท่านั้น เช่น Example_01"
			} else if e.Field() == "Password" && (e.Tag() == "min" || e.Tag() == "max") {
				fieldErrors[strings.ToLower(e.Field())] = "ความยาว 6-20 อักษร"
			} else if e.Field() == "WebName" && (e.Tag() == "min" || e.Tag() == "max") {
				fieldErrors[strings.ToLower(e.Field())] = "ความยาว 2-30 อักษร"
			} else if e.Field() == "WebName" && (e.Tag() == "web_validate") {
				fieldErrors[strings.ToLower(e.Field())] = "ใช้อักษรภาษาอังกฤษตัวเล็ก (a-z), ตัวเลข (0-9) ห้ามใช้เครื่องหมายอักขระพิเศษยกเว้นขีด (-) ห้ามเว้นวรรค และห้ามใช้ภาษาไทย"
			} else {
				fieldErrors[strings.ToLower(e.Field())] = e.Field() + " is required"
			}
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation errors occurred",
			"errors":  fieldErrors,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": user})
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet
			DogID: v.DogID, //112
			Type:  typeStr, //no color
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	type ResultData struct {
		Data  []m.DogsRes `json:"data"`
		Name  string    `json:"name"`
		Count int       `json:"count"`
	}
	r := ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs), //หาผลรวม,
	}
	return c.Status(200).JSON(r)
}

func GetDeletedDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at is NOT NULL").Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDogsIdMoreThan(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs
	db.Where("dog_id > 50 AND dog_id < 100").Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDogsJsonSummary(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	sumRed := 0
	sumGreen := 0
	sumPink := 0
	sumNone := 0

	var dataResults []m.DogsRes
	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sumRed++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sumGreen++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sumPink++
		} else {
			typeStr = "no color"
			sumNone++
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:        dataResults,
		Name:        "golang-test",
		// Count:       len(dogs),
		Sum_red:     sumRed,
		Sum_green:   sumGreen,
		Sum_pink:    sumPink,
		Sum_noColor: sumNone,
	}
	return c.Status(200).JSON(r)
}

//Companies CRUD
func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var companies []m.Companies
	db.Find(&companies)
	return c.Status(200).JSON(companies)
}

func AddCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Companies
	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&company)
	return c.Status(201).JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Companies
	id := c.Params("id")
	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Companies
	result := db.Delete(&company, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
