package erply

type GetVerifyUserResponse struct {
	Status  Status     `json:"status"`
	Records []UserInfo `json:"records"`
}

type GetCustomerResponse struct {
	Status  Status     `json:"status"`
	Records []Customer `json:"records"`
}

type Status struct {
	Request           string  `json:"request"`
	RequestUnixTime   int64   `json:"requestUnixTime"`
	ResponseStatus    string  `json:"responseStatus"`
	ErrorCode         int     `json:"errorCode"`
	GenerationTime    float64 `json:"generationTime"`
	RecordsTotal      int     `json:"recordsTotal"`
	RecordsInResponse int     `json:"recordsInResponse"`
}

type UserInfo struct {
	UserID       string `json:"UserID"`
	EmployeeName string `json:"employeeName"`
	SessionKey   string `json:"sessionKey"`
}

type Customer struct {
	ID                      int           `json:"id" db:"id"`
	CustomerID              int           `json:"customerID" db:"customerid"`
	FullName                string        `json:"fullName" db:"fullname"`
	CompanyName             string        `json:"companyName,omitempty" db:"companyname"`
	CompanyTypeID           int           `json:"companyTypeID" db:"companytypeid"`
	FirstName               string        `json:"firstName" db:"firstname"`
	LastName                string        `json:"lastName" db:"lastname"`
	PersonTitleID           int           `json:"personTitleID" db:"persontitleid"`
	EInvoiceEmail           string        `json:"eInvoiceEmail" db:"einvoiceemail"`
	EInvoiceReference       string        `json:"eInvoiceReference" db:"einvoicereference"`
	EmailEnabled            int           `json:"emailEnabled" db:"emailenabled"`
	EInvoiceEnabled         int           `json:"eInvoiceEnabled" db:"einvoiceenabled"`
	DocuraEDIEnabled        int           `json:"docuraEDIEnabled" db:"docuraedienabled"`
	MailEnabled             int           `json:"mailEnabled" db:"mailenabled"`
	OperatorIdentifier      string        `json:"operatorIdentifier" db:"operatoridentifier"`
	EDI                     string        `json:"EDI" db:"edi"`
	DoNotSell               int           `json:"doNotSell" db:"donotsell"`
	PartialTaxExemption     int           `json:"partialTaxExemption" db:"partialtaxexemption"`
	GroupID                 int           `json:"groupID" db:"groupid"`
	CountryID               string        `json:"countryID" db:"countryid"`
	PayerID                 int           `json:"payerID" db:"payerid"`
	Phone                   string        `json:"phone" db:"phone"`
	Mobile                  string        `json:"mobile" db:"mobile"`
	Email                   string        `json:"email" db:"email"`
	Fax                     string        `json:"fax" db:"fax"`
	Code                    string        `json:"code" db:"code"`
	Birthday                string        `json:"birthday" db:"birthday"`
	IntegrationCode         string        `json:"integrationCode" db:"integrationcode"`
	FlagStatus              int           `json:"flagStatus" db:"flagstatus"`
	ColorStatus             string        `json:"colorStatus" db:"colorstatus"`
	Credit                  int           `json:"credit" db:"credit"`
	SalesBlocked            int           `json:"salesBlocked" db:"salesblocked"`
	ReferenceNumber         string        `json:"referenceNumber" db:"referencenumber"`
	CustomerCardNumber      string        `json:"customerCardNumber" db:"customercardnumber"`
	FactoringContractNumber string        `json:"factoringContractNumber" db:"factoringcontractnumber"`
	GroupName               string        `json:"groupName" db:"groupname"`
	CustomerType            string        `json:"customerType" db:"customertype"`
	Address                 string        `json:"address" db:"address"`
	Street                  string        `json:"street" db:"street"`
	Address2                string        `json:"address2" db:"address2"`
	City                    string        `json:"city" db:"city"`
	PostalCode              string        `json:"postalCode" db:"postalcode"`
	Country                 string        `json:"country" db:"country"`
	State                   string        `json:"state" db:"state"`
	AddressTypeID           int           `json:"addressTypeID" db:"addresstypeid"`
	AddressTypeName         string        `json:"addressTypeName" db:"addresstypename"`
	IsPOSDefaultCustomer    int           `json:"isPOSDefaultCustomer" db:"isposdefaultcustomer"`
	EUCustomerType          string        `json:"euCustomerType" db:"eucustomertype"`
	EDIType                 string        `json:"ediType" db:"editype"`
	LastModifierUsername    string        `json:"lastModifierUsername" db:"lastmodifierusername"`
	LastModifierEmployeeID  int           `json:"lastModifierEmployeeID" db:"lastmodifieremployeeid"`
	TaxExempt               int           `json:"taxExempt" db:"taxexempt"`
	PaysViaFactoring        int           `json:"paysViaFactoring" db:"paysviafactoring"`
	RewardPoints            int           `json:"rewardPoints" db:"rewardpoints"`
	TwitterID               string        `json:"twitterID,omitempty" db:"twitterid"`
	FacebookName            string        `json:"facebookName,omitempty" db:"facebookname"`
	CreditCardLastNumbers   string        `json:"creditCardLastNumbers,omitempty" db:"creditcardlastnumbers"`
	GLN                     string        `json:"GLN,omitempty" db:"gln"`
	DeliveryTypeID          int           `json:"deliveryTypeID,omitempty" db:"deliverytypeid"`
	Image                   string        `json:"image,omitempty" db:"image"`
	CustomerBalanceDisabled int           `json:"customerBalanceDisabled" db:"customerbalancedisabled"`
	RewardPointsDisabled    int           `json:"rewardPointsDisabled" db:"rewardpointsdisabled"`
	PosCouponsDisabled      int           `json:"posCouponsDisabled" db:"poscouponsdisabled"`
	EmailOptOut             int           `json:"emailOptOut" db:"emailoptout"`
	SignUpStoreID           int           `json:"signUpStoreID" db:"signupstoreid"`
	HomeStoreID             int           `json:"homeStoreID" db:"homestoreid"`
	Gender                  string        `json:"gender,omitempty" db:"gender"`
	PeppolID                string        `json:"PeppolID,omitempty" db:"peppolid"`
	ExternalIDs             []interface{} `json:"externalIDs,omitempty" db:"externalids"`
}
