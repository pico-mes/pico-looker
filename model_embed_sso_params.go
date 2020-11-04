/*
 * Looker API 3.1 Reference
 *
 * ### Authorization  The Looker API uses Looker **API3** credentials for authorization and access control. Looker admins can create API3 credentials on Looker's **Admin/Users** page. Pass API3 credentials to the **_/login** endpoint to obtain a temporary access_token. Include that access_token in the Authorization header of Looker API requests. For details, see [Looker API Authorization](https://looker.com/docs/r/api/authorization)  ### Client SDKs  The Looker API is a RESTful system that should be usable by any programming language capable of making HTTPS requests. Client SDKs for a variety of programming languages can be generated from the Looker API's Swagger JSON metadata to streamline use of the Looker API in your applications. A client SDK for Ruby is available as an example. For more information, see [Looker API Client SDKs](https://looker.com/docs/r/api/client_sdks)  ### Try It Out!  The 'api-docs' page served by the Looker instance includes 'Try It Out!' buttons for each API method. After logging in with API3 credentials, you can use the \"Try It Out!\" buttons to call the API directly from the documentation page to interactively explore API features and responses.  Note! With great power comes great responsibility: The \"Try It Out!\" button makes API calls to your live Looker instance. Be especially careful with destructive API operations such as `delete_user` or similar. There is no \"undo\" for API operations.  ### Versioning  Future releases of Looker will expand this API release-by-release to securely expose more and more of the core power of Looker to API client applications. API endpoints marked as \"beta\" may receive breaking changes without warning (but we will try to avoid doing that). Stable (non-beta) API endpoints should not receive breaking changes in future releases. For more information, see [Looker API Versioning](https://looker.com/docs/r/api/versioning)  ### In This Release  The following are a few examples of noteworthy items that have changed between API 3.0 and API 3.1. For more comprehensive coverage of API changes, please see the release notes for your Looker release.  ### Examples of new things added in API 3.1 (compared to API 3.0):  * [Dashboard construction](#!/3.1/Dashboard/) APIs * [Themes](#!/3.1/Theme/) and [custom color collections](#!/3.1/ColorCollection) APIs * Create and run [SQL Runner](#!/3.1/Query/run_sql_query) queries * Create and run [merged results](#!/3.1/Query/create_merge_query) queries * Create and modify [dashboard filters](#!/3.1/Dashboard/create_dashboard_filter) * Create and modify [password requirements](#!/3.1/Auth/password_config)  ### Deprecated in API 3.0  The following functions and properties have been deprecated in API 3.0.  They continue to exist and work in API 3.0 for the next several Looker releases but they have not been carried forward to API 3.1:  * Dashboard Prefetch functions * User access_filter functions * User API 1.0 credentials functions * Space.is_root and Space.is_user_root properties. Use Space.is_shared_root and Space.is_users_root instead.  ### Semantic changes in API 3.1:  * [all_looks()](#!/3.1/Look/all_looks) no longer includes soft-deleted looks, matching [all_dashboards()](#!/3.1/Dashboard/all_dashboards) behavior. You can find soft-deleted looks using [search_looks()](#!/3.1/Look/search_looks) with the `deleted` param set to True. * [all_spaces()](#!/3.1/Space/all_spaces) no longer includes duplicate items * [search_users()](#!/3.1/User/search_users) no longer accepts Y,y,1,0,N,n for Boolean params, only \"true\" and \"false\". * For greater client and network compatibility, [render_task_results](#!/3.1/RenderTask/render_task_results) now returns HTTP status **202 Accepted** instead of HTTP status **102 Processing** * [all_running_queries()](#!/3.1/Query/all_running_queries) and [kill_query](#!/3.1/Query/kill_query) functions have moved into the [Query](#!/3.1/Query/) function group.   If you have application code which relies on the old behavior of the APIs above, you may continue using the API 3.0 functions in this Looker release. We strongly suggest you update your code to use API 3.1 analogs as soon as possible.  
 *
 * API version: 3.1.0
 * Contact: support@looker.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package looker
// EmbedSsoParams struct for EmbedSsoParams
type EmbedSsoParams struct {
	// The complete URL of the Looker UI page to display in the embed context. For example, to display the dashboard with id 34, `target_url` would look like: `https://mycompany.looker.com:9999/dashboards/34`. `target_uri` MUST contain a scheme (HTTPS), domain name, and URL path. Port must be included if it is required to reach the Looker server from browser clients. If the Looker instance is behind a load balancer or other proxy, `target_uri` must be the public-facing domain name and port required to reach the Looker instance, not the actual internal network machine name of the Looker instance.
	TargetUrl string `json:"target_url"`
	// Number of seconds the SSO embed session will be valid after the embed session is started. Defaults to 300 seconds. Maximum session length accepted is 2592000 seconds (30 days).
	SessionLength int64 `json:"session_length,omitempty"`
	// When true, the embed session will purge any residual Looker login state (such as in browser cookies) before creating a new login state with the given embed user info. Defaults to true.
	ForceLogoutLogin bool `json:"force_logout_login,omitempty"`
	// A value from an external system that uniquely identifies the embed user. Since the user_ids of Looker embed users may change with every embed session, external_user_id provides a way to assign a known, stable user identifier across multiple embed sessions.
	ExternalUserId string `json:"external_user_id,omitempty"`
	// First name of the embed user. Defaults to 'Embed' if not specified
	FirstName string `json:"first_name,omitempty"`
	// Last name of the embed user. Defaults to 'User' if not specified
	LastName string `json:"last_name,omitempty"`
	// Sets the user timezone for the embed user session, if the User Specific Timezones setting is enabled in the Looker admin settings. A value of `null` forces the embed user to use the Looker Application Default Timezone. You MUST omit this property from the request if the User Specific Timezones setting is disabled. Timezone values are validated against the IANA Timezone standard and can be seen in the Application Time Zone dropdown list on the Looker General Settings admin page.
	UserTimezone string `json:"user_timezone,omitempty"`
	// List of Looker permission names to grant to the embed user. Requested permissions will be filtered to permissions allowed for embed sessions.
	Permissions []string `json:"permissions,omitempty"`
	// List of model names that the embed user may access
	Models []string `json:"models,omitempty"`
	// List of Looker group ids in which to enroll the embed user
	GroupIds []int64 `json:"group_ids,omitempty"`
	// A unique value identifying an embed-exclusive group. Multiple embed users using the same `external_group_id` value will be able to share Looker content with each other. Content and embed users associated with the `external_group_id` will not be accessible to normal Looker users or embed users not associated with this `external_group_id`.
	ExternalGroupId int64 `json:"external_group_id,omitempty"`
	// A dictionary of name-value pairs associating a Looker user attribute name with a value.
	UserAttributes map[string]string `json:"user_attributes,omitempty"`
	// Id of the embed secret to use to sign this SSO url. If specified, the value must be an id of a valid (active) secret defined in the Looker instance. If not specified, the URL will be signed with the newest active embed secret defined in the Looker instance.
	SecretId int64 `json:"secret_id,omitempty"`
}