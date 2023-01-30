package registry

import (
	"context"
	"time"

	"github.com/driver005/gateway/billing/credit"
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/billing/invoiceItem"
	"github.com/driver005/gateway/billing/plan"
	"github.com/driver005/gateway/billing/quote"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/billing/subscriptionItem"
	"github.com/driver005/gateway/billing/subscriptionSchedule"
	"github.com/driver005/gateway/billing/usageRecord"
	"github.com/driver005/gateway/database"
	"github.com/driver005/gateway/driver"
	"github.com/driver005/gateway/handler"
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/dispute"
	"github.com/driver005/gateway/internal/file"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/internal/mandate"
	"github.com/driver005/gateway/internal/payout"
	"github.com/driver005/gateway/internal/refund"
	setup_attempt "github.com/driver005/gateway/internal/setup/attempt"
	setup_intent "github.com/driver005/gateway/internal/setup/intent"
	"github.com/driver005/gateway/logger"
	"github.com/driver005/gateway/payment/bank"
	"github.com/driver005/gateway/payment/card"
	"github.com/driver005/gateway/payment/cash"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/payment/source"
	"github.com/driver005/gateway/pdf"
	"github.com/driver005/gateway/products/coupon"
	"github.com/driver005/gateway/products/discount"
	"github.com/driver005/gateway/products/price"
	"github.com/driver005/gateway/products/product"
	"github.com/driver005/gateway/products/promotion"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/service"
	"github.com/driver005/gateway/utils"

	// "github.com/driver005/gateway/service"
	"github.com/driver005/gateway/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Base struct {
	l            *logger.Logger
	al           *logger.Logger
	h            *handler.Handler
	d            *driver.Handler
	s            *service.Handler
	buildVersion string
	buildHash    string
	buildDate    string
	r            Registry
	trc          trace.Tracer
	database     sql.Database
	db           *database.Handler
	rp           repository.TransactionRepository
	//Payment
	method *method.Handler
	bank   *bank.Handler
	card   *card.Handler
	cash   *cash.Handler
	source *source.Handler

	//internal
	balance       *balance.Handler
	charge        *charge.Handler
	customer      *customer.Handler
	dispute       *dispute.Handler
	file          *file.Handler
	intent        *intent.Handler
	mandate       *mandate.Handler
	payout        *payout.Handler
	refund        *refund.Handler
	setup_attempt *setup_attempt.Handler
	setup_intent  *setup_intent.Handler

	//billing
	credit               *credit.Handler
	invoice              *invoice.Handler
	invoiceItem          *invoiceItem.Handler
	plan                 *plan.Handler
	quote                *quote.Handler
	subscription         *subscription.Handler
	subscriptionItem     *subscriptionItem.Handler
	subscriptionSchedule *subscriptionSchedule.Handler
	usageRecord          *usageRecord.Handler

	//products
	coupon    *coupon.Handler
	discount  *discount.Handler
	price     *price.Handler
	product   *product.Handler
	promotion *promotion.Handler

	//utils
	utils *utils.Handler
}

func (m *Base) with(r Registry) *Base {
	m.r = r
	return m
}

func (m *Base) WithBuildInfo(version, hash, date string) Registry {
	m.buildVersion = version
	m.buildHash = hash
	m.buildDate = date
	return m.r
}

func (m *Base) BuildVersion() string {
	return m.buildVersion
}

func (m *Base) BuildDate() string {
	return m.buildDate
}

func (m *Base) BuildHash() string {
	return m.buildHash
}

func (m *Base) WithLogger(l *logger.Logger) Registry {
	m.l = l
	return m.r
}

func (m *Base) Logger() *logger.Logger {
	if m.l == nil {
		m.l = logger.New("ORY Hydra", m.BuildVersion())
	}
	return m.l
}

func (m *Base) Tracer(ctx context.Context) trace.Tracer {
	if m.trc == nil {
		tp := otel.GetTracerProvider()
		m.trc = tp.Tracer("github.com/driver005/gateway", trace.WithInstrumentationVersion(m.BuildVersion()))
	}

	return m.trc
}

func (m *Base) AuditLogger() *logger.Logger {
	if m.al == nil {
		m.al = logger.NewAudit("ORY Hydra", m.BuildVersion())
	}
	return m.al
}

func (m *Base) RegisterRoutes(router *fiber.App) {
	group := router.Group("/api/v1")

	m.Handler().Routes(group)
	m.Driver().SetRoutes(group)

	m.PaymentMethod().SetRoutes(group)
	m.Bank().SetRoutes(group)
	m.Card().SetRoutes(group)
	m.Card().SetRoutes(group)
	m.Source().SetRoutes(group)
	m.Balance().SetRoutes(group)
	m.Charge().SetRoutes(group)
	m.Customer().SetRoutes(group)
	m.Dispute().SetRoutes(group)
	m.File().SetRoutes(group)
	m.PaymentIntent().SetRoutes(group)
	m.Mandate().SetRoutes(group)
	m.Payout().SetRoutes(group)
	m.Refund().SetRoutes(group)
	m.SetupAttempt().SetRoutes(group)
	m.SetupIntent().SetRoutes(group)
	m.Credit().SetRoutes(group)
	m.Invoice().SetRoutes(group)
	m.InvoiceItem().SetRoutes(group)
	m.Plan().SetRoutes(group)
	m.Quote().SetRoutes(group)
	m.Subscription().SetRoutes(group)
	m.SubscriptionItem().SetRoutes(group)
	m.SubscriptionSchedule().SetRoutes(group)
	m.UsageRecord().SetRoutes(group)
	m.Coupon().SetRoutes(group)
	m.Discount().SetRoutes(group)
	m.Price().SetRoutes(group)
	m.Product().SetRoutes(group)
	m.Promotion().SetRoutes(group)
}

func (m *Base) Handler() *handler.Handler {
	if m.h == nil {
		m.h = handler.NewHandler(m.r, "localhost:32838")
	}
	return m.h
}

func (m *Base) Db() *database.Handler {
	if m.db == nil {
		m.db = database.NewHandler(m.r)
	}
	return m.db
}

func (m *Base) Repository() repository.TransactionRepository {
	if m.rp == nil {
		m.rp = repository.NewRepositories(m.Database().DB(context.Background()))
	}
	return m.rp
}

func (m *Base) Driver() *driver.Handler {
	if m.d == nil {
		m.d = driver.NewHandler(m.r)
	}
	return m.d
}

func (m *Base) Pay(ctx context.Context, i *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	r, err := m.Driver().Pay(ctx, i)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (m *Base) Service() *service.Handler {
	if m.s == nil {
		m.s = service.NewHandler()
	}
	return m.s
}

func (m *Base) Database() sql.Database {
	return m.database
}

func (m *Base) Setup() {
	pdf.Initialize()

	public := fiber.New(fiber.Config{
		ServerHeader:   "Fiber",
		AppName:        "Test App v1.0.1",
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		StrictRouting:  true,
		ReadBufferSize: 4096 * 10,
		Views:          html.New("./views", ".html"),
	})

	m.RegisterRoutes(public)

	public.Use(cors.New())
	// public.Use(csrf.New())
	public.Use(fiber_logger.New())
	// public.Use(limiter.New())

	public.Listen("localhost:80")
}
