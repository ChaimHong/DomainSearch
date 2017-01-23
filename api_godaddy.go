package ds

import (
	"encoding/json"
	"fmt"
)

type Godaddy struct{}

func (g *Godaddy) GetUrl(v string) string {
	return fmt.Sprintf("https://sg.godaddy.com/zh/domainsapi/v1/search/exact?q=%s&key=dlp_offer_com&pc=91720155&ptl=1", v)
}

type GodaddyRet struct {
	Products []struct {
		PriceInfo struct {
			CurrentPrice float64
		}
	}
	ExactMatchDomain struct {
		IsAvailable bool
		Fqdn        string
	}
}

// {
// "Products":[{"Tld":"com",
// 	"PhaseId":28,
// 	"PhaseCode":"GA",
// 	"TierId":0,
// 	"ProductId":101,
// 	"RenewalProductId":10101,
// 	"HasIcannFee":true,
// 	"PriceInfo":{
// 		"ListPrice":102.14,
// 		"ListPriceDisplay":"¥102.14",
// 		"CurrentPrice":7.0,
// 		"CurrentPriceDisplay":"¥7.00",
// 		"OverridePriceUsd":0,
// 		"VatFees":18,
// 		"RenewalPriceDisplay":"¥102.14",
// 		"IsPromoDiscount":true,
// 		"PromoRegLengthFlag":2},
// 	}
// ],
// "ExactMatchDomain":{"AvailabilityStatus":1000,"SyntaxStatus":1000,"SyntaxMessage":"No errors","IsBackorderable":false,"IsDbsAvailable":false,"IsPurchasable":true,"AuctionId":0,"AuctionTypeId":0,"IsFree":true,"IsAvailable":true,"IsValid":true,"IsNxd":false,"IdnScript":"ENG","Index":0,"Fqdn":"a111dfa1s.com","Extension":"com","NameWithoutExtension":"a111dfa1s","DomainScore":0.0,"TierId":0,"IsPremiumTier":false,"PhaseId":28,"PhaseCode":"GA","Price":0.0,"UsdPrice":0,"ProductId":101,"RenewalProductId":0,"VendorId":0,"CommissionPercent":0.0,"IsIdn":false},"RecommendedDomains":[],"Tlds":["com"],"CurrencyDecimalSeparator":".","IsVatCountry":true}
// }

func (g *Godaddy) ParseBody(body []byte) (ret []Domain) {
	v := new(GodaddyRet)

	json.Unmarshal(body, v)
	if v.ExactMatchDomain.IsAvailable && v.Products[0].PriceInfo.CurrentPrice < 10000 {
		ret = append(ret, Domain{Name: v.ExactMatchDomain.Fqdn, Price: v.Products[0].PriceInfo.CurrentPrice})
	}
	return
}
