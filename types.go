package qiwi

type ProfileParams struct {
	AuthInfoEnabled     bool `json:"authInfoEnabled,omitempty" url:"authInfoEnabled,omitempty"`
	ContractInfoEnabled bool `json:"contractInfoEnabled,omitempty" url:"contractInfoEnabled,omitempty"`
	UserInfoEnabled     bool `json:"userInfoEnabled,omitempty" url:"userInfoEnabled,omitempty"`
}

type UserProfile struct {
	AuthInfo     AuthInfo     `json:"authInfo"`
	ContractInfo ContractInfo `json:"contractInfo"`
	UserInfo     UserInfo     `json:"userInfo"`
}

type AuthInfo struct {
	PersonId         int64         `json:"personId"`
	RegistrationDate string        `json:"registrationDate"`
	BoundEmail       string        `json:"boundEmail"`
	IP               string        `json:"ip"`
	LastLoginDate    string        `json:"lastLoginDate"`
	MobilePinInfo    MobilePinInfo `json:"mobilePinInfo"`
	PassInfo         PassInfo      `json:"passInfo"`
}

type MobilePinInfo struct {
	LastMobilePinChange string `json:"lastMobilePinChange"`
	MobilePinUsed       bool   `json:"mobilePinUsed"`
	NextMobilePinChange string `json:"nextMobilePinChange"`
}

type PassInfo struct {
	LastPassChange string `json:"lastPassChange"`
	NextPassChange string `json:"nextPassChange"`
	PasswordUsed   bool   `json:"passwordUsed"`
}

type PinInfo struct {
	PinUsed bool `json:"pinUsed"`
}

type ContractInfo struct {
	Blocked            bool                 `json:"blocked"`
	ContractID         int64                `json:"contractId"`
	CreationDate       string               `json:"creationDate"`
	Features           []interface{}        `json:"features"`
	IdentificationInfo []IdentificationInfo `json:"identificationInfo"`
}

type IdentificationInfo struct {
	BankAlias           string `json:"bankAlias"`
	IdentificationLevel string `json:"identificationLevel"`
}

type UserInfo struct {
	DefaultPayCurrency int    `json:"defaultPayCurrency"`
	DefaultPaySource   int    `json:"defaultPaySource"`
	Email              string `json:"email"`
	FirstTxnID         int64  `json:"firstTxnId"`
	Language           string `json:"language"`
	Operator           string `json:"operator"`
	PhoneHash          string `json:"phoneHash"`
	PromoEnabled       string `json:"promoEnabled"`
}

// all dates should be formatted in RFC3339
type HistoryParams struct {
	Rows        int      `json:"rows,omitempty" url:"rows,omitempty"`
	Operation   string   `json:"operation,omitempty" url:"operation,omitempty"`
	Sources     []string `json:"sources,omitempty" url:"sources,omitempty"`
	StartDate   string   `json:"startDate,omitempty" url:"startDate,omitempty"`
	EndDate     string   `json:"endDate,omitempty" url:"endDate,omitempty"`
	NextTxnDate string   `json:"nextTxnDate,omitempty" url:"nextTxnDate,omitempty"`
	NextTxnId   int64    `json:"nextTxnId,omitempty" url:"nextTxnId,omitempty"`
}

type History struct {
	Data        []Transaction `json:"data"`
	NextTxnId   int64         `json:"nextTxnId"`
	NextTxnDate string        `json:"nextTxnDate"`
}

type Transaction struct {
	TxnId                  int64         `json:"txnId"`
	PersonId               int64         `json:"personId"`
	Date                   string        `json:"date"`
	ErrorCode              int           `json:"errorCode"`
	Error                  string        `json:"error"`
	Status                 string        `json:"status"`
	Type                   string        `json:"type"`
	StatusText             string        `json:"statusText"`
	TrmTxnId               string        `json:"trmTxnId"`
	Account                string        `json:"account"`
	Sum                    SumNumeric    `json:"sum"`
	Commission             SumNumeric    `json:"commission"`
	Total                  SumNumeric    `json:"total"`
	Provider               Provider      `json:"provider"`
	Comment                string        `json:"comment"`
	CurrencyRate           float64       `json:"currencyRate"`
	Extras                 []interface{} `json:"extras"`
	ChequeReady            bool          `json:"chequeReady"`
	BankDocumentAvailable  bool          `json:"bankDocumentAvailable"`
	BankDocumentReady      bool          `json:"bankDocumentReady"`
	RepeatPaymentEnabled   bool          `json:"repeatPaymentEnabled"`
	FavoritePaymentEnabled bool          `json:"favoritePaymentEnabled"`
	RegularPaymentEnabled  bool          `json:"regularPaymentEnabled"`
}

type Sum struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type SumNumeric struct {
	Amount   float64 `json:"amount"`
	Currency int     `json:"currency"`
}

type Provider struct {
	ID          int64  `json:"id"`
	ShortName   string `json:"shortName"`
	LongName    string `json:"longName"`
	LogoUrl     string `json:"logoUrl"`
	Description string `json:"description"`
	Keys        string `json:"keys"`
	SiteUrl     string `json:"siteUrl"`
}

// all dates should be formatted in RFC3339
type PaymentStatisticParams struct {
	StartDate string   `json:"startDate,omitempty" url:"startDate,omitempty"`
	EndDate   string   `json:"endDate,omitempty" url:"endDate,omitempty"`
	Operation string   `json:"operation,omitempty" url:"operation,omitempty"`
	Sources   []string `json:"sources,omitempty" url:"sources,omitempty"`
}

type PaymentStatistic struct {
	IncomingTotal []Sum `json:"incomingTotal"`
	OutgoingTotal []Sum `json:"outgoingTotal"`
}

type Balances struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Alias      string     `json:"alias"`
	FsAlias    string     `json:"fsAlias"`
	Title      string     `json:"title"`
	HasBalance bool       `json:"hasBalance"`
	Currency   int        `json:"currency"`
	Type       Type       `json:"type"`
	Balance    SumNumeric `json:"balance"`
}

type Type struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type StandardRate struct {
	Content Content `json:"content"`
}

type Content struct {
	Terms Terms `json:"terms"`
}

type Terms struct {
	Cashbacks         []interface{}  `json:"cashbacks"`
	Commission        Commission     `json:"commission"`
	Description       string         `json:"description"`
	ID                string         `json:"id"`
	Identification    Identification `json:"identification"`
	Limits            []Limit        `json:"limits"`
	Overpayment       bool           `json:"overpayment"`
	RepeatablePayment bool           `json:"repeatablePayment"`
	Type              string         `json:"type"`
	Underpayment      bool           `json:"underpayment"`
}

type Identification struct {
	Required bool `json:"required"`
}

type Limit struct {
	Currency string  `json:"currency"`
	Max      float64 `json:"max"`
	Min      float64 `json:"min"`
}

type Commission struct {
	Ranges []Range `json:"ranges"`
}

type Range struct {
	Bound float32 `json:"bound"`
	Rate  float32 `json:"rate"`
	Min   float32 `json:"min"`
	Max   float32 `json:"max"`
	Fixed float32 `json:"fixed"`
}

type SpecialRateParams struct {
	Account       string        `json:"account"`
	PaymentMethod PaymentMethod `json:"paymentMethod"`
	PurchaseTotal PurchaseTotal `json:"purchaseTotal"`
}

type PurchaseTotal struct {
	Total Sum `json:"total"`
}

type PaymentMethod struct {
	Type      string `json:"type"`
	AccountID string `json:"accountId"`
}

type SpecialRate struct {
	ProviderId               string `json:"providerId"`
	WithdrawSum              Sum    `json:"withdrawSum"`
	EnrollmentSum            Sum    `json:"enrollmentSum"`
	QwCommission             Sum    `json:"qwCommission"`
	FundingSourceCommission  Sum    `json:"fundingSourceCommission"`
	WithdrawToEnrollmentRate int    `json:"withdrawToEnrollmentRate"`
}

type PaymentParams struct {
	ID            string        `json:"id"`
	Sum           Sum           `json:"sum"`
	Source        string        `json:"source"`
	PaymentMethod PaymentMethod `json:"paymentMethod"`
	Fields        Fields        `json:"fields"`
	Comment       string        `json:"comment"`
}

type Fields struct {
	Account     string `json:"account"`
	AccountType string `json:"account_type"`
	ExpDate     string `json:"exp_date"`
}

type Payment struct {
	ID          string         `json:"id"`
	Terms       int            `json:"terms"`
	Fields      Fields         `json:"fields"`
	Sum         Sum            `json:"sum"`
	Source      string         `json:"source"`
	Comment     string         `json:"comment"`
	Transaction TransactionMin `json:"transaction"`
}

type TransactionMin struct {
	ID    string `json:"id"`
	State State  `json:"state"`
}

type State struct {
	Code string `json:"code"`
}

type DetermineOperatorParams struct {
	Phone string `json:"phone,omitempty" url:"phone,omitempty"`
}

type DeterminedProvider struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

type Code struct {
	Value string `json:"value"`
	Name  string `json:"_name"`
}
