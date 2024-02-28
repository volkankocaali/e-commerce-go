package parser

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
)

// TagParser is the default implementation of TagParser interface
type TagParser struct{}

// NewTagParser returns a new instance of TagParser
func NewTagParser() *TagParser {
	return &TagParser{}
}

// Parse implements the Parse method of the TagParser interface
func (t *TagParser) Parse(pt []models.ProductTags) []schema.TagsResponseSchema {
	// Implement tag parsing logic here
	var tags []schema.TagsResponseSchema

	for _, v := range pt {
		if v.Tags != nil {
			tags = append(tags, schema.TagsResponseSchema{
				ID:        v.Tags.ID,
				Name:      v.Tags.TagName,
				CreatedAt: v.Tags.CreatedAt,
				UpdatedAt: v.Tags.UpdatedAt,
			})
		}
	}

	return tags
}
