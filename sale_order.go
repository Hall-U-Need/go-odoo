package odoo

// SaleOrder represents sale.order model.
type SaleOrder struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                 *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                   *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning               *String    `xmlrpc:"access_warning,omitempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AmountTax                   *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTotal                 *Float     `xmlrpc:"amount_total,omitempty"`
	AmountUndiscounted          *Float     `xmlrpc:"amount_undiscounted,omitempty"`
	AmountUnpaid                *Float     `xmlrpc:"amount_unpaid,omitempty"`
	AmountUntaxed               *Float     `xmlrpc:"amount_untaxed,omitempty"`
	AnalyticAccountId           *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	AuthorizedTransactionIds    *Relation  `xmlrpc:"authorized_transaction_ids,omitempty"`
	CampaignId                  *Many2One  `xmlrpc:"campaign_id,omitempty"`
	CartExpireDate              *Time      `xmlrpc:"cart_expire_date,omitempty"`
	CartQuantity                *Int       `xmlrpc:"cart_quantity,omitempty"`
	CartRecoveryEmailSent       *Bool      `xmlrpc:"cart_recovery_email_sent,omitempty"`
	ClientOrderRef              *String    `xmlrpc:"client_order_ref,omitempty"`
	CommitmentDate              *Time      `xmlrpc:"commitment_date,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryCode                 *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrencyRate                *Float     `xmlrpc:"currency_rate,omitempty"`
	DateOrder                   *Time      `xmlrpc:"date_order,omitempty"`
	DeliveryCount               *Int       `xmlrpc:"delivery_count,omitempty"`
	DeliveryStatus              *Selection `xmlrpc:"delivery_status,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	EffectiveDate               *Time      `xmlrpc:"effective_date,omitempty"`
	ExpectedDate                *Time      `xmlrpc:"expected_date,omitempty"`
	FiscalPositionId            *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	Grid                        *String    `xmlrpc:"grid,omitempty"`
	GridProductTmplId           *Many2One  `xmlrpc:"grid_product_tmpl_id,omitempty"`
	GridUpdate                  *Bool      `xmlrpc:"grid_update,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	Incoterm                    *Many2One  `xmlrpc:"incoterm,omitempty"`
	IncotermLocation            *String    `xmlrpc:"incoterm_location,omitempty"`
	InternalNote                *String    `xmlrpc:"internal_note,omitempty"`
	InvoiceCount                *Int       `xmlrpc:"invoice_count,omitempty"`
	InvoiceIds                  *Relation  `xmlrpc:"invoice_ids,omitempty"`
	InvoiceStatus               *Selection `xmlrpc:"invoice_status,omitempty"`
	IsAbandonedCart             *Bool      `xmlrpc:"is_abandoned_cart,omitempty"`
	IsExpired                   *Bool      `xmlrpc:"is_expired,omitempty"`
	IsProductMilestone          *Bool      `xmlrpc:"is_product_milestone,omitempty"`
	JsonPopover                 *String    `xmlrpc:"json_popover,omitempty"`
	MediumId                    *Many2One  `xmlrpc:"medium_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MilestoneCount              *Int       `xmlrpc:"milestone_count,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	NbWristbands                *Int       `xmlrpc:"nb_wristbands,omitempty"`
	Note                        *String    `xmlrpc:"note,omitempty"`
	OnlyServices                *Bool      `xmlrpc:"only_services,omitempty"`
	OrderLine                   *Relation  `xmlrpc:"order_line,omitempty"`
	Origin                      *String    `xmlrpc:"origin,omitempty"`
	PartnerCreditWarning        *String    `xmlrpc:"partner_credit_warning,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerInvoiceId            *Many2One  `xmlrpc:"partner_invoice_id,omitempty"`
	PartnerShippingId           *Many2One  `xmlrpc:"partner_shipping_id,omitempty"`
	PaymentTermId               *Many2One  `xmlrpc:"payment_term_id,omitempty"`
	PickingIds                  *Relation  `xmlrpc:"picking_ids,omitempty"`
	PickingPolicy               *Selection `xmlrpc:"picking_policy,omitempty"`
	PosOrderCount               *Int       `xmlrpc:"pos_order_count,omitempty"`
	PosOrderLineIds             *Relation  `xmlrpc:"pos_order_line_ids,omitempty"`
	PricelistId                 *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	ProcurementGroupId          *Many2One  `xmlrpc:"procurement_group_id,omitempty"`
	ProjectCount                *Int       `xmlrpc:"project_count,omitempty"`
	ProjectId                   *Many2One  `xmlrpc:"project_id,omitempty"`
	ProjectIds                  *Relation  `xmlrpc:"project_ids,omitempty"`
	PurchaseOrderCount          *Int       `xmlrpc:"purchase_order_count,omitempty"`
	Reference                   *String    `xmlrpc:"reference,omitempty"`
	ReportGrids                 *Bool      `xmlrpc:"report_grids,omitempty"`
	RequirePayment              *Bool      `xmlrpc:"require_payment,omitempty"`
	RequireSignature            *Bool      `xmlrpc:"require_signature,omitempty"`
	ResourceBookingCount        *Int       `xmlrpc:"resource_booking_count,omitempty"`
	ResourceBookingIds          *Relation  `xmlrpc:"resource_booking_ids,omitempty"`
	SaleOrderOptionIds          *Relation  `xmlrpc:"sale_order_option_ids,omitempty"`
	SaleOrderTemplateId         *Many2One  `xmlrpc:"sale_order_template_id,omitempty"`
	ShopWarning                 *String    `xmlrpc:"shop_warning,omitempty"`
	ShowJsonPopover             *Bool      `xmlrpc:"show_json_popover,omitempty"`
	ShowUpdateFpos              *Bool      `xmlrpc:"show_update_fpos,omitempty"`
	ShowUpdatePricelist         *Bool      `xmlrpc:"show_update_pricelist,omitempty"`
	Signature                   *String    `xmlrpc:"signature,omitempty"`
	SignedBy                    *String    `xmlrpc:"signed_by,omitempty"`
	SignedOn                    *Time      `xmlrpc:"signed_on,omitempty"`
	SourceId                    *Many2One  `xmlrpc:"source_id,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty"`
	TasksCount                  *Int       `xmlrpc:"tasks_count,omitempty"`
	TasksIds                    *Relation  `xmlrpc:"tasks_ids,omitempty"`
	TaxCountryId                *Many2One  `xmlrpc:"tax_country_id,omitempty"`
	TaxTotals                   *String    `xmlrpc:"tax_totals,omitempty"`
	TeamId                      *Many2One  `xmlrpc:"team_id,omitempty"`
	TermsType                   *Selection `xmlrpc:"terms_type,omitempty"`
	TicketCount                 *Int       `xmlrpc:"ticket_count,omitempty"`
	TimesheetCount              *Float     `xmlrpc:"timesheet_count,omitempty"`
	TimesheetEncodeUomId        *Many2One  `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	TimesheetIds                *Relation  `xmlrpc:"timesheet_ids,omitempty"`
	TimesheetTotalDuration      *Int       `xmlrpc:"timesheet_total_duration,omitempty"`
	TransactionIds              *Relation  `xmlrpc:"transaction_ids,omitempty"`
	TypeName                    *String    `xmlrpc:"type_name,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	UserIsSaleManager           *Bool      `xmlrpc:"user_is_sale_manager,omitempty"`
	ValidityDate                *Time      `xmlrpc:"validity_date,omitempty"`
	VisibleProject              *Bool      `xmlrpc:"visible_project,omitempty"`
	WarehouseId                 *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WebsiteDescription          *String    `xmlrpc:"website_description,omitempty"`
	WebsiteId                   *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WebsiteOrderLine            *Relation  `xmlrpc:"website_order_line,omitempty"`
	WristbandsPrinted           *Bool      `xmlrpc:"wristbands_printed,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
	XStudioCharFieldHX6Eb       *String    `xmlrpc:"x_studio_char_field_HX6eb,omitempty"`
	XStudioDateDeDbut           *Time      `xmlrpc:"x_studio_date_de_dbut,omitempty"`
	XStudioDateDeFin            *Time      `xmlrpc:"x_studio_date_de_fin,omitempty"`
	XStudioDatetimeFieldCa4I4   *Time      `xmlrpc:"x_studio_datetime_field_Ca4I4,omitempty"`
	XStudioFraisDeGestion       *Int       `xmlrpc:"x_studio_frais_de_gestion,omitempty"`
	XStudioFraisDeGestion1      *Float     `xmlrpc:"x_studio_frais_de_gestion_1,omitempty"`
	XStudioFraisDeGestionHt     *Float     `xmlrpc:"x_studio_frais_de_gestion_ht,omitempty"`
	XStudioHeureDeDbut          *Time      `xmlrpc:"x_studio_heure_de_dbut,omitempty"`
	XStudioHeureDeFin           *Time      `xmlrpc:"x_studio_heure_de_fin,omitempty"`
	XStudioIntegerFieldX2Zxq    *Int       `xmlrpc:"x_studio_integer_field_x2zxq,omitempty"`
	XStudioMany2ManyField5QHx7  *Relation  `xmlrpc:"x_studio_many2many_field_5QHx7,omitempty"`
	XStudioMany2ManyFieldXIafJ  *Relation  `xmlrpc:"x_studio_many2many_field_XIafJ,omitempty"`
	XStudioMany2ManyFieldB72F8  *Relation  `xmlrpc:"x_studio_many2many_field_b72F8,omitempty"`
	XStudioNbDePersonnes        *Int       `xmlrpc:"x_studio_nb_de_personnes,omitempty"`
	XStudioNomDeLvnement        *String    `xmlrpc:"x_studio_nom_de_lvnement,omitempty"`
	XStudioRelatedField54Zsh    *String    `xmlrpc:"x_studio_related_field_54Zsh,omitempty"`
	XStudioRelatedFieldNpcUd    *Selection `xmlrpc:"x_studio_related_field_NpcUd,omitempty"`
	XStudioRelatedFieldQ8DEg    *Int       `xmlrpc:"x_studio_related_field_Q8DEg,omitempty"`
	XStudioRelatedFieldZYAIr    *String    `xmlrpc:"x_studio_related_field_ZYAIr,omitempty"`
	XStudioRelatedFieldZsriR    *String    `xmlrpc:"x_studio_related_field_ZsriR,omitempty"`
}

// SaleOrders represents array of sale.order model.
type SaleOrders []SaleOrder

// SaleOrderModel is the odoo model name.
const SaleOrderModel = "sale.order"

// Many2One convert SaleOrder to *Many2One.
func (so *SaleOrder) Many2One() *Many2One {
	return NewMany2One(so.Id.Get(), "")
}

// CreateSaleOrder creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrder(so *SaleOrder) (int64, error) {
	ids, err := c.CreateSaleOrders([]*SaleOrder{so})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrder creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrders(sos []*SaleOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range sos {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderModel, vv, nil)
}

// UpdateSaleOrder updates an existing sale.order record.
func (c *Client) UpdateSaleOrder(so *SaleOrder) error {
	return c.UpdateSaleOrders([]int64{so.Id.Get()}, so)
}

// UpdateSaleOrders updates existing sale.order records.
// All records (represented by ids) will be updated by so values.
func (c *Client) UpdateSaleOrders(ids []int64, so *SaleOrder) error {
	return c.Update(SaleOrderModel, ids, so, nil)
}

// DeleteSaleOrder deletes an existing sale.order record.
func (c *Client) DeleteSaleOrder(id int64) error {
	return c.DeleteSaleOrders([]int64{id})
}

// DeleteSaleOrders deletes existing sale.order records.
func (c *Client) DeleteSaleOrders(ids []int64) error {
	return c.Delete(SaleOrderModel, ids)
}

// GetSaleOrder gets sale.order existing record.
func (c *Client) GetSaleOrder(id int64) (*SaleOrder, error) {
	sos, err := c.GetSaleOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sos)[0]), nil
}

// GetSaleOrders gets sale.order existing records.
func (c *Client) GetSaleOrders(ids []int64) (*SaleOrders, error) {
	sos := &SaleOrders{}
	if err := c.Read(SaleOrderModel, ids, nil, sos); err != nil {
		return nil, err
	}
	return sos, nil
}

// FindSaleOrder finds sale.order record by querying it with criteria.
func (c *Client) FindSaleOrder(criteria *Criteria) (*SaleOrder, error) {
	sos := &SaleOrders{}
	if err := c.SearchRead(SaleOrderModel, criteria, NewOptions().Limit(1), sos); err != nil {
		return nil, err
	}
	return &((*sos)[0]), nil
}

// FindSaleOrders finds sale.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrders(criteria *Criteria, options *Options) (*SaleOrders, error) {
	sos := &SaleOrders{}
	if err := c.SearchRead(SaleOrderModel, criteria, options, sos); err != nil {
		return nil, err
	}
	return sos, nil
}

// FindSaleOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleOrderModel, criteria, options)
}

// FindSaleOrderId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
