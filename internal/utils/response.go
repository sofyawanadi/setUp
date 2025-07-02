package utils

const SuccessMessage = "Operation completed successfully"
const ErrorMessage = "An error occurred during the operation"

func ResponseSuccess(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

func ResponseError(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "error",
		"message": message,
	}
}