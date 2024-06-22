package bi

const (
	// EventTypeUserSet MixPanel: https://developer.mixpanel.com/docs/http#section--set
	//ThinkingData:overwriting one or more user attributes,
	// overwriting the previous value if the attribute already exists
	EventTypeUserSet EventType = "#user.set"
	// EventTypeUserSetOnce MixPanel:https://developer.mixpanel.com/docs/http#section--set_once
	//ThinkingData: initialize one or more user attributes, if the attribute already exists, ignore the operation
	EventTypeUserSetOnce EventType = "#user.setOnce"
	//MixPanel:https://developer.mixpanel.com/docs/http#section--add
	//ThinkingData: adding calculations to one or more numeric user attributes
	EventTypeUserAdd EventType = "#user.add"
	//MixPanel:https://developer.mixpanel.com/docs/http#section--delete
	//ThinkingData: delete this user table
	EventTypeUserDel EventType = "#user.delete"

	// EventTypeUserAppend (Unimplemented)
	// MixPanel:https://developer.mixpanel.com/docs/http#section--append
	// /ThinkingData: not support
	EventTypeUserAppend EventType = "#user.append"
	// EventTypeUserUnion (Unimplemented)
	// MixPanel:https://developer.mixpanel.com/docs/http#section--union
	// ThinkingData: not support
	EventTypeUserUnion EventType = "#user.union"
	// EventTypeUserRemove (Unimplemented)
	// MixPanel: https://developer.mixpanel.com/docs/http#section--remove
	// ThinkingData: not support
	EventTypeUserRemove EventType = "#user.remove"
	//(Unimplemented)
	// MixPanel: https://developer.mixpanel.com/docs/http#section--unset
	// ThinkingData:not support
	EventTypeUserUnset EventType = "#user.unset"
)

type EventType string

func (et EventType) String() string {
	return string(et)
}
