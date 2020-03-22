// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// List undocumented
type List struct {
	// BaseItem is the base model of List
	BaseItem
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// List undocumented
	List *ListInfo `json:"list,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// System undocumented
	System *SystemFacet `json:"system,omitempty"`
	// Activities undocumented
	Activities []ItemActivityOLD `json:"activities,omitempty"`
	// Columns undocumented
	Columns []ColumnDefinition `json:"columns,omitempty"`
	// ContentTypes undocumented
	ContentTypes []ContentType `json:"contentTypes,omitempty"`
	// Drive undocumented
	Drive *Drive `json:"drive,omitempty"`
	// Items undocumented
	Items []ListItem `json:"items,omitempty"`
	// Subscriptions undocumented
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

// ListInfo undocumented
type ListInfo struct {
	// Object is the base model of ListInfo
	Object
	// ContentTypesEnabled undocumented
	ContentTypesEnabled *bool `json:"contentTypesEnabled,omitempty"`
	// Hidden undocumented
	Hidden *bool `json:"hidden,omitempty"`
	// Template undocumented
	Template *string `json:"template,omitempty"`
}

// ListItem undocumented
type ListItem struct {
	// BaseItem is the base model of ListItem
	BaseItem
	// ContentType undocumented
	ContentType *ContentTypeInfo `json:"contentType,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// Activities undocumented
	Activities []ItemActivityOLD `json:"activities,omitempty"`
	// Analytics undocumented
	Analytics *ItemAnalytics `json:"analytics,omitempty"`
	// DriveItem undocumented
	DriveItem *DriveItem `json:"driveItem,omitempty"`
	// Fields undocumented
	Fields *FieldValueSet `json:"fields,omitempty"`
	// Versions undocumented
	Versions []ListItemVersion `json:"versions,omitempty"`
}

// ListItemVersion undocumented
type ListItemVersion struct {
	// BaseItemVersion is the base model of ListItemVersion
	BaseItemVersion
	// Fields undocumented
	Fields *FieldValueSet `json:"fields,omitempty"`
}