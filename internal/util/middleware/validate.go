package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func ValidateStruct(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, validationErr := range validationErrors {
				fieldName := validationErr.Field()
				tag := validationErr.Tag()

				var errorMessage string
				switch tag {
				case "required":
					errorMessage = fieldName + " is required"
				case "tags":
					errorMessage = fieldName + " cannot have more than 5 tags"
				case "duplicated":
					errorMessage = fieldName + " has duplicated values"
				case "not_blank":
					errorMessage = fieldName + " cannot be empty or contain only spaces"
				case "max":
					errorMessage = fieldName + " must not exceed " + validationErr.Param() + " characters"
				case "min":
					errorMessage = fieldName + " must have at least " + validationErr.Param() + " characters"
				default:
					errorMessage = fieldName + " is invalid"
				}

				errorMessages = append(errorMessages, errorMessage)
			}

			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return false
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	return true
}

func NotBlank(fl validator.FieldLevel) bool {
	if str, ok := fl.Field().Interface().(string); ok {
		return strings.TrimSpace(str) != ""
	}
	return false
}

func ValidateTags(fl validator.FieldLevel) bool {
	if slice, ok := fl.Field().Interface().([]string); ok {
		return len(slice) <= 5
	}
	return false
}

func ValidateDuplicatedTags(fl validator.FieldLevel) bool {
	if slice, ok := fl.Field().Interface().([]string); ok {
		return !hasDuplicates(slice)
	}
	return false
}

func hasDuplicates(strings []string) bool {
	seen := make(map[string]bool)

	for _, str := range strings {
		if seen[str] {
			return true
		}
		seen[str] = true
	}
	return false
}
