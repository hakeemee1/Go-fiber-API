package controllers

import (
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
