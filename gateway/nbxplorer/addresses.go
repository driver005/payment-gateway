package nbxplorer

import (
	"errors"
	"strconv"
)

// TrackAddress
func (c *Client) TrackAddress(address string) error {
	var r ErrorResponse
	resp, err := c.httpClient.R().SetError(&r).Post("/addresses/" + address)
	if err != nil {
		return err
	}

	if len(resp.Body()) == 0 {
		return nil
	}

	return errors.New(r.Message)
}

// GetAddressTransactions
func (c *Client) GetAddressTransactions(address string, hex bool) ([]TransactionVerbose, error) {
	var txs []TransactionVerbose
	var r ErrorResponse

	_, err := c.httpClient.R().SetResult(&txs).SetError(&r).SetQueryParam("includeTransaction", strconv.FormatBool(hex)).Get("/addresses/" + address + "/transactions")
	return txs, err
}

// GetAddressTransaction
func (c *Client) GetAddressTransaction(address string, txid string, hex bool) (TransactionVerbose, error) {
	var tx TransactionVerbose
	var r ErrorResponse

	_, err := c.httpClient.R().SetResult(&tx).SetError(&r).SetQueryParam("includeTransaction", strconv.FormatBool(hex)).Get("/addresses/" + address + "/transactions/" + txid)
	return tx, err
}

type Feature string

const (
	Deposit Feature = "deposit"
	Change  Feature = "change"
	Direct  Feature = "direct"
	Custom  Feature = "custom"
)

// UnusedAddress struct
type UnusedAddress struct {
	TrackedSource      string  `json:"trackedSource"`
	Feature            Feature `json:"feature"`
	DerivationStrategy string  `json:"derivationStrategy"`
	KeyPath            string  `json:"keyPath"`
	ScriptPubKey       string  `json:"scriptPubKey"`
	Address            string  `json:"address"`
	Redeem             string  `json:"redeem"`
}

func (c *Client) NewUnusedAddress(derivationScheme string, feature Feature, skip int, reserve bool) (UnusedAddress, error) {
	var address UnusedAddress
	var r *ErrorResponse

	_, err := c.httpClient.R().
		SetResult(&address).
		SetError(&r).
		SetQueryParam("feature", string(feature)).
		SetQueryParam("skip", strconv.Itoa(skip)).
		SetQueryParam("reserve", strconv.FormatBool(reserve)).
		Get("/derivations/" + derivationScheme + "/addresses/unused")

	if r != nil {
		return address, errors.New(r.Message)
	}

	if address.Address == "" {
		return address, errors.New("derivation scheme not found")
	}

	return address, err
}

// GetAddressUTXOs of derivation scheme
func (c *Client) GetAddressUTXOs(address string) (UTXOInfos, error) {
	var infos UTXOInfos
	var r *ErrorResponse

	_, err := c.httpClient.R().SetResult(&infos).SetError(&r).Get("/addresses/" + address + "/utxos")
	if r != nil {
		return infos, errors.New(r.Message)
	}

	return infos, err
}
