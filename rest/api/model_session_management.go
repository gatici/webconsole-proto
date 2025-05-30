/*
 * WebConsole NFConfig API
 *
 * API for managing access, mobility, policy, session and PLMN information.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SessionManagement struct {

	SliceName string `json:"sliceName,omitempty"`

	PlmnId PlmnId `json:"plmnId,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	IpDomain []IpDomain `json:"ipDomain,omitempty"`

	Upf Upf `json:"upf,omitempty"`

	GnbNames []string `json:"gnbNames,omitempty"`
}
