package helpers

import (
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/crud"
)

// Check the submited user's infos while loggin
func VerifyUserInfo(response *model.Reponse, db *model.DB, isCheckPasswordNeed bool, condition string, datas map[string]interface{}) model.User {
	columns := []string{}
	conditions := map[string]interface{}{
		condition: datas[condition],
	}

	var usr model.User
	SelectionErr := crud.Select(db.Instance, "User", "", conditions, columns, &usr.ID, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Username, &usr.Bio, &usr.Avatar, &usr.Password, &usr.CreationDate)

	if SelectionErr != nil || (isCheckPasswordNeed && !PasswordDecrypter(usr.Password, datas["password"].(string))) {
		ErrorWriter(response, "Invalid email or password", 401)
		return usr
	}

	return usr
}
