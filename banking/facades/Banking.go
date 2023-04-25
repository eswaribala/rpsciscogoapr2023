package facades

import (
	"CiscoApr2023/banking/models"
)

type IBank interface {
	Add(date *models.Date)
	View(permission bool)
	ViewById(id *int64)
	Update(runningTotal *int64)
	Delete(id *int64)
}
