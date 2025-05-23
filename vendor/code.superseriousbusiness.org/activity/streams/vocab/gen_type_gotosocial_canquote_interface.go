// Code generated by astool. DO NOT EDIT.

package vocab

type GoToSocialCanQuote interface {
	// GetGoToSocialAlways returns the "always" property if it exists, and nil
	// otherwise.
	GetGoToSocialAlways() GoToSocialAlwaysProperty
	// GetGoToSocialApprovalRequired returns the "approvalRequired" property
	// if it exists, and nil otherwise.
	GetGoToSocialApprovalRequired() GoToSocialApprovalRequiredProperty
	// GetGoToSocialAutomaticApproval returns the "automaticApproval" property
	// if it exists, and nil otherwise.
	GetGoToSocialAutomaticApproval() GoToSocialAutomaticApprovalProperty
	// GetGoToSocialManualApproval returns the "manualApproval" property if it
	// exists, and nil otherwise.
	GetGoToSocialManualApproval() GoToSocialManualApprovalProperty
	// GetJSONLDId returns the "id" property if it exists, and nil otherwise.
	GetJSONLDId() JSONLDIdProperty
	// GetTypeName returns the name of this type.
	GetTypeName() string
	// GetUnknownProperties returns the unknown properties for the CanQuote
	// type. Note that this should not be used by app developers. It is
	// only used to help determine which implementation is LessThan the
	// other. Developers who are creating a different implementation of
	// this type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// IsExtending returns true if the CanQuote type extends from the other
	// type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this CanQuote is lesser, with an arbitrary but
	// stable determination.
	LessThan(o GoToSocialCanQuote) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetGoToSocialAlways sets the "always" property.
	SetGoToSocialAlways(i GoToSocialAlwaysProperty)
	// SetGoToSocialApprovalRequired sets the "approvalRequired" property.
	SetGoToSocialApprovalRequired(i GoToSocialApprovalRequiredProperty)
	// SetGoToSocialAutomaticApproval sets the "automaticApproval" property.
	SetGoToSocialAutomaticApproval(i GoToSocialAutomaticApprovalProperty)
	// SetGoToSocialManualApproval sets the "manualApproval" property.
	SetGoToSocialManualApproval(i GoToSocialManualApprovalProperty)
	// SetJSONLDId sets the "id" property.
	SetJSONLDId(i JSONLDIdProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}
