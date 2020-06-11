package globalUtils

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	//AuditTopic: Topic to be used when publishing and subscribing to audit related message in the broker
	AuditTopic = "Audit"
	//AuditQueueInsert: Queue to be used when retrieve audit messages from the broker via a subscription to allow multiple instances processing messages in parallel
	AuditQueueInsert = "Audit.Insert"
)

//auditMsg: Structure of the audit messages to be sent to the broker
type auditMsg struct {
	topic        string
	objectToSend []byte
	header       AuditMsgHeader
}

func (a auditMsg) Header() AuditMsgHeader {
	return a.header
}

func (a auditMsg) ObjectToSend() []byte {
	return a.objectToSend
}

func (a auditMsg) Topic() string {
	return a.topic
}

//AuditMsgHeader: structure of the header portion of the audit message to be sent to the broker
type AuditMsgHeader map[string]string

//Struct version of the AuditMsgHeader (map) to allow for easier handling
type AuditMsgHeaderStruct struct {
	ServiceName string
	ActionFunc  string
	ActionType  string
	ObjectId    int64
	PerformedBy int64
	ActionTime  time.Time
	ObjectName  string
}

func (a *AuditMsgHeaderStruct) GetObjectName() string {
	return a.ObjectName
}

func (a *AuditMsgHeaderStruct) GetActionTime() time.Time {
	return a.ActionTime
}

func (a *AuditMsgHeaderStruct) GetPerformedBy() int64 {
	return a.PerformedBy
}

func (a *AuditMsgHeaderStruct) GetObjectId() int64 {
	return a.ObjectId
}

func (a *AuditMsgHeaderStruct) GetActionType() string {
	return a.ActionType
}

func (a *AuditMsgHeaderStruct) GetActionFunc() string {
	return a.ActionFunc
}

func (a *AuditMsgHeaderStruct) GetServiceName() string {
	return a.ServiceName
}

//NewAuditMsg: Validate and create a new Audit message that is ready to be sent out to the broker
func NewAuditMsg(serviceName, actionFunc, actionType string, performedBy int64, objectName string, objectId int64, objectToSend []byte) (*auditMsg, error) {
	var missingFields string
	if serviceName == "" {
		missingFields += " serviceName,"
	}
	if actionFunc == "" {
		missingFields += " actionFunc,"
	}
	if actionType == "" {
		missingFields += " actionType,"
	}
	if performedBy == 0 {
		missingFields += " performedBy,"
	}
	if objectId == 0 {
		missingFields += " objectId,"
	}
	if objectToSend == nil {
		missingFields += " objectToSend,"
	}
	if missingFields != "" {
		return nil, fmt.Errorf("all fields must be filled in audit messages. The following fields are empty: %s",
			missingFields[1:len(missingFields)-1])
	}
	aud := auditMsg{
		topic:        AuditTopic,
		objectToSend: objectToSend,
		header: AuditMsgHeader{
			"service":     serviceName,
			"actionFunc":  actionFunc,
			"actionType":  actionType,
			"objectId":    strconv.FormatInt(objectId, 10),
			"performedBy": strconv.FormatInt(performedBy, 10),
			"actionTime":  time.Now().Format(DateLayoutISO),
			"objectName":  objectName,
		},
	}
	//log.Printf(" objid : %v, objidstr: %s ",  objectId, strconv.FormatInt(performedBy, 10))
	return &aud, nil
}

//AuditMsgHeaderToStruct: convert the AuditMsgHeader to its struct based counterpart AuditMsgHeaderStruct
func AuditMsgHeaderToStruct(header AuditMsgHeader) (*AuditMsgHeaderStruct, error) {
	if header == nil {
		return nil, fmt.Errorf("message header cannot be nil when trying to convert to struct")
	}
	objectId, err := strconv.ParseInt(header["objectId"], 10, 64)
	if err != nil {
		return nil, err
	}
	performedBy, err := strconv.ParseInt(header["performedBy"], 10, 64)
	if err != nil {
		return nil, err
	}
	actionTime, err := time.Parse(DateLayoutISO, header["actionTime"])
	if err != nil {
		log.Printf("Unable to Format date %v\n", header["actionTime"])
	}
	headerStruct := &AuditMsgHeaderStruct{
		ServiceName: header["service"],
		ActionFunc:  header["actionFunc"],
		ActionType:  header["actionType"],
		ObjectId:    objectId,
		PerformedBy: performedBy,
		ActionTime:  actionTime,
		ObjectName:  header["objectName"],
	}

	return headerStruct, nil
}