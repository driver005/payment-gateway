package intent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/driver005/database"
	"github.com/driver005/database/clause"
	"github.com/driver005/gateway/internal/intent/methods"
	"github.com/driver005/gateway/internal/intent/next"
	"github.com/google/uuid"
)

func (h *Handler) Migrate() {
	methods.Migrate(h.r)
	next.Migrate(h.r)

	err := h.r.Context().AutoMigrate(
		&PaymentIntent{},
		&PaymentIntentCardProcessing{},
		&PaymentIntentProcessing{},
		&PaymentIntentProcessingCustomerNotification{},
	)
	if err != nil {
		panic(err)
	}
}

func (s *Handler) Retrive(ctx context.Context, Id uuid.UUID) (*PaymentIntent, error) {
	var m PaymentIntent

	if err := s.r.Manager(ctx).Preload(clause.Associations).Where("id = ?", Id).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Handler) List(ctx context.Context, Offset int, Size int) ([]PaymentIntent, *int64, error) {
	var m = make([]PaymentIntent, 0)

	r := s.r.Manager(ctx).Preload(clause.Associations).Offset(Offset).Limit(Size).Order("id").Find(&m)
	if r.Error != nil {
		return nil, nil, r.Error
	}

	return m, &r.RowsAffected, nil
}

func (s *Handler) Create(ctx context.Context, m *PaymentIntent) (*PaymentIntent, error) {
	if err := s.r.Manager(ctx).Session(&database.Session{FullSaveAssociations: true}).Save(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Handler) Update(ctx context.Context, m *PaymentIntent) (*PaymentIntent, error) {
	var o *PaymentIntent

	o, err := s.Retrive(ctx, m.Id)
	if err != nil {
		return nil, err
	}

	m.Id = o.Id
	if err := s.r.Manager(ctx).Model(&o).Updates(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Handler) Upsert(ctx context.Context, m *PaymentIntent) (*PaymentIntent, error) {
	res := r.r.Manager(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(&m)
	return m, res.Error
}

func (s *Handler) Delete(ctx context.Context, Id uuid.UUID) error {
	if err := s.r.Manager(ctx).Where("id = ?", Id).Delete(&PaymentIntent{}).Error; err != nil {
		return err
	}

	return nil
}

func (h *Handler) Unmarshal(data []byte, m *PaymentIntent) error {
	type Alias PaymentIntent

	model := struct {
		*Alias
		AmountDetails           string `json:"amount_details,omitempty"`
		AutomaticPaymentMethods string `json:"automatic_payment_methods,omitempty"`
		Customer                string `json:"customer,omitempty"`
		Invoice                 string `json:"invoice,omitempty"`
		LastPaymentError        string `json:"last_payment_error,omitempty"`
		LatestCharge            string `json:"latest_charge,omitempty"`
		NextAction              string `json:"next_action,omitempty"`
		PaymentMethod           string `json:"payment_method,omitempty"`
		PaymentMethodOptions    string `json:"payment_method_options,omitempty"`
		Processing              string `json:"processing,omitempty"`
		Shipping                string `json:"shipping,omitempty"`
		Review                  string `json:"review,omitempty"`
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, &model); err != nil {
		return err
	}

	if model.PaymentMethod != "" {
		if v, err := uuid.Parse(model.PaymentMethod); err == nil {
			r, err := h.r.PaymentMethod().Retrive(context.Background(), v)
			if err != nil {
				return err
			}
			m.PaymentMethod = r

			fmt.Println(m.PaymentMethod.Id.String())
		}
	}

	return nil
}
