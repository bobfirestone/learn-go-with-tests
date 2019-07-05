package store

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// Server umm yeah
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // TODO: log the errorss somewhere
		}
		fmt.Fprint(w, data)
	}
}

// Store interface
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// SpyStore object
type SpyStore struct {
	response string
	t        *testing.T
}

// Fetch implemented for SpyStore
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spystore got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}

		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}
