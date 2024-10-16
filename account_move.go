package odoo

// AccountMove represents account.move model.
type AccountMove struct {
	LastUpdate                            *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                           *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                             *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning                         *String    `xmlrpc:"access_warning,omitempty"`
	ActivityCalendarEventId               *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                  *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration           *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                 *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                           *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                         *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                       *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                      *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                        *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                        *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AlwaysTaxExigible                     *Bool      `xmlrpc:"always_tax_exigible,omitempty"`
	AmountPaid                            *Float     `xmlrpc:"amount_paid,omitempty"`
	AmountResidual                        *Float     `xmlrpc:"amount_residual,omitempty"`
	AmountResidualSigned                  *Float     `xmlrpc:"amount_residual_signed,omitempty"`
	AmountTax                             *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTaxSigned                       *Float     `xmlrpc:"amount_tax_signed,omitempty"`
	AmountTotal                           *Float     `xmlrpc:"amount_total,omitempty"`
	AmountTotalInCurrencySigned           *Float     `xmlrpc:"amount_total_in_currency_signed,omitempty"`
	AmountTotalSigned                     *Float     `xmlrpc:"amount_total_signed,omitempty"`
	AmountUntaxed                         *Float     `xmlrpc:"amount_untaxed,omitempty"`
	AmountUntaxedSigned                   *Float     `xmlrpc:"amount_untaxed_signed,omitempty"`
	AssetAssetType                        *Selection `xmlrpc:"asset_asset_type,omitempty"`
	AssetDepreciatedValue                 *Float     `xmlrpc:"asset_depreciated_value,omitempty"`
	AssetDepreciationBeginningDate        *Time      `xmlrpc:"asset_depreciation_beginning_date,omitempty"`
	AssetId                               *Many2One  `xmlrpc:"asset_id,omitempty"`
	AssetIdDisplayName                    *String    `xmlrpc:"asset_id_display_name,omitempty"`
	AssetIds                              *Relation  `xmlrpc:"asset_ids,omitempty"`
	AssetNumberDays                       *Int       `xmlrpc:"asset_number_days,omitempty"`
	AssetRemainingValue                   *Float     `xmlrpc:"asset_remaining_value,omitempty"`
	AssetValueChange                      *Bool      `xmlrpc:"asset_value_change,omitempty"`
	AttachmentIds                         *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorizedTransactionIds              *Relation  `xmlrpc:"authorized_transaction_ids,omitempty"`
	AutoPost                              *Selection `xmlrpc:"auto_post,omitempty"`
	AutoPostOriginId                      *Many2One  `xmlrpc:"auto_post_origin_id,omitempty"`
	AutoPostUntil                         *Time      `xmlrpc:"auto_post_until,omitempty"`
	BankPartnerId                         *Many2One  `xmlrpc:"bank_partner_id,omitempty"`
	CampaignId                            *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CommercialPartnerId                   *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyCurrencyId                     *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryCode                           *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                                  *Time      `xmlrpc:"date,omitempty"`
	DepreciationValue                     *Float     `xmlrpc:"depreciation_value,omitempty"`
	DirectionSign                         *Int       `xmlrpc:"direction_sign,omitempty"`
	DisplayInactiveCurrencyWarning        *Bool      `xmlrpc:"display_inactive_currency_warning,omitempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omitempty"`
	DisplayQrCode                         *Bool      `xmlrpc:"display_qr_code,omitempty"`
	DraftAssetIds                         *Bool      `xmlrpc:"draft_asset_ids,omitempty"`
	DuplicatedRefIds                      *Relation  `xmlrpc:"duplicated_ref_ids,omitempty"`
	EdiBlockingLevel                      *Selection `xmlrpc:"edi_blocking_level,omitempty"`
	EdiDocumentIds                        *Relation  `xmlrpc:"edi_document_ids,omitempty"`
	EdiErrorCount                         *Int       `xmlrpc:"edi_error_count,omitempty"`
	EdiErrorMessage                       *String    `xmlrpc:"edi_error_message,omitempty"`
	EdiShowAbandonCancelButton            *Bool      `xmlrpc:"edi_show_abandon_cancel_button,omitempty"`
	EdiShowCancelButton                   *Bool      `xmlrpc:"edi_show_cancel_button,omitempty"`
	EdiState                              *Selection `xmlrpc:"edi_state,omitempty"`
	EdiWebServicesToProcess               *String    `xmlrpc:"edi_web_services_to_process,omitempty"`
	ExtractAttachmentId                   *Many2One  `xmlrpc:"extract_attachment_id,omitempty"`
	ExtractCanShowBanners                 *Bool      `xmlrpc:"extract_can_show_banners,omitempty"`
	ExtractCanShowResendButton            *Bool      `xmlrpc:"extract_can_show_resend_button,omitempty"`
	ExtractCanShowSendButton              *Bool      `xmlrpc:"extract_can_show_send_button,omitempty"`
	ExtractErrorMessage                   *String    `xmlrpc:"extract_error_message,omitempty"`
	ExtractRemoteId                       *Int       `xmlrpc:"extract_remote_id,omitempty"`
	ExtractState                          *Selection `xmlrpc:"extract_state,omitempty"`
	ExtractStatusCode                     *Int       `xmlrpc:"extract_status_code,omitempty"`
	ExtractWordIds                        *Relation  `xmlrpc:"extract_word_ids,omitempty"`
	FiscalPositionId                      *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	HasMessage                            *Bool      `xmlrpc:"has_message,omitempty"`
	HasReconciledEntries                  *Bool      `xmlrpc:"has_reconciled_entries,omitempty"`
	HidePostButton                        *Bool      `xmlrpc:"hide_post_button,omitempty"`
	HighestName                           *String    `xmlrpc:"highest_name,omitempty"`
	Id                                    *Int       `xmlrpc:"id,omitempty"`
	InalterableHash                       *String    `xmlrpc:"inalterable_hash,omitempty"`
	InvoiceCashRoundingId                 *Many2One  `xmlrpc:"invoice_cash_rounding_id,omitempty"`
	InvoiceDate                           *Time      `xmlrpc:"invoice_date,omitempty"`
	InvoiceDateDue                        *Time      `xmlrpc:"invoice_date_due,omitempty"`
	InvoiceFilterTypeDomain               *String    `xmlrpc:"invoice_filter_type_domain,omitempty"`
	InvoiceHasOutstanding                 *Bool      `xmlrpc:"invoice_has_outstanding,omitempty"`
	InvoiceIncotermId                     *Many2One  `xmlrpc:"invoice_incoterm_id,omitempty"`
	InvoiceLineIds                        *Relation  `xmlrpc:"invoice_line_ids,omitempty"`
	InvoiceOrigin                         *String    `xmlrpc:"invoice_origin,omitempty"`
	InvoiceOutstandingCreditsDebitsWidget *String    `xmlrpc:"invoice_outstanding_credits_debits_widget,omitempty"`
	InvoicePartnerDisplayName             *String    `xmlrpc:"invoice_partner_display_name,omitempty"`
	InvoicePaymentTermId                  *Many2One  `xmlrpc:"invoice_payment_term_id,omitempty"`
	InvoicePaymentsWidget                 *String    `xmlrpc:"invoice_payments_widget,omitempty"`
	InvoiceSourceEmail                    *String    `xmlrpc:"invoice_source_email,omitempty"`
	InvoiceUserId                         *Many2One  `xmlrpc:"invoice_user_id,omitempty"`
	InvoiceVendorBillId                   *Many2One  `xmlrpc:"invoice_vendor_bill_id,omitempty"`
	IsMoveSent                            *Bool      `xmlrpc:"is_move_sent,omitempty"`
	IsStorno                              *Bool      `xmlrpc:"is_storno,omitempty"`
	JournalId                             *Many2One  `xmlrpc:"journal_id,omitempty"`
	LineIds                               *Relation  `xmlrpc:"line_ids,omitempty"`
	LinkedAssetType                       *String    `xmlrpc:"linked_asset_type,omitempty"`
	MadeSequenceHole                      *Bool      `xmlrpc:"made_sequence_hole,omitempty"`
	MediumId                              *Many2One  `xmlrpc:"medium_id,omitempty"`
	MessageAttachmentCount                *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                    *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                       *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                    *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                            *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                     *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId               *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction                     *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter              *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                     *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MoveType                              *Selection `xmlrpc:"move_type,omitempty"`
	MyActivityDateDeadline                *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                  *String    `xmlrpc:"name,omitempty"`
	Narration                             *String    `xmlrpc:"narration,omitempty"`
	NeededTerms                           *String    `xmlrpc:"needed_terms,omitempty"`
	NeededTermsDirty                      *Bool      `xmlrpc:"needed_terms_dirty,omitempty"`
	NumberAssetIds                        *Int       `xmlrpc:"number_asset_ids,omitempty"`
	PartnerBankId                         *Many2One  `xmlrpc:"partner_bank_id,omitempty"`
	PartnerCreditWarning                  *String    `xmlrpc:"partner_credit_warning,omitempty"`
	PartnerId                             *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerShippingId                     *Many2One  `xmlrpc:"partner_shipping_id,omitempty"`
	PaymentId                             *Many2One  `xmlrpc:"payment_id,omitempty"`
	PaymentIds                            *Relation  `xmlrpc:"payment_ids,omitempty"`
	PaymentReference                      *String    `xmlrpc:"payment_reference,omitempty"`
	PaymentState                          *Selection `xmlrpc:"payment_state,omitempty"`
	PaymentStateBeforeSwitch              *String    `xmlrpc:"payment_state_before_switch,omitempty"`
	PaymentTermDetails                    *String    `xmlrpc:"payment_term_details,omitempty"`
	PosOrderIds                           *Relation  `xmlrpc:"pos_order_ids,omitempty"`
	PosPaymentIds                         *Relation  `xmlrpc:"pos_payment_ids,omitempty"`
	PostedBefore                          *Bool      `xmlrpc:"posted_before,omitempty"`
	PurchaseId                            *Many2One  `xmlrpc:"purchase_id,omitempty"`
	PurchaseOrderCount                    *Int       `xmlrpc:"purchase_order_count,omitempty"`
	PurchaseVendorBillId                  *Many2One  `xmlrpc:"purchase_vendor_bill_id,omitempty"`
	QrCodeMethod                          *Selection `xmlrpc:"qr_code_method,omitempty"`
	QuickEditMode                         *Bool      `xmlrpc:"quick_edit_mode,omitempty"`
	QuickEditTotalAmount                  *Float     `xmlrpc:"quick_edit_total_amount,omitempty"`
	QuickEncodingVals                     *String    `xmlrpc:"quick_encoding_vals,omitempty"`
	Ref                                   *String    `xmlrpc:"ref,omitempty"`
	RestrictModeHashTable                 *Bool      `xmlrpc:"restrict_mode_hash_table,omitempty"`
	ReversalMoveId                        *Relation  `xmlrpc:"reversal_move_id,omitempty"`
	ReversedEntryId                       *Many2One  `xmlrpc:"reversed_entry_id,omitempty"`
	SaleOrderCount                        *Int       `xmlrpc:"sale_order_count,omitempty"`
	SecureSequenceNumber                  *Int       `xmlrpc:"secure_sequence_number,omitempty"`
	SequenceNumber                        *Int       `xmlrpc:"sequence_number,omitempty"`
	SequencePrefix                        *String    `xmlrpc:"sequence_prefix,omitempty"`
	ShowDiscountDetails                   *Bool      `xmlrpc:"show_discount_details,omitempty"`
	ShowNameWarning                       *Bool      `xmlrpc:"show_name_warning,omitempty"`
	ShowPaymentTermDetails                *Bool      `xmlrpc:"show_payment_term_details,omitempty"`
	ShowResetToDraftButton                *Bool      `xmlrpc:"show_reset_to_draft_button,omitempty"`
	SourceId                              *Many2One  `xmlrpc:"source_id,omitempty"`
	State                                 *Selection `xmlrpc:"state,omitempty"`
	StatementLineId                       *Many2One  `xmlrpc:"statement_line_id,omitempty"`
	StatementLineIds                      *Relation  `xmlrpc:"statement_line_ids,omitempty"`
	StockMoveId                           *Many2One  `xmlrpc:"stock_move_id,omitempty"`
	StockValuationLayerIds                *Relation  `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	StringToHash                          *String    `xmlrpc:"string_to_hash,omitempty"`
	SuitableJournalIds                    *Relation  `xmlrpc:"suitable_journal_ids,omitempty"`
	SuspenseStatementLineId               *Many2One  `xmlrpc:"suspense_statement_line_id,omitempty"`
	TaxCashBasisCreatedMoveIds            *Relation  `xmlrpc:"tax_cash_basis_created_move_ids,omitempty"`
	TaxCashBasisOriginMoveId              *Many2One  `xmlrpc:"tax_cash_basis_origin_move_id,omitempty"`
	TaxCashBasisRecId                     *Many2One  `xmlrpc:"tax_cash_basis_rec_id,omitempty"`
	TaxClosingEndDate                     *Time      `xmlrpc:"tax_closing_end_date,omitempty"`
	TaxCountryCode                        *String    `xmlrpc:"tax_country_code,omitempty"`
	TaxCountryId                          *Many2One  `xmlrpc:"tax_country_id,omitempty"`
	TaxLockDateMessage                    *String    `xmlrpc:"tax_lock_date_message,omitempty"`
	TaxReportControlError                 *Bool      `xmlrpc:"tax_report_control_error,omitempty"`
	TaxTotals                             *String    `xmlrpc:"tax_totals,omitempty"`
	TeamId                                *Many2One  `xmlrpc:"team_id,omitempty"`
	TimesheetCount                        *Int       `xmlrpc:"timesheet_count,omitempty"`
	TimesheetEncodeUomId                  *Many2One  `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	TimesheetIds                          *Relation  `xmlrpc:"timesheet_ids,omitempty"`
	TimesheetTotalDuration                *Int       `xmlrpc:"timesheet_total_duration,omitempty"`
	ToCheck                               *Bool      `xmlrpc:"to_check,omitempty"`
	TransactionIds                        *Relation  `xmlrpc:"transaction_ids,omitempty"`
	TransferModelId                       *Many2One  `xmlrpc:"transfer_model_id,omitempty"`
	TypeName                              *String    `xmlrpc:"type_name,omitempty"`
	UserId                                *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteId                             *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds                     *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omitempty"`
	XStudioDateDeDbut                     *Time      `xmlrpc:"x_studio_date_de_dbut,omitempty"`
	XStudioDateDeDbutDvnement             *Time      `xmlrpc:"x_studio_date_de_dbut_dvnement,omitempty"`
	XStudioDateDeFin                      *Time      `xmlrpc:"x_studio_date_de_fin,omitempty"`
	XStudioDateDeLevenement               *Time      `xmlrpc:"x_studio_date_de_levenement,omitempty"`
	XStudioDateFinDvnement                *Time      `xmlrpc:"x_studio_date_fin_dvnement,omitempty"`
	XStudioEtatFacture                    *Selection `xmlrpc:"x_studio_etat_facture,omitempty"`
	XStudioFactureYooz                    *String    `xmlrpc:"x_studio_facture_yooz,omitempty"`
	XStudioInformationsComplmentaires     *String    `xmlrpc:"x_studio_informations_complmentaires,omitempty"`
	XStudioJournalType                    *Selection `xmlrpc:"x_studio_journal_type,omitempty"`
	XStudioNPiceYooz                      *String    `xmlrpc:"x_studio_n_pice_yooz,omitempty"`
	XStudioNfacture                       *String    `xmlrpc:"x_studio_nfacture,omitempty"`
	XStudioNumroDeDevis                   *String    `xmlrpc:"x_studio_numro_de_devis,omitempty"`
	XStudioOne2ManyFieldHFxD2             *Relation  `xmlrpc:"x_studio_one2many_field_hFxD2,omitempty"`
	XStudioRelatedFieldQ8UEh              *String    `xmlrpc:"x_studio_related_field_Q8uEh,omitempty"`
	XStudioRelatedFieldQqfPp              *String    `xmlrpc:"x_studio_related_field_qqfPp,omitempty"`
	XStudioRelatedFieldSo6NV              *String    `xmlrpc:"x_studio_related_field_so6nV,omitempty"`
}

// AccountMoves represents array of account.move model.
type AccountMoves []AccountMove

// AccountMoveModel is the odoo model name.
const AccountMoveModel = "account.move"

// Many2One convert AccountMove to *Many2One.
func (am *AccountMove) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAccountMove creates a new account.move model and returns its id.
func (c *Client) CreateAccountMove(am *AccountMove) (int64, error) {
	ids, err := c.CreateAccountMoves([]*AccountMove{am})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMove creates a new account.move model and returns its id.
func (c *Client) CreateAccountMoves(ams []*AccountMove) ([]int64, error) {
	var vv []interface{}
	for _, v := range ams {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveModel, vv, nil)
}

// UpdateAccountMove updates an existing account.move record.
func (c *Client) UpdateAccountMove(am *AccountMove) error {
	return c.UpdateAccountMoves([]int64{am.Id.Get()}, am)
}

// UpdateAccountMoves updates existing account.move records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAccountMoves(ids []int64, am *AccountMove) error {
	return c.Update(AccountMoveModel, ids, am, nil)
}

// DeleteAccountMove deletes an existing account.move record.
func (c *Client) DeleteAccountMove(id int64) error {
	return c.DeleteAccountMoves([]int64{id})
}

// DeleteAccountMoves deletes existing account.move records.
func (c *Client) DeleteAccountMoves(ids []int64) error {
	return c.Delete(AccountMoveModel, ids)
}

// GetAccountMove gets account.move existing record.
func (c *Client) GetAccountMove(id int64) (*AccountMove, error) {
	ams, err := c.GetAccountMoves([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// GetAccountMoves gets account.move existing records.
func (c *Client) GetAccountMoves(ids []int64) (*AccountMoves, error) {
	ams := &AccountMoves{}
	if err := c.Read(AccountMoveModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountMove finds account.move record by querying it with criteria.
func (c *Client) FindAccountMove(criteria *Criteria) (*AccountMove, error) {
	ams := &AccountMoves{}
	if err := c.SearchRead(AccountMoveModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// FindAccountMoves finds account.move records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoves(criteria *Criteria, options *Options) (*AccountMoves, error) {
	ams := &AccountMoves{}
	if err := c.SearchRead(AccountMoveModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountMoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveModel, criteria, options)
}

// FindAccountMoveId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
