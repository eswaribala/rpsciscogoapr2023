package facades

import (
	"CiscoApr2023/banking/stores"
)

type IBank interface {
	Add(date *stores.Date)
	View(permission bool)
	ViewById(id *int64)
	Update(runningTotal *int64)
	Delete(id *int64)
}
