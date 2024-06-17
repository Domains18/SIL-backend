package functions

import (
	"fmt"
	"os"

	"github.com/Domains18/SIL-backend/internal/core/models"
	"github.com/Domains18/SIL-backend/internal/core/repositories"
	"github.com/Domains18/SIL-backend/pkg/db"
	"github.com/MikeMwita/africastalking-go/pkg/sms"
)

func SendSMS(order models.Order) error {
	// Send SMS
	client := sms.SmsSender{
		ApiKey: os.Getenv("AFRICASTALKING_API_KEY"),
		ApiUser: os.Getenv("AFRICASTALKING_API_USER"),
		Sender: os.Getenv("AFRICASTALKING_SENDER"),
	}
	customerRepo := repositories.NewCustomerRepo(pkg.DB)
	customer, err := customerRepo.GetCustomerByID(order.CustomerID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("Hello %s, your order for %s has been received. You will be notified when it is ready for pickup.", customer.Name, order.Item)
	client.Recipients = []string{customer.Phone}

	client.Message = message

	response, err := client.SendSMS()
	if err != nil {
		return err
	}
	fmt.Sprintf("SMS sent to %s with response %v", customer.Phone, response)

	return nil
}