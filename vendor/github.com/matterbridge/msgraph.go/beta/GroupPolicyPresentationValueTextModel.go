// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupPolicyPresentationValueText The entity represents a string value for a drop-down list, combo box, or text box presentation on a policy definition.
type GroupPolicyPresentationValueText struct {
	// GroupPolicyPresentationValue is the base model of GroupPolicyPresentationValueText
	GroupPolicyPresentationValue
	// Value A string value for the associated presentation.
	Value *string `json:"value,omitempty"`
}