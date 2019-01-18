package gmo

import "fmt"

type RegisterMemberOutput struct {
	MemberID string
	ErrCode  string
	ErrInfo  string
}

func (gmo *GMO) RegisterMember(id, name string) (output RegisterMemberOutput, err error) {
	var params = Params{"MemberID": id, "MemberName": name}
	err = gmo.HandleSiteRequest("/payment/SaveMember.idPass", params, &output)
	return
}

type UpdateMemberOutput struct {
	MemberID string
	ErrCode  string
	ErrInfo  string
}

func (gmo *GMO) UpdateMember(id, name string) (output UpdateMemberOutput, err error) {
	var params = Params{"MemberID": id, "MemberName": name}
	err = gmo.HandleSiteRequest("/payment/UpdateMember.idPass", params, &output)
	return
}

type SearchMemberOutput struct {
	MemberID   string
	MemberName string
	DeleteFlag string
	ErrCode    string
	ErrInfo    string
}

func (gmo *GMO) SearchMember(id string) (output SearchMemberOutput, err error) {
	var params = Params{"MemberID": id}
	err = gmo.HandleSiteRequest("/payment/SearchMember.idPass", params, &output)
	return
}

type DeleteMemberOutput struct {
	MemberID string
	ErrCode  string
	ErrInfo  string
}

func (gmo *GMO) DeleteMember(id string) (output DeleteMemberOutput, err error) {
	var params = Params{"MemberID": id}
	err = gmo.HandleSiteRequest("/payment/DeleteMember.idPass", params, &output)
	return
}

type SaveCardOutput struct {
	CardSeq string
	CardNo  string
	Forward string
	ErrCode string
	ErrInfo string
}

func (gmo *GMO) SaveCard(memberID, cardNo, expire, holderName string) (output SaveCardOutput, err error) {
	var params = Params{"MemberID": memberID, "CardNo": cardNo, "Expire": expire, "HolderName": holderName, "SeqMode": "1"}
	err = gmo.HandleSiteRequest("/payment/SaveCard.idPass", params, &output)
	return
}

type SearchCardOutput struct {
	CardSeq     string
	DefaultFlag string
	CardName    string
	CardNo      string
	Expire      string
	HolderName  string
	DeleteFlag  string
	ErrCode     string
	ErrInfo     string
}

func (gmo *GMO) SearchCard(memberID, cardSeq string) (output SearchCardOutput, err error) {
	var params = Params{"MemberID": memberID, "CardSeq": cardSeq, "SeqMode": "1"}
	err = gmo.HandleSiteRequest("/payment/SearchCard.idPass", params, &output)
	return
}

type DeleteCardOutput struct {
	CardSeq string
	ErrCode string
	ErrInfo string
}

func (gmo *GMO) DeleteCard(memberID, cardSeq string) (output DeleteCardOutput, err error) {
	var params = Params{"MemberID": memberID, "CardSeq": cardSeq, "SeqMode": "1"}
	err = gmo.HandleSiteRequest("/payment/DeleteCard.idPass", params, &output)
	return
}

type EntryTranOutput struct {
	AccessID   string
	AccessPass string
	ErrCode    string
	ErrInfo    string
}

const (
	JobCdUnprocessed   = "UNPROCESSED"
	JobCdAuthenticated = "AUTHENTICATED"
	JobCdCheck         = "CHECK"
	JobCdCapture       = "CAPTURE"
	JobCdAuth          = "AUTH"
	JobCdSales         = "SALES"
	JobCdVoid          = "VOID"
	JobCdReturn        = "RETURN"
	JobCdReturnx       = "RETURNX"
	JobCdSauth         = "SAUTH"
)

func (gmo *GMO) EntryTran(orderID, amount, tax string, jobcd string) (output EntryTranOutput, err error) {
	// var params = Params{"OrderID": orderID, "JobCd": "CAPTURE", "Amount": amount, "Tax": tax}
	var params = Params{"OrderID": orderID, "JobCd": jobcd, "Amount": amount, "Tax": tax}
	err = gmo.HandleShopRequest("/payment/EntryTran.idPass", params, &output)
	return
}

type ExecTranOutput struct {
	ACS          string
	OrderID      string
	Forward      string
	Method       string
	PayTimes     string
	Approve      string
	TranID       string
	TranDate     string
	CheckString  string
	ClientField1 string
	ClientField2 string
	ClientField3 string
	ErrCode      string
	ErrInfo      string
}

func (gmo *GMO) ExecTran(accessID, accessPass, orderID, memberID, cardSeq, securityCode string) (output ExecTranOutput, err error) {
	var params = Params{"AccessID": accessID, "AccessPass": accessPass, "OrderID": orderID, "MemberID": memberID, "CardSeq": cardSeq, "SecurityCode": securityCode, "SeqMode": "1", "Method": "1"}
	err = gmo.HandleSiteRequest("/payment/ExecTran.idPass", params, &output)
	return
}

const (
	CardStatusUnprocessed   = "UNPROCESSED"   // unsettled
	CardStatusAuthenticated = "AUTHENTICATED" // unsettled (3D registration complete)
	CardStatusCheck         = "CHECK"         // validity check
	CardStatusCapture       = "CAPTURE"       // instant sales
	CardStatusAuth          = "AUTH"          // authorization
	CardStatusSales         = "SALES"         // sales capture
	CardStatusVoid          = "VOID"          // cancellation
	CardStatusReturn        = "RETURN"        // return goods
	CardStatusReturnx       = "RETURNX"       // return goods next month
	CardStatusSauth         = "SAUTH"         // simple authority

	PayPalStatusUnprocessed = "UNPROCESSED" // unsettled
	PayPalStatusReqsuccess  = "REQSUCCESS"  // request succeed
	PayPalStatusPaysuccess  = "PAYSUCCESS"  // settlement completed
	PayPalStatusPayfail     = "PAYFAIL"     // settlement failure
	PayPalStatusExpired     = "EXPIRED"     // expired
	PayPalStatusCapture     = "CAPTURE"     // instant sales
	PayPalStatusCancel      = "CANCEL"      // cancellation
)

func (gmo *GMO) GetURLToPaypalPage(accessID string) string {
	return fmt.Sprintf(
		"%s/payment/PaypalStart.idPass?ShopID=%s&AccessID=%s",
		gmo.Endpoint,
		gmo.ShopID,
		accessID,
	)
}

type EntryTranPaypalOutput struct {
	OrderID    string
	AccessID   string
	AccessPass string
	ErrCode    string
	ErrInfo    string
}

func (gmo *GMO) EntryTranPaypal(orderID, amount, tax, currency string) (output EntryTranPaypalOutput, err error) {
	var params = Params{"OrderID": orderID, "JobCd": "CAPTURE", "Amount": amount, "Tax": tax, "Currency": currency}
	err = gmo.HandleShopRequest("/payment/EntryTranPaypal.idPass", params, &output)
	return
}

type ExecTranPaypalOutput struct {
	OrderID      string
	ClientField1 string
	ClientField2 string
	ClientField3 string
	ErrCode      string
	ErrInfo      string
}

func (gmo *GMO) ExecTranPaypal(accessID, accessPass, orderID, itemName, redirectURL string) (output ExecTranPaypalOutput, err error) {
	var params = Params{"AccessID": accessID, "AccessPass": accessPass, "OrderID": orderID, "ItemName": itemName, "RedirectURL": redirectURL}
	err = gmo.HandleShopRequest("/payment/ExecTranPaypal.idPass", params, &output)
	return
}

// const (
// 	PaypalStatusPaySuccess = "CAPTURE"
// 	PaypalStatusPayFail    = "PAYFAIL"
// )

type PaypalReturnOutput struct {
	ShopID   string
	OrderID  string
	Status   string
	TranID   string
	TranDate string
	ErrCode  string
	ErrInfo  string
}

type ChangeTranOutput struct {
	AccessID   string
	AccessPass string
	Forward    string
	Approve    string
	TranID     string
	TranDate   string
	ErrCode    string
	ErrInfo    string
}

func (gmo *GMO) ChangeTran(accessID, accessPass, amount, tax, jobcd string) (output ChangeTranOutput, err error) {
	var params = Params{"AccessID": accessID, "AccessPass": accessPass, "JobCd": jobcd, "Amount": amount, "Tax": tax}
	err = gmo.HandleShopRequest("/payment/ChangeTran.idPass", params, &output)
	return
}

type CaptureSalesOutput struct {
	AccessID   string
	AccessPass string
	Forward    string
	Approve    string
	TranID     string
	TranDate   string
	ErrCode    string
	ErrInfo    string
}

func (gmo *GMO) CaptureSales(accessID, accessPass, amount string) (output CaptureSalesOutput, err error) {
	var params = Params{"AccessID": accessID, "AccessPass": accessPass, "JobCd": "SALES", "Amount": amount}
	err = gmo.HandleShopRequest("/payment/AlterTran.idPass", params, &output)
	return
}

type CancelTranOutput struct {
	AccessID   string
	AccessPass string
	Forward    string
	Approve    string
	TranID     string
	TranDate   string
	ErrCode    string
	ErrInfo    string
}

func (gmo *GMO) CancelTran(accessID, accessPass string) (output CancelTranOutput, err error) {
	var params = Params{"AccessID": accessID, "AccessPass": accessPass, "JobCd": "VOID"}
	err = gmo.HandleShopRequest("/payment/AlterTran.idPass", params, &output)
	return
}

type SearchTradeOutput struct {
	OrderID      string
	Status       string
	ProcessDate  string
	JobCd        string
	AccessID     string
	AccessPass   string
	ItemCode     string
	Amount       string
	Tax          string
	SiteID       string
	MemberID     string
	CardNo       string
	Expire       string
	Method       string
	PayTimes     string
	Forward      string
	TranID       string
	Approve      string
	ClientField1 string
	ClientField2 string
	ClientField3 string
	ErrCode      string
	ErrInfo      string
}

func (gmo *GMO) SearchTrade(orderID string) (output SearchTradeOutput, err error) {
	var params = Params{"OrderID": orderID}
	err = gmo.HandleShopRequest("/payment/SearchTrade.idPass", params, &output)
	return
}

type CancelTranPaypalOutput struct {
	OrderID  string
	TranID   string
	TranDate string
	ErrCode  string
	ErrInfo  string
}

func (gmo *GMO) CancelTranPaypal(accessID, accessPass, orderID, amount, tax string) (output CancelTranPaypalOutput, err error) {
	var params = Params{"AccessID": accessID, "AccessPass": accessPass, "OrderID": orderID, "Amount": amount, "Tax": tax}
	err = gmo.HandleShopRequest("/payment/CancelTranPaypal.idPass", params, &output)
	return
}

type SearchTradeMultiOutput struct {
	Status           string
	ProcessDate      string
	JobCd            string
	AccessID         string
	AccessPass       string
	ItemCode         string
	Amount           string
	Tax              string
	SiteID           string
	MemberID         string
	CardNo           string
	Expire           string
	Method           string
	PayTimes         string
	Forward          string
	TranID           string
	Approve          string
	ClientField1     string
	ClientField2     string
	ClientField3     string
	PayType          string
	CvsCode          string
	CvsConfNo        string
	CvsReceiptNo     string
	EdyReceiptNo     string
	EdyOrderNo       string
	SuicaReceiptNo   string
	SuicaOrderNo     string
	CustID           string
	BkCode           string
	ConfNo           string
	EncryptReceiptNo string
	PaymentTerm      string
	ErrCode          string
	ErrInfo          string
}

type PayType string

const (
	PayTypeCredit                  PayType = "0"
	PayTypeMobileSuica                     = "1"
	PayTypeMobileEdy                       = "2"
	PayTypeConvenienceStorePayment         = "3"
	PayTypePayEasy                         = "4"
	PayTypePayPal                          = "5"
	PayTypeIDNet                           = "6"
	PayTypeWebMoney                        = "7"
)

func (gmo *GMO) SearchTradeMulti(orderID string, typ PayType) (output SearchTradeMultiOutput, err error) {
	var params = Params{"OrderID": orderID, "PayType": string(typ)}
	err = gmo.HandleShopRequest("/payment/SearchTradeMulti.idPass", params, &output)
	return
}

type ExecTranDocomoOutput struct {
	AccessID       string
	Token          string
	StartURL       string
	StartLimitDate string
	ErrCode        string
	ErrInfo        string
}

type DocomoStartOutput struct {
	ShopID               string
	OrderID              string
	Status               string
	TranDate             string
	DocomoSettlementCode string
	ErrCode              string
	ErrInfo              string
}

type DocomoSalesOutput struct {
	OrderID string
	Status  string
	Amount  string
	Tax     string
	ErrCode string
	ErrInfo string
}

type ExecTranSoftbankOutput struct {
	AccessID       string
	Token          string
	StartURL       string
	StartLimitDate string
	ErrCode        string
	ErrInfo        string
}

type SoftbankStartOutput struct {
	ShopID       string
	OrderID      string
	Status       string
	TranDate     string
	SbTrackingId string
	ErrCode      string
	ErrInfo      string
}

type SoftbankSalesOutput struct {
	OrderID string
	Status  string
	Amount  string
	Tax     string
	ErrCode string
	ErrInfo string
}

type ExecTranAuOutput struct {
	AccessID       string
	Token          string
	StartURL       string
	StartLimitDate string
	ErrCode        string
	ErrInfo        string
}

type AuStartOutput struct {
	ShopID    string
	OrderID   string
	Status    string
	TranDate  string
	PayInfoNo string
	PayMethod string
	ErrCode   string
	ErrInfo   string
}

type AuSalesOutput struct {
	OrderID string
	Status  string
	Amount  string
	Tax     string
	ErrCode string
	ErrInfo string
}
