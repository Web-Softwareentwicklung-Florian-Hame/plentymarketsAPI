package plentymarketsAPI

type ShippingProfileName string

const (
	AmazonPrime        ShippingProfileName = "Amazon Prime"
	DHLPackage         ShippingProfileName = "DHL Paketversand"
	DPLetter           ShippingProfileName = "DP Brief"
	DPDCloudWebService ShippingProfileName = "DPD Cloud Webservice"
	FBA                ShippingProfileName = "FBA"
	Hermes             ShippingProfileName = "Hermes"
	PosLP              ShippingProfileName = "POS-LP"
)

var ShippingProfileIdMap = map[ShippingProfileName]int32{
	AmazonPrime:        12,
	DHLPackage:         8,
	DPLetter:           9,
	DPDCloudWebService: 10,
	FBA:                7,
	Hermes:             11,
	PosLP:              13,
}
