/*
 * WebConsole NFConfig API
 *
 * API for managing access, mobility, policy, session and PLMN information.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DnnQos struct {

	DnnName string `json:"dnnName,omitempty"`

	MbrUplink string `json:"mbrUplink,omitempty"`

	MbrDownlink string `json:"mbrDownlink,omitempty"`

	FiveQi int32 `json:"fiveQi,omitempty"`

	ArpPriorityLevel int32 `json:"arpPriorityLevel,omitempty"`
}
