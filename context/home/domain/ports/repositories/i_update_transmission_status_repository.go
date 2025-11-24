package irepositories

type IUpdateTransmissionStatusRepository interface {
	MarkSynchronized(id int) error
}
