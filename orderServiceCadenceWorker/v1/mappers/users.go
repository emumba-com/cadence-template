package mappers

import "orderServiceCadenceWorker/v1/models"

func UsersToUsersMap(users []models.User) map[string]models.User {
	usersMap := make(map[string]models.User)

	for _, user := range users {
		usersMap[user.ID] = user
	}

	return usersMap
}
