package registry

import (
	"context"
	"os"
	"sync"

	"github.com/driver005/database"
	"github.com/driver005/gateway/billing/credit"
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/billing/invoiceItem"
	"github.com/driver005/gateway/billing/plan"
	"github.com/driver005/gateway/billing/quote"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/billing/subscriptionItem"
	"github.com/driver005/gateway/billing/subscriptionSchedule"
	"github.com/driver005/gateway/billing/usageRecord"
	"github.com/driver005/gateway/checkout"
	db "github.com/driver005/gateway/database"
	"github.com/driver005/gateway/driver"
	"github.com/driver005/gateway/handler"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/dispute"
	"github.com/driver005/gateway/internal/file"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/internal/mandate"
	"github.com/driver005/gateway/internal/payout"
	"github.com/driver005/gateway/internal/refund"
	"github.com/driver005/gateway/internal/setup/attempt"
	setup_intent "github.com/driver005/gateway/internal/setup/intent"
	"github.com/driver005/gateway/link"
	"github.com/driver005/gateway/logger"
	"github.com/driver005/gateway/payment/bank"
	"github.com/driver005/gateway/payment/card"
	"github.com/driver005/gateway/payment/cash"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/payment/source"
	"github.com/driver005/gateway/products/coupon"
	"github.com/driver005/gateway/products/discount"
	"github.com/driver005/gateway/products/item"
	"github.com/driver005/gateway/products/price"
	"github.com/driver005/gateway/products/product"
	"github.com/driver005/gateway/products/promotion"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	drivers = make([]func() Driver, 0)
	dmtx    sync.Mutex

	// ErrNoResponsibleDriverFound is returned when no driver was found for the provided DSN.
	ErrNoResponsibleDriverFound = errors.New("dsn value requested an unknown driver")
)

type Driver interface {
	// CanHandle returns true if the driver is capable of handling the given DSN or false otherwise.
	CanHandle(dsn string) bool

	Manager(ctx context.Context) *database.DB

	Context() *database.DB

	DeleteTabels(db *database.DB) error

	Migrate() bool
}

type Registry interface {
	Driver

	Setup()
	Init(ctx context.Context, migrate bool) error
	WithLogger(l *logger.Logger) Registry
	Logger() *logger.Logger
	Handler() *handler.Handler
	RegisterRoutes(router *fiber.App)
	ClientManager() *db.Handler
	Repository() repository.TransactionRepository
	Driver() *driver.Handler
	Pay(ctx context.Context, i *intent.PaymentIntent) (*intent.PaymentIntent, error)

	//Payment
	PaymentMethod() *method.Handler
	Bank() *bank.Handler
	Card() *card.Handler
	Cash() *cash.Handler
	Source() *source.Handler

	//Internal
	Balance() *balance.Handler
	Charge() *charge.Handler
	Customer() *customer.Handler
	Dispute() *dispute.Handler
	File() *file.Handler
	PaymentIntent() *intent.Handler
	Mandate() *mandate.Handler
	Payout() *payout.Handler
	Refund() *refund.Handler
	SetupAttempt() *attempt.Handler
	SetupIntent() *setup_intent.Handler

	//billing
	Credit() *credit.Handler
	Invoice() *invoice.Handler
	InvoiceItem() *invoiceItem.Handler
	Plan() *plan.Handler
	Quote() *quote.Handler
	Subscription() *subscription.Handler
	SubscriptionItem() *subscriptionItem.Handler
	SubscriptionSchedule() *subscriptionSchedule.Handler
	UsageRecord() *usageRecord.Handler

	//product
	Coupon() *coupon.Handler
	Discount() *discount.Handler
	Price() *price.Handler
	Product() *product.Handler
	Promotion() *promotion.Handler
	LineItem() *item.Handler

	//checkout
	Checkout() *checkout.Handler

	//link
	Link() *link.Handler

	//utils
	Utils() *utils.Handler

	sql.Provider
}

// RegisterDriver registers a driver
func RegisterDriver(d func() Driver) {
	dmtx.Lock()
	drivers = append(drivers, d)
	dmtx.Unlock()
}

// GetDriverFor returns a driver for the given DSN or ErrNoResponsibleDriverFound if no driver was found.
func GetDriverFor(dsn string) (Driver, error) {
	for _, f := range drivers {
		driver := f()
		if driver.CanHandle(dsn) {
			return driver, nil
		}
	}
	return nil, ErrNoResponsibleDriverFound
}

func NewRegistryFromDSN(ctx context.Context, l *logger.Logger, migrate bool) (Registry, error) {
	driver, err := GetDriverFor("postgres://jzyozqtc:Hdh78SBKXkvgs6-Z5fpVQAw_la4Iln__@batyr.db.elephantsql.com/jzyozqt?=columnsWithAlias=true")
	if err != nil {
		return nil, helper.WithStack(err)
	}

	registry, ok := driver.(Registry)
	if !ok {
		return nil, errors.Errorf("driver of type %T does not implement interface Registry", driver)
	}

	registry = registry.WithLogger(l)

	if err := registry.Init(ctx, migrate); err != nil {
		return nil, err
	}

	return registry, nil
}

func CallRegistry(r Registry) {
	w := wow.New(os.Stdout, spin.Get(spin.Earth), "    Initialize tabels + handlers")
	w.Start()

	//Payment
	r.PaymentMethod()
	r.Bank()
	r.Card()
	r.Cash()
	r.Source()

	//Internal
	r.Balance()
	r.Charge()
	r.Customer()
	r.Dispute()
	r.File()
	r.PaymentIntent()
	r.Mandate()
	r.Payout()
	r.Refund()
	r.SetupAttempt()
	r.SetupIntent()

	//utils
	r.Utils()

	//billing
	r.Credit()
	r.Invoice()
	r.InvoiceItem()
	r.Plan()
	r.Quote()
	r.Subscription()
	r.SubscriptionItem()
	r.SubscriptionSchedule()
	r.UsageRecord()

	//product
	r.Coupon()
	r.Discount()
	r.Price()
	r.Product()
	r.Promotion()
	r.LineItem()

	//checkout
	r.Checkout()

	//link
	r.Link()

	w.PersistWith(spin.Spinner{Frames: pq.StringArray{"✅"}}, " Finished")
}
