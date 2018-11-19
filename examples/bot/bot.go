package examples

import (
	"github.com/golbex/SDKGo"
	"github.com/op/go-logging"
)


var sdk sdkGO.Sdk

func Start(token string){
	var log = logging.MustGetLogger("main")
	sdk = sdkGO.NewDev(token)

	var products, err = sdk.Products()
	var wallets, err = sdk.Wallets()

	if err.Code > 0   {
		log.Error(err.Msg)
	}



}