package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response - umumiy javob modeli
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

// DataWrapper - har doim obyekt shaklida qaytariladigan data tuzilmasi
type DataWrapper struct {
	Items interface{} `json:"items,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
}

// Send - umumiy yuboruvchi
func Send(c *gin.Context, status int, data interface{}, err interface{}) {
	c.JSON(status, Response{
		Status: status,
		Data:   data,
		Error:  err,
	})
}

// Success - muvaffaqiyatli javob uchun (avtomatik obyekt shaklida)
func Success(c *gin.Context, items interface{}) {
	wrapped := DataWrapper{
		Items: items,
	}
	Send(c, http.StatusOK, wrapped, nil)
}

// Created - yangi resurs yaratildi
func Created(c *gin.Context, item interface{}) {
	wrapped := DataWrapper{
		Items: item,
	}
	Send(c, http.StatusCreated, wrapped, nil)
}

// SuccessWithMeta - agar pagination yoki meta ma’lumot bo‘lsa
func SuccessWithMeta(c *gin.Context, items interface{}, meta interface{}) {
	wrapped := DataWrapper{
		Items: items,
		Meta:  meta,
	}
	Send(c, http.StatusOK, wrapped, nil)
}

// Fail - xatoliklar uchun
func Fail(c *gin.Context, status int, message string) {
	err := map[string]interface{}{
		"message": message,
	}
	Send(c, status, nil, err)
}
