package paymentgateway

import (
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/refund"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/sub"
)

type StripeCard struct {
	Secret string
	Key string
	Currency string
}

func (c *StripeCard) SubscribeToPlan(customerID,plan,email,cardType string) (string,error) {
	items := []*stripe.SubscriptionItemsParams{
		{Plan: stripe.String(plan)},
	}

	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerID),
		Items: items,
	}

	params.AddMetadata("card_type",cardType)
	params.AddExpand("latest_invoice.payment_intent")

	subscription,err := sub.New(params)
	if err != nil {
		return "",err
	}

	return subscription.ID,nil
}

//criar um customer no dashbioard do stripe
func (c *StripeCard) CreateCustomer(paymentMethod, email string) (string, string, error) {
	stripe.Key = c.Secret
	customerParams := &stripe.CustomerParams{
		PaymentMethod: stripe.String(paymentMethod),
		Email: stripe.String(email),
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(paymentMethod),
		},
	}

	cust,err := customer.New(customerParams)
	if err != nil {
		msg:= ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMsg(stripeErr.Code)
		}
		return "", msg, err
	}
	return cust.ID, "", nil
}

func(c *StripeCard) Refunds(pi string, amount int) error {
	stripe.Key = c.Secret
	amountToRefund := int64(amount)

	refundParams := &stripe.RefundParams{
		Amount: &amountToRefund,
		PaymentIntent: &pi,
	}

	_, err := refund.New(refundParams)
	if err != nil {
		return err
	}

	return nil
}

func (c *StripeCard) CancelSubscription(subID string) error {
	stripe.Key = c.Secret

	params := &stripe.SubscriptionParams{
		CancelAtPeriodEnd: stripe.Bool(true),
	}

	_, err := sub.Update(subID, params)
	if err != nil {
		return err
	}
	return nil
}

func cardErrorMsg(code any) string {
	var msg = ""

	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Your Card was declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "Your Card is expired"
	case stripe.ErrorCodeIncorrectCVC:
		msg = "Incorrect CVC code"
	case stripe.ErrorCodeIncorrectZip:
		msg = "Incorrect Zip/Postal code"
	case stripe.ErrorCodeAmountTooLarge:
		msg = "Amount too large to charge to your card"
	case stripe.ErrorCodeAmountTooSmall:
		msg = "Amount too smal to charge to your card"
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "balance insuficient"
	case stripe.ErrorCodePostalCodeInvalid:
		msg = "your postal code is invalid"
	default:
		msg = "Your Card was declined"
	}
	return msg
}