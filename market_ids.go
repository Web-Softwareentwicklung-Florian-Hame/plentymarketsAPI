package plentymarketsAPI

type MarketName string

const (
	Ebay              MarketName = "Ebay"
	Amazon            MarketName = "Amazon"
	AmazonFrance      MarketName = "Amazon France"
	AmazonGermany     MarketName = "Amazon Germany"
	AmazonItaly       MarketName = "Amazon Italy"
	AmazonNetherlands MarketName = "Amazon Netherlands"
	AmazonPoland      MarketName = "Amazon Poland"
	AmazonSpain       MarketName = "Amazon Spain"
	AmazonSweden      MarketName = "Amazon Sweden"
	Froogle           MarketName = "Froogle"
	Kaufland          MarketName = "Kaufland"
	KauflandDe        MarketName = "Kaufland.de"
	Limango           MarketName = "Limango"
	ManuelleEingabe   MarketName = "Manuelle Eingabe"
	MyToys            MarketName = "MyToys"
	OttoMarket        MarketName = "Otto Market"
	POSLP             MarketName = "POS-LP"
	Shopify           MarketName = "Shopify"
	WebApi            MarketName = "WebApi"
	WebShop           MarketName = "WebShop"
	Zalando           MarketName = "Zalando"

	/* Values below are not checked if their name matches to API. Names are from general plenty manual */
	EBayUnitedStates               MarketName = "eBay United States"
	EBayCanadaEnglisch             MarketName = "eBay Canada (Englisch)"
	EBayUK                         MarketName = "eBay UK"
	EBayAustralia                  MarketName = "eBay Australia"
	EBayAustria                    MarketName = "eBay Austria"
	EBayBelgiumFranzösisch         MarketName = "eBay Belgium (Französisch)"
	EBayFrance                     MarketName = "eBay France"
	EBayGermany                    MarketName = "eBay Germany"
	EBayMotors                     MarketName = "eBay Motors"
	EBayItaly                      MarketName = "eBay Italy"
	EBayBelgiumNiederländisch      MarketName = "eBay Belgium (Niederländisch)"
	EBayNetherlands                MarketName = "eBay Netherlands"
	EBaySpain                      MarketName = "eBay Spain"
	EBaySwitzerland                MarketName = "eBay Switzerland"
	EBayHongKong                   MarketName = "eBay Hong Kong"
	EBayIndia                      MarketName = "eBay India"
	EBayIreland                    MarketName = "eBay Ireland"
	EBayMalaysia                   MarketName = "eBay Malaysia"
	EBayCanadaFranzösisch          MarketName = "eBay Canada (Französisch)"
	EBayPhilippines                MarketName = "eBay Philippines"
	EBayPoland                     MarketName = "eBay Poland"
	EBaySingapore                  MarketName = "eBay Singapore"
	Elmr                           MarketName = "Elm@r"
	AmazonUK                       MarketName = "Amazon UK"
	AmazonUSA                      MarketName = "Amazon USA"
	AmazonCanada                   MarketName = "Amazon Canada"
	AmazonMexico                   MarketName = "Amazon Mexico"
	AmazonTurkey                   MarketName = "Amazon Turkey"
	AmazonUnitedArabEmiratesUAE    MarketName = "Amazon United Arab Emirates (U.A.E.)"
	AmazonBelgium                  MarketName = "Amazon Belgium"
	AmazonGermanyBB                MarketName = "Amazon Germany B2B"
	AmazonUKBB                     MarketName = "Amazon UK B2B"
	AmazonFranceBB                 MarketName = "Amazon France B2B"
	AmazonItalyBB                  MarketName = "Amazon Italy B2B"
	AmazonSpainBB                  MarketName = "Amazon Spain B2B"
	Yatego                         MarketName = "Yatego"
	Kelkoo                         MarketName = "Kelkoo"
	GoogleShopping                 MarketName = "Google Shopping"
	Wish                           MarketName = "Wish"
	ricardo                        MarketName = "ricardo"
	KauflandCZ                     MarketName = "Kaufland CZ"
	KauflandSK                     MarketName = "Kaufland SK"
	Kassensystem                   MarketName = "Kassensystem"
	AmazonFBA                      MarketName = "Amazon FBA"
	AmazonFBAGermany               MarketName = "Amazon FBA Germany"
	AmazonFBAUK                    MarketName = "Amazon FBA UK"
	AmazonFBAUSA                   MarketName = "Amazon FBA USA"
	AmazonFBAFrance                MarketName = "Amazon FBA France"
	AmazonFBAItaly                 MarketName = "Amazon FBA Italy"
	AmazonFBASpain                 MarketName = "Amazon FBA Spain"
	AmazonFBACanada                MarketName = "Amazon FBA Canada"
	AmazonFBAMexico                MarketName = "Amazon FBA Mexico"
	AmazonFBANetherlands           MarketName = "Amazon FBA Netherlands"
	AmazonFBAPoland                MarketName = "Amazon FBA Poland"
	AmazonFBASweden                MarketName = "Amazon FBA Sweden"
	AmazonFBATurkey                MarketName = "Amazon FBA Turkey"
	AmazonFBAUnitedArabEmiratesUAE MarketName = "Amazon FBA United Arab Emirates (U.A.E.)"
	AmazonFBAGermanyBB             MarketName = "Amazon FBA Germany B2B"
	AmazonFBAUKBB                  MarketName = "Amazon FBA UK B2B"
	AmazonFBAFranceBB              MarketName = "Amazon FBA France B2B"
	AmazonFBAItalyBB               MarketName = "Amazon FBA Italy B2B"
	AmazonFBASpainBB               MarketName = "Amazon FBA Spain B2B"
	Zentralverkaufde               MarketName = "Zentralverkauf.de"
	Rakutende                      MarketName = "Rakuten.de"
	Rakutencouk                    MarketName = "Rakuten.co.uk"
	NeckermanndeEnterprise         MarketName = "Neckermann.de Enterprise"
	OTTOPreDropshipment            MarketName = "OTTO PreDropshipment"
	Shopgate                       MarketName = "Shopgate"
	Gimahhot                       MarketName = "Gimahhot"
	Shopperella                    MarketName = "Shopperella"
	ShopShare                      MarketName = "ShopShare"
	Quelle                         MarketName = "Quelle"
	Restposten                     MarketName = "Restposten"
	Kauflux                        MarketName = "Kauflux"
	Home                           MarketName = "Home24"
	ZalandoDE                      MarketName = "Zalando DE"
	ZalandoNL                      MarketName = "Zalando NL"
	ZalandoFR                      MarketName = "Zalando FR"
	ZalandoIT                      MarketName = "Zalando IT"
	ZalandoUK                      MarketName = "Zalando UK"
	ZalandoAT                      MarketName = "Zalando AT"
	ZalandoCH                      MarketName = "Zalando CH"
	ZalandoPL                      MarketName = "Zalando PL"
	ZalandoBE                      MarketName = "Zalando BE"
	ZalandoSE                      MarketName = "Zalando SE"
	ZalandoFI                      MarketName = "Zalando FI"
	ZalandoDK                      MarketName = "Zalando DK"
	ZalandoES                      MarketName = "Zalando ES"
	ZalandoNO                      MarketName = "Zalando NO"
	ZalandoCZ                      MarketName = "Zalando CZ"
	ZalandoIE                      MarketName = "Zalando IE"
	ZalandoPT                      MarketName = "Zalando PT"
	ZalandoSK                      MarketName = "Zalando SK"
	ZalandoSI                      MarketName = "Zalando SI"
	ZalandoLT                      MarketName = "Zalando LT"
	ZalandoLV                      MarketName = "Zalando LV"
	ZalandoEE                      MarketName = "Zalando EE"
	ZalandoHR                      MarketName = "Zalando HR"
	ZalandoHU                      MarketName = "Zalando HU"
	ZalandoRO                      MarketName = "Zalando RO"
	NeckermannatEnterprise         MarketName = "Neckermann.at Enterprise"
	NeckermannatCrossDocking       MarketName = "Neckermann.at CrossDocking"
	Idealo                         MarketName = "idealo"
	IdealoDirektkauf               MarketName = "idealo Direktkauf"
	Laary                          MarketName = "Laary"
	SumoNet                        MarketName = "SumoNet"
	Hood                           MarketName = "Hood"
	ParfumDEAL                     MarketName = "ParfumDEAL"
	BeezUP                         MarketName = "BeezUP"
	Tracdelight                    MarketName = "tracdelight"
	Plusde                         MarketName = "Plus.de"
	GartenXXLde                    MarketName = "GartenXXL.de"
	Twenga                         MarketName = "Twenga"
	SporTrade                      MarketName = "sporTrade"
	NewsletterGo                   MarketName = "Newsletter2Go"
	Playcom                        MarketName = "Play.com"
	Grosshandeleu                  MarketName = "grosshandel.eu"
	Hertie                         MarketName = "Hertie"
	CouchCommerce                  MarketName = "CouchCommerce"
	Cdiscountcom                   MarketName = "Cdiscount.com"
	CdiscountcomCLogistique        MarketName = "Cdiscount.com C Logistique"
	Fruugo                         MarketName = "Fruugo"
	Flubit                         MarketName = "Flubit"
	WebAPI                         MarketName = "WebAPI"
	Mercateo                       MarketName = "Mercateo"
	Check                          MarketName = "Check24"
	Bolcom                         MarketName = "bol.com"
	BolcomFBB                      MarketName = "bol.com FBB"
	Criteo                         MarketName = "Criteo"
	Netto                          MarketName = "Netto"
	GartenXXLat                    MarketName = "GartenXXL.at"
	Etsy                           MarketName = "Etsy"
	OTTO                           MarketName = "OTTO"
	KauflandFBK                    MarketName = "Kaufland FBK"
	KauflandDEFBK                  MarketName = "Kaufland DE FBK"
	KauflandCZFBK                  MarketName = "Kaufland CZ FBK"
	KauflandSKFBK                  MarketName = "Kaufland SK FBK"
	Marktkauf                      MarketName = "Marktkauf"
)

var SkuMarketsMapId = map[MarketName]int32{
	Amazon:     4,
	Kaufland:   102,
	OttoMarket: 160,
}

var MarketIdMap = map[MarketName]float32{
	Ebay:              2,
	AmazonFrance:      4.04,
	AmazonGermany:     4.01,
	AmazonItaly:       4.05,
	AmazonNetherlands: 4.09,
	AmazonPoland:      4.11,
	AmazonSpain:       4.06,
	AmazonSweden:      4.12,
	Froogle:           7,
	KauflandDe:        102,
	Limango:           12,
	ManuelleEingabe:   0,
	MyToys:            14,
	OttoMarket:        160.1,
	POSLP:             103.01,
	Shopify:           13,
	WebApi:            148,
	WebShop:           1,
	Zalando:           118,
}

var MarketNameByIdMap = map[float32]MarketName{
	0:      ManuelleEingabe,
	1:      WebShop,
	2:      Ebay,
	2.01:   EBayUnitedStates,
	2.02:   EBayCanadaEnglisch,
	2.03:   EBayUK,
	2.04:   EBayAustralia,
	2.05:   EBayAustria,
	2.06:   EBayBelgiumFranzösisch,
	2.07:   EBayFrance,
	2.08:   EBayGermany,
	2.09:   EBayMotors,
	2.1:    EBayItaly,
	2.11:   EBayBelgiumNiederländisch,
	2.12:   EBayNetherlands,
	2.13:   EBaySpain,
	2.14:   EBaySwitzerland,
	2.15:   EBayHongKong,
	2.16:   EBayIndia,
	2.17:   EBayIreland,
	2.18:   EBayMalaysia,
	2.19:   EBayCanadaFranzösisch,
	2.2:    EBayPhilippines,
	2.21:   EBayPoland,
	2.22:   EBaySingapore,
	3:      Elmr,
	4:      Amazon,
	4.01:   AmazonGermany,
	4.02:   AmazonUK,
	4.03:   AmazonUSA,
	4.04:   AmazonFrance,
	4.05:   AmazonItaly,
	4.06:   AmazonSpain,
	4.07:   AmazonCanada,
	4.08:   AmazonMexico,
	4.09:   AmazonNetherlands,
	4.11:   AmazonPoland,
	4.12:   AmazonSweden,
	4.13:   AmazonTurkey,
	4.14:   AmazonUnitedArabEmiratesUAE,
	4.15:   AmazonBelgium,
	4.21:   AmazonGermanyBB,
	4.22:   AmazonUKBB,
	4.24:   AmazonFranceBB,
	4.25:   AmazonItalyBB,
	4.26:   AmazonSpainBB,
	5:      Yatego,
	6:      Kelkoo,
	7:      GoogleShopping,
	9:      Wish,
	12:     Limango,
	101:    ricardo,
	102:    Kaufland,
	102.01: KauflandDe,
	102.02: KauflandCZ,
	102.03: KauflandSK,
	103:    Kassensystem,
	104:    AmazonFBA,
	104.01: AmazonFBAGermany,
	104.02: AmazonFBAUK,
	104.03: AmazonFBAUSA,
	104.04: AmazonFBAFrance,
	104.05: AmazonFBAItaly,
	104.06: AmazonFBASpain,
	104.07: AmazonFBACanada,
	104.08: AmazonFBAMexico,
	104.09: AmazonFBANetherlands,
	104.11: AmazonFBAPoland,
	104.12: AmazonFBASweden,
	104.13: AmazonFBATurkey,
	104.14: AmazonFBAUnitedArabEmiratesUAE,
	104.21: AmazonFBAGermanyBB,
	104.22: AmazonFBAUKBB,
	104.24: AmazonFBAFranceBB,
	104.25: AmazonFBAItalyBB,
	104.26: AmazonFBASpainBB,
	105:    Zentralverkaufde,
	106:    Rakutende,
	106.02: Rakutencouk,
	107:    NeckermanndeEnterprise,
	108.04: OTTOPreDropshipment,
	109:    Shopgate,
	111:    Gimahhot,
	112:    Shopperella,
	113:    ShopShare,
	114:    Quelle,
	115:    Restposten,
	116:    Kauflux,
	117:    Home,
	118:    Zalando,
	118.01: ZalandoDE,
	118.02: ZalandoNL,
	118.03: ZalandoFR,
	118.04: ZalandoIT,
	118.05: ZalandoUK,
	118.06: ZalandoAT,
	118.07: ZalandoCH,
	118.08: ZalandoPL,
	118.09: ZalandoBE,
	118.1:  ZalandoSE,
	118.11: ZalandoFI,
	118.12: ZalandoDK,
	118.13: ZalandoES,
	118.14: ZalandoNO,
	118.15: ZalandoCZ,
	118.16: ZalandoIE,
	118.17: ZalandoPT,
	118.18: ZalandoSK,
	118.19: ZalandoSI,
	118.2:  ZalandoLT,
	118.21: ZalandoLV,
	118.22: ZalandoEE,
	118.23: ZalandoHR,
	118.24: ZalandoHU,
	118.25: ZalandoRO,
	119:    NeckermannatEnterprise,
	120:    NeckermannatCrossDocking,
	121:    Idealo,
	121.02: IdealoDirektkauf,
	123:    Laary,
	124:    SumoNet,
	125:    Hood,
	126:    ParfumDEAL,
	127:    BeezUP,
	130:    Tracdelight,
	131:    Plusde,
	132:    GartenXXLde,
	133:    Twenga,
	134:    SporTrade,
	135:    NewsletterGo,
	136:    Playcom,
	137:    Grosshandeleu,
	138:    Hertie,
	139:    CouchCommerce,
	143:    Cdiscountcom,
	143.02: CdiscountcomCLogistique,
	145:    Fruugo,
	147:    Flubit,
	148:    WebAPI,
	149:    Mercateo,
	150:    Check,
	152:    Bolcom,
	152.01: BolcomFBB,
	153:    Criteo,
	154:    Netto,
	155:    GartenXXLat,
	156.00: Etsy,
	160:    OTTO,
	160.1:  OttoMarket,
	170:    KauflandFBK,
	170.01: KauflandDEFBK,
	170.02: KauflandCZFBK,
	170.03: KauflandSKFBK,
	171.00: Marktkauf,
}
