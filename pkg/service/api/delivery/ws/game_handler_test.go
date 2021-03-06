package ws

import (
	"bufio"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	userID = 2
)

func TestGameHandler(t *testing.T) {
	s := NewHandlerTestSuite(t)

	players := []*model.Player{{UserID: userID}}
	foods := []model.Food{{ID: 1, Position: model.Position{0, 100}}}
	s.auth.EXPECT().GetUserID(test.Session).Return(userID, nil)
	s.game.EXPECT().StartGame(userID).Return(nil)
	s.game.EXPECT().StopGame(userID).Return(nil)
	s.game.EXPECT().GetPlayers(userID).Return(players, nil)
	s.game.EXPECT().GetFood(userID).Return(foods, nil)
	s.game.EXPECT().ListenEvents(userID).Return(make(chan map[string]interface{}), nil)

	conn, err := s.Connect()
	assert.NoError(t, err)

	assert.NoError(t, conn.WriteJSON(map[string]interface{}{"type": game.EventStart}))

	var response map[string]interface{}
	assert.NoError(t, conn.ReadJSON(&response))
	assert.Equal(t, game.EventStart, response["type"])
	assert.NotEmpty(t, response["players"])
	assert.NotEmpty(t, response["food"])

	assert.NoError(t, conn.WriteJSON(map[string]interface{}{"type": game.EventStop}))
}

func TestGameHandler_UnknownRequest(t *testing.T) {
	s := NewHandlerTestSuite(t)

	s.auth.EXPECT().GetUserID(test.Session).Return(userID, nil)

	conn, err := s.Connect()
	assert.NoError(t, err)
	assert.NoError(t, conn.WriteJSON(map[string]interface{}{"type": "unknown"}))
	assert.Equal(t, errUnknownRequestType, s.ListenError())
}

type HandlerTestSuite struct {
	handler *GameHandler
	game    *api.MockGameUsecase
	auth    *api.MockAuthUsecase

	server net.Conn
	dialer *websocket.Dialer
	conn   *websocket.Conn
	e      *echo.Echo
	t      *testing.T
}

func NewHandlerTestSuite(t *testing.T) *HandlerTestSuite {
	controller := gomock.NewController(t)
	suite := HandlerTestSuite{
		e:    echo.New(),
		t:    t,
		game: api.NewMockGameUsecase(controller),
		auth: api.NewMockAuthUsecase(controller),
	}
	logger := mock.NewMockLogger(t)
	suite.handler = NewGameHandler(suite.e, suite.game, suite.auth, logger)

	client, server := net.Pipe()
	suite.dialer = &websocket.Dialer{NetDial: func(network, addr string) (net.Conn, error) { return client, nil }}
	suite.server = server
	go func() {
		err := suite.runServer()
		assert.NoError(t, err)
	}()

	return &suite
}

func (s *HandlerTestSuite) Connect() (*websocket.Conn, error) {
	conn, _, err := s.dialer.Dial("ws://whatever", nil)
	if err != nil {
		return nil, err
	}
	s.conn = conn
	return conn, nil
}

func (s *HandlerTestSuite) ListenError() error {
	timeout := time.After(100 * time.Millisecond)
	err := make(chan error)
	go func() {
		var response map[string]interface{}
		s.conn.ReadJSON(&response)
		if response["type"] == game.EventError {
			err <- errors.New(response["message"].(string))
		}
	}()
	select {
	case e := <-err:
		return e
	case <-timeout:
		return nil
	}
}

func (s *HandlerTestSuite) runServer() error {
	request, err := http.ReadRequest(bufio.NewReader(s.server))
	if err != nil {
		return err
	}
	request.AddCookie(&http.Cookie{
		Name:  delivery.SessionCookieName,
		Value: test.Session,
	})
	response := newResponseRecorder(s.server)
	return s.handler.handleGame(s.e.NewContext(request, response))
}

type responseRecorder struct {
	*httptest.ResponseRecorder
	server net.Conn
}

func newResponseRecorder(server net.Conn) *responseRecorder {
	return &responseRecorder{
		ResponseRecorder: httptest.NewRecorder(),
		server:           server,
	}
}

func (r *responseRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	rw := bufio.NewReadWriter(bufio.NewReader(r.server), bufio.NewWriter(r.server))
	return r.server, rw, nil
}
