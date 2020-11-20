/*
 * Looker API 4.0 (Experimental) Reference
 *
 *  Welcome to the future! This is an early preview of our next-generation Looker API 4.0. API 4.0 runs alongside APIs 3.1 and 3.0. We've tagged 4.0 as \"experimental\" to reflect that we have more work planned for API 4.0 which may include breaking changes. Please pardon our dust while we remodel a few rooms!  ### In This Release  We're spinning up this new API 4.0 version so that we can make adjustments to our API functions, parameters, and response types to fix bugs and inconsistencies. These changes fall outside the bounds of non-breaking additive changes we can make to our stable API 3.1.  One benefit of these type adjustments in API 4.0 is dramatically better support for strongly typed languages like TypeScript, Kotlin, Java, and more. Looker is also creating client SDKs to call the Looker API from these and other languages. These client SDKs will be available as pre-built packages for download from public repositories such as npmjs.org, RubyGems.org, PyPi.org. If you use an IDE for software development, you will soon be able to install a Looker SDK for your programming language with the click of a button!  While API 3.1 is still the defacto Looker API (\"current\", \"stable\", \"default\", etc), the bulk of our development activity will gradually shift to API 4.0.  
 *
 * API version: 4.0.7.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package looker
// ConnectionFeatures struct for ConnectionFeatures
type ConnectionFeatures struct {
	// Name of the dialect for this connection
	DialectName string `json:"dialect_name,omitempty"`
	// True for cost estimating support
	CostEstimate bool `json:"cost_estimate,omitempty"`
	// True for multiple database support
	MultipleDatabases bool `json:"multiple_databases,omitempty"`
	// True for cost estimating support
	ColumnSearch bool `json:"column_search,omitempty"`
	// True for secondary index support
	PersistentTableIndexes bool `json:"persistent_table_indexes,omitempty"`
	// True for persistent derived table support
	PersistentDerivedTables bool `json:"persistent_derived_tables,omitempty"`
	// True for turtles support
	Turtles bool `json:"turtles,omitempty"`
	// True for percentile support
	Percentile bool `json:"percentile,omitempty"`
	// True for distinct percentile support
	DistinctPercentile bool `json:"distinct_percentile,omitempty"`
	// True for stable views support
	StableViews bool `json:"stable_views,omitempty"`
	// True for millisecond support
	Milliseconds bool `json:"milliseconds,omitempty"`
	// True for microsecond support
	Microseconds bool `json:"microseconds,omitempty"`
	// True for subtotal support
	Subtotals bool `json:"subtotals,omitempty"`
	// True for geographic location support
	Location bool `json:"location,omitempty"`
	// True for timezone conversion in query support
	Timezone bool `json:"timezone,omitempty"`
	// True for connection pooling support
	ConnectionPooling bool `json:"connection_pooling,omitempty"`
}
