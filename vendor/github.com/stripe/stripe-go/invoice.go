package stripe

import "encoding/json"

// InvoiceLineType is the list of allowed values for the invoice line's type.
type InvoiceLineType string

// List of values that InvoiceLineType can take.
const (
	InvoiceLineTypeInvoiceItem  InvoiceLineType = "invoiceitem"
	InvoiceLineTypeSubscription InvoiceLineType = "subscription"
)

// InvoiceBilling is the type of billing method for this invoice.
type InvoiceBilling string

// List of values that InvoiceBilling can take.
const (
	InvoiceBillingChargeAutomatically InvoiceBilling = "charge_automatically"
	InvoiceBillingSendInvoice         InvoiceBilling = "send_invoice"
)

// InvoiceBillingReason is the reason why a given invoice was created
type InvoiceBillingReason string

// List of values that InvoiceBillingReason can take.
const (
	InvoiceBillingReasonManual                InvoiceBillingReason = "manual"
	InvoiceBillingReasonSubscription          InvoiceBillingReason = "subscription"
	InvoiceBillingReasonSubscriptionCreate    InvoiceBillingReason = "subscription_create"
	InvoiceBillingReasonSubscriptionCycle     InvoiceBillingReason = "subscription_cycle"
	InvoiceBillingReasonSubscriptionThreshold InvoiceBillingReason = "subscription_threshold"
	InvoiceBillingReasonSubscriptionUpdate    InvoiceBillingReason = "subscription_update"
	InvoiceBillingReasonUpcoming              InvoiceBillingReason = "upcoming"
)

// InvoiceBillingStatus is the reason why a given invoice was created
type InvoiceBillingStatus string

// List of values that InvoiceBillingStatus can take.
const (
	InvoiceBillingStatusDraft         InvoiceBillingStatus = "draft"
	InvoiceBillingStatusOpen          InvoiceBillingStatus = "open"
	InvoiceBillingStatusPaid          InvoiceBillingStatus = "paid"
	InvoiceBillingStatusUncollectible InvoiceBillingStatus = "uncollectible"
	InvoiceBillingStatusVoid          InvoiceBillingStatus = "void"
)

// InvoiceUpcomingInvoiceItemPeriodParams represents the period associated with that invoice item
type InvoiceUpcomingInvoiceItemPeriodParams struct {
	End   *int64 `form:"end"`
	Start *int64 `form:"start"`
}

// InvoiceUpcomingInvoiceItemParams is the set of parameters that can be used when adding or modifying
// invoice items on an upcoming invoice.
// For more details see https://stripe.com/docs/api#upcoming_invoice-invoice_items.
type InvoiceUpcomingInvoiceItemParams struct {
	Amount       *int64                                  `form:"amount"`
	Currency     *string                                 `form:"currency"`
	Description  *string                                 `form:"description"`
	Discountable *bool                                   `form:"discountable"`
	InvoiceItem  *string                                 `form:"invoiceitem"`
	Period       *InvoiceUpcomingInvoiceItemPeriodParams `form:"period"`
	Quantity     *int64                                  `form:"quantity"`
	UnitAmount   *int64                                  `form:"unit_amount"`
}

// InvoiceCustomFieldParams represents the parameters associated with one custom field on an invoice.
type InvoiceCustomFieldParams struct {
	Name  *string `form:"name"`
	Value *string `form:"value"`
}

// InvoiceTransferDataParams is the set of parameters allowed for the transfer_data hash.
type InvoiceTransferDataParams struct {
	Destination *string `form:"destination"`
}

// InvoiceParams is the set of parameters that can be used when creating or updating an invoice.
// For more details see https://stripe.com/docs/api#create_invoice, https://stripe.com/docs/api#update_invoice.
type InvoiceParams struct {
	Params               `form:"*"`
	AutoAdvance          *bool                       `form:"auto_advance"`
	ApplicationFeeAmount *int64                      `form:"application_fee_amount"`
	Billing              *string                     `form:"billing"`
	CustomFields         []*InvoiceCustomFieldParams `form:"custom_fields"`
	Customer             *string                     `form:"customer"`
	DaysUntilDue         *int64                      `form:"days_until_due"`
	DefaultSource        *string                     `form:"default_source"`
	Description          *string                     `form:"description"`
	DueDate              *int64                      `form:"due_date"`
	Footer               *string                     `form:"footer"`
	Paid                 *bool                       `form:"paid"`
	StatementDescriptor  *string                     `form:"statement_descriptor"`
	Subscription         *string                     `form:"subscription"`
	TaxPercent           *float64                    `form:"tax_percent"`
	TransferData         *InvoiceTransferDataParams  `form:"transfer_data"`

	// These are all for exclusive use by GetNext.

	Coupon                         *string                             `form:"coupon"`
	InvoiceItems                   []*InvoiceUpcomingInvoiceItemParams `form:"invoice_items"`
	SubscriptionBillingCycleAnchor *int64                              `form:"subscription_billing_cycle_anchor"`
	SubscriptionCancelAtPeriodEnd  *bool                               `form:"subscription_cancel_at_period_end"`
	SubscriptionItems              []*SubscriptionItemsParams          `form:"subscription_items"`
	SubscriptionPlan               *string                             `form:"subscription_plan"`
	SubscriptionProrate            *bool                               `form:"subscription_prorate"`
	SubscriptionProrationDate      *int64                              `form:"subscription_proration_date"`
	SubscriptionQuantity           *int64                              `form:"subscription_quantity"`
	SubscriptionTaxPercent         *float64                            `form:"subscription_tax_percent"`
	SubscriptionTrialEnd           *int64                              `form:"subscription_trial_end"`
	SubscriptionTrialFromPlan      *bool                               `form:"subscription_trial_from_plan"`

	// This parameter is considered deprecated. Prefer using ApplicationFeeAmount
	ApplicationFee *int64 `form:"application_fee"`
}

// InvoiceListParams is the set of parameters that can be used when listing invoices.
// For more details see https://stripe.com/docs/api#list_customer_invoices.
type InvoiceListParams struct {
	ListParams   `form:"*"`
	Billing      *string           `form:"billing"`
	Customer     *string           `form:"customer"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created_range"`
	DueDate      *int64            `form:"due_date"`
	Subscription *string           `form:"subscription"`

	// Those parameters are deprecated. Prefer using Created or CreatedRange
	Date      *int64            `form:"date"`
	DateRange *RangeQueryParams `form:"date"`
}

// InvoiceLineListParams is the set of parameters that can be used when listing invoice line items.
// For more details see https://stripe.com/docs/api#invoice_lines.
type InvoiceLineListParams struct {
	ListParams `form:"*"`

	Customer *string `form:"customer"`

	// ID is the invoice ID to list invoice lines for.
	ID *string `form:"-"` // Goes in the URL

	Subscription *string `form:"subscription"`
}

// InvoiceFinalizeParams is the set of parameters that can be used when finalizing invoices.
type InvoiceFinalizeParams struct {
	Params      `form:"*"`
	AutoAdvance *bool `form:"auto_advance"`
}

// InvoiceMarkUncollectibleParams is the set of parameters that can be used when marking
// invoices as uncollectible.
type InvoiceMarkUncollectibleParams struct {
	Params `form:"*"`
}

// InvoicePayParams is the set of parameters that can be used when
// paying invoices. For more details, see:
// https://stripe.com/docs/api#pay_invoice.
type InvoicePayParams struct {
	Params        `form:"*"`
	Forgive       *bool   `form:"forgive"`
	PaidOutOfBand *bool   `form:"paid_out_of_band"`
	Source        *string `form:"source"`
}

// InvoiceSendParams is the set of parameters that can be used when sending invoices.
type InvoiceSendParams struct {
	Params `form:"*"`
}

// InvoiceVoidParams is the set of parameters that can be used when voiding invoices.
type InvoiceVoidParams struct {
	Params `form:"*"`
}

// Invoice is the resource representing a Stripe invoice.
// For more details see https://stripe.com/docs/api#invoice_object.
type Invoice struct {
	AmountDue                 int64                    `json:"amount_due"`
	AmountPaid                int64                    `json:"amount_paid"`
	AmountRemaining           int64                    `json:"amount_remaining"`
	ApplicationFeeAmount      int64                    `json:"application_fee_amount"`
	AttemptCount              int64                    `json:"attempt_count"`
	Attempted                 bool                     `json:"attempted"`
	AutoAdvance               bool                     `json:"auto_advance"`
	Billing                   InvoiceBilling           `json:"billing"`
	BillingReason             InvoiceBillingReason     `json:"billing_reason"`
	Charge                    *Charge                  `json:"charge"`
	Created                   int64                    `json:"created"`
	Currency                  Currency                 `json:"currency"`
	CustomFields              []*InvoiceCustomField    `json:"custom_fields"`
	Customer                  *Customer                `json:"customer"`
	DefaultSource             *PaymentSource           `json:"default_source"`
	Description               string                   `json:"description"`
	Discount                  *Discount                `json:"discount"`
	DueDate                   int64                    `json:"due_date"`
	EndingBalance             int64                    `json:"ending_balance"`
	FinalizedAt               int64                    `json:"finalized_at"`
	Footer                    string                   `json:"footer"`
	HostedInvoiceURL          string                   `json:"hosted_invoice_url"`
	ID                        string                   `json:"id"`
	InvoicePDF                string                   `json:"invoice_pdf"`
	Lines                     *InvoiceLineList         `json:"lines"`
	Livemode                  bool                     `json:"livemode"`
	Metadata                  map[string]string        `json:"metadata"`
	NextPaymentAttempt        int64                    `json:"next_payment_attempt"`
	Number                    string                   `json:"number"`
	Paid                      bool                     `json:"paid"`
	PeriodEnd                 int64                    `json:"period_end"`
	PeriodStart               int64                    `json:"period_start"`
	ReceiptNumber             string                   `json:"receipt_number"`
	StartingBalance           int64                    `json:"starting_balance"`
	StatementDescriptor       string                   `json:"statement_descriptor"`
	Status                    InvoiceBillingStatus     `json:"status"`
	StatusTransitions         InvoiceStatusTransitions `json:"status_transitions"`
	Subscription              string                   `json:"subscription"`
	SubscriptionProrationDate int64                    `json:"subscription_proration_date"`
	Subtotal                  int64                    `json:"subtotal"`
	Tax                       int64                    `json:"tax"`
	TaxPercent                float64                  `json:"tax_percent"`
	ThreasholdReason          *InvoiceThresholdReason  `json:"threshold_reason"`
	Total                     int64                    `json:"total"`
	TransferData              *InvoiceTransferData     `json:"transfer_data"`
	WebhooksDeliveredAt       int64                    `json:"webhooks_delivered_at"`

	// This property is considered deprecated. Prefer using ApplicationFeeAmount
	ApplicationFee int64 `json:"application_fee"`

	// This property is considered deprecated. Prefer using created
	Date int64 `json:"date"`
}

// InvoiceCustomField is a structure representing a custom field on an Invoice.
type InvoiceCustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// InvoiceThresholdReason is a structure representing a reason for a billing threshold.
type InvoiceThresholdReason struct {
	AmountGTE   int64                               `json:"amount_gte"`
	ItemReasons []*InvoiceThresholdReasonItemReason `json:"item_reasons"`
}

// InvoiceThresholdReasonItemReason is a structure representing the line items that
// triggered an invoice.
type InvoiceThresholdReasonItemReason struct {
	LineItemIDs []string `json:"line_item_ids"`
	UsageGTE    int64    `json:"usage_gte"`
}

// InvoiceList is a list of invoices as retrieved from a list endpoint.
type InvoiceList struct {
	ListMeta
	Data []*Invoice `json:"data"`
}

// InvoiceLine is the resource representing a Stripe invoice line item.
// For more details see https://stripe.com/docs/api#invoice_line_item_object.
type InvoiceLine struct {
	Amount           int64             `json:"amount"`
	Currency         Currency          `json:"currency"`
	Description      string            `json:"description"`
	Discountable     bool              `json:"discountable"`
	ID               string            `json:"id"`
	Livemode         bool              `json:"live_mode"`
	Metadata         map[string]string `json:"metadata"`
	Period           *Period           `json:"period"`
	Plan             *Plan             `json:"plan"`
	Proration        bool              `json:"proration"`
	Quantity         int64             `json:"quantity"`
	Subscription     string            `json:"subscription"`
	SubscriptionItem string            `json:"subscription_item"`
	Type             InvoiceLineType   `json:"type"`
}

// InvoiceTransferData represents the information for the transfer_data associated with an invoice.
type InvoiceTransferData struct {
	Destination *Account `json:"destination"`
}

// Period is a structure representing a start and end dates.
type Period struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

// InvoiceLineList is a list object for invoice line items.
type InvoiceLineList struct {
	ListMeta
	Data []*InvoiceLine `json:"data"`
}

// InvoiceStatusTransitions are the timestamps at which the invoice status was updated.
type InvoiceStatusTransitions struct {
	FinalizedAt           int64 `json:"finalized_at"`
	MarkedUncollectibleAt int64 `json:"marked_uncollectible_at"`
	PaidAt                int64 `json:"paid_at"`
	VoidedAt              int64 `json:"voided_at"`
}

// UnmarshalJSON handles deserialization of an Invoice.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *Invoice) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type invoice Invoice
	var v invoice
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = Invoice(v)
	return nil
}