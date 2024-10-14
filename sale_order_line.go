package odoo

// SaleOrderLine represents sale.order.line model.
type SaleOrderLine struct {
	LastUpdate                        *Time       `xmlrpc:"__last_update,omitempty"`
	AnalyticDistribution              interface{} `xmlrpc:"analytic_distribution,omitempty"`
	AnalyticDistributionSearch        interface{} `xmlrpc:"analytic_distribution_search,omitempty"`
	AnalyticLineIds                   *Relation   `xmlrpc:"analytic_line_ids,omitempty"`
	AnalyticPrecision                 *Int        `xmlrpc:"analytic_precision,omitempty"`
	CompanyId                         *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                        *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                         *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                        *Many2One   `xmlrpc:"currency_id,omitempty"`
	CustomerLead                      *Float      `xmlrpc:"customer_lead,omitempty"`
	Discount                          *Float      `xmlrpc:"discount,omitempty"`
	DisplayName                       *String     `xmlrpc:"display_name,omitempty"`
	DisplayQtyWidget                  *Bool       `xmlrpc:"display_qty_widget,omitempty"`
	DisplayType                       *Selection  `xmlrpc:"display_type,omitempty"`
	ForceSaleMaxQty                   *Bool       `xmlrpc:"force_sale_max_qty,omitempty"`
	ForceSaleMinQty                   *Bool       `xmlrpc:"force_sale_min_qty,omitempty"`
	ForecastExpectedDate              *Time       `xmlrpc:"forecast_expected_date,omitempty"`
	FreeQtyToday                      *Float      `xmlrpc:"free_qty_today,omitempty"`
	HasDisplayedWarningUpsell         *Bool       `xmlrpc:"has_displayed_warning_upsell,omitempty"`
	Id                                *Int        `xmlrpc:"id,omitempty"`
	InvoiceLines                      *Relation   `xmlrpc:"invoice_lines,omitempty"`
	InvoiceStatus                     *Selection  `xmlrpc:"invoice_status,omitempty"`
	IsConfigurableProduct             *Bool       `xmlrpc:"is_configurable_product,omitempty"`
	IsDownpayment                     *Bool       `xmlrpc:"is_downpayment,omitempty"`
	IsExpense                         *Bool       `xmlrpc:"is_expense,omitempty"`
	IsMto                             *Bool       `xmlrpc:"is_mto,omitempty"`
	IsQtyBiggerMaxQty                 *Bool       `xmlrpc:"is_qty_bigger_max_qty,omitempty"`
	IsQtyLessMinQty                   *Bool       `xmlrpc:"is_qty_less_min_qty,omitempty"`
	IsQtyNotMultipleQty               *Bool       `xmlrpc:"is_qty_not_multiple_qty,omitempty"`
	IsService                         *Bool       `xmlrpc:"is_service,omitempty"`
	LinkedLineId                      *Many2One   `xmlrpc:"linked_line_id,omitempty"`
	MoveIds                           *Relation   `xmlrpc:"move_ids,omitempty"`
	Name                              *String     `xmlrpc:"name,omitempty"`
	NameShort                         *String     `xmlrpc:"name_short,omitempty"`
	OptionLineIds                     *Relation   `xmlrpc:"option_line_ids,omitempty"`
	OrderId                           *Many2One   `xmlrpc:"order_id,omitempty"`
	OrderPartnerId                    *Many2One   `xmlrpc:"order_partner_id,omitempty"`
	PartnerId                         *Many2One   `xmlrpc:"partner_id,omitempty"`
	PosOrderLineIds                   *Relation   `xmlrpc:"pos_order_line_ids,omitempty"`
	PriceReduce                       *Float      `xmlrpc:"price_reduce,omitempty"`
	PriceReduceTaxexcl                *Float      `xmlrpc:"price_reduce_taxexcl,omitempty"`
	PriceReduceTaxinc                 *Float      `xmlrpc:"price_reduce_taxinc,omitempty"`
	PriceSubtotal                     *Float      `xmlrpc:"price_subtotal,omitempty"`
	PriceTax                          *Float      `xmlrpc:"price_tax,omitempty"`
	PriceTotal                        *Float      `xmlrpc:"price_total,omitempty"`
	PriceUnit                         *Float      `xmlrpc:"price_unit,omitempty"`
	PricelistItemId                   *Many2One   `xmlrpc:"pricelist_item_id,omitempty"`
	ProductAddMode                    *Selection  `xmlrpc:"product_add_mode,omitempty"`
	ProductCustomAttributeValueIds    *Relation   `xmlrpc:"product_custom_attribute_value_ids,omitempty"`
	ProductId                         *Many2One   `xmlrpc:"product_id,omitempty"`
	ProductNoVariantAttributeValueIds *Relation   `xmlrpc:"product_no_variant_attribute_value_ids,omitempty"`
	ProductPackagingId                *Many2One   `xmlrpc:"product_packaging_id,omitempty"`
	ProductPackagingQty               *Float      `xmlrpc:"product_packaging_qty,omitempty"`
	ProductTemplateAttributeValueIds  *Relation   `xmlrpc:"product_template_attribute_value_ids,omitempty"`
	ProductTemplateId                 *Many2One   `xmlrpc:"product_template_id,omitempty"`
	ProductType                       *Selection  `xmlrpc:"product_type,omitempty"`
	ProductUom                        *Many2One   `xmlrpc:"product_uom,omitempty"`
	ProductUomCategoryId              *Many2One   `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomQty                     *Float      `xmlrpc:"product_uom_qty,omitempty"`
	ProductUomReadonly                *Bool       `xmlrpc:"product_uom_readonly,omitempty"`
	ProductUpdatable                  *Bool       `xmlrpc:"product_updatable,omitempty"`
	ProjectId                         *Many2One   `xmlrpc:"project_id,omitempty"`
	PurchaseLineCount                 *Int        `xmlrpc:"purchase_line_count,omitempty"`
	PurchaseLineIds                   *Relation   `xmlrpc:"purchase_line_ids,omitempty"`
	QtyAvailableToday                 *Float      `xmlrpc:"qty_available_today,omitempty"`
	QtyDelivered                      *Float      `xmlrpc:"qty_delivered,omitempty"`
	QtyDeliveredMethod                *Selection  `xmlrpc:"qty_delivered_method,omitempty"`
	QtyInvoiced                       *Float      `xmlrpc:"qty_invoiced,omitempty"`
	QtyToDeliver                      *Float      `xmlrpc:"qty_to_deliver,omitempty"`
	QtyToInvoice                      *Float      `xmlrpc:"qty_to_invoice,omitempty"`
	ReachedMilestonesIds              *Relation   `xmlrpc:"reached_milestones_ids,omitempty"`
	RemainingHours                    *Float      `xmlrpc:"remaining_hours,omitempty"`
	RemainingHoursAvailable           *Bool       `xmlrpc:"remaining_hours_available,omitempty"`
	ResourceBookingId                 *Many2One   `xmlrpc:"resource_booking_id,omitempty"`
	ResourceBookingIds                *Relation   `xmlrpc:"resource_booking_ids,omitempty"`
	ResourceBookingInfo               *String     `xmlrpc:"resource_booking_info,omitempty"`
	RouteId                           *Many2One   `xmlrpc:"route_id,omitempty"`
	SaleMaxQty                        *Float      `xmlrpc:"sale_max_qty,omitempty"`
	SaleMinQty                        *Float      `xmlrpc:"sale_min_qty,omitempty"`
	SaleMultipleQty                   *Float      `xmlrpc:"sale_multiple_qty,omitempty"`
	SaleOrderOptionIds                *Relation   `xmlrpc:"sale_order_option_ids,omitempty"`
	SalesmanId                        *Many2One   `xmlrpc:"salesman_id,omitempty"`
	ScheduledDate                     *Time       `xmlrpc:"scheduled_date,omitempty"`
	Sequence                          *Int        `xmlrpc:"sequence,omitempty"`
	ShopWarning                       *String     `xmlrpc:"shop_warning,omitempty"`
	State                             *Selection  `xmlrpc:"state,omitempty"`
	TaskId                            *Many2One   `xmlrpc:"task_id,omitempty"`
	TaxId                             *Relation   `xmlrpc:"tax_id,omitempty"`
	TemporalType                      *Selection  `xmlrpc:"temporal_type,omitempty"`
	TimesheetIds                      *Relation   `xmlrpc:"timesheet_ids,omitempty"`
	UntaxedAmountInvoiced             *Float      `xmlrpc:"untaxed_amount_invoiced,omitempty"`
	UntaxedAmountToInvoice            *Float      `xmlrpc:"untaxed_amount_to_invoice,omitempty"`
	VirtualAvailableAtDate            *Float      `xmlrpc:"virtual_available_at_date,omitempty"`
	WarehouseId                       *Many2One   `xmlrpc:"warehouse_id,omitempty"`
	WebsiteDescription                *String     `xmlrpc:"website_description,omitempty"`
	WriteDate                         *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                          *Many2One   `xmlrpc:"write_uid,omitempty"`
	XStudioFonction                   *String     `xmlrpc:"x_studio_fonction,omitempty"`
	XStudioHeureDeDbut                *Time       `xmlrpc:"x_studio_heure_de_dbut,omitempty"`
	XStudioSelectionFieldQ9JAw        *Selection  `xmlrpc:"x_studio_selection_field_q9JAw,omitempty"`
}

// SaleOrderLines represents array of sale.order.line model.
type SaleOrderLines []SaleOrderLine

// SaleOrderLineModel is the odoo model name.
const SaleOrderLineModel = "sale.order.line"

// Many2One convert SaleOrderLine to *Many2One.
func (sol *SaleOrderLine) Many2One() *Many2One {
	return NewMany2One(sol.Id.Get(), "")
}

// CreateSaleOrderLine creates a new sale.order.line model and returns its id.
func (c *Client) CreateSaleOrderLine(sol *SaleOrderLine) (int64, error) {
	ids, err := c.CreateSaleOrderLines([]*SaleOrderLine{sol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderLine creates a new sale.order.line model and returns its id.
func (c *Client) CreateSaleOrderLines(sols []*SaleOrderLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range sols {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderLineModel, vv, nil)
}

// UpdateSaleOrderLine updates an existing sale.order.line record.
func (c *Client) UpdateSaleOrderLine(sol *SaleOrderLine) error {
	return c.UpdateSaleOrderLines([]int64{sol.Id.Get()}, sol)
}

// UpdateSaleOrderLines updates existing sale.order.line records.
// All records (represented by ids) will be updated by sol values.
func (c *Client) UpdateSaleOrderLines(ids []int64, sol *SaleOrderLine) error {
	return c.Update(SaleOrderLineModel, ids, sol, nil)
}

// DeleteSaleOrderLine deletes an existing sale.order.line record.
func (c *Client) DeleteSaleOrderLine(id int64) error {
	return c.DeleteSaleOrderLines([]int64{id})
}

// DeleteSaleOrderLines deletes existing sale.order.line records.
func (c *Client) DeleteSaleOrderLines(ids []int64) error {
	return c.Delete(SaleOrderLineModel, ids)
}

// GetSaleOrderLine gets sale.order.line existing record.
func (c *Client) GetSaleOrderLine(id int64) (*SaleOrderLine, error) {
	sols, err := c.GetSaleOrderLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sols)[0]), nil
}

// GetSaleOrderLines gets sale.order.line existing records.
func (c *Client) GetSaleOrderLines(ids []int64) (*SaleOrderLines, error) {
	sols := &SaleOrderLines{}
	if err := c.Read(SaleOrderLineModel, ids, nil, sols); err != nil {
		return nil, err
	}
	return sols, nil
}

// FindSaleOrderLine finds sale.order.line record by querying it with criteria.
func (c *Client) FindSaleOrderLine(criteria *Criteria) (*SaleOrderLine, error) {
	sols := &SaleOrderLines{}
	if err := c.SearchRead(SaleOrderLineModel, criteria, NewOptions().Limit(1), sols); err != nil {
		return nil, err
	}
	return &((*sols)[0]), nil
}

// FindSaleOrderLines finds sale.order.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderLines(criteria *Criteria, options *Options) (*SaleOrderLines, error) {
	sols := &SaleOrderLines{}
	if err := c.SearchRead(SaleOrderLineModel, criteria, options, sols); err != nil {
		return nil, err
	}
	return sols, nil
}

// FindSaleOrderLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleOrderLineModel, criteria, options)
}

// FindSaleOrderLineId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
