package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type TeamSpecialization int

const (
    NONE_TEAMSPECIALIZATION TeamSpecialization = iota
    EDUCATIONSTANDARD_TEAMSPECIALIZATION
    EDUCATIONCLASS_TEAMSPECIALIZATION
    EDUCATIONPROFESSIONALLEARNINGCOMMUNITY_TEAMSPECIALIZATION
    EDUCATIONSTAFF_TEAMSPECIALIZATION
    HEALTHCARESTANDARD_TEAMSPECIALIZATION
    HEALTHCARECARECOORDINATION_TEAMSPECIALIZATION
    UNKNOWNFUTUREVALUE_TEAMSPECIALIZATION
)

func (i TeamSpecialization) String() string {
    return []string{"none", "educationStandard", "educationClass", "educationProfessionalLearningCommunity", "educationStaff", "healthcareStandard", "healthcareCareCoordination", "unknownFutureValue"}[i]
}
func ParseTeamSpecialization(v string) (interface{}, error) {
    result := NONE_TEAMSPECIALIZATION
    switch v {
        case "none":
            result = NONE_TEAMSPECIALIZATION
        case "educationStandard":
            result = EDUCATIONSTANDARD_TEAMSPECIALIZATION
        case "educationClass":
            result = EDUCATIONCLASS_TEAMSPECIALIZATION
        case "educationProfessionalLearningCommunity":
            result = EDUCATIONPROFESSIONALLEARNINGCOMMUNITY_TEAMSPECIALIZATION
        case "educationStaff":
            result = EDUCATIONSTAFF_TEAMSPECIALIZATION
        case "healthcareStandard":
            result = HEALTHCARESTANDARD_TEAMSPECIALIZATION
        case "healthcareCareCoordination":
            result = HEALTHCARECARECOORDINATION_TEAMSPECIALIZATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMSPECIALIZATION
        default:
            return 0, errors.New("Unknown TeamSpecialization value: " + v)
    }
    return &result, nil
}
func SerializeTeamSpecialization(values []TeamSpecialization) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
