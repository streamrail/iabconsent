package iabconsent

// ParsedConsentInterface provides a common type for both V1 and V2 parsed consent structures.
type ParsedConsentInterface interface {
	CheckAllowedPurposesExists() bool
	VendorAllowed(vendor int) bool
}
