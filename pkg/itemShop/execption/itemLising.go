package execption

type ItemListing struct{}

func (e *ItemListing) Error() string {
	return "Failed to list items"
}