package contract

type TransactionManager interface {
	Run(callback func(tx any) error) error
}
