package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/moke-game/platform/api/gen/profile/api"
)

type PrivateDao struct {
	collection *mongo.Collection
}

func (d *PrivateDao) Init(db *mongo.Database) error {
	key, e := NewProfileCollectionName()
	if e != nil {
		return e
	}
	d.collection = db.Collection(key.String())
	return nil
}

// GetProfiles get profile by uid
func (d *PrivateDao) GetProfiles(platformId, channel string, uid ...string) ([]*pb.Profile, error) {
	ks := make([]string, 0)
	for _, v := range uid {
		if k, err := NewProfileKey(v); err != nil {
			return nil, err
		} else {
			ks = append(ks, k.String())
		}
	}
	filters := bson.M{"_id": bson.M{"$in": ks}}
	if platformId != "" {
		filters["data.platformid"] = platformId
	}
	if channel != "" {
		filters["data.channel"] = channel
	}
	cur, err := d.collection.Find(context.Background(), filters)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	profiles := make([]*pb.Profile, 0)
	for cur.Next(context.Background()) {
		data := cur.Current.Lookup("data")
		profile := &pb.Profile{}
		if err := data.Unmarshal(profile); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

// GetProfileByNickname find profile name by regex
// https://www.mongodb.com/docs/manual/reference/operator/query/regex/
func (d *PrivateDao) GetProfileByNickname(
	platformId string,
	channel string,
	name string,
	isRegex bool,
	page, size int64,
) ([]*pb.Profile, error) {
	var filter bson.M
	option := &options.FindOptions{}
	if isRegex {
		filter = bson.M{
			"data.nickname": bson.M{"$regex": "^" + name, "$options": "i"},
		}
		if platformId != "" {
			filter["data.platformid"] = platformId
		}
		if channel != "" {
			filter["data.channel"] = channel
		}
		skip := (page - 1) * size
		option.Skip = &skip
		option.Limit = &size
	} else {
		filter = bson.M{"data.nickname": name}
	}
	cur, err := d.collection.Find(context.Background(), filter, option)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	profiles := make([]*pb.Profile, 0)
	for cur.Next(context.Background()) {
		data := cur.Current.Lookup("data")
		profile := &pb.Profile{}
		if err := data.Unmarshal(profile); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

// GetAllProfiles get all profiles by page
// https://www.mongodb.com/docs/atlas/atlas-search/tutorial/pagination-tutorial/
func (d *PrivateDao) GetAllProfiles(
	platformId string,
	channel string, page, size int64) ([]*pb.Profile, error) {
	skip := (page - 1) * size
	filters := bson.M{}
	if platformId != "" {
		filters["data.platformid"] = platformId
	}
	if channel != "" {
		filters["data.channel"] = channel
	}
	cur, err := d.collection.Find(context.Background(), filters, &options.FindOptions{
		Skip:  &skip,
		Limit: &size,
	})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	profiles := make([]*pb.Profile, 0)
	for cur.Next(context.Background()) {
		data := cur.Current.Lookup("data")
		profile := &pb.Profile{}
		if err := data.Unmarshal(profile); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil

}

func (d *PrivateDao) GetProfilesByAccount(val string) (*pb.Profile, error) {
	filter := bson.M{"data.account": val}
	cur, err := d.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	profile := &pb.Profile{}
	for cur.Next(context.Background()) {
		data := cur.Current.Lookup("data")
		if err := data.Unmarshal(profile); err != nil {
			return nil, err
		}
	}
	return profile, nil
}
