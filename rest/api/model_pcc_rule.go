/*
 * WebConsole NFConfig API
 *
 * API for managing access, mobility, policy, session and PLMN information.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PccRule struct {

	RuleId string `json:"ruleId,omitempty"`

	Flows []PccFlow `json:"flows,omitempty"`

	Qos PccQos `json:"qos,omitempty"`

	Precedence int32 `json:"precedence,omitempty"`
}
