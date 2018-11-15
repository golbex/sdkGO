package sdkGO

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"
)

type Sdk struct {
	token  string
	host   string
	client *http.Client
}

func NewProd(token string) Sdk {
	host := "https://api.golbex.com"
	return Sdk{
		token: token,
		host:  host,
		client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func NewDev(token string) Sdk {
	host := "https://api.golbex.co"
	return Sdk{
		token: token,
		host:  host,
		client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func (s Sdk) get(uri string, target interface{}, apiError *Error) (err error) {
	req, _ := http.NewRequest("GET", s.host+uri, nil)
	req.Header.Add("_t", s.token)

	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		json.NewDecoder(res.Body).Decode(apiError)
		return
	}

	err = json.NewDecoder(res.Body).Decode(target)

	return
}

func (s Sdk) post(uri string, body interface{}, target interface{}, apiError *Error) (err error) {
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", s.host+uri, bytes.NewReader(b))
	req.Header.Add("_t", s.token)

	res, err := s.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		json.NewDecoder(res.Body).Decode(apiError)
		return
	}
	err = json.NewDecoder(res.Body).Decode(target)
	return
}

func (s Sdk) Products() (products []Product, err Error){
	s.get(":3003/products", &products, &err)
	return
}
func (s Sdk) Stat(product string) (stat Stat, err Error){
	s.get(":3003/stat/"+product, &stat, &err)
	return
}
func (s Sdk) Candles(product, interval string) (candles []Candle, err Error){
	s.get(":3003/candles/current/"+product+"/"+interval, &candles, &err)
	return
}
func (s Sdk) LastCandle(product, interval string) (candle Candle, err Error){
	s.get(":3003/candles/last/"+product, &product, &err)
	return
}
func (s Sdk) Wallets() (wallets []wallet, err Error){
	s.get(":3005/user/wallets", &wallets, &err)
	return
}
func (s Sdk) OpenOrders() (orders []Order, err Error){
	s.get(":3005/user/orders", &orders, &err)
	return
}
func (s Sdk) CancelOrder(id string) (order Order, err Error){
	s.get(":3005/user/orders/cancel/"+id, &order, &err)
	return
}
func (s Sdk) AddOrder(newOrder NewOrder) (order Order, err Error){
	s.post(":3005/user/orders/add", newOrder, &order, &err)
	return
}