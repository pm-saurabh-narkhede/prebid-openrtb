package adcom

import "encoding/json"

// ImageAssetFormat object is used to provide native asset format specifications for an image element.
// Image elements are typically used for the actual creative image and icons.
type ImageAssetFormat struct {
	// Attribute:
	//   type
	// Type:
	//   integer
	// Definition:
	//   The type of image asset supported.
	//   Refer to List: Native Image Asset Types.
	Type NativeImageAssetType `json:"type,omitempty"`

	// Attribute:
	//   mime
	// Type:
	//   string array
	// Definition:
	//   Array of supported mime types (e.g., “image/jpeg”, “image/gif”).
	//   If omitted, all types are assumed.
	MIME []string `json:"mime,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer
	// Definition:
	//   Absolute width of the image asset in device independent pixels (DIPS).
	//   Note that mixing absolute and relative sizes is not recommended.
	W int `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Definition:
	//   Absolute height of the image asset in device independent pixels (DIPS).
	//   Note that mixing absolute and relative sizes is not recommended.
	H int `json:"h,omitempty"`

	// Attribute:
	//   wmin
	// Type:
	//   integer
	// Definition:
	//   The minimum requested absolute width of the image in device independent pixels (DIPS).
	//   This option should be used for any scaling of images by the client.
	WMin int `json:"wmin,omitempty"`

	// Attribute:
	//   hmin
	// Type:
	//   integer
	// Definition:
	//   The minimum requested absolute height of the image in device independent pixels (DIPS).
	//   This option should be used for any scaling of images by the client.
	HMin int `json:"hmin,omitempty"`

	// Attribute:
	//   wratio
	// Type:
	//   integer
	// Definition:
	//   Relative width of the image asset when expressing size as a ratio.
	//   Note that mixing absolute and relative sizes is not recommended.
	WRatio int `json:"wratio,omitempty"`

	// Attribute:
	//   hratio
	// Type:
	//   integer
	// Definition:
	//   Relative height of the image asset when expressing size as a ratio.
	//   Note that mixing absolute and relative sizes is not recommended.
	HRatio int `json:"hratio,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
