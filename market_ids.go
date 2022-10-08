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
	Shopify           MarketName = "Shopify"
	WebApi            MarketName = "WebApi"
	WebShop           MarketName = "WebShop"
	Zalando           MarketName = "Zalando"
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
	Shopify:           13,
	WebApi:            148,
	WebShop:           1,
	Zalando:           118,
}
