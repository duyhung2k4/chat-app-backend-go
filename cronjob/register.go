package cronjob

import (
	"chatapp/config"
	"chatapp/model"
	"time"
)

func DeleteExpiredTemporaryEmail() {
	db := config.GetDB()

	db.Where("time_end < ?", time.Now()).Unscoped().Delete(&model.TemporaryCredential{})
}
