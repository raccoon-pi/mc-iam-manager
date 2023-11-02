package handler

import (
	"log"
	"mc_iam_manager/models"
	"net/http"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func MappingWsUserRole(tx *pop.Connection, bindModel *models.MCIamWsUserRoleMapping) map[string]interface{} {

	err := tx.Create(bindModel)

	if err != nil {
		return map[string]interface{}{
			"message": err,
			"status":  http.StatusBadRequest,
		}
	}
	return map[string]interface{}{
		"message": "success",
		"status":  http.StatusOK,
	}
}

func GetWsUserRole(tx *pop.Connection, bindModel *models.MCIamWsUserRoleMapping) *models.MCIamWsUserRoleMappings {

	respModel := &models.MCIamWsUserRoleMappings{}

	if user_id := bindModel.UserID; user_id != uuid.Nil {
		q := tx.Eager().Where("user_id = ?", user_id)
		err := q.All(respModel)
		if err != nil {

		}
	}

	if role_id := bindModel.RoleID; role_id != uuid.Nil {
		q := tx.Eager().Where("role_id = ?", role_id)
		err := q.All(respModel)
		if err != nil {

		}
	}
	if ws_id := bindModel.WsID; ws_id != uuid.Nil {
		q := tx.Eager().Where("ws_id = ?", ws_id)
		err := q.All(respModel)
		if err != nil {

		}
	}
	return respModel
}

func MappingWsProject(tx *pop.Connection, bindModel *models.MCIamWsProjectMapping) map[string]interface{} {

	log.Println("======== mapping ws project bind model =====")
	log.Println(bindModel)
	log.Println("======== mapping ws project bind model =====")
	err := tx.Create(bindModel)

	if err != nil {
		log.Println("======== mapping ws project =====")
		log.Println(err)
		log.Println("======== mapping ws project =====")
		return map[string]interface{}{
			"message": err,
			"status":  http.StatusBadRequest,
		}
	}
	return map[string]interface{}{
		"message": "success",
		"status":  http.StatusOK,
	}
}

func MappingGetProjectByWorkspace(tx *pop.Connection, wsId string) *models.MCIamWsProjectMappings {
	ws := &models.MCIamWsProjectMappings{}

	err := tx.Eager().Where("ws_id =?", wsId).All(ws)
	if err != nil {

	}
	return ws

}

func MappingUserRole(tx *pop.Connection, bindModel *models.MCIamUserRoleMapping) map[string]interface{} {

	err := tx.Create(bindModel)

	if err != nil {
		return map[string]interface{}{
			"message": err,
			"status":  http.StatusBadRequest,
		}
	}
	return map[string]interface{}{
		"message": "success",
		"status":  http.StatusOK,
	}
}
