package response

type IResponder interface {
	SingleResponder(successCode int, data interface{})
	ErrorResponder(errorCode int, appErrorCode string, errorMessage string)
	SetContext(ctx interface{}) IResponder
}
