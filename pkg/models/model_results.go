/*
 * Unified Policy Engine I/O Formats
 *
 * Documentation for the input and output formats used in Unified Policy Engine
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

// This is the top-level output from the Unified Policy Engine.
type Results struct {
	Format        string `json:"format"`
	FormatVersion string `json:"format_version"`
	// The inputs that produced these results. This field is optional.
	Inputs []State `json:"inputs,omitempty"`
	// A map of rule package to a rule results object
	RuleResults map[string]RuleResults `json:"rule_results"`
}
