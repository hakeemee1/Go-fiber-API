package controllers

import (
	"strconv"

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
