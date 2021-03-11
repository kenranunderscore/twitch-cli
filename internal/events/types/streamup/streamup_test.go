// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package streamup

import (
	"encoding/json"
	"testing"

	"github.com/twitchdev/twitch-cli/internal/events"
	"github.com/twitchdev/twitch-cli/internal/models"
	"github.com/twitchdev/twitch-cli/internal/util"
)

var fromUser = "1234"
var toUser = "4567"

func TestEventSub(t *testing.T) {
	a := util.SetupTestEnv(t)

	params := *&events.MockEventParameters{
		FromUserID: fromUser,
		ToUserID:   toUser,
		Transport:  models.TransportEventSub,
		Trigger:    "streamup",
	}

	r, err := Event{}.GenerateEvent(params)
	a.Nil(err)

	var body models.StreamUpEventSubResponse
	err = json.Unmarshal(r.JSON, &body)
	a.Nil(err)

	// write actual tests here (making sure you set appropriate values and the like) for eventsub
}

func TestWebSub(t *testing.T) {
	a := util.SetupTestEnv(t)

	params := *&events.MockEventParameters{
		FromUserID: fromUser,
		ToUserID:   toUser,
		Transport:  models.TransportWebSub,
		Trigger:    "unsubscribe",
	}

	r, err := Event{}.GenerateEvent(params)
	a.Nil(err)

	var body models.StreamUpWebSubResponse
	err = json.Unmarshal(r.JSON, &body)
	a.Nil(err)

	// write tests here for websub
}

func TestFakeTransport(t *testing.T) {
	a := util.SetupTestEnv(t)

	params := *&events.MockEventParameters{
		FromUserID: fromUser,
		ToUserID:   toUser,
		Transport:  "fake_transport",
		Trigger:    "unsubscribe",
	}

	r, err := Event{}.GenerateEvent(params)
	a.Nil(err)
	a.Empty(r)
}
func TestValidTrigger(t *testing.T) {
	a := util.SetupTestEnv(t)

	r := Event{}.ValidTrigger("streamup")
	a.Equal(true, r)

	r = Event{}.ValidTrigger("notgift")
	a.Equal(false, r)
}

func TestValidTransport(t *testing.T) {
	a := util.SetupTestEnv(t)

	r := Event{}.ValidTransport(models.TransportEventSub)
	a.Equal(true, r)

	r = Event{}.ValidTransport("noteventsub")
	a.Equal(false, r)
}
func TestGetTopic(t *testing.T) {
	a := util.SetupTestEnv(t)

	r := Event{}.GetTopic(models.TransportEventSub, "streamup")
	a.NotNil(r)
}
