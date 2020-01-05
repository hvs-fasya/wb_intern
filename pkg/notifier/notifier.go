package notifier

import "fmt"

// Cfg notifier instance config struct
type Cfg struct {
	Tmpl string
}

// Notifier common notifier interface
type Notifier interface {
	NotifyCreateUser([]interface{})
}

type notifier struct {
	Cfg
}

// NotifyCreateUser create user notiifcation method
func (n *notifier) NotifyCreateUser(subst []interface{}) {
	fmt.Printf(n.Tmpl, subst...)
}

// NewNotifier notifier constructor
func NewNotifier(cfg Cfg) Notifier {
	return &notifier{Cfg: cfg}
}
