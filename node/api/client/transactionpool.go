package client

import (
	"encoding/base64"
	"net/url"

	"github.com/NebulousLabs/Sia/encoding"
	"github.com/NebulousLabs/Sia/node/api"
	"github.com/NebulousLabs/Sia/types"
)

// TransactionPoolFeeGet uses the /tpool/fee endpoint to get a fee estimation.
func (c *Client) TransactionPoolFeeGet() (tfg api.TpoolFeeGET, err error) {
	err = c.get("/tpool/fee", &tfg)
	return
}

// TransactionpoolRawPost uses the /tpool/raw endpoint to broadcast a
// transaction by adding it to the transactionpool.
func (c *Client) TransactionpoolRawPost(parents []types.Transaction, txn types.Transaction) (err error) {
	parentsBytes := encoding.Marshal(parents)
	txnBytes := encoding.Marshal(txn)
	values := url.Values{}
	values.Set("parents", base64.StdEncoding.EncodeToString(parentsBytes))
	values.Set("transaction", base64.StdEncoding.EncodeToString(txnBytes))
	err = c.post("/tpool/raw", values.Encode(), nil)
	return
}
