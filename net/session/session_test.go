package session

import "testing"

func TestNewSession(t *testing.T) {
	s := New(nil)
	if s.ID() < 1 {
		t.Fail()
	}
}

func TestSession_Bind(t *testing.T) {
	s := New(nil)
	userIDs := []int64{100, 1000, 10000000}
	for i, userId := range userIDs {
		s.Bind(userId)
		if s.UserID() != userIDs[i] {
			t.Fail()
		}
	}
}
func TestSession_HasKey(t *testing.T) {
	s := New(nil)
	key := "hello"
	value := "world"
	s.Set(key, value)
	if !s.HasKey(key) {
		t.Fail()
	}
}

func TestSession_String(t *testing.T) {
	s := New(nil)
	key := "hello"
	value := "world"
	s.Set(key, value)
	getValue := s.String(key)
	if value != getValue {
		t.Fail()
	}
}

func TestSession_Int64(t *testing.T) {
	s := New(nil)
	key := "testkey"
	value := int64(444454)
	s.Set(key, value)
	if value != s.Int64(key) {
		t.Fail()
	}
}
