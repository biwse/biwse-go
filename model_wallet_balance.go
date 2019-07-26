/*
 * Biwse API V1 Reference
 *
 * This API documentation generated from OpenAPI specification. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package biwse

type WalletBalance struct {
	// Total wallet balance.
	Total float32 `json:"total,omitempty"`
	// Fund in wallet available to withdrawal.
	Confirmed float32 `json:"confirmed,omitempty"`
}
