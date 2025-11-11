package presentation_api_middlewares

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func BindAndValidateQuery[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj T
		objType := reflect.TypeOf(obj)
		objValue := reflect.ValueOf(&obj).Elem()

		if objType.Kind() != reflect.Struct {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "El tipo proporcionado no es una estructura"})
			c.Abort()
			return
		}

		for i := 0; i < objType.NumField(); i++ {
			field := objType.Field(i)
			valueField := objValue.Field(i)

			queryName := field.Tag.Get("query")
			if queryName == "" {
				continue
			}

			rawValue := c.Query(queryName)
			rules := field.Tag.Get("validate")

			if rules != "" && contains(rules, "required") && rawValue == "" {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("El parámetro %s es obligatorio", queryName)})
				c.Abort()
				return
			}

			if rawValue == "" {
				continue
			}

			switch {
			case contains(rules, "int"):
				parsed, err := strconv.Atoi(rawValue)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("El parámetro %s debe ser numérico", queryName)})
					c.Abort()
					return
				}
				valueField.SetInt(int64(parsed))

			case contains(rules, "datetime"):
				t, err := time.Parse("2006-01-02 15:04:05", rawValue)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("El parámetro %s debe tener el formato YYYY-MM-DD HH:MM:SS", queryName)})
					c.Abort()
					return
				}
				valueField.Set(reflect.ValueOf(t))

			case contains(rules, "string"):
				valueField.SetString(rawValue)
			}
		}

		c.Set(reflect.TypeOf(obj).Name(), obj)
		c.Next()
	}
}

func contains(rules string, keyword string) bool {
	for _, r := range splitAndTrim(rules) {
		if r == keyword {
			return true
		}
	}
	return false
}

func splitAndTrim(s string) []string {
	var res []string
	current := ""
	for _, c := range s {
		if c == ',' {
			if current != "" {
				res = append(res, current)
				current = ""
			}
		} else {
			current += string(c)
		}
	}
	if current != "" {
		res = append(res, current)
	}
	return res
}
