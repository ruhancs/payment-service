package entity

type Customer struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IsActive  bool   `json:"is_active"`
}

func NewCustomer() *Customer {
	return &Customer{}
}

func(c *Customer) Active() {
	c.IsActive = true
}

func(c *Customer) Deactive() {
	c.IsActive = false
}