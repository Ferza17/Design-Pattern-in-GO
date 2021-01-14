package main

import (
	"github.com/ferza17/Design_Pattern/Builder/Builder_Parameter/Model/Email_Model"
)

func main() {
	Email_Model.SendEmail(func(b *Email_Model.EmailBuilder) {
		b.
			From("foo@bar.com").
			To("bar@bar.com").
			Subject("Meeting").
			Body("Hello do you want to meet ?")
	})
}
