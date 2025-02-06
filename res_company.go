package odoo

// ResCompany represents res.company model.
type ResCompany struct {
	LastUpdate                                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccountCashBasisBaseAccountId               *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omitempty"`
	AccountDashboardOnboardingState             *Selection `xmlrpc:"account_dashboard_onboarding_state,omitempty"`
	AccountDefaultPosReceivableAccountId        *Many2One  `xmlrpc:"account_default_pos_receivable_account_id,omitempty"`
	AccountDisplayRepresentativeField           *Bool      `xmlrpc:"account_display_representative_field,omitempty"`
	AccountEnabledTaxCountryIds                 *Relation  `xmlrpc:"account_enabled_tax_country_ids,omitempty"`
	AccountFiscalCountryId                      *Many2One  `xmlrpc:"account_fiscal_country_id,omitempty"`
	AccountFolder                               *Many2One  `xmlrpc:"account_folder,omitempty"`
	AccountInvoiceOnboardingState               *Selection `xmlrpc:"account_invoice_onboarding_state,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountJournalPaymentCreditAccountId        *Many2One  `xmlrpc:"account_journal_payment_credit_account_id,omitempty"`
	AccountJournalPaymentDebitAccountId         *Many2One  `xmlrpc:"account_journal_payment_debit_account_id,omitempty"`
	AccountJournalSuspenseAccountId             *Many2One  `xmlrpc:"account_journal_suspense_account_id,omitempty"`
	AccountOnboardingCreateInvoiceState         *Selection `xmlrpc:"account_onboarding_create_invoice_state,omitempty"`
	AccountOnboardingCreateInvoiceStateFlag     *Bool      `xmlrpc:"account_onboarding_create_invoice_state_flag,omitempty"`
	AccountOnboardingInvoiceLayoutState         *Selection `xmlrpc:"account_onboarding_invoice_layout_state,omitempty"`
	AccountOnboardingSaleTaxState               *Selection `xmlrpc:"account_onboarding_sale_tax_state,omitempty"`
	AccountOpeningDate                          *Time      `xmlrpc:"account_opening_date,omitempty"`
	AccountOpeningJournalId                     *Many2One  `xmlrpc:"account_opening_journal_id,omitempty"`
	AccountOpeningMoveId                        *Many2One  `xmlrpc:"account_opening_move_id,omitempty"`
	AccountPurchaseTaxId                        *Many2One  `xmlrpc:"account_purchase_tax_id,omitempty"`
	AccountRepresentativeId                     *Many2One  `xmlrpc:"account_representative_id,omitempty"`
	AccountRevaluationExpenseProvisionAccountId *Many2One  `xmlrpc:"account_revaluation_expense_provision_account_id,omitempty"`
	AccountRevaluationIncomeProvisionAccountId  *Many2One  `xmlrpc:"account_revaluation_income_provision_account_id,omitempty"`
	AccountRevaluationJournalId                 *Many2One  `xmlrpc:"account_revaluation_journal_id,omitempty"`
	AccountSaleTaxId                            *Many2One  `xmlrpc:"account_sale_tax_id,omitempty"`
	AccountSepaLei                              *String    `xmlrpc:"account_sepa_lei,omitempty"`
	AccountSetupBankDataState                   *Selection `xmlrpc:"account_setup_bank_data_state,omitempty"`
	AccountSetupBillState                       *Selection `xmlrpc:"account_setup_bill_state,omitempty"`
	AccountSetupCoaState                        *Selection `xmlrpc:"account_setup_coa_state,omitempty"`
	AccountSetupFyDataState                     *Selection `xmlrpc:"account_setup_fy_data_state,omitempty"`
	AccountSetupTaxesState                      *Selection `xmlrpc:"account_setup_taxes_state,omitempty"`
	AccountStorno                               *Bool      `xmlrpc:"account_storno,omitempty"`
	AccountTaxPeriodicity                       *Selection `xmlrpc:"account_tax_periodicity,omitempty"`
	AccountTaxPeriodicityJournalId              *Many2One  `xmlrpc:"account_tax_periodicity_journal_id,omitempty"`
	AccountTaxPeriodicityReminderDay            *Int       `xmlrpc:"account_tax_periodicity_reminder_day,omitempty"`
	AccountTaxUnitIds                           *Relation  `xmlrpc:"account_tax_unit_ids,omitempty"`
	AccountUseCreditLimit                       *Bool      `xmlrpc:"account_use_credit_limit,omitempty"`
	Active                                      *Bool      `xmlrpc:"active,omitempty"`
	AnalyticPlanId                              *Many2One  `xmlrpc:"analytic_plan_id,omitempty"`
	AngloSaxonAccounting                        *Bool      `xmlrpc:"anglo_saxon_accounting,omitempty"`
	AnnualInventoryDay                          *Int       `xmlrpc:"annual_inventory_day,omitempty"`
	AnnualInventoryMonth                        *Selection `xmlrpc:"annual_inventory_month,omitempty"`
	Ape                                         *String    `xmlrpc:"ape,omitempty"`
	AutomaticEntryDefaultJournalId              *Many2One  `xmlrpc:"automatic_entry_default_journal_id,omitempty"`
	BackgroundImage                             *String    `xmlrpc:"background_image,omitempty"`
	BankAccountCodePrefix                       *String    `xmlrpc:"bank_account_code_prefix,omitempty"`
	BankIds                                     *Relation  `xmlrpc:"bank_ids,omitempty"`
	BankJournalIds                              *Relation  `xmlrpc:"bank_journal_ids,omitempty"`
	BaseOnboardingCompanyState                  *Selection `xmlrpc:"base_onboarding_company_state,omitempty"`
	CashAccountCodePrefix                       *String    `xmlrpc:"cash_account_code_prefix,omitempty"`
	CatchallEmail                               *String    `xmlrpc:"catchall_email,omitempty"`
	CatchallFormatted                           *String    `xmlrpc:"catchall_formatted,omitempty"`
	ChartTemplateId                             *Many2One  `xmlrpc:"chart_template_id,omitempty"`
	ChildIds                                    *Relation  `xmlrpc:"child_ids,omitempty"`
	City                                        *String    `xmlrpc:"city,omitempty"`
	CompanyDetails                              *String    `xmlrpc:"company_details,omitempty"`
	CompanyRegistry                             *String    `xmlrpc:"company_registry,omitempty"`
	CountryCode                                 *String    `xmlrpc:"country_code,omitempty"`
	CountryId                                   *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyExchangeJournalId                   *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrencyIntervalUnit                        *Selection `xmlrpc:"currency_interval_unit,omitempty"`
	CurrencyNextExecutionDate                   *Time      `xmlrpc:"currency_next_execution_date,omitempty"`
	CurrencyProvider                            *Selection `xmlrpc:"currency_provider,omitempty"`
	DaysToPurchase                              *Float     `xmlrpc:"days_to_purchase,omitempty"`
	DefaultCashDifferenceExpenseAccountId       *Many2One  `xmlrpc:"default_cash_difference_expense_account_id,omitempty"`
	DefaultCashDifferenceIncomeAccountId        *Many2One  `xmlrpc:"default_cash_difference_income_account_id,omitempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omitempty"`
	DocumentsAccountSettings                    *Bool      `xmlrpc:"documents_account_settings,omitempty"`
	DocumentsHrFolder                           *Many2One  `xmlrpc:"documents_hr_folder,omitempty"`
	DocumentsHrSettings                         *Bool      `xmlrpc:"documents_hr_settings,omitempty"`
	DocumentsProductSettings                    *Bool      `xmlrpc:"documents_product_settings,omitempty"`
	DocumentsSpreadsheetFolderId                *Many2One  `xmlrpc:"documents_spreadsheet_folder_id,omitempty"`
	EarlyPayDiscountComputation                 *Selection `xmlrpc:"early_pay_discount_computation,omitempty"`
	Email                                       *String    `xmlrpc:"email,omitempty"`
	EmailFormatted                              *String    `xmlrpc:"email_formatted,omitempty"`
	ExpectsChartOfAccounts                      *Bool      `xmlrpc:"expects_chart_of_accounts,omitempty"`
	ExpenseAccrualAccountId                     *Many2One  `xmlrpc:"expense_accrual_account_id,omitempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	ExternalReportLayoutId                      *Many2One  `xmlrpc:"external_report_layout_id,omitempty"`
	ExtractInInvoiceDigitalizationMode          *Selection `xmlrpc:"extract_in_invoice_digitalization_mode,omitempty"`
	ExtractOutInvoiceDigitalizationMode         *Selection `xmlrpc:"extract_out_invoice_digitalization_mode,omitempty"`
	ExtractSingleLinePerTax                     *Bool      `xmlrpc:"extract_single_line_per_tax,omitempty"`
	Favicon                                     *String    `xmlrpc:"favicon,omitempty"`
	FiscalPositionIds                           *Relation  `xmlrpc:"fiscal_position_ids,omitempty"`
	FiscalyearLastDay                           *Int       `xmlrpc:"fiscalyear_last_day,omitempty"`
	FiscalyearLastMonth                         *Selection `xmlrpc:"fiscalyear_last_month,omitempty"`
	FiscalyearLockDate                          *Time      `xmlrpc:"fiscalyear_lock_date,omitempty"`
	Font                                        *Selection `xmlrpc:"font,omitempty"`
	GainAccountId                               *Many2One  `xmlrpc:"gain_account_id,omitempty"`
	HasMessage                                  *Bool      `xmlrpc:"has_message,omitempty"`
	HasReceivedWarningStockSms                  *Bool      `xmlrpc:"has_received_warning_stock_sms,omitempty"`
	HrPresenceControlEmailAmount                *Int       `xmlrpc:"hr_presence_control_email_amount,omitempty"`
	HrPresenceControlIpList                     *String    `xmlrpc:"hr_presence_control_ip_list,omitempty"`
	IapEnrichAutoDone                           *Bool      `xmlrpc:"iap_enrich_auto_done,omitempty"`
	Id                                          *Int       `xmlrpc:"id,omitempty"`
	IncomeCurrencyExchangeAccountId             *Many2One  `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	IncotermId                                  *Many2One  `xmlrpc:"incoterm_id,omitempty"`
	InternalProjectId                           *Many2One  `xmlrpc:"internal_project_id,omitempty"`
	InternalTransitLocationId                   *Many2One  `xmlrpc:"internal_transit_location_id,omitempty"`
	InvoiceIsEmail                              *Bool      `xmlrpc:"invoice_is_email,omitempty"`
	InvoiceIsPrint                              *Bool      `xmlrpc:"invoice_is_print,omitempty"`
	InvoiceIsSnailmail                          *Bool      `xmlrpc:"invoice_is_snailmail,omitempty"`
	InvoiceTerms                                *String    `xmlrpc:"invoice_terms,omitempty"`
	InvoiceTermsHtml                            *String    `xmlrpc:"invoice_terms_html,omitempty"`
	InvoicingSwitchThreshold                    *Time      `xmlrpc:"invoicing_switch_threshold,omitempty"`
	IsCompanyDetailsEmpty                       *Bool      `xmlrpc:"is_company_details_empty,omitempty"`
	IsFranceCountry                             *Bool      `xmlrpc:"is_france_country,omitempty"`
	L10NFrClosingSequenceId                     *Many2One  `xmlrpc:"l10n_fr_closing_sequence_id,omitempty"`
	LayoutBackground                            *Selection `xmlrpc:"layout_background,omitempty"`
	LayoutBackgroundImage                       *String    `xmlrpc:"layout_background_image,omitempty"`
	Logo                                        *String    `xmlrpc:"logo,omitempty"`
	LogoWeb                                     *String    `xmlrpc:"logo_web,omitempty"`
	LossAccountId                               *Many2One  `xmlrpc:"loss_account_id,omitempty"`
	MessageAttachmentCount                      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId                     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction                           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter                    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Mobile                                      *String    `xmlrpc:"mobile,omitempty"`
	MultiVatForeignCountryIds                   *Relation  `xmlrpc:"multi_vat_foreign_country_ids,omitempty"`
	Name                                        *String    `xmlrpc:"name,omitempty"`
	Nic                                         *String    `xmlrpc:"nic,omitempty"`
	NomenclatureId                              *Many2One  `xmlrpc:"nomenclature_id,omitempty"`
	PaperformatId                               *Many2One  `xmlrpc:"paperformat_id,omitempty"`
	ParentId                                    *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerGid                                  *Int       `xmlrpc:"partner_gid,omitempty"`
	PartnerId                                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PaymentOnboardingPaymentMethod              *Selection `xmlrpc:"payment_onboarding_payment_method,omitempty"`
	PaymentProviderOnboardingState              *Selection `xmlrpc:"payment_provider_onboarding_state,omitempty"`
	PeriodLockDate                              *Time      `xmlrpc:"period_lock_date,omitempty"`
	Phone                                       *String    `xmlrpc:"phone,omitempty"`
	PoDoubleValidation                          *Selection `xmlrpc:"po_double_validation,omitempty"`
	PoDoubleValidationAmount                    *Float     `xmlrpc:"po_double_validation_amount,omitempty"`
	PoLead                                      *Float     `xmlrpc:"po_lead,omitempty"`
	PoLock                                      *Selection `xmlrpc:"po_lock,omitempty"`
	PointOfSaleUpdateStockQuantities            *Selection `xmlrpc:"point_of_sale_update_stock_quantities,omitempty"`
	PointOfSaleUseTicketQrCode                  *Bool      `xmlrpc:"point_of_sale_use_ticket_qr_code,omitempty"`
	PortalConfirmationPay                       *Bool      `xmlrpc:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                      *Bool      `xmlrpc:"portal_confirmation_sign,omitempty"`
	PreventOldTimesheetsEncoding                *Bool      `xmlrpc:"prevent_old_timesheets_encoding,omitempty"`
	PrimaryColor                                *String    `xmlrpc:"primary_color,omitempty"`
	ProductFolder                               *Many2One  `xmlrpc:"product_folder,omitempty"`
	ProductTags                                 *Relation  `xmlrpc:"product_tags,omitempty"`
	ProjectTimeModeId                           *Many2One  `xmlrpc:"project_time_mode_id,omitempty"`
	PropertyStockAccountInputCategId            *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omitempty"`
	PropertyStockAccountOutputCategId           *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omitempty"`
	PropertyStockValuationAccountId             *Many2One  `xmlrpc:"property_stock_valuation_account_id,omitempty"`
	QrCode                                      *Bool      `xmlrpc:"qr_code,omitempty"`
	QuickEditMode                               *Selection `xmlrpc:"quick_edit_mode,omitempty"`
	QuotationValidityDays                       *Int       `xmlrpc:"quotation_validity_days,omitempty"`
	ReportFooter                                *String    `xmlrpc:"report_footer,omitempty"`
	ReportHeader                                *String    `xmlrpc:"report_header,omitempty"`
	ResourceCalendarId                          *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceCalendarIds                         *Relation  `xmlrpc:"resource_calendar_ids,omitempty"`
	RevenueAccrualAccountId                     *Many2One  `xmlrpc:"revenue_accrual_account_id,omitempty"`
	SaleOnboardingOrderConfirmationState        *Selection `xmlrpc:"sale_onboarding_order_confirmation_state,omitempty"`
	SaleOnboardingPaymentMethod                 *Selection `xmlrpc:"sale_onboarding_payment_method,omitempty"`
	SaleOnboardingSampleQuotationState          *Selection `xmlrpc:"sale_onboarding_sample_quotation_state,omitempty"`
	SaleOrderTemplateId                         *Many2One  `xmlrpc:"sale_order_template_id,omitempty"`
	SaleQuotationOnboardingState                *Selection `xmlrpc:"sale_quotation_onboarding_state,omitempty"`
	SecondaryColor                              *String    `xmlrpc:"secondary_color,omitempty"`
	SecurityLead                                *Float     `xmlrpc:"security_lead,omitempty"`
	SepaInitiatingPartyName                     *String    `xmlrpc:"sepa_initiating_party_name,omitempty"`
	SepaOrgidId                                 *String    `xmlrpc:"sepa_orgid_id,omitempty"`
	SepaOrgidIssr                               *String    `xmlrpc:"sepa_orgid_issr,omitempty"`
	Sequence                                    *Int       `xmlrpc:"sequence,omitempty"`
	Siren                                       *String    `xmlrpc:"siren,omitempty"`
	Siret                                       *String    `xmlrpc:"siret,omitempty"`
	SnailmailColor                              *Bool      `xmlrpc:"snailmail_color,omitempty"`
	SnailmailCover                              *Bool      `xmlrpc:"snailmail_cover,omitempty"`
	SnailmailDuplex                             *Bool      `xmlrpc:"snailmail_duplex,omitempty"`
	SocialFacebook                              *String    `xmlrpc:"social_facebook,omitempty"`
	SocialGithub                                *String    `xmlrpc:"social_github,omitempty"`
	SocialInstagram                             *String    `xmlrpc:"social_instagram,omitempty"`
	SocialLinkedin                              *String    `xmlrpc:"social_linkedin,omitempty"`
	SocialTwitter                               *String    `xmlrpc:"social_twitter,omitempty"`
	SocialYoutube                               *String    `xmlrpc:"social_youtube,omitempty"`
	StateId                                     *Many2One  `xmlrpc:"state_id,omitempty"`
	StockMailConfirmationTemplateId             *Many2One  `xmlrpc:"stock_mail_confirmation_template_id,omitempty"`
	StockMoveEmailValidation                    *Bool      `xmlrpc:"stock_move_email_validation,omitempty"`
	StockMoveSmsValidation                      *Bool      `xmlrpc:"stock_move_sms_validation,omitempty"`
	StockSmsConfirmationTemplateId              *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omitempty"`
	Street                                      *String    `xmlrpc:"street,omitempty"`
	Street2                                     *String    `xmlrpc:"street2,omitempty"`
	TaxCalculationRoundingMethod                *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                       *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                              *Bool      `xmlrpc:"tax_exigibility,omitempty"`
	TaxLockDate                                 *Time      `xmlrpc:"tax_lock_date,omitempty"`
	TermsType                                   *Selection `xmlrpc:"terms_type,omitempty"`
	TimesheetEncodeUomId                        *Many2One  `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	TimesheetMailEmployeeAllow                  *Bool      `xmlrpc:"timesheet_mail_employee_allow,omitempty"`
	TimesheetMailEmployeeDelay                  *Int       `xmlrpc:"timesheet_mail_employee_delay,omitempty"`
	TimesheetMailEmployeeInterval               *Selection `xmlrpc:"timesheet_mail_employee_interval,omitempty"`
	TimesheetMailEmployeeNextdate               *Time      `xmlrpc:"timesheet_mail_employee_nextdate,omitempty"`
	TimesheetMailManagerAllow                   *Bool      `xmlrpc:"timesheet_mail_manager_allow,omitempty"`
	TimesheetMailManagerDelay                   *Int       `xmlrpc:"timesheet_mail_manager_delay,omitempty"`
	TimesheetMailManagerInterval                *Selection `xmlrpc:"timesheet_mail_manager_interval,omitempty"`
	TimesheetMailManagerNextdate                *Time      `xmlrpc:"timesheet_mail_manager_nextdate,omitempty"`
	TotalsBelowSections                         *Bool      `xmlrpc:"totals_below_sections,omitempty"`
	TransferAccountCodePrefix                   *String    `xmlrpc:"transfer_account_code_prefix,omitempty"`
	TransferAccountId                           *Many2One  `xmlrpc:"transfer_account_id,omitempty"`
	UserIds                                     *Relation  `xmlrpc:"user_ids,omitempty"`
	Vat                                         *String    `xmlrpc:"vat,omitempty"`
	VatCheckVies                                *Bool      `xmlrpc:"vat_check_vies,omitempty"`
	Website                                     *String    `xmlrpc:"website,omitempty"`
	WebsiteId                                   *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds                           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WebsiteSaleDashboardOnboardingState         *Selection `xmlrpc:"website_sale_dashboard_onboarding_state,omitempty"`
	WebsiteSaleOnboardingPaymentProviderState   *Selection `xmlrpc:"website_sale_onboarding_payment_provider_state,omitempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omitempty"`
	XStudioIntercardId                          *Int       `xmlrpc:"x_studio_intercard_id,omitempty"`
	XStudioPiedDePageFacture                    *String    `xmlrpc:"x_studio_pied_de_page_facture,omitempty"`
	XStudioQrcodeAreaId                         *Int       `xmlrpc:"x_studio_qrcode_area_id,omitempty"`
	XStudioShortName                            *String    `xmlrpc:"x_studio_short_name,omitempty"`
	Zip                                         *String    `xmlrpc:"zip,omitempty"`
}

// ResCompanys represents array of res.company model.
type ResCompanys []ResCompany

// ResCompanyModel is the odoo model name.
const ResCompanyModel = "res.company"

// Many2One convert ResCompany to *Many2One.
func (rc *ResCompany) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompany(rc *ResCompany) (int64, error) {
	ids, err := c.CreateResCompanys([]*ResCompany{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompanys(rcs []*ResCompany) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResCompanyModel, vv, nil)
}

// UpdateResCompany updates an existing res.company record.
func (c *Client) UpdateResCompany(rc *ResCompany) error {
	return c.UpdateResCompanys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCompanys updates existing res.company records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCompanys(ids []int64, rc *ResCompany) error {
	return c.Update(ResCompanyModel, ids, rc, nil)
}

// DeleteResCompany deletes an existing res.company record.
func (c *Client) DeleteResCompany(id int64) error {
	return c.DeleteResCompanys([]int64{id})
}

// DeleteResCompanys deletes existing res.company records.
func (c *Client) DeleteResCompanys(ids []int64) error {
	return c.Delete(ResCompanyModel, ids)
}

// GetResCompany gets res.company existing record.
func (c *Client) GetResCompany(id int64) (*ResCompany, error) {
	rcs, err := c.GetResCompanys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// GetResCompanys gets res.company existing records.
func (c *Client) GetResCompanys(ids []int64) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.Read(ResCompanyModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompany finds res.company record by querying it with criteria.
func (c *Client) FindResCompany(criteria *Criteria) (*ResCompany, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// FindResCompanys finds res.company records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanys(criteria *Criteria, options *Options) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompanyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResCompanyModel, criteria, options)
}

// FindResCompanyId finds record id by querying it with criteria.
func (c *Client) FindResCompanyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
