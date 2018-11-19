package bot

import (
	"../.."
	"github.com/op/go-logging"
	"github.com/pkg/errors"
	"sort"
	"strings"
)


var sdk sdkGO.Sdk

func Start(token string){
	var log = logging.MustGetLogger("main")
	var stopped = false
	sdk = sdkGO.NewDev(token)

	products, err := sdk.Products()
	wallets, err := sdk.Wallets()
	log.Debug(wallets)
	if err.Code > 0   {
		log.Error(err.Msg)
	}

	for stopped != true{
		sort.Slice(wallets, func(i, j int) bool {
			return wallets[i].Value.In > wallets[j].Value.In
		})


	currentWallet := wallets[0]

	currentProduct, e := getCurrentProduct(currentWallet.CurrencyCode, products)
		if e != nil {
			log.Error(e.Error())
			break
		}
	productStat, err := sdk.Stat(currentProduct.Uid)
		if err.Code > 0   {
			log.Error(err.Msg)
			break
		}
		size := currentWallet.Value.In/100*20

		if size < currentProduct.Min {
			size = currentProduct.Min
		}
		if size > currentProduct.Max {
			size = currentProduct.Max
		}

		price := productStat.Ticker.Price + 1 * sdkGO.Random(0,2) * currentProduct.Increment

		side := "buy"
		if currentWallet.Type == "crypto"{
			side = "sell"
		}

		newOrd := sdkGO.NewOrder{
			Product: currentProduct.Uid,
			Type:    "limit",
			Side:    side,
			Price:   price,
			Size:    size,
		}

		_, err = sdk.AddOrder(newOrd)
		if err.Code > 0   {
			log.Error(err.Msg)
		}
	}






}

func getCurrentProduct(currency string, products []sdkGO.Product) (product sdkGO.Product, err error){

	for _, v := range products{

		if strings.Contains(v.Uid, currency){
			product = v
			return
		}
	}
	err = errors.New("product not found")
	return
}