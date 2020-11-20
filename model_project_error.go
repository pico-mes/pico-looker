/*
 * Looker API 4.0 (Experimental) Reference
 *
 *  Welcome to the future! This is an early preview of our next-generation Looker API 4.0. API 4.0 runs alongside APIs 3.1 and 3.0. We've tagged 4.0 as \"experimental\" to reflect that we have more work planned for API 4.0 which may include breaking changes. Please pardon our dust while we remodel a few rooms!  ### In This Release  We're spinning up this new API 4.0 version so that we can make adjustments to our API functions, parameters, and response types to fix bugs and inconsistencies. These changes fall outside the bounds of non-breaking additive changes we can make to our stable API 3.1.  One benefit of these type adjustments in API 4.0 is dramatically better support for strongly typed languages like TypeScript, Kotlin, Java, and more. Looker is also creating client SDKs to call the Looker API from these and other languages. These client SDKs will be available as pre-built packages for download from public repositories such as npmjs.org, RubyGems.org, PyPi.org. If you use an IDE for software development, you will soon be able to install a Looker SDK for your programming language with the click of a button!  While API 3.1 is still the defacto Looker API (\"current\", \"stable\", \"default\", etc), the bulk of our development activity will gradually shift to API 4.0.  
 *
 * API version: 4.0.7.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package looker
// ProjectError struct for ProjectError
type ProjectError struct {
	// A stable token that uniquely identifies this class of error, ignoring parameter values. Error message text may vary due to parameters or localization, but error codes do not. For example, a \"File not found\" error will have the same error code regardless of the filename in question or the user's display language
	Code string `json:"code,omitempty"`
	// Severity: fatal, error, warning, info, success
	Severity string `json:"severity,omitempty"`
	// Error classification: syntax, deprecation, model_configuration, etc
	Kind string `json:"kind,omitempty"`
	// Error message which may contain information such as dashboard or model names that may be considered sensitive in some use cases. Avoid storing or sending this message outside of Looker
	Message string `json:"message,omitempty"`
	// The field associated with this error
	FieldName string `json:"field_name,omitempty"`
	// Name of the file containing this error
	FilePath string `json:"file_path,omitempty"`
	// Line number in the file of this error
	LineNumber int64 `json:"line_number,omitempty"`
	// The model associated with this error
	ModelId string `json:"model_id,omitempty"`
	// The explore associated with this error
	Explore string `json:"explore,omitempty"`
	// A link to Looker documentation about this error
	HelpUrl string `json:"help_url,omitempty"`
	// Error parameters
	Params map[string]string `json:"params,omitempty"`
	// A version of the error message that does not contain potentially sensitive information. Suitable for situations in which messages are stored or sent to consumers outside of Looker, such as external logs. Sanitized messages will display \"(?)\" where sensitive information would appear in the corresponding non-sanitized message
	SanitizedMessage string `json:"sanitized_message,omitempty"`
}
