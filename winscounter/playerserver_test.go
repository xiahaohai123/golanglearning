package poker_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	poker "summersea.top/golanglearning/winscounter"
	"testing"
)

const jsonContentType = "application/json"

func TestGetScore(t *testing.T) {
	store := poker.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := poker.NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStringEquals(t, response.Body.String(), "20")
		assertIntEquals(t, response.Code, http.StatusOK)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStringEquals(t, response.Body.String(), "10")
		assertIntEquals(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertIntEquals(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := poker.StubPlayerStore{
		Scores:   map[string]int{},
		WinCalls: []string{},
	}
	server := poker.NewPlayerServer(&store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertIntEquals(t, response.Code, http.StatusAccepted)

		assertPlayerWin(t, &store, player)
	})
}

// 集成测试
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	file, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	store, err := poker.NewFileSystemPlayerStore(file)
	assertNoError(t, err)
	server := poker.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertIntEquals(t, response.Code, http.StatusOK)

		assertStringEquals(t, response.Body.String(), "3")
	})

	t.Run("get leagues", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		assertIntEquals(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		got := getLeagueFromResponse(t, response.Body)
		wantedLeague := poker.League{
			{"Pepper", 3},
		}
		assertPlayerSliceEquals(t, got, wantedLeague)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := poker.League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := poker.StubPlayerStore{League: wantedLeague}
		server := poker.NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertIntEquals(t, response.Code, http.StatusOK)
		got := getLeagueFromResponse(t, response.Body)
		assertPlayerSliceEquals(t, got, wantedLeague)

		assertContentType(t, response, jsonContentType)
	})
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not hava content-type of '%s', got %v", want, response.Header())
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) poker.League {
	t.Helper()
	var got poker.League

	err := json.NewDecoder(body).Decode(&got)
	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}
	return got
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertStringEquals(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertIntEquals(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertPlayerSliceEquals(t *testing.T, got, wantedLeague poker.League) {
	t.Helper()
	if !reflect.DeepEqual(got, wantedLeague) {
		t.Errorf("got %v but want %v", got, wantedLeague)
	}
}

func assertPlayerWin(t *testing.T, store *poker.StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin but want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner,got '%s', want '%s'", store.WinCalls[0], winner)
	}
}
