package adcom

import "encoding/json"

// App object is used to define an ad supported non-browser application, in contrast to a typical website, example.
// As a derived class, an App object inherits all DistributionChannel attributes and adds those defined below.
type App struct {
	DistributionChannel

	// Attribute:
	//   domain
	// Type:
	//   string
	// Definition:
	//   Domain of the app (e.g., “mygame.foo.com”).
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Definition:
	//   Array of content categories describing the app using IDs from the taxonomy indicated in cattax.
	Cat []string `json:"cat,omitempty"`

	// Attribute:
	//   sectcat
	// Type:
	//   string array
	// Definition:
	//   Array of content categories describing the current section of the app using IDs from the taxonomy indicated in cattax.
	SectCat []string `json:"sectcat,omitempty"`

	// Attribute:
	//   pagecat
	// Type:
	//   string array
	// Definition:
	//   Array of content categories describing the current page or view of the app using IDs from the taxonomy indicated in cattax.
	PageCat []string `json:"pagecat,omitempty"`

	// Attribute:
	//   cattax
	// Type:
	//   integer
	// Definition:
	//   The taxonomy in use for the cat, sectcat and pagecat attributes.
	//   Refer to List: Category Taxonomies.
	CatTax CategoryTaxonomy `json:"cattax,omitempty"`

	// Attribute:
	//   privpolicy
	// Type:
	//   integer
	// Definition:
	//   Indicates if the app has a privacy policy, where 0 = no, 1 = yes.
	PrivPolicy int8 `json:"privpolicy,omitempty"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Definition:
	//   Comma separated list of keywords about the app.
	Keywords string `json:"keywords,omitempty"`

	// Attribute:
	//   bundle
	// Type:
	//   string
	// Definition:
	//   Bundle or package name of the app (e.g., “com.foo.mygame”) and should NOT be app store IDs (e.g., not iTunes store IDs).
	Bundle string `json:"bundle,omitempty"`

	// Attribute:
	//   storeid
	// Type:
	//   string
	// Definition:
	//   The ID of the app in an app store (e.g., Apple iTunes, Google Play).
	StoreID string `json:"storeid,omitempty"`

	// Attribute:
	//   storeurl
	// Type:
	//   string
	// Definition:
	//   App store URL for an installed app; for IQG 2.1 compliance.
	StoreURL string `json:"storeurl,omitempty"`

	// Attribute:
	//   ver
	// Type:
	//   string
	// Definition:
	//   Application version.
	Ver string `json:"ver,omitempty"`

	// Attribute:
	//   paid
	// Type:
	//   integer; default 0
	// Definition:
	//   Indicator of whether or not this is a paid app, where 0 = free, 1 = paid.
	Paid int8 `json:"paid,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
