package message

import (
	"errors"
	"reflect"
	"testing"
)

type mockRepo struct{}

func (repo *mockRepo) Get(id string) (*Message, error) {
	if id == "error" {
		return nil, errors.New("Mocked error")
	}
	return &Message{Text: "TestMessage"}, nil
}

func NewInmemRepository() Repository {
	return new(mockRepo)
}

func Test_messageSrv_Get(t *testing.T) {
	// creates in memory message repository
	repo := NewInmemRepository()
	// create new service instance in order to test
	srv := NewServiceWith(repo)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *Message
		wantErr bool
	}{
		{
			"test OK",
			args{"myMessageID"},
			&Message{Text: "TestMessage"},
			false,
		},
		{
			"test fail",
			args{"error"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("messageSrv.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageSrv.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
