package odoo

// BlogPost represents blog.post model.
type BlogPost struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	Active                   *Bool     `xmlrpc:"active,omitempty"`
	AuthorAvatar             *String   `xmlrpc:"author_avatar,omitempty"`
	AuthorId                 *Many2One `xmlrpc:"author_id,omitempty"`
	AuthorName               *String   `xmlrpc:"author_name,omitempty"`
	BlogId                   *Many2One `xmlrpc:"blog_id,omitempty"`
	CanPublish               *Bool     `xmlrpc:"can_publish,omitempty"`
	Content                  *String   `xmlrpc:"content,omitempty"`
	CoverProperties          *String   `xmlrpc:"cover_properties,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	IsPublished              *Bool     `xmlrpc:"is_published,omitempty"`
	IsSeoOptimized           *Bool     `xmlrpc:"is_seo_optimized,omitempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	PostDate                 *Time     `xmlrpc:"post_date,omitempty"`
	PublishedDate            *Time     `xmlrpc:"published_date,omitempty"`
	SeoName                  *String   `xmlrpc:"seo_name,omitempty"`
	Subtitle                 *String   `xmlrpc:"subtitle,omitempty"`
	TagIds                   *Relation `xmlrpc:"tag_ids,omitempty"`
	Teaser                   *String   `xmlrpc:"teaser,omitempty"`
	TeaserManual             *String   `xmlrpc:"teaser_manual,omitempty"`
	Visits                   *Int      `xmlrpc:"visits,omitempty"`
	WebsiteId                *Many2One `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WebsiteMetaDescription   *String   `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords      *String   `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg         *String   `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle         *String   `xmlrpc:"website_meta_title,omitempty"`
	WebsitePublished         *Bool     `xmlrpc:"website_published,omitempty"`
	WebsiteUrl               *String   `xmlrpc:"website_url,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BlogPosts represents array of blog.post model.
type BlogPosts []BlogPost

// BlogPostModel is the odoo model name.
const BlogPostModel = "blog.post"

// Many2One convert BlogPost to *Many2One.
func (bp *BlogPost) Many2One() *Many2One {
	return NewMany2One(bp.Id.Get(), "")
}

// CreateBlogPost creates a new blog.post model and returns its id.
func (c *Client) CreateBlogPost(bp *BlogPost) (int64, error) {
	ids, err := c.CreateBlogPosts([]*BlogPost{bp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBlogPost creates a new blog.post model and returns its id.
func (c *Client) CreateBlogPosts(bps []*BlogPost) ([]int64, error) {
	var vv []interface{}
	for _, v := range bps {
		vv = append(vv, v)
	}
	return c.Create(BlogPostModel, vv, nil)
}

// UpdateBlogPost updates an existing blog.post record.
func (c *Client) UpdateBlogPost(bp *BlogPost) error {
	return c.UpdateBlogPosts([]int64{bp.Id.Get()}, bp)
}

// UpdateBlogPosts updates existing blog.post records.
// All records (represented by ids) will be updated by bp values.
func (c *Client) UpdateBlogPosts(ids []int64, bp *BlogPost) error {
	return c.Update(BlogPostModel, ids, bp, nil)
}

// DeleteBlogPost deletes an existing blog.post record.
func (c *Client) DeleteBlogPost(id int64) error {
	return c.DeleteBlogPosts([]int64{id})
}

// DeleteBlogPosts deletes existing blog.post records.
func (c *Client) DeleteBlogPosts(ids []int64) error {
	return c.Delete(BlogPostModel, ids)
}

// GetBlogPost gets blog.post existing record.
func (c *Client) GetBlogPost(id int64) (*BlogPost, error) {
	bps, err := c.GetBlogPosts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bps)[0]), nil
}

// GetBlogPosts gets blog.post existing records.
func (c *Client) GetBlogPosts(ids []int64) (*BlogPosts, error) {
	bps := &BlogPosts{}
	if err := c.Read(BlogPostModel, ids, nil, bps); err != nil {
		return nil, err
	}
	return bps, nil
}

// FindBlogPost finds blog.post record by querying it with criteria.
func (c *Client) FindBlogPost(criteria *Criteria) (*BlogPost, error) {
	bps := &BlogPosts{}
	if err := c.SearchRead(BlogPostModel, criteria, NewOptions().Limit(1), bps); err != nil {
		return nil, err
	}
	return &((*bps)[0]), nil
}

// FindBlogPosts finds blog.post records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBlogPosts(criteria *Criteria, options *Options) (*BlogPosts, error) {
	bps := &BlogPosts{}
	if err := c.SearchRead(BlogPostModel, criteria, options, bps); err != nil {
		return nil, err
	}
	return bps, nil
}

// FindBlogPostIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBlogPostIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BlogPostModel, criteria, options)
}

// FindBlogPostId finds record id by querying it with criteria.
func (c *Client) FindBlogPostId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BlogPostModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
