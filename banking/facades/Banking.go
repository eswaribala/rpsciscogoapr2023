package facades

type IBank interface {
	Add(date *string)
	View(permission bool)
	ViewById(id *int64)
	Update(runningTotal *int64)
	Delete(id *int64)
}
