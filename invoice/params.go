package invoice

import (
	"net/url"
	"strconv"
	"time"

	"github.com/xendit/xendit-go/utils"
)

// CreateParams is parameters that are needed to Create
type CreateParams struct {
	ExternalID               string   `json:"external_id" validate:"required"`
	PayerEmail               string   `json:"payer_email" validate:"required"`
	Description              string   `json:"description" validate:"required"`
	Amount                   float64  `json:"amount" validate:"required"`
	ShouldSendEmail          bool     `json:"should_send_email,omitempty"`
	CallbackVirtualAccountID string   `json:"callback_virtual_account_id,omitempty"`
	InvoiceDuration          int      `json:"invoice_duration,omitempty"`
	SuccessRedirectURL       string   `json:"success_redirect_url,omitempty"`
	FailureRedirectURL       string   `json:"failure_redirect_url,omitempty"`
	PaymentMethods           []string `json:"payment_methods,omitempty"`
	MidLabel                 string   `json:"mid_label,omitempty"`
	Currency                 string   `json:"currency,omitempty"`
	FixedVA                  bool     `json:"fixed_va,omitempty"`
}

// GetAllParams is parameters that are needed to GetAll
type GetAllParams struct {
	Statuses           []string  `json:"statuses,omitempty"`
	Limit              int       `json:"limit,omitempty"`
	CreatedAfter       time.Time `json:"created_after,omitempty"`
	CreatedBefore      time.Time `json:"created_before,omitempty"`
	PaidAfter          time.Time `json:"paid_after,omitempty"`
	PaidBefore         time.Time `json:"paid_before,omitempty"`
	ExpiredAfter       time.Time `json:"expired_after,omitempty"`
	ExpiredBefore      time.Time `json:"expired_before,omitempty"`
	LastInvoiceID      string    `json:"last_invoice_id,omitempty"`
	ClientTypes        []string  `json:"client_types,omitempty"`
	PaymentChannels    []string  `json:"payment_channels,omitempty"`
	OnDemandLink       string    `json:"on_demand_link,omitempty"`
	RecurringPaymentID string    `json:"recurring_payment_id,omitempty"`
}

// QueryString create query string from GetAllParams, ignore nil values
func (p *GetAllParams) QueryString() string {
	urlValues := &url.Values{}

	utils.AddStringSliceToURLValues(urlValues, p.Statuses, "statuses")
	if p.Limit > 0 {
		urlValues.Add("limit", strconv.Itoa(p.Limit))
	}
	utils.AddTimeToURLValues(urlValues, p.CreatedAfter, "created_after")
	utils.AddTimeToURLValues(urlValues, p.CreatedBefore, "created_before")
	utils.AddTimeToURLValues(urlValues, p.PaidAfter, "paid_after")
	utils.AddTimeToURLValues(urlValues, p.PaidBefore, "paid_before")
	utils.AddTimeToURLValues(urlValues, p.PaidBefore, "paid_before")
	utils.AddTimeToURLValues(urlValues, p.ExpiredAfter, "expired_after")
	utils.AddTimeToURLValues(urlValues, p.ExpiredBefore, "expired_before")
	utils.AddStringSliceToURLValues(urlValues, p.ClientTypes, "client_types")
	utils.AddStringSliceToURLValues(urlValues, p.PaymentChannels, "payment_channels")
	if p.OnDemandLink != "" {
		urlValues.Add("on_demand", p.OnDemandLink)
	}
	if p.RecurringPaymentID != "" {
		urlValues.Add("recurring_payment_id", p.RecurringPaymentID)
	}

	return urlValues.Encode()
}
