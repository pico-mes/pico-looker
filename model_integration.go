/*
 * Looker API 4.0 (Experimental) Reference
 *
 *  Welcome to the future! This is an early preview of our next-generation Looker API 4.0. API 4.0 runs alongside APIs 3.1 and 3.0. We've tagged 4.0 as \"experimental\" to reflect that we have more work planned for API 4.0 which may include breaking changes. Please pardon our dust while we remodel a few rooms!  ### In This Release  We're spinning up this new API 4.0 version so that we can make adjustments to our API functions, parameters, and response types to fix bugs and inconsistencies. These changes fall outside the bounds of non-breaking additive changes we can make to our stable API 3.1.  One benefit of these type adjustments in API 4.0 is dramatically better support for strongly typed languages like TypeScript, Kotlin, Java, and more. Looker is also creating client SDKs to call the Looker API from these and other languages. These client SDKs will be available as pre-built packages for download from public repositories such as npmjs.org, RubyGems.org, PyPi.org. If you use an IDE for software development, you will soon be able to install a Looker SDK for your programming language with the click of a button!  While API 3.1 is still the defacto Looker API (\"current\", \"stable\", \"default\", etc), the bulk of our development activity will gradually shift to API 4.0.  
 *
 * API version: 4.0.7.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package looker
// Integration struct for Integration
type Integration struct {
	// Operations the current user is able to perform on this object
	Can map[string]bool `json:"can,omitempty"`
	// ID of the integration.
	Id string `json:"id,omitempty"`
	// ID of the integration hub.
	IntegrationHubId int64 `json:"integration_hub_id,omitempty"`
	// Label for the integration.
	Label string `json:"label,omitempty"`
	// Description of the integration.
	Description string `json:"description,omitempty"`
	// Whether the integration is available to users.
	Enabled bool `json:"enabled,omitempty"`
	// Array of params for the integration.
	Params []IntegrationParam `json:"params,omitempty"`
	// A list of data formats the integration supports. If unspecified, the default is all data formats. Valid values are: \"txt\", \"csv\", \"inline_json\", \"json\", \"json_label\", \"json_detail\", \"json_detail_lite_stream\", \"xlsx\", \"html\", \"wysiwyg_pdf\", \"assembled_pdf\", \"wysiwyg_png\", \"csv_zip\".
	SupportedFormats []string `json:"supported_formats,omitempty"`
	// A list of action types the integration supports. Valid values are: \"cell\", \"query\", \"dashboard\".
	SupportedActionTypes []string `json:"supported_action_types,omitempty"`
	// A list of formatting options the integration supports. If unspecified, defaults to all formats. Valid values are: \"formatted\", \"unformatted\".
	SupportedFormattings []string `json:"supported_formattings,omitempty"`
	// A list of visualization formatting options the integration supports. If unspecified, defaults to all formats. Valid values are: \"apply\", \"noapply\".
	SupportedVisualizationFormattings []string `json:"supported_visualization_formattings,omitempty"`
	// A list of all the download mechanisms the integration supports. The order of values is not significant: Looker will select the most appropriate supported download mechanism for a given query. The integration must ensure it can handle any of the mechanisms it claims to support. If unspecified, this defaults to all download setting values. Valid values are: \"push\", \"url\".
	SupportedDownloadSettings []string `json:"supported_download_settings,omitempty"`
	// URL to an icon for the integration.
	IconUrl string `json:"icon_url,omitempty"`
	// Whether the integration uses oauth.
	UsesOauth bool `json:"uses_oauth,omitempty"`
	// A list of descriptions of required fields that this integration is compatible with. If there are multiple entries in this list, the integration requires more than one field. If unspecified, no fields will be required.
	RequiredFields []IntegrationRequiredField `json:"required_fields,omitempty"`
	// Whether the integration uses delegate oauth, which allows federation between an integration installation scope specific entity (like org, group, and team, etc.) and Looker.
	DelegateOauth bool `json:"delegate_oauth,omitempty"`
	// Whether the integration is available to users.
	InstalledDelegateOauthTargets []int64 `json:"installed_delegate_oauth_targets,omitempty"`
}
