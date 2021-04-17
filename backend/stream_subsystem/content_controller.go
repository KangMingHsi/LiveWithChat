package stream_subsystem

// ContentController is a domain service for video contents.
type ContentController interface {
	// Save content with the id and type
	Save(id, contentType string, content interface{}) error
	// Get content information
	GetContentInfo(id string) (map[string]interface{}, error)
	// Delete content by id
	Delete(id string) error
}