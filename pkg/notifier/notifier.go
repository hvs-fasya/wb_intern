package notifier

//NotifierIface common notifier interface
type NotifierIface interface {
	NotifyCreateUser(string)
}

type notifier struct{}

//NotifyCreateUser create user notiifcation method
func (n *notifier) NotifyCreateUser(userName string) {
	println("user '" + userName + "' created")
}

//NewNotifier notifier constructor
func NewNotifier() NotifierIface {
	return &notifier{}
}
