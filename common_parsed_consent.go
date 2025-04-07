package iabconsent

// ParsedConsentData provides a common type for both V1 and V2 parsed consent structures.
type ParsedConsentData interface {
	CheckIfPurposesNotAllowed() bool
	VendorAllowed(vendor int) bool
	SuitableToProcess(ps []int, vendor int) bool
	GetVersion() int
	IsEmpty() bool
}
