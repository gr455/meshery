package mesh

import "github.com/layer5io/meshkit/errors"

const (
	ErrGettingSessionDataCode                = "1009"
	ErrNoAdaptersCode                        = "1010"
	ErrPromptCode                            = "1011"
	ErrCreatingDeployRequestCode             = "1012"
	ErrCreatingDeployResponseRequestCode     = "1013"
	ErrAddingAuthDetailsCode                 = "1014"
	ErrCreatingDeployResponseStreamCode      = "1015"
	ErrTimeoutWaitingForDeployResponseCode   = "1016"
	ErrFailedDeployingMeshCode               = "1017"
	ErrCreatingValidateRequestCode           = "1018"
	ErrCreatingValidateResponseRequestCode   = "1019"
	ErrTimeoutWaitingForValidateResponseCode = "1020"
	ErrSMIConformanceTestsFailedCode         = "1021"
	ErrCreatingRemoveRequestCode             = "1022"
	ErrCreatingRemoveResponseRequestCode     = "1023"
	ErrCreatingRemoveResponseStreamCode      = "1024"
	ErrTimeoutWaitingForRemoveResponseCode   = "1025"
	ErrFailedRemovingMeshCode                = "1026"
)

var (
	// no adapters found
	ErrNoAdapters = errors.New(ErrNoAdaptersCode, errors.Fatal, []string{"Adapter for required mesh not found"}, []string{"Adapter for required mesh not found"}, []string{""}, []string{"Deploy the proper Meshery Adapter for your service mesh"})

	ErrFailedDeployingMesh = errors.New(ErrFailedDeployingMeshCode, errors.Fatal, []string{"Failed to deploy the service mesh"}, []string{"Failed to deploy the service mesh"}, []string{}, []string{"Check your environment and try again"})

	ErrTimeoutWaitingForDeployResponse = errors.New(ErrTimeoutWaitingForDeployResponseCode, errors.Fatal, []string{"Timed out waiting for deploy event"}, []string{"Timed out waiting for deployment"}, []string{}, []string{"Check your environment and try again"})

	ErrTimeoutWaitingForValidateResponse = errors.New(ErrTimeoutWaitingForValidateResponseCode, errors.Fatal, []string{"Timed out waiting for validate response"}, []string{"Timed out waiting for validate response"}, []string{""}, []string{"Check your environment and try again"})

	ErrSMIConformanceTestsFailed = errors.New(ErrSMIConformanceTestsFailedCode, errors.Fatal, []string{"SMI conformance tests failed"}, []string{"SMI conformance tests failed"}, []string{}, []string{"Join https://layer5io.slack.com/archives/C010H0HE2E6"})

	ErrTimeoutWaitingForRemoveResponse = errors.New(ErrTimeoutWaitingForRemoveResponseCode, errors.Fatal, []string{"Timed out waiting for remove event"}, []string{}, []string{}, []string{"Check your environment and try again"})

	ErrFailedRemovingMesh = errors.New(ErrFailedRemovingMeshCode, errors.Fatal, []string{"Failed to remove the service mesh"}, []string{""}, []string{}, []string{"Check your environment and try again"})
)

// When unable to get release data
func ErrGettingSessionData(err error) error {
	return errors.New(ErrGettingSessionDataCode, errors.Fatal, []string{"Unable to fetch session data"}, []string{err.Error()}, []string{}, []string{})
}

func ErrPrompt(err error) error {
	return errors.New(ErrPromptCode, errors.Fatal, []string{"Error while reading selected option"}, []string{err.Error()}, []string{}, []string{})
}

func ErrCreatingDeployRequest(err error) error {
	return errors.New(ErrCreatingDeployRequestCode, errors.Fatal, []string{"Error sending deploy request"}, []string{err.Error()}, []string{}, []string{})
}

func ErrCreatingDeployResponseRequest(err error) error {
	return errors.New(ErrCreatingDeployResponseRequestCode, errors.Fatal, []string{"Error creating request for deploy response"}, []string{err.Error()}, []string{}, []string{})
}

func ErrAddingAuthDetails(err error) error {
	return errors.New(ErrAddingAuthDetailsCode, errors.Fatal, []string{"Error adding auth details"}, []string{err.Error()}, []string{}, []string{})
}

func ErrCreatingDeployResponseStream(err error) error {
	return errors.New(ErrCreatingDeployResponseStreamCode, errors.Fatal, []string{"Error creating deploy event response stream"}, []string{err.Error()}, []string{}, []string{})
}

func ErrCreatingValidateRequest(err error) error {
	return errors.New(ErrCreatingValidateRequestCode, errors.Fatal, []string{"Error sending Validate request"}, []string{err.Error()}, []string{}, []string{})
}

func ErrCreatingValidateResponseRequest(err error) error {
	return errors.New(ErrCreatingValidateResponseRequestCode, errors.Fatal, []string{"Error creating request for validate response"}, []string{err.Error()}, []string{}, []string{})
}

func ErrCreatingValidateResponseStream(err error) error {
	return errors.New(ErrCreatingDeployResponseStreamCode, errors.Fatal, []string{"Error creating validate event response stream"}, []string{err.Error()}, []string{}, []string{})
}

func ErrCreatingRemoveRequest(err error) error {
	return errors.New(ErrCreatingRemoveRequestCode, errors.Fatal, []string{"Error sending remove request"}, []string{err.Error()}, []string{}, []string{})
}

func ErrCreatingRemoveResponseRequest(err error) error {
	return errors.New(ErrCreatingRemoveResponseRequestCode, errors.Fatal, []string{"Error creating request for remove response"}, []string{err.Error()}, []string{}, []string{})
}

func ErrCreatingRemoveResponseStream(err error) error {
	return errors.New(ErrCreatingRemoveResponseStreamCode, errors.Fatal, []string{"Error creating deploy event response stream"}, []string{err.Error()}, []string{}, []string{})
}
