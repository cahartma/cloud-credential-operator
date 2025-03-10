package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type ProvisioningAction int

const (
    OTHER_PROVISIONINGACTION ProvisioningAction = iota
    CREATE_PROVISIONINGACTION
    DELETE_PROVISIONINGACTION
    DISABLE_PROVISIONINGACTION
    UPDATE_PROVISIONINGACTION
    STAGEDDELETE_PROVISIONINGACTION
    UNKNOWNFUTUREVALUE_PROVISIONINGACTION
)

func (i ProvisioningAction) String() string {
    return []string{"other", "create", "delete", "disable", "update", "stagedDelete", "unknownFutureValue"}[i]
}
func ParseProvisioningAction(v string) (interface{}, error) {
    result := OTHER_PROVISIONINGACTION
    switch v {
        case "other":
            result = OTHER_PROVISIONINGACTION
        case "create":
            result = CREATE_PROVISIONINGACTION
        case "delete":
            result = DELETE_PROVISIONINGACTION
        case "disable":
            result = DISABLE_PROVISIONINGACTION
        case "update":
            result = UPDATE_PROVISIONINGACTION
        case "stagedDelete":
            result = STAGEDDELETE_PROVISIONINGACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PROVISIONINGACTION
        default:
            return 0, errors.New("Unknown ProvisioningAction value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningAction(values []ProvisioningAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
