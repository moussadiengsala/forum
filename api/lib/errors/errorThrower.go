package lib

import "learn.zone01dakar.sn/forum-rest-api/lib"


func ErrorWriter(response *lib.Response, message string, statusCode int) {
	response.Message = message
	response.Code = statusCode
}
