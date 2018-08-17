package db

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Cursor devuelve un apuntador con la conexión a la bd y la colección especifica
func Cursor(session *mgo.Session, Collection string) *mgo.Collection {
	mongoDB := beego.AppConfig.String("mongo_db")
	c := session.DB(mongoDB).C(Collection)
	return c
}

// GetSession crea una sesión con las credenciales
func GetSession() (*mgo.Session, error) {

	mongoHost := beego.AppConfig.String("mongo_host")
	mongoUser := beego.AppConfig.String("mongo_user")
	mongoPassword := beego.AppConfig.String("mongo_pass")
	mongoDatabase := beego.AppConfig.String("mongo_db_connect")

	info := &mgo.DialInfo{
		Addrs:    []string{mongoHost},
		Timeout:  60 * time.Second,
		Database: mongoDatabase,
		Username: mongoUser,
		Password: mongoPassword,
	}

	// beego.Info("mongoHost: ", mongoHost, "mongoUser: ", mongoUser, "mongoPassword: ", mongoPassword, "mongoDatabase: ", mongoDatabase)

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println("Helo this is an error!")
		panic(err)
	} else {
		session.SetMode(mgo.Monotonic, true)
	}

	return session, err
}

func GetAll(session *mgo.Session, collection string) []bson.M {
	c := Cursor(session, collection)
	defer session.Close()
	var records []bson.M
	err := c.Find(bson.M{}).All(&records)
	if err != nil {
		fmt.Println(err)
	}

	return records
}

func getType(i interface{}) interface{} {
	return reflect.New(reflect.TypeOf(i))
}
