package neutrino

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

var (
	defaultConnectionString = "localhost:27017"
	defaultDatabase = "test"
	defaultCollection = "test"
)

func init() {
	Initialize(defaultConnectionString)
}

func getDefaultDbService() DbService {
	return getDbService(defaultCollection)
}

func getDbService(collection string) DbService {
	return NewDbService(defaultDatabase, collection, mgo.Index{})
}

func ErrorTestField(t *testing.T, field string, expected, actual interface{}) {
	t.Error("expected", field, "to", "equal", expected, "got", actual);
}

//type fakeMessageService struct {
//	broadcastCalledTimes, broadcastJsonCalledTimes int
//	lastMessage string
//	lastJson map[string]interface{}
//}
//
//func (m *fakeMessageService) InitSocketHandler() http.Handler {
//	return nil
//}
//
//func (m *fakeMessageService) GetSessions() []sockjs.Session {
//	return nil
//}
//
//func (m *fakeMessageService) Broadcast(message string) {
//	m.broadcastCalledTimes++
//	m.lastMessage = message
//}
//
//func (m *fakeMessageService) BroadcastJSON(message map[string]interface{}) {
//	m.broadcastJsonCalledTimes++
//	m.lastJson = message
//}

func TestDbServiceGetSettings(t *testing.T) {
	d := getDefaultDbService()
	s := d.GetSettings();

	if s["ConnectionString"] != defaultConnectionString {
		ErrorTestField(t, "ConnectionString", defaultConnectionString, s["ConnectionString"])
	}

	if s["DbName"] != defaultDatabase {
		ErrorTestField(t, "DbName", defaultDatabase, s["DbName"])
	}

	if s["ColName"] != defaultCollection {
		ErrorTestField(t, "ColName", defaultCollection, s["ColName"])
	}
}

func TestDbServiceGetSession(t *testing.T) {
	d := getDefaultDbService()
	s := d.GetSession()

	if s == nil {
		t.Error("Database session is nil");
	}
}

func TestDbServiceGetDb(t *testing.T) {
	d := getDefaultDbService()
	session, db := d.GetDb()

	if db == nil {
		t.Error("Database is nil")
	}

	session.Close()
}

func TestDbServicGetCollection(t *testing.T) {
	d := getDefaultDbService()
	session, c := d.GetCollection()

	if c == nil {
		t.Error("Collection is nil")
	}

	session.Close()
}

func TestDbServiceInsertAndFindId(t *testing.T) {
	d := getDefaultDbService()
	doc := bson.M{"_id": "pesho", "number": 5}
	d.Insert(doc)

	res, err := d.FindId("pesho", nil)

	if err != nil {
		t.Fatal(err)
	}

	number := res["number"]

	if number != 5 {
		 ErrorTestField(t, "number", 5, number)
	}
}
