package odoo

// ResourceBookingCombination represents resource.booking.combination model.
type ResourceBookingCombination struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omitempty"`
	Active           *Bool     `xmlrpc:"active,omitempty"`
	BookingCount     *Int      `xmlrpc:"booking_count,omitempty"`
	BookingIds       *Relation `xmlrpc:"booking_ids,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	ForcedCalendarId *Many2One `xmlrpc:"forced_calendar_id,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	Name             *String   `xmlrpc:"name,omitempty"`
	ResourceIds      *Relation `xmlrpc:"resource_ids,omitempty"`
	TypeCount        *Int      `xmlrpc:"type_count,omitempty"`
	TypeRelIds       *Relation `xmlrpc:"type_rel_ids,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResourceBookingCombinations represents array of resource.booking.combination model.
type ResourceBookingCombinations []ResourceBookingCombination

// ResourceBookingCombinationModel is the odoo model name.
const ResourceBookingCombinationModel = "resource.booking.combination"

// Many2One convert ResourceBookingCombination to *Many2One.
func (rbc *ResourceBookingCombination) Many2One() *Many2One {
	return NewMany2One(rbc.Id.Get(), "")
}

// CreateResourceBookingCombination creates a new resource.booking.combination model and returns its id.
func (c *Client) CreateResourceBookingCombination(rbc *ResourceBookingCombination) (int64, error) {
	ids, err := c.CreateResourceBookingCombinations([]*ResourceBookingCombination{rbc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResourceBookingCombination creates a new resource.booking.combination model and returns its id.
func (c *Client) CreateResourceBookingCombinations(rbcs []*ResourceBookingCombination) ([]int64, error) {
	var vv []interface{}
	for _, v := range rbcs {
		vv = append(vv, v)
	}
	return c.Create(ResourceBookingCombinationModel, vv, nil)
}

// UpdateResourceBookingCombination updates an existing resource.booking.combination record.
func (c *Client) UpdateResourceBookingCombination(rbc *ResourceBookingCombination) error {
	return c.UpdateResourceBookingCombinations([]int64{rbc.Id.Get()}, rbc)
}

// UpdateResourceBookingCombinations updates existing resource.booking.combination records.
// All records (represented by ids) will be updated by rbc values.
func (c *Client) UpdateResourceBookingCombinations(ids []int64, rbc *ResourceBookingCombination) error {
	return c.Update(ResourceBookingCombinationModel, ids, rbc, nil)
}

// DeleteResourceBookingCombination deletes an existing resource.booking.combination record.
func (c *Client) DeleteResourceBookingCombination(id int64) error {
	return c.DeleteResourceBookingCombinations([]int64{id})
}

// DeleteResourceBookingCombinations deletes existing resource.booking.combination records.
func (c *Client) DeleteResourceBookingCombinations(ids []int64) error {
	return c.Delete(ResourceBookingCombinationModel, ids)
}

// GetResourceBookingCombination gets resource.booking.combination existing record.
func (c *Client) GetResourceBookingCombination(id int64) (*ResourceBookingCombination, error) {
	rbcs, err := c.GetResourceBookingCombinations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rbcs)[0]), nil
}

// GetResourceBookingCombinations gets resource.booking.combination existing records.
func (c *Client) GetResourceBookingCombinations(ids []int64) (*ResourceBookingCombinations, error) {
	rbcs := &ResourceBookingCombinations{}
	if err := c.Read(ResourceBookingCombinationModel, ids, nil, rbcs); err != nil {
		return nil, err
	}
	return rbcs, nil
}

// FindResourceBookingCombination finds resource.booking.combination record by querying it with criteria.
func (c *Client) FindResourceBookingCombination(criteria *Criteria) (*ResourceBookingCombination, error) {
	rbcs := &ResourceBookingCombinations{}
	if err := c.SearchRead(ResourceBookingCombinationModel, criteria, NewOptions().Limit(1), rbcs); err != nil {
		return nil, err
	}
	return &((*rbcs)[0]), nil
}

// FindResourceBookingCombinations finds resource.booking.combination records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceBookingCombinations(criteria *Criteria, options *Options) (*ResourceBookingCombinations, error) {
	rbcs := &ResourceBookingCombinations{}
	if err := c.SearchRead(ResourceBookingCombinationModel, criteria, options, rbcs); err != nil {
		return nil, err
	}
	return rbcs, nil
}

// FindResourceBookingCombinationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceBookingCombinationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResourceBookingCombinationModel, criteria, options)
}

// FindResourceBookingCombinationId finds record id by querying it with criteria.
func (c *Client) FindResourceBookingCombinationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResourceBookingCombinationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
