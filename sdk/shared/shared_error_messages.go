package shared

// Type refers to string representation of target object type. I.e buildings, policies, computergroups

const (
	// Pagination - type: string, error: any
	errorMsgFailedPaginatedGet = "failed to get paginated %s, error: %v"

	// CRUD - format always type: string, id/name: any, error: any
	errorMsgFailedGet            = "failed to get %s, error: %v"
	errorMsgFailedGetByID        = "failed to get %s by id: %v, error: %v"
	errorMsgFailedGetByName      = "failed to get %s by name: %s, error: %v"
	errorMsgFailedCreate         = "failed to create %s, error: %v"
	errorMsgFailedUpdate         = "failed to update %s, error: %v"
	errorMsgFailedUpdateByID     = "failed to update %s by id: %v, error: %v"
	errorMsgFailedUpdateByName   = "failed to update %s by name: %s, error: %v"
	errorMsgFailedDeleteByID     = "failed to delete %s by id: %v, error: %v"
	errorMsgFailedDeleteByName   = "failed to delete %s by name: %s, error: %v"
	errorMsgFailedDeleteMultiple = "failed to delete multiple %s, by ids: %v, error: %v"

	// Mapstructure - type: string, error: any
	errorMsgFailedMapstruct = "failed to map interfaced %s to structs, error: %v"

	// JSON Marshalling
	errorMsgFailedJsonMarshal = "failed to marshal %s, error: %v"

	// Client Credentials
	errorMsgFailedRefreshClientCreds = "failed to refresh client credentials at id: %s, error :%v"
)
