package internal

type MPTopicType string

const (
	AnalyticsTopicMPTrack        MPTopicType = "analytics.mp.track"
	AnalyticsTopicMPUserProfiles MPTopicType = "analytics.mp.engage"
)

func (mt MPTopicType) String() string {
	return string(mt)
}
