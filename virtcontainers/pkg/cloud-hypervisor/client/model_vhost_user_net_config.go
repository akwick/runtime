/*
 * Cloud Hypervisor API
 *
 * Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// VhostUserNetConfig struct for VhostUserNetConfig
type VhostUserNetConfig struct {
	Sock string `json:"sock"`
	NumQueues int32 `json:"num_queues,omitempty"`
	QueueSize int32 `json:"queue_size,omitempty"`
	Mac string `json:"mac,omitempty"`
}