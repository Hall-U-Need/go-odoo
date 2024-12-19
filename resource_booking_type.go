package odoo

// ResourceBookingType represents resource.booking.type model.
type ResourceBookingType struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
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
	AlarmIds                    *Relation  `xmlrpc:"alarm_ids,omitempty"`
	BookingCount                *Int       `xmlrpc:"booking_count,omitempty"`
	BookingIds                  *Relation  `xmlrpc:"booking_ids,omitempty"`
	CategIds                    *Relation  `xmlrpc:"categ_ids,omitempty"`
	CombinationAssignment       *Selection `xmlrpc:"combination_assignment,omitempty"`
	CombinationRelIds           *Relation  `xmlrpc:"combination_rel_ids,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Duration                    *Float     `xmlrpc:"duration,omitempty"`
	DurationInTimeline          *Float     `xmlrpc:"duration_in_timeline,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	Location                    *String    `xmlrpc:"location,omitempty"`
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
	ModificationsDeadline       *Float     `xmlrpc:"modifications_deadline,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	ProductIds                  *Relation  `xmlrpc:"product_ids,omitempty"`
	RequesterAdvice             *String    `xmlrpc:"requester_advice,omitempty"`
	ResourceCalendarId          *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	SlotDuration                *Float     `xmlrpc:"slot_duration,omitempty"`
	UnavailableAfterDuration    *Float     `xmlrpc:"unavailable_after_duration,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResourceBookingTypes represents array of resource.booking.type model.
type ResourceBookingTypes []ResourceBookingType

// ResourceBookingTypeModel is the odoo model name.
const ResourceBookingTypeModel = "resource.booking.type"

// Many2One convert ResourceBookingType to *Many2One.
func (rbt *ResourceBookingType) Many2One() *Many2One {
	return NewMany2One(rbt.Id.Get(), "")
}

// CreateResourceBookingType creates a new resource.booking.type model and returns its id.
func (c *Client) CreateResourceBookingType(rbt *ResourceBookingType) (int64, error) {
	ids, err := c.CreateResourceBookingTypes([]*ResourceBookingType{rbt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResourceBookingType creates a new resource.booking.type model and returns its id.
func (c *Client) CreateResourceBookingTypes(rbts []*ResourceBookingType) ([]int64, error) {
	var vv []interface{}
	for _, v := range rbts {
		vv = append(vv, v)
	}
	return c.Create(ResourceBookingTypeModel, vv, nil)
}

// UpdateResourceBookingType updates an existing resource.booking.type record.
func (c *Client) UpdateResourceBookingType(rbt *ResourceBookingType) error {
	return c.UpdateResourceBookingTypes([]int64{rbt.Id.Get()}, rbt)
}

// UpdateResourceBookingTypes updates existing resource.booking.type records.
// All records (represented by ids) will be updated by rbt values.
func (c *Client) UpdateResourceBookingTypes(ids []int64, rbt *ResourceBookingType) error {
	return c.Update(ResourceBookingTypeModel, ids, rbt, nil)
}

// DeleteResourceBookingType deletes an existing resource.booking.type record.
func (c *Client) DeleteResourceBookingType(id int64) error {
	return c.DeleteResourceBookingTypes([]int64{id})
}

// DeleteResourceBookingTypes deletes existing resource.booking.type records.
func (c *Client) DeleteResourceBookingTypes(ids []int64) error {
	return c.Delete(ResourceBookingTypeModel, ids)
}

// GetResourceBookingType gets resource.booking.type existing record.
func (c *Client) GetResourceBookingType(id int64) (*ResourceBookingType, error) {
	rbts, err := c.GetResourceBookingTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rbts)[0]), nil
}

// GetResourceBookingTypes gets resource.booking.type existing records.
func (c *Client) GetResourceBookingTypes(ids []int64) (*ResourceBookingTypes, error) {
	rbts := &ResourceBookingTypes{}
	if err := c.Read(ResourceBookingTypeModel, ids, nil, rbts); err != nil {
		return nil, err
	}
	return rbts, nil
}

// FindResourceBookingType finds resource.booking.type record by querying it with criteria.
func (c *Client) FindResourceBookingType(criteria *Criteria) (*ResourceBookingType, error) {
	rbts := &ResourceBookingTypes{}
	if err := c.SearchRead(ResourceBookingTypeModel, criteria, NewOptions().Limit(1), rbts); err != nil {
		return nil, err
	}
	return &((*rbts)[0]), nil
}

// FindResourceBookingTypes finds resource.booking.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceBookingTypes(criteria *Criteria, options *Options) (*ResourceBookingTypes, error) {
	rbts := &ResourceBookingTypes{}
	if err := c.SearchRead(ResourceBookingTypeModel, criteria, options, rbts); err != nil {
		return nil, err
	}
	return rbts, nil
}

// FindResourceBookingTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceBookingTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResourceBookingTypeModel, criteria, options)
}

// FindResourceBookingTypeId finds record id by querying it with criteria.
func (c *Client) FindResourceBookingTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResourceBookingTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
