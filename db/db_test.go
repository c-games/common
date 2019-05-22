package db

import (
	"testing"
	"github.com/golang/mock/gomock"
)


func TestConnect(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
}