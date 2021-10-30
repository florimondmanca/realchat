package app

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	s := NewChatServer()
	s.Clock = NewFakeClock(1635605898)

	s.Start()
	defer s.Stop()

	ts := httptest.NewServer(s.Handler())
	defer ts.Close()

	url := fmt.Sprintf(
		"%s?EIO=1&transport=websocket&t=%v",
		strings.Replace(ts.URL, "http", "ws", 1),
		time.Now().Unix(),
	)

	// Connect first user
	cAlice := connect(t, url)
	defer cAlice.Close()

	// Send 'join' event
	err := cAlice.Emit("join", "alice")
	require.Nil(t, err)

	// Receive 'join' broadcast
	m, err := cAlice.RecvRaw()
	require.Nil(t, err)
	require.Equal(t, "42[\"join\",\"alice\"]\n", string(m))

	// Send 'msg' event
	err = cAlice.Emit("msg", "hello")
	require.Nil(t, err)

	// Receive 'msg' broadcast
	m, err = cAlice.RecvRaw()
	require.Nil(t, err)
	require.Equal(t, "42[\"msg\",1,1635605898,\"alice\",\"hello\"]\n", string(m))

	// Connect second user
	cBob := connect(t, url)
	defer cBob.Close()

	// Send 'join' event for second user
	err = cBob.Emit("join", "bob")
	require.Nil(t, err)

	// Receive 'join' broadcast for first user
	m, err = cAlice.RecvRaw()
	require.Nil(t, err)
	require.Equal(t, "42[\"join\",\"bob\"]\n", string(m))
	// Receive history for second user
	m, err = cBob.RecvRaw()
	require.Nil(t, err)
	require.Equal(t, "42[\"join\",\"alice\"]\n", string(m))
	m, err = cBob.RecvRaw()
	require.Nil(t, err)
	require.Equal(t, "42[\"msg\",1,1635605898,\"alice\",\"hello\"]\n", string(m))
	// Receive 'join' broadcast for second user
	m, err = cBob.RecvRaw()
	require.Nil(t, err)
	require.Equal(t, "42[\"join\",\"bob\"]\n", string(m))

	// Send 'msg' event for second user
	err = cBob.Emit("msg", "hi")
	require.Nil(t, err)

	// Receive 'msg' broadcast for first user
	m, err = cAlice.RecvRaw()
	require.Nil(t, err)
	require.Equal(t, "42[\"msg\",2,1635605898,\"bob\",\"hi\"]\n", string(m))
	// Receive 'msg' broadcast for second user
	m, err = cBob.RecvRaw()
	require.Nil(t, err)
	require.Equal(t, "42[\"msg\",2,1635605898,\"bob\",\"hi\"]\n", string(m))

	// Disconnect first user
	cAlice.Close()

	// Receive 'leave' broadcast for second user
	m, err = cBob.RecvRaw()
	require.Nil(t, err)
	require.Equal(t, "42[\"leave\",\"alice\"]\n", string(m))
}

type client struct {
	ws *websocket.Conn
}

func connect(t *testing.T, url string) *client {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	require.Nil(t, err)

	// Follow socket.io session protocol.
	// https://github.com/socketio/socket.io-protocol#sample-session

	// Receive OPEN packet
	_, m, err := ws.ReadMessage()
	require.Nil(t, err)
	require.Equal(t, byte('0'), m[0])

	// Receive connection request approval
	_, m, err = ws.ReadMessage()
	require.Nil(t, err)
	require.Equal(t, "40", string(m))

	return &client{ws}
}

func (c *client) Emit(event string, args ...string) error {
	// Send EVENT socket.io packet
	// See: https://github.com/socketio/socket.io-protocol#sample-session
	elems := []string{event}
	elems = append(elems, args...)
	quotedElems := []string{}
	for _, e := range elems {
		quotedElems = append(quotedElems, fmt.Sprintf("\"%s\"", e))
	}
	content := fmt.Sprintf("42[%s]", strings.Join(quotedElems, ","))

	return c.ws.WriteMessage(websocket.TextMessage, []byte(content))
}

func (c *client) RecvRaw() ([]byte, error) {
	_, m, err := c.ws.ReadMessage()
	return m, err
}

func (c *client) Close() {
	c.ws.Close()
}

type fakeClock struct {
	Clock
	ts int64
}

func NewFakeClock(ts int64) *fakeClock {
	return &fakeClock{ts: ts}
}

func (c *fakeClock) Now() int64 {
	return c.ts
}
