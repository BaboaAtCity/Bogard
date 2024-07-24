package handlers

import "strconv"

func IsAuth(id int64, allowed string) bool {
	idStr := strconv.FormatInt(id, 10)
	return idStr == allowed
}
