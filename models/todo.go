// models/todo.go
package models

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Task      string             `bson:"task"`
    CreatedAt time.Time          `bson:"created_at"`
}

type TodoModel struct {
    Collection *mongo.Collection
}

func (m *TodoModel) Insert(todo Todo) error {
    _, err := m.Collection.InsertOne(context.TODO(), todo)
    return err
}

func (m *TodoModel) GetAll() ([]Todo, error) {
    var todos []Todo
    cursor, err := m.Collection.Find(context.TODO(), bson.D{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var todo Todo
        err := cursor.Decode(&todo)
        if err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

func (m *TodoModel) Delete(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = m.Collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    return err
}
