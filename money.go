package money

// Currency describes the minor-unit precision of an ISO 4217 currency.
// Minor is the number of digits after the decimal point used by that
// currency's minor unit: 2 for USD, 0 for JPY, 3 for KWD.
type Currency struct {
	Code  string
	Minor int
}

func (c Currency) String() string {
	return c.Code
}

// ISO 4217 currencies, covering the common case (2 decimal places) and
// the currencies whose minor-unit precision breaks code written only
// against USD: zero decimal places (JPY and the other currencies below
// that have no minor unit in circulation) and three decimal places
// (BHD, IQD, JOD, KWD, LYD, OMR, TND).
var (
	USD = Currency{Code: "USD", Minor: 2}
	EUR = Currency{Code: "EUR", Minor: 2}
	GBP = Currency{Code: "GBP", Minor: 2}
	JPY = Currency{Code: "JPY", Minor: 0}
	KWD = Currency{Code: "KWD", Minor: 3}
	BHD = Currency{Code: "BHD", Minor: 3}

	AED = Currency{Code: "AED", Minor: 2}
	ARS = Currency{Code: "ARS", Minor: 2}
	AUD = Currency{Code: "AUD", Minor: 2}
	BDT = Currency{Code: "BDT", Minor: 2}
	BGN = Currency{Code: "BGN", Minor: 2}
	BIF = Currency{Code: "BIF", Minor: 0}
	BRL = Currency{Code: "BRL", Minor: 2}
	BYN = Currency{Code: "BYN", Minor: 2}
	CAD = Currency{Code: "CAD", Minor: 2}
	CHF = Currency{Code: "CHF", Minor: 2}
	CLP = Currency{Code: "CLP", Minor: 0}
	CNY = Currency{Code: "CNY", Minor: 2}
	COP = Currency{Code: "COP", Minor: 2}
	CZK = Currency{Code: "CZK", Minor: 2}
	DJF = Currency{Code: "DJF", Minor: 0}
	DKK = Currency{Code: "DKK", Minor: 2}
	DZD = Currency{Code: "DZD", Minor: 2}
	EGP = Currency{Code: "EGP", Minor: 2}
	GHS = Currency{Code: "GHS", Minor: 2}
	GNF = Currency{Code: "GNF", Minor: 0}
	HKD = Currency{Code: "HKD", Minor: 2}
	HUF = Currency{Code: "HUF", Minor: 2}
	IDR = Currency{Code: "IDR", Minor: 2}
	ILS = Currency{Code: "ILS", Minor: 2}
	INR = Currency{Code: "INR", Minor: 2}
	IQD = Currency{Code: "IQD", Minor: 3}
	ISK = Currency{Code: "ISK", Minor: 0}
	JOD = Currency{Code: "JOD", Minor: 3}
	KES = Currency{Code: "KES", Minor: 2}
	KHR = Currency{Code: "KHR", Minor: 2}
	KMF = Currency{Code: "KMF", Minor: 0}
	KRW = Currency{Code: "KRW", Minor: 0}
	KZT = Currency{Code: "KZT", Minor: 2}
	LAK = Currency{Code: "LAK", Minor: 2}
	LKR = Currency{Code: "LKR", Minor: 2}
	LYD = Currency{Code: "LYD", Minor: 3}
	MAD = Currency{Code: "MAD", Minor: 2}
	MMK = Currency{Code: "MMK", Minor: 2}
	MNT = Currency{Code: "MNT", Minor: 2}
	MXN = Currency{Code: "MXN", Minor: 2}
	MYR = Currency{Code: "MYR", Minor: 2}
	NGN = Currency{Code: "NGN", Minor: 2}
	NOK = Currency{Code: "NOK", Minor: 2}
	NPR = Currency{Code: "NPR", Minor: 2}
	NZD = Currency{Code: "NZD", Minor: 2}
	OMR = Currency{Code: "OMR", Minor: 3}
	PEN = Currency{Code: "PEN", Minor: 2}
	PHP = Currency{Code: "PHP", Minor: 2}
	PKR = Currency{Code: "PKR", Minor: 2}
	PLN = Currency{Code: "PLN", Minor: 2}
	PYG = Currency{Code: "PYG", Minor: 0}
	QAR = Currency{Code: "QAR", Minor: 2}
	RON = Currency{Code: "RON", Minor: 2}
	RUB = Currency{Code: "RUB", Minor: 2}
	RWF = Currency{Code: "RWF", Minor: 0}
	SAR = Currency{Code: "SAR", Minor: 2}
	SEK = Currency{Code: "SEK", Minor: 2}
	SGD = Currency{Code: "SGD", Minor: 2}
	THB = Currency{Code: "THB", Minor: 2}
	TND = Currency{Code: "TND", Minor: 3}
	TRY = Currency{Code: "TRY", Minor: 2}
	TWD = Currency{Code: "TWD", Minor: 2}
	UAH = Currency{Code: "UAH", Minor: 2}
	UGX = Currency{Code: "UGX", Minor: 0}
	UZS = Currency{Code: "UZS", Minor: 2}
	VND = Currency{Code: "VND", Minor: 0}
	VUV = Currency{Code: "VUV", Minor: 0}
	XAF = Currency{Code: "XAF", Minor: 0}
	XOF = Currency{Code: "XOF", Minor: 0}
	XPF = Currency{Code: "XPF", Minor: 0}
	ZAR = Currency{Code: "ZAR", Minor: 2}
)

// byCode looks up one of the currencies above by its ISO 4217 code, so
// that UnmarshalJSON can turn a code back into a Currency with the
// right minor-unit precision.
var byCode = map[string]Currency{
	USD.Code: USD,
	EUR.Code: EUR,
	GBP.Code: GBP,
	JPY.Code: JPY,
	KWD.Code: KWD,
	BHD.Code: BHD,
	AED.Code: AED,
	ARS.Code: ARS,
	AUD.Code: AUD,
	BDT.Code: BDT,
	BGN.Code: BGN,
	BIF.Code: BIF,
	BRL.Code: BRL,
	BYN.Code: BYN,
	CAD.Code: CAD,
	CHF.Code: CHF,
	CLP.Code: CLP,
	CNY.Code: CNY,
	COP.Code: COP,
	CZK.Code: CZK,
	DJF.Code: DJF,
	DKK.Code: DKK,
	DZD.Code: DZD,
	EGP.Code: EGP,
	GHS.Code: GHS,
	GNF.Code: GNF,
	HKD.Code: HKD,
	HUF.Code: HUF,
	IDR.Code: IDR,
	ILS.Code: ILS,
	INR.Code: INR,
	IQD.Code: IQD,
	ISK.Code: ISK,
	JOD.Code: JOD,
	KES.Code: KES,
	KHR.Code: KHR,
	KMF.Code: KMF,
	KRW.Code: KRW,
	KZT.Code: KZT,
	LAK.Code: LAK,
	LKR.Code: LKR,
	LYD.Code: LYD,
	MAD.Code: MAD,
	MMK.Code: MMK,
	MNT.Code: MNT,
	MXN.Code: MXN,
	MYR.Code: MYR,
	NGN.Code: NGN,
	NOK.Code: NOK,
	NPR.Code: NPR,
	NZD.Code: NZD,
	OMR.Code: OMR,
	PEN.Code: PEN,
	PHP.Code: PHP,
	PKR.Code: PKR,
	PLN.Code: PLN,
	PYG.Code: PYG,
	QAR.Code: QAR,
	RON.Code: RON,
	RUB.Code: RUB,
	RWF.Code: RWF,
	SAR.Code: SAR,
	SEK.Code: SEK,
	SGD.Code: SGD,
	THB.Code: THB,
	TND.Code: TND,
	TRY.Code: TRY,
	TWD.Code: TWD,
	UAH.Code: UAH,
	UGX.Code: UGX,
	UZS.Code: UZS,
	VND.Code: VND,
	VUV.Code: VUV,
	XAF.Code: XAF,
	XOF.Code: XOF,
	XPF.Code: XPF,
	ZAR.Code: ZAR,
}
