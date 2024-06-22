package data

type RecentProfile struct {
	ID      string
	AddTime int64
}

func (bq *BuddyQueue) AddRecentProfiles() {

}

func (bq *BuddyQueue) DeleteRecentProfiles(ids ...string) {
	for _, v := range ids {
		for i := 0; i < len(bq.RecentMet); i++ {
			if bq.RecentMet[i].ID == v {
				bq.RecentMet = append(bq.RecentMet[:i], bq.RecentMet[i+1:]...)
				i--
			}
		}
	}
}
