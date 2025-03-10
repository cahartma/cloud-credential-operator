package models
import (
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type VppTokenSyncStatus int

const (
    // Default status.
    NONE_VPPTOKENSYNCSTATUS VppTokenSyncStatus = iota
    // Last Sync in progress.
    INPROGRESS_VPPTOKENSYNCSTATUS
    // Last Sync completed successfully.
    COMPLETED_VPPTOKENSYNCSTATUS
    // Last Sync failed.
    FAILED_VPPTOKENSYNCSTATUS
)

func (i VppTokenSyncStatus) String() string {
    return []string{"none", "inProgress", "completed", "failed"}[i]
}
func ParseVppTokenSyncStatus(v string) (interface{}, error) {
    result := NONE_VPPTOKENSYNCSTATUS
    switch v {
        case "none":
            result = NONE_VPPTOKENSYNCSTATUS
        case "inProgress":
            result = INPROGRESS_VPPTOKENSYNCSTATUS
        case "completed":
            result = COMPLETED_VPPTOKENSYNCSTATUS
        case "failed":
            result = FAILED_VPPTOKENSYNCSTATUS
        default:
            return 0, errors.New("Unknown VppTokenSyncStatus value: " + v)
    }
    return &result, nil
}
func SerializeVppTokenSyncStatus(values []VppTokenSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
