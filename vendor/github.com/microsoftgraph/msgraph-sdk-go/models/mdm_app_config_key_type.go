package models
import (
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type MdmAppConfigKeyType int

const (
    STRINGTYPE_MDMAPPCONFIGKEYTYPE MdmAppConfigKeyType = iota
    INTEGERTYPE_MDMAPPCONFIGKEYTYPE
    REALTYPE_MDMAPPCONFIGKEYTYPE
    BOOLEANTYPE_MDMAPPCONFIGKEYTYPE
    TOKENTYPE_MDMAPPCONFIGKEYTYPE
)

func (i MdmAppConfigKeyType) String() string {
    return []string{"stringType", "integerType", "realType", "booleanType", "tokenType"}[i]
}
func ParseMdmAppConfigKeyType(v string) (interface{}, error) {
    result := STRINGTYPE_MDMAPPCONFIGKEYTYPE
    switch v {
        case "stringType":
            result = STRINGTYPE_MDMAPPCONFIGKEYTYPE
        case "integerType":
            result = INTEGERTYPE_MDMAPPCONFIGKEYTYPE
        case "realType":
            result = REALTYPE_MDMAPPCONFIGKEYTYPE
        case "booleanType":
            result = BOOLEANTYPE_MDMAPPCONFIGKEYTYPE
        case "tokenType":
            result = TOKENTYPE_MDMAPPCONFIGKEYTYPE
        default:
            return 0, errors.New("Unknown MdmAppConfigKeyType value: " + v)
    }
    return &result, nil
}
func SerializeMdmAppConfigKeyType(values []MdmAppConfigKeyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
