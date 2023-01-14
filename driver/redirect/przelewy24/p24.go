package przelewy24

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sync"
)

// APIVersion Current version of przelewy24 API
const APIVersion = "3.2"

var ipRange = []net.IP{
	net.ParseIP("91.216.191.181"),
	net.ParseIP("91.216.191.185"),
}

type P24 struct {
	MerchantID int
	PosID      int
	CRCKey     string

	SessionID string

	Currency string

	Description string
	Email       string
	Country     string

	ReturnURL string
	StatusURL string

	Sandbox bool

	ShoppingList ShoppingList

	registered bool
	verified   bool
}

// AddProduct adds product
func (p *P24) AddProduct(product Product) {
	p.ShoppingList = append(p.ShoppingList, product)
}

// Amount calculates sum of products
func (p *P24) Amount() int {
	amount := 0
	for _, prod := range p.ShoppingList {
		amount = amount + prod.Price
	}
	return amount
}

func (p *P24) apiCall(endPoint string, values url.Values, sandbox bool) (url.Values, error) {
	var resp *http.Response
	var err error
	if sandbox {
		resp, err = http.PostForm("https://sandbox.przelewy24.pl/"+endPoint, values)
	} else {
		resp, err = http.PostForm("https://secure.przelewy24.pl/"+endPoint, values)
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	m, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Register registers transaction
func (p *P24) Register() (string, error) {
	if p.registered {
		return "", fmt.Errorf("Transaction already registered")
	}
	v := p.Values()

	m, err := p.apiCall("trnRegister", v, p.Sandbox)
	if err != nil {
		return "", err
	}
	token := m.Get("token")
	if token == "" {
		errMsg := m.Get("errorMessage")
		return "", fmt.Errorf(errMsg)
	}

	p.registered = true

	if p.Sandbox {
		return "https://sandbox.przelewy24.pl/trnRequest/" + token, err
	}
	return "https://secure.przelewy24.pl/trnRequest/" + token, err

}

// Signature calculates digital signature for transaction
func (p *P24) Signature() string {
	data := fmt.Sprintf("%s|%d|%d|%s|%s", p.SessionID, p.MerchantID, p.Amount(), p.Currency, p.CRCKey)
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

// Verify verifies transaction
// map[p24_currency:[PLN] p24_method:[25] p24_sign:[a5e954561fb3fa3401c09fc39714a333] p24_session_id:[SESSIONID] p24_merchant_id:[12345] p24_pos_id:[12345] p24_statement:[p24-111-111-111] p24_amount:[123] p24_order_id:[1212345132]]
func (p *P24) Verify(values url.Values) (bool, error) {
	if p.verified {
		return true, fmt.Errorf("Transaction already verified")
	}

	remoteSign := values.Get("p24_sign")
	if remoteSign == "" {
		return false, fmt.Errorf("Empty p24_sign")
	}
	orderID := values.Get("p24_order_id")
	if orderID == "" {
		return false, fmt.Errorf("Empty p24_order_id")
	}

	localSign := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s|%s|%d|%s|%s", p.SessionID, orderID, p.Amount(), p.Currency, p.CRCKey))))
	if localSign != remoteSign {
		return false, fmt.Errorf("Invalid p24_sign %s != %s", localSign, remoteSign)
	}

	v := url.Values{}
	v.Add("p24_merchant_id", fmt.Sprint(p.MerchantID))
	v.Add("p24_pos_id", fmt.Sprint(p.PosID))
	v.Add("p24_session_id", p.SessionID)

	v.Add("p24_amount", fmt.Sprint(p.Amount()))
	v.Add("p24_currency", p.Currency)
	v.Add("p24_order_id", values.Get("p24_order_id"))
	v.Add("p24_sign", localSign)

	m, err := p.apiCall("trnVerify", v, p.Sandbox)
	if err != nil {
		return false, err
	}

	if m.Get("error") != "0" {
		errMsg := m.Get("errorMessage")
		return false, fmt.Errorf(errMsg)
	}

	p.verified = true

	return true, nil
}

// Values exports all values for requests
func (p *P24) Values() url.Values {
	v := url.Values{}

	v.Add("p24_merchant_id", fmt.Sprint(p.MerchantID))
	v.Add("p24_pos_id", fmt.Sprint(p.PosID))
	v.Add("p24_session_id", p.SessionID)

	v.Add("p24_amount", fmt.Sprint(p.Amount()))
	v.Add("p24_currency", p.Currency)

	v.Add("p24_description", p.Description)
	v.Add("p24_email", p.Email)
	v.Add("p24_country", p.Country)

	v.Add("p24_url_return", p.ReturnURL)
	if p.StatusURL != "" {
		v.Add("p24_url_status", p.StatusURL)
	}

	v.Add("p24_api_version", APIVersion)
	v.Add("p24_sign", p.Signature())

	for key, values := range p.ShoppingList.Values() {
		for _, value := range values {
			v.Add(key, value)
		}
	}
	return v
}

type ShoppingList []Product

// Values exports all values for requests
func (l ShoppingList) Values() url.Values {
	v := url.Values{}
	for i, value := range l {
		v.Add(fmt.Sprintf("p24_name_%d", i), value.Name)
		v.Add(fmt.Sprintf("p24_description_%d", i), value.Description)
		v.Add(fmt.Sprintf("p24_quantity_%d", i), fmt.Sprint(value.Quantity))
		v.Add(fmt.Sprintf("p24_price_%d", i), fmt.Sprint(value.Price))
	}
	return v
}

// Product represents single item
type Product struct {
	Name        string
	Description string
	Quantity    int
	Price       int
}

func NewTransaction(MerchantID int, PosID int, CRCKey string, SessionID string, Description string, Currency string, Email string, ReturnURL string, Sandbox ...bool) *P24 {
	return &P24{
		MerchantID:  MerchantID,
		PosID:       PosID,
		CRCKey:      CRCKey,
		SessionID:   SessionID,
		Currency:    Currency,
		Description: Description,

		Email:     Email,
		ReturnURL: ReturnURL,

		Sandbox: Sandbox[0],
	}
}

// Checks IP to filter out the requests outside Przelewy24 API
func CheckIP(ipString string) bool {
	ip := net.ParseIP(ipString)
	if ip.To4() == nil {
		return false
	}

	return bytes.Compare(ip, ipRange[0]) >= 0 && bytes.Compare(ip, ipRange[1]) <= 0
}

type Store interface {
	// Get should return a cached *P24
	Get(SessionID string) *P24

	// Save should persist *P24 to the underlying store implementation
	Save(p *P24) error
}

func NewSimpleStore() *SimpleStore {
	return &SimpleStore{
		m: sync.Mutex{},
		c: make(map[string]*P24),
	}
}

type SimpleStore struct {
	m sync.Mutex
	c map[string]*P24
}

func (s *SimpleStore) Get(SessionID string) *P24 {
	s.m.Lock()
	defer s.m.Unlock()

	return s.c[SessionID]
}
func (s *SimpleStore) Save(p *P24) error {
	s.m.Lock()
	defer s.m.Unlock()

	if _, ok := s.c[p.SessionID]; ok {
		return fmt.Errorf("%s already in store", p.SessionID)
	}
	s.c[p.SessionID] = p
	return nil
}
func (s *SimpleStore) Remove(p *P24) error {
	s.m.Lock()
	defer s.m.Unlock()

	if _, ok := s.c[p.SessionID]; !ok {
		return fmt.Errorf("%s not found in store", p.SessionID)
	}

	delete(s.c, p.SessionID)
	return nil
}

/*
func main() {
	p := NewTransaction(12345, 12345, "CRCKey", "UNIQUE SESSION ID", "Descriptn", "PLN", "p24@example.com", "http://example.com/p24/success")
	p.StatusURL = "http://example.com/api/p24_status"
	p.AddProduct(Product{"Test product", "product description", 1, 123})
	redirect_url, err := p.Register()
	status, err := p.Verify(req.From)
}
*/
