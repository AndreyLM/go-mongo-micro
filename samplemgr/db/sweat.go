package db

import (
	"context"
	"time"

	"github.com/andreylm/go-mongo-micro/sqmplemgr/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SweatTable - collection name
const SweatTable string = "sweat"

// Sweat - sweat
type Sweat struct {
	// Database specific fields
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`

	// Potential disease Diagnosis
	Glucose  float32 `bson:"glucose" json:"glucose,omitempty"`
	Chloride float32 `bson:"chloride" json:"chloride,omitempty"`

	// Electrolytes
	Sodium    float32 `bson:"sodium" json:"sodium,omitempty"`
	Potassium float32 `bson:"potassium" json:"potassium,omitempty"`
	Magnesium float32 `bson:"magnesium" json:"magnesium,omitempty"`
	Calcium   float32 `bson:"calcium" json:"calcium,omitempty"`

	// Environmetal conditions and determining criteria
	Humidity        float32 `bson:"humidity" json:"humidity,omitempty"`
	RootTemperature float32 `bson:"root_temperature" json:"root_temperature,omitempty"`
	BodyTemperature float32 `bson:"body_temperature" json:"body_temperature,omitempty"`
	HeartBeat       int32   `bson:"heart_beat" json:"heart_beat,omitempty"`
}

// ListAllSweat - gets all
func ListAllSweat() (sweats []Sweat, err error) {
	db := GetDB()

	collection := db.Collection(SweatTable)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		logger.Get().Infof("Find error: %v", err)
		return
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &sweats); err != nil {
		logger.Get().Infof("Error getting data: ", err)
		return
	}

	return
}

// GetByID - gets sweat by id
func GetByID(ID string) (sweat Sweat, err error) {
	db := GetDB()

	collection := db.Collection(SweatTable)
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		logger.Get().Infof("Invalid id: %v", err)
		return
	}

	sr := collection.FindOne(context.TODO(), bson.M{"_id": objID})
	err = sr.Decode(&sweat)
	if err != nil {
		logger.Get().Infof("Error decoding data: ", err)
	}

	return
}

// Create - creates new row in collection
func (s *Sweat) Create(ctx context.Context) (err error) {
	var tempUserID primitive.ObjectID
	userID := ""

	if ctx.Value("UserID") != nil {
		userID = ctx.Value("UserID").(string)
	}

	if userID == "" {
		logger.Get().Error("User not specified in context")
		tempUserID = primitive.NewObjectID()
	}
	s.UserID = tempUserID
	s.CreatedAt = time.Now()

	db := GetDB()

	collection := db.Collection(SweatTable)
	if _, err = collection.InsertOne(ctx, s); err != nil {
		logger.Get().Infof("Error inserting sweat: %v", s)
		return
	}

	logger.Get().Info("Inserted sweat into collection")
	return
}

// Delete - deletes sweat
func (s *Sweat) Delete() (err error) {
	db := GetDB()

	collection := db.Collection(SweatTable)

	if _, err = collection.DeleteOne(context.TODO(), s); err != nil {
		logger.Get().Infof("Error deleting sweat: %v", s)
		return
	}

	logger.Get().Info("Deleted sweat from collection")
	return
}
