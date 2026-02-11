package pichak

type Repository interface {
	FindBySayadNo(id string) (*Cheque, error)
	Update(cheque *Cheque) error
}
