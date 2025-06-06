/*
 * WebConsole NFConfig API
 *
 * API for managing access, mobility, policy, session and PLMN information.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PolicyControl struct {

	PlmnId PlmnId `json:"plmnId"`

	Snssai Snssai `json:"snssai"`

	DnnQos []DnnQos `json:"dnnQos,omitempty"`

	PccRules []PccRule `json:"pccRules"`
}
