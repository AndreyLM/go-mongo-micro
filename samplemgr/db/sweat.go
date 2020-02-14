package db

import (
	"context"
	"fmt"
	"time"

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
	db, err := GetDB()
	if err != nil {
		return
	}
	collection := db.Collection(SweatTable)
	cursor, err := collection.Find(context.TODO(), struct{}{})
	if err != nil {
		return
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &sweats); err != nil {
		return
	}

	return
}

// Create - creates new row in collection
func (s *Sweat) Create() (err error) {
	db, err := GetDB()
	if err != nil {
		fmt.Println("No Database connection: ", err)
		return
	}

	s.CreatedAt = time.Now()
	collection := db.Collection(SweatTable)
	if _, err = collection.InsertOne(context.TODO(), s); err != nil {
		fmt.Printf("Error inserting sweat: %v", s)
		return
	}

	fmt.Println("Inserted sweat into collection")
	return
}

// Delete - deletes sweat
func (s *Sweat) Delete() (err error) {
	db, err := GetDB()
	if err != nil {
		fmt.Println("No Database connection: ", err)
		return
	}
	collection := db.Collection(SweatTable)

	if _, err = collection.DeleteOne(context.TODO(), s); err != nil {
		fmt.Printf("Error deleting sweat: %v", s)
		return
	}

	fmt.Println("Deleted sweat from collection")
	return
}
