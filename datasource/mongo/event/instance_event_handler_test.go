package event

import (
	"bou.ke/monkey"
	"reflect"
	"testing"

	"github.com/apache/servicecomb-service-center/datasource/mongo/client"
	"github.com/apache/servicecomb-service-center/datasource/mongo/sd"
	"github.com/apache/servicecomb-service-center/server/syncernotify"
	"github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/go-chassis/v2/storage"
	"github.com/stretchr/testify/assert"
)

func init() {
	config := storage.Options{
		URI: "mongodb://localhost:27017",
	}
	client.NewMongoClient(config)
}

func TestInstanceEventHandler_OnEvent(t *testing.T) {

	t.Run("microservice not nil after query database", func(t *testing.T) {
		h := InstanceEventHandler{}
		h.OnEvent(mongoAssign())
		assert.NotNil(t, discovery.MicroService{})
	})
	t.Run("when there is no such a service in database", func(t *testing.T) {
		h := InstanceEventHandler{}
		h.OnEvent(mongoEventWronServiceId())
		assert.Error(t, assert.AnError, "get from db failed")
	})
	t.Run("OnEvent test when syncer notify center closed", func(t *testing.T) {
		h := InstanceEventHandler{}
		h.OnEvent(mongoAssign())
		assert.Error(t, assert.AnError)
	})
	t.Run("OnEvent test when syncer notify center open", func(t *testing.T) {
		defer monkey.UnpatchAll()
		monkey.PatchInstanceMethod(reflect.TypeOf((*syncernotify.Service)(nil)), "Closed",
			func(service *syncernotify.Service) bool {
				return false
			})
		h := InstanceEventHandler{}
		h.OnEvent(mongoAssign())
		assert.Equal(t, false, t.Failed(), "add event succeed")
	})
}

func TestNotifySyncerInstanceEvent(t *testing.T) {
	t.Run("test when data is ok", func(t *testing.T) {
		mongoEvent := mongoAssign()
		microService := getMicroService()
		NotifySyncerInstanceEvent(mongoEvent, microService)
		assert.Equal(t, false, t.Failed())
	})
}

func mongoAssign() sd.MongoEvent {
	sd.Store().Service().Cache()
	endPoints := []string{"127.0.0.1:27017"}
	instance := discovery.MicroServiceInstance{
		InstanceId: "f73dceb440f711eba63ffa163e7cdcb8",
		ServiceId:  "2a20507274fc71c925d138341517dce14b600744",
		Endpoints:  endPoints,
	}
	mongoInstance := sd.Instance{}
	mongoInstance.InstanceInfo = &instance
	mongoInstance.Domain = "default"
	mongoInstance.Project = "default"
	mongoEvent := sd.MongoEvent{}
	mongoEvent.DocumentID = "5fdc483b4a885f69317e3505"
	mongoEvent.Value = mongoInstance
	mongoEvent.Type = discovery.EVT_CREATE
	mongoEvent.ResourceID = "f73dceb440f711eba63ffa163e7cdcb8"
	return mongoEvent
}

func mongoEventWronServiceId() sd.MongoEvent {
	sd.Store().Service().Cache()
	endPoints := []string{"127.0.0.1:27017"}
	instance := discovery.MicroServiceInstance{
		InstanceId: "f73dceb440f711eba63ffa163e7cdcb8",
		ServiceId:  "2a20507274fc71c925d138341517dce14b6007443333",
		Endpoints:  endPoints,
	}
	mongoInstance := sd.Instance{}
	mongoInstance.InstanceInfo = &instance
	mongoInstance.Domain = "default"
	mongoInstance.Project = "default"
	mongoEvent := sd.MongoEvent{}
	mongoEvent.DocumentID = "5fdc483b4a885f69317e3505"
	mongoEvent.Value = mongoInstance
	mongoEvent.Type = discovery.EVT_CREATE
	mongoEvent.ResourceID = "f73dceb440f711eba63ffa163e7cdcb8"
	return mongoEvent
}

func getMicroService() *discovery.MicroService {
	microService := discovery.MicroService{
		ServiceId:    "1efe8be8eacce1efbf67967978133572fb8b5667",
		AppId:        "default",
		ServiceName:  "ProviderDemoService1-2",
		Version:      "1.0.0",
		Level:        "BACK",
		Status:       "UP",
		Timestamp:    "1608260891",
		ModTimestamp: "1608260891",
	}
	return &microService
}
