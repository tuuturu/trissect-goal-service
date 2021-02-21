/*
 * Trissect Goal Service
 *
 * Handles CRUD operations on Goal resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Goal struct {
	Id string `json:"id,omitempty"`

	Parent string `json:"parent,omitempty"`

	Title string `json:"title"`

	Reasoning string `json:"reasoning,omitempty"`

	Complete bool `json:"complete,omitempty"`
}
