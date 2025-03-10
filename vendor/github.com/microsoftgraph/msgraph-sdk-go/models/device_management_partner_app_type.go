package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type DeviceManagementPartnerAppType int

const (
    // Partner App type is unknown.
    UNKNOWN_DEVICEMANAGEMENTPARTNERAPPTYPE DeviceManagementPartnerAppType = iota
    // Partner App is Single tenant in AAD.
    SINGLETENANTAPP_DEVICEMANAGEMENTPARTNERAPPTYPE
    // Partner App is Multi tenant in AAD.
    MULTITENANTAPP_DEVICEMANAGEMENTPARTNERAPPTYPE
)

func (i DeviceManagementPartnerAppType) String() string {
    return []string{"unknown", "singleTenantApp", "multiTenantApp"}[i]
}
func ParseDeviceManagementPartnerAppType(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEMANAGEMENTPARTNERAPPTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEMANAGEMENTPARTNERAPPTYPE
        case "singleTenantApp":
            result = SINGLETENANTAPP_DEVICEMANAGEMENTPARTNERAPPTYPE
        case "multiTenantApp":
            result = MULTITENANTAPP_DEVICEMANAGEMENTPARTNERAPPTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementPartnerAppType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementPartnerAppType(values []DeviceManagementPartnerAppType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
