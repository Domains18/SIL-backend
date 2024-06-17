package functions

import (
	"testing"

	"github.com/Domains18/SIL-backend/internal/core/models"
)


func TestSms(t *testing.T) {
	type args struct {
		order models.Order
	}
	tests := []struct {
		name   string
		args   args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendSMS(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("SendSMS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}