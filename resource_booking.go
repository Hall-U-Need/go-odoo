package odoo

// ResourceBooking represents resource.booking model.
type ResourceBooking struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                 *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                   *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning               *String    `xmlrpc:"access_warning,omitempty"`
	Active                      *Bool      `xmlrpc:"active,omitempty"`
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
	CategIds                    *Relation  `xmlrpc:"categ_ids,omitempty"`
	CombinationAutoAssign       *Bool      `xmlrpc:"combination_auto_assign,omitempty"`
	CombinationId               *Many2One  `xmlrpc:"combination_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Duration                    *Float     `xmlrpc:"duration,omitempty"`
	DurationOnly                *Float     `xmlrpc:"duration_only,omitempty"`
	Expiration                  *Time      `xmlrpc:"expiration,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsModifiable                *Bool      `xmlrpc:"is_modifiable,omitempty"`
	IsOverdue                   *Bool      `xmlrpc:"is_overdue,omitempty"`
	IsPos                       *Bool      `xmlrpc:"is_pos,omitempty"`
	Location                    *String    `xmlrpc:"location,omitempty"`
	MeetingId                   *Many2One  `xmlrpc:"meeting_id,omitempty"`
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
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerIds                  *Relation  `xmlrpc:"partner_ids,omitempty"`
	PosOrderLineId              *Many2One  `xmlrpc:"pos_order_line_id,omitempty"`
	PrereservedEmail            *String    `xmlrpc:"prereserved_email,omitempty"`
	PrereservedName             *String    `xmlrpc:"prereserved_name,omitempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductTemplateIds          *Relation  `xmlrpc:"product_template_ids,omitempty"`
	RequesterAdvice             *String    `xmlrpc:"requester_advice,omitempty"`
	SaleOrderId                 *Many2One  `xmlrpc:"sale_order_id,omitempty"`
	SaleOrderLineId             *Many2One  `xmlrpc:"sale_order_line_id,omitempty"`
	SaleOrderLineIds            *Relation  `xmlrpc:"sale_order_line_ids,omitempty"`
	SaleOrderState              *Selection `xmlrpc:"sale_order_state,omitempty"`
	Start                       *Time      `xmlrpc:"start,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	Stop                        *Time      `xmlrpc:"stop,omitempty"`
	TypeId                      *Many2One  `xmlrpc:"type_id,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResourceBookings represents array of resource.booking model.
type ResourceBookings []ResourceBooking

// ResourceBookingModel is the odoo model name.
const ResourceBookingModel = "resource.booking"

// Many2One convert ResourceBooking to *Many2One.
func (rb *ResourceBooking) Many2One() *Many2One {
	return NewMany2One(rb.Id.Get(), "")
}

// CreateResourceBooking creates a new resource.booking model and returns its id.
func (c *Client) CreateResourceBooking(rb *ResourceBooking) (int64, error) {
	ids, err := c.CreateResourceBookings([]*ResourceBooking{rb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResourceBooking creates a new resource.booking model and returns its id.
func (c *Client) CreateResourceBookings(rbs []*ResourceBooking) ([]int64, error) {
	var vv []interface{}
	for _, v := range rbs {
		vv = append(vv, v)
	}
	return c.Create(ResourceBookingModel, vv, nil)
}

// UpdateResourceBooking updates an existing resource.booking record.
func (c *Client) UpdateResourceBooking(rb *ResourceBooking) error {
	return c.UpdateResourceBookings([]int64{rb.Id.Get()}, rb)
}

// UpdateResourceBookings updates existing resource.booking records.
// All records (represented by ids) will be updated by rb values.
func (c *Client) UpdateResourceBookings(ids []int64, rb *ResourceBooking) error {
	return c.Update(ResourceBookingModel, ids, rb, nil)
}

// DeleteResourceBooking deletes an existing resource.booking record.
func (c *Client) DeleteResourceBooking(id int64) error {
	return c.DeleteResourceBookings([]int64{id})
}

// DeleteResourceBookings deletes existing resource.booking records.
func (c *Client) DeleteResourceBookings(ids []int64) error {
	return c.Delete(ResourceBookingModel, ids)
}

// GetResourceBooking gets resource.booking existing record.
func (c *Client) GetResourceBooking(id int64) (*ResourceBooking, error) {
	rbs, err := c.GetResourceBookings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rbs)[0]), nil
}

// GetResourceBookings gets resource.booking existing records.
func (c *Client) GetResourceBookings(ids []int64) (*ResourceBookings, error) {
	rbs := &ResourceBookings{}
	if err := c.Read(ResourceBookingModel, ids, nil, rbs); err != nil {
		return nil, err
	}
	return rbs, nil
}

// FindResourceBooking finds resource.booking record by querying it with criteria.
func (c *Client) FindResourceBooking(criteria *Criteria) (*ResourceBooking, error) {
	rbs := &ResourceBookings{}
	if err := c.SearchRead(ResourceBookingModel, criteria, NewOptions().Limit(1), rbs); err != nil {
		return nil, err
	}
	return &((*rbs)[0]), nil
}

// FindResourceBookings finds resource.booking records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceBookings(criteria *Criteria, options *Options) (*ResourceBookings, error) {
	rbs := &ResourceBookings{}
	if err := c.SearchRead(ResourceBookingModel, criteria, options, rbs); err != nil {
		return nil, err
	}
	return rbs, nil
}

// FindResourceBookingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceBookingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResourceBookingModel, criteria, options)
}

// FindResourceBookingId finds record id by querying it with criteria.
func (c *Client) FindResourceBookingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResourceBookingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
