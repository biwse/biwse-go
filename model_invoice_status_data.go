/*
 * Biwse API V1 Reference
 *
 * This API documentation generated from OpenAPI specification. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package biwse

type InvoiceStatusData struct {
	// Invoice's amount
	Amount float32 `json:"amount,omitempty"`
	// Number of transaction confirmations
	Confs float32 `json:"confs,omitempty"`
	// Payment receieved on address
	Address string `json:"address,omitempty"`
	// Transaction hash
	TxHash string `json:"tx_hash,omitempty"`
	// Currency type
	Type bool `json:"type,omitempty"`
	// Payment accepted or not
	Accepted bool `json:"accepted,omitempty"`
	// Verification hash
	VerificationHash string `json:"verification_hash,omitempty"`
}
