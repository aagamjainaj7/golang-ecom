package user

import (
	"net/http"
	"testing"

	"github.com/aagamjainaj7/ecom/types"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run("should fail if user payload is invalid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/register", payload)
	})
}

type mockUserStore struct {}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m * mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}