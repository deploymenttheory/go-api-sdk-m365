// shared_error_messages.go
package shared

// Type refers to string representation of target object type. I.e buildings, policies, computergroups

const (
	// Pagination - type: string, error: any
	ErrorMsgFailedPaginatedGet = "failed to get paginated %s, error: %v"

	// Graph operations - format always type: string, id/name: any, error: any
	ErrorMsgFailedGet            = "failed to get %s, error: %v"
	ErrorMsgFailedGetByID        = "failed to get %s by id: %v, error: %v"
	ErrorMsgFailedGetByName      = "failed to get %s by name: %s, error: %v"
	ErrorMsgFailedCreate         = "failed to create %s, error: %v"
	ErrorMsgFailedUpdate         = "failed to update %s, error: %v"
	ErrorMsgFailedUpdateByID     = "failed to update %s by id: %v, error: %v"
	ErrorMsgFailedUpdateByName   = "failed to update %s by name: %s, error: %v"
	ErrorMsgFailedDeleteByID     = "failed to delete %s by id: %v, error: %v"
	ErrorMsgFailedDeleteByName   = "failed to delete %s by name: %s, error: %v"
	ErrorMsgFailedDeleteMultiple = "failed to delete multiple %s, by ids: %v, error: %v"
	ErrorMsgFailedAssign         = "failed to assign %s by id: %v, error: %v"
	ErrorMsgFailedCreateCopy     = "failed to copy %s with id: %v, error: %v"
	ErrorMsgFailedReorder        = "failed to set the priority of %s to id: %v, error: %v"

	// Mapstructure - type: string, error: any
	ErrorMsgFailedMapstruct = "failed to map interfaced %s to structs, error: %v"

	// JSON Marshalling
	ErrorMsgFailedJsonMarshal = "failed to marshal %s, error: %v"

	// Client Credentials
	ErrorMsgFailedRefreshClientCreds = "failed to refresh client credentials at id: %s, error :%v"

	// Logging
	// matched configuration
	LogMsgFoundMatchedConfigID = "found matched configuration ID: %v for %s with search name: %s"
)
