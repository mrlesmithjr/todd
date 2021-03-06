/*
   ToDD agent responses

   These are asynchronous responses sent back to the server, usually as a response to a task.
   Not all tasks result in an agent sending a response to the server. Responses are for highly sensitive operations like test distribution and execution.

    Copyright 2016 Matt Oswalt. Use or modification of this
    source code is governed by the license provided here:
    https://github.com/toddproject/todd/blob/master/LICENSE
*/

package responses

// Response type constants.
const (
	TypeSetAgentStatus = "AgentStatus"
	TypeUploadTestData = "TestData"
)

// Response is an interface to define Response behavior
type Response interface{}

// Base is a struct that is intended to be embedded by specific response structs. Both of these in conjunction
// are used primarily to house the JSON message for passing responses over the comms package (i.e. message queue), but may also contain important
// dependencies of the response, such as an HTTP handler.
type Base struct {
	AgentUUID string `json:"agent_uuid"`
	Type      string `json:"type"`
}

// SetAgentStatus defines this particular response.
type SetAgentStatus struct {
	Base
	TestUUID string `json:"test_uuid"`
	Status   string `json:"status"`
}

// NewSetAgentStatus returns a new SetAgentStatus response.
func NewSetAgentStatus(agentUUID, testUUID, status string) SetAgentStatus {
	return SetAgentStatus{
		Base: Base{
			AgentUUID: agentUUID,
			Type:      TypeSetAgentStatus,
		},
		TestUUID: testUUID,
		Status:   status,
	}
}

// UploadTestData defines this particular response.
type UploadTestData struct {
	Base
	TestUUID string `json:"test_uuid"`
	TestData string `json:"test_data"`
}

// NewUploadTestData returns a new UploadTestData response.
func NewUploadTestData(agentUUID, testUUID, testData string) UploadTestData {
	return UploadTestData{
		Base: Base{
			AgentUUID: agentUUID,
			Type:      TypeUploadTestData,
		},
		TestUUID: testUUID,
		TestData: testData,
	}
}
