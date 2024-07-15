package controllers

import (
	"regexp"
	"strconv"
	

	models "go-fiber-test/models"

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

// CustomValidator เป็นฟังก์ชันสำหรับ validate ชื่อเว็บไซต์
func WebsiteDomainValidator(fl validator.FieldLevel) bool {
	website := fl.Field().String()
	regex := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-]{0,28}[a-zA-Z0-9]$`)
	return regex.MatchString(website)
}

var validate = validator.New()

func init() {
	validate.RegisterValidation("websiteDomain", WebsiteDomainValidator)
}

func Register(c *fiber.Ctx) error {
	var user models.Register

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": formatValidationErrors(errors),
		})
	}

	// TODO: ดำเนินการลงทะเบียนผู้ใช้ในฐานข้อมูล

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func formatValidationErrors(errors validator.ValidationErrors) map[string]string {
	errorMap := make(map[string]string)

	for _, err := range errors {
		switch err.Tag() {
		case "required":
			errorMap[err.Field()] = "กรุณากรอกข้อมูลในช่องนี้"
		case "email":
			errorMap[err.Field()] = "กรุณากรอกอีเมลให้ถูกต้อง"
		case "alphanum":
			errorMap[err.Field()] = "กรุณาใช้ตัวอักษรภาษาอังกฤษหรือตัวเลขเท่านั้น"
		case "min":
			errorMap[err.Field()] = "ข้อมูลต้องมีความยาวอย่างน้อย " + err.Param() + " ตัวอักษร"
		case "max":
			errorMap[err.Field()] = "ข้อมูลต้องมีความยาวไม่เกิน " + err.Param() + " ตัวอักษร"
		case "numeric":
			errorMap[err.Field()] = "กรุณากรอกตัวเลขเท่านั้น"
		case "len":
			errorMap[err.Field()] = "ข้อมูลต้องมีความยาว " + err.Param() + " ตัวอักษร"
		case "websiteDomain":
			errorMap[err.Field()] = "ชื่อเว็บไซต์ไม่ถูกต้อง"
		default:
			errorMap[err.Field()] = "ข้อมูลไม่ถูกต้อง"
		}
	}

	return errorMap
}
