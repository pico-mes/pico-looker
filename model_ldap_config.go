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
// LdapConfig struct for LdapConfig
type LdapConfig struct {
	// Operations the current user is able to perform on this object
	Can map[string]bool `json:"can,omitempty"`
	// Allow alternate email-based login via '/login/email' for admins and for specified users with the 'login_special_email' permission. This option is useful as a fallback during ldap setup, if ldap config problems occur later, or if you need to support some users who are not in your ldap directory. Looker email/password logins are always disabled for regular users when ldap is enabled.
	AlternateEmailLoginAllowed bool `json:"alternate_email_login_allowed,omitempty"`
	// (Write-Only)  Password for the LDAP account used to access the LDAP server
	AuthPassword string `json:"auth_password,omitempty"`
	// Users will not be allowed to login at all unless a role for them is found in LDAP if set to true
	AuthRequiresRole bool `json:"auth_requires_role,omitempty"`
	// Distinguished name of LDAP account used to access the LDAP server
	AuthUsername string `json:"auth_username,omitempty"`
	// LDAP server hostname
	ConnectionHost string `json:"connection_host,omitempty"`
	// LDAP host port
	ConnectionPort string `json:"connection_port,omitempty"`
	// Use Transport Layer Security
	ConnectionTls bool `json:"connection_tls,omitempty"`
	// Do not verify peer when using TLS
	ConnectionTlsNoVerify bool `json:"connection_tls_no_verify,omitempty"`
	// (Write-Only)  Array of ids of groups that will be applied to new users the first time they login via LDAP
	DefaultNewUserGroupIds []int64 `json:"default_new_user_group_ids,omitempty"`
	// (Read-only) Groups that will be applied to new users the first time they login via LDAP
	DefaultNewUserGroups []Group `json:"default_new_user_groups,omitempty"`
	// (Write-Only)  Array of ids of roles that will be applied to new users the first time they login via LDAP
	DefaultNewUserRoleIds []int64 `json:"default_new_user_role_ids,omitempty"`
	// (Read-only) Roles that will be applied to new users the first time they login via LDAP
	DefaultNewUserRoles []Role `json:"default_new_user_roles,omitempty"`
	// Enable/Disable LDAP authentication for the server
	Enabled bool `json:"enabled,omitempty"`
	// Don't attempt to do LDAP search result paging (RFC 2696) even if the LDAP server claims to support it.
	ForceNoPage bool `json:"force_no_page,omitempty"`
	// (Read-only) Array of mappings between LDAP Groups and Looker Roles
	Groups []LdapGroupRead `json:"groups,omitempty"`
	// Base dn for finding groups in LDAP searches
	GroupsBaseDn string `json:"groups_base_dn,omitempty"`
	// Identifier for a strategy for how Looker will search for groups in the LDAP server
	GroupsFinderType string `json:"groups_finder_type,omitempty"`
	// LDAP Group attribute that signifies the members of the groups. Most commonly 'member'
	GroupsMemberAttribute string `json:"groups_member_attribute,omitempty"`
	// Optional comma-separated list of supported LDAP objectclass for groups when doing groups searches
	GroupsObjectclasses string `json:"groups_objectclasses,omitempty"`
	// LDAP Group attribute that signifies the user in a group. Most commonly 'dn'
	GroupsUserAttribute string `json:"groups_user_attribute,omitempty"`
	// (Read/Write) Array of mappings between LDAP Groups and arrays of Looker Role ids
	GroupsWithRoleIds []LdapGroupWrite `json:"groups_with_role_ids,omitempty"`
	// (Read-only) Has the password been set for the LDAP account used to access the LDAP server
	HasAuthPassword bool `json:"has_auth_password,omitempty"`
	// Merge first-time ldap login to existing user account by email addresses. When a user logs in for the first time via ldap this option will connect this user into their existing account by finding the account with a matching email address. Otherwise a new user account will be created for the user.
	MergeNewUsersByEmail bool `json:"merge_new_users_by_email,omitempty"`
	// When this config was last modified
	ModifiedAt string `json:"modified_at,omitempty"`
	// User id of user who last modified this config
	ModifiedBy string `json:"modified_by,omitempty"`
	// Set user roles in Looker based on groups from LDAP
	SetRolesFromGroups bool `json:"set_roles_from_groups,omitempty"`
	// (Write-Only)  Test LDAP user password. For ldap tests only.
	TestLdapPassword string `json:"test_ldap_password,omitempty"`
	// (Write-Only)  Test LDAP user login id. For ldap tests only.
	TestLdapUser string `json:"test_ldap_user,omitempty"`
	// Name of user record attributes used to indicate email address field
	UserAttributeMapEmail string `json:"user_attribute_map_email,omitempty"`
	// Name of user record attributes used to indicate first name
	UserAttributeMapFirstName string `json:"user_attribute_map_first_name,omitempty"`
	// Name of user record attributes used to indicate last name
	UserAttributeMapLastName string `json:"user_attribute_map_last_name,omitempty"`
	// Name of user record attributes used to indicate unique record id
	UserAttributeMapLdapId string `json:"user_attribute_map_ldap_id,omitempty"`
	// (Read-only) Array of mappings between LDAP User Attributes and Looker User Attributes
	UserAttributes []LdapUserAttributeRead `json:"user_attributes,omitempty"`
	// (Read/Write) Array of mappings between LDAP User Attributes and arrays of Looker User Attribute ids
	UserAttributesWithIds []LdapUserAttributeWrite `json:"user_attributes_with_ids,omitempty"`
	// Distinguished name of LDAP node used as the base for user searches
	UserBindBaseDn string `json:"user_bind_base_dn,omitempty"`
	// (Optional) Custom RFC-2254 filter clause for use in finding user during login. Combined via 'and' with the other generated filter clauses.
	UserCustomFilter string `json:"user_custom_filter,omitempty"`
	// Name(s) of user record attributes used for matching user login id (comma separated list)
	UserIdAttributeNames string `json:"user_id_attribute_names,omitempty"`
	// (Optional) Name of user record objectclass used for finding user during login id
	UserObjectclass string `json:"user_objectclass,omitempty"`
	// Allow LDAP auth'd users to be members of non-reflected Looker groups. If 'false', user will be removed from non-reflected groups on login.
	AllowNormalGroupMembership bool `json:"allow_normal_group_membership,omitempty"`
	// LDAP auth'd users will be able to inherit roles from non-reflected Looker groups.
	AllowRolesFromNormalGroups bool `json:"allow_roles_from_normal_groups,omitempty"`
	// Allows roles to be directly assigned to LDAP auth'd users.
	AllowDirectRoles bool `json:"allow_direct_roles,omitempty"`
	// Link to get this item
	Url string `json:"url,omitempty"`
}
