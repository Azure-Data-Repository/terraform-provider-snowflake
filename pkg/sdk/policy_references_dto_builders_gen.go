// Code generated by dto builder generator; DO NOT EDIT.

package sdk

import ()

func NewGetForEntityPolicyReferenceRequest(
	RefEntityName ObjectIdentifier,
	RefEntityDomain PolicyEntityDomain,
) *GetForEntityPolicyReferenceRequest {
	s := GetForEntityPolicyReferenceRequest{}
	s.RefEntityName = RefEntityName
	s.RefEntityDomain = RefEntityDomain
	return &s
}
