package transloadit

import (
	"testing"
)

var notificationAssemblyId string

func TestListNotifications(t *testing.T) {
	client := setup(t)

	notification, err := client.ListNotifications(&ListOptions{
		PageSize: 3,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(notification.Notifications) < 3 {
		t.Fatal("wrong number of notification")
	}

	if notification.Count == 0 {
		t.Fatal("wrong count")
	}

	if notification.Notifications[0].Id == "" {
		t.Fatal("wrong notification name")
	}

	notificationAssemblyId = notification.Notifications[0].AssemblyId
}

func TestReplayNotification(t *testing.T) {
	client := setup(t)

	res, err := client.ReplayNotification(notificationAssemblyId, "http://jsfiddle.net/echo/json/")
	if err != nil {
		t.Fatal(err)
	}

	if res["ok"].(string) != "ASSEMBLY_NOTIFICATION_REPLAYED" {
		t.Fatal("wrong status code returned")
	}
}
