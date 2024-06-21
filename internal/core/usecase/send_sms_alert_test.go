package Usecase

import (
	"testing"

	"github.com/Domains18/food-delivery/internal/core/models"
)

func TestSendSmsAlert(t *testing.T) {
	type args struct {
		order models.Order
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendSmsAlert(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("SendSmsAlert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
