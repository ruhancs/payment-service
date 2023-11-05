package paymentgateway

type PaymentGatewayInterface interface {
	CreateCustomer(paymentMethod, email string) (string, string, error)
	SubscribeToPlan(customerID,plan,email,cardType string) (string,error)
	Refunds(pi string, amount int) error
	CancelSubscription(subID string) error
}