package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	"github.com/azcov/evermos-flash-sale/domain/entity"
	"github.com/azcov/evermos-flash-sale/domain/request"
	"github.com/azcov/evermos-flash-sale/domain/response"
	appInit "github.com/azcov/evermos-flash-sale/init"
	"github.com/azcov/evermos-flash-sale/service/order/repository"
	"github.com/stretchr/testify/require"
)

const (
	testProductID = 2
)

func TestCreateOrder(t *testing.T) {

	payloads := []request.CreateOrderRequest{
		{
			UserEmail: "asep@gmail.com",
			Products: []request.CreateOrderProduct{
				{
					ProductID: testProductID,
					Qty:       1,
				},
			},
		},
		{
			UserEmail: "dadang@yahoo.com",
			Products: []request.CreateOrderProduct{
				{
					ProductID: testProductID,
					Qty:       1,
				},
			},
		},
	}

	config, _ := appInit.StartAppInit()
	pgDb, _ := appInit.ConnectToPGServer(config)
	orderRepo := repository.NewRepository(pgDb)

	ctx := context.TODO()
	// make sure product id only have 1 qty
	orderRepo.UpdateProductQty(ctx, &entity.UpdateProductQtyParams{
		ProductID: testProductID,
		Qty:       1,
	})

	var wg sync.WaitGroup
	var numErrRes, numSucRes int
	responses := make(chan response.Base)

	for idx := range payloads {
		wg.Add(1)
		go func(idx int) {
			payload, _ := json.Marshal(payloads[idx])

			client := &http.Client{}
			req, err := http.NewRequest(http.MethodPost, "http://localhost:8081/evermos/v1/checkout", bytes.NewBuffer(payload))
			if err != nil {
				fmt.Print(err.Error())
			}
			req.Header.Add("Accept", "application/json")
			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				fmt.Print(err.Error())
			}
			defer resp.Body.Close()
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Print(err.Error())
			}
			apiRes := response.Base{}
			err = json.Unmarshal(bodyBytes, &apiRes)
			if err != nil {
				fmt.Print(err.Error())
			}
			fmt.Println(apiRes)
			responses <- apiRes
			wg.Done()
		}(idx)
	}

	for range payloads {
		apiRes := <-responses
		if apiRes.StatusCode == http.StatusCreated || apiRes.StatusCode == http.StatusOK {
			numSucRes++
		} else {
			numErrRes++
		}
	}

	close(responses)

	products, _ := orderRepo.GetProductByProductIDForUpdate(ctx, testProductID)

	// products qty must be 0
	require.Equal(t, int32(0), products.Qty)
	// only 1 request for last product's qty to get response succes
	require.Equal(t, 1, numSucRes)
	// number request to get error response must be total request - 1
	require.Equal(t, numErrRes, len(payloads)-1)
}
