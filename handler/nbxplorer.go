package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/nbxplorer"
	"github.com/julienschmidt/httprouter"
)

type Nbxplorer struct {
	r    Registry
	host string
}

func NewNbxplorer(r Registry, host string) *Nbxplorer {
	return &Nbxplorer{
		r:    r,
		host: host,
	}
}

func (n *Nbxplorer) query(r *http.Request, name string) string {
	return r.FormValue(name)
}

func (n *Nbxplorer) state(r *http.Request, name string, defaultState bool) bool {
	if n.query(r, name) == "false" {
		return false
	} else if n.query(r, name) != "true" {
		return true
	} else {
		return defaultState
	}
}

func (n *Nbxplorer) helper(m int, _ error) int {
	return m
}

// Basic operations

func (n *Nbxplorer) GetStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetStatus()
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) RescanTransactions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m []nbxplorer.TransactionID

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	if err := client.RescanTransactions(m); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, nil)
}

// Derivation

func (n *Nbxplorer) CreateWallet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m nbxplorer.CreateWalletOptions

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.CreateWallet(m)
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) TrackDerivationScheme(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m nbxplorer.TrackDerivationSchemeOptions

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	if err := client.TrackDerivationScheme(ps.ByName("derivationScheme"), &m); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, nil)
}

func (n *Nbxplorer) GetDerivationSchemeTransactions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetDerivationSchemeTransactions(ps.ByName("derivationScheme"))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) GetDerivationSchemeTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetDerivationSchemeTransaction(ps.ByName("derivationScheme"), ps.ByName("txId"))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) GetCurrentBalance(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetCurrentBalance(ps.ByName("derivationScheme"))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

type newUnusedAddress struct {
	feature nbxplorer.Feature
	skip    int
	reserve bool
}

func (n *Nbxplorer) NewUnusedAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.NewUnusedAddress(
		ps.ByName("derivationScheme"),
		nbxplorer.Feature(n.query(r, "feature")),
		n.helper(strconv.Atoi(n.query(r, "skip"))),
		n.state(r, "reserve", false),
	)
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) GetScriptPubKeyInfos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetScriptPubKeyInfos(ps.ByName("derivationScheme"), ps.ByName("script"))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) GetDerivationSchemeUTXOs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetDerivationSchemeUTXOs(ps.ByName("derivationScheme"))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) ScanUTXOSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	if err := client.ScanUTXOSet(
		ps.ByName("address"),
		n.helper(strconv.Atoi(n.query(r, "batchSize"))),
		n.helper(strconv.Atoi(n.query(r, "gapLimit"))),
		n.helper(strconv.Atoi(n.query(r, "from"))),
	); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, nil)
}

type metadataScheme struct {
	metadata interface{}
}

func (n *Nbxplorer) AttachDerivationSchemeMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m metadataScheme

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	if err := client.AttachDerivationSchemeMetadata(ps.ByName("derivationScheme"), ps.ByName("key"), m.metadata); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, nil)
}

func (n *Nbxplorer) DetachDerivationSchemeMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	if err := client.DetachDerivationSchemeMetadata(ps.ByName("derivationScheme"), ps.ByName("key")); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, nil)
}

func (n *Nbxplorer) GetDerivationSchemeMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetDerivationSchemeMetadata(ps.ByName("derivationScheme"), ps.ByName("key"))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) PruneUTXOSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.PruneUTXOSet(ps.ByName("derivationScheme"), n.helper(strconv.Atoi(n.query(r, "daysToKeep"))))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

// Address

func (n *Nbxplorer) TrackAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	if err := client.TrackAddress(ps.ByName("address")); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, nil)
}

func (n *Nbxplorer) GetAddressTransactions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetAddressTransactions(ps.ByName("address"), n.state(r, "includeTransaction", true))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) GetAddressTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetAddressTransaction(ps.ByName("address"), ps.ByName("txId"), n.state(r, "includeTransaction", true))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) GetAddressUTXOs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetAddressUTXOs(ps.ByName("address"))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

// Transaction

func (n *Nbxplorer) GetTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetTransaction(ps.ByName("txId"), n.state(r, "includeTransaction", true))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

type Broadcast struct {
	tx []byte
}

func (n *Nbxplorer) BroadcastTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m Broadcast

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	if err := client.BroadcastTransaction(m.tx, n.state(r, "testMempoolAccept", false)); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, nil)
}

// Fees

func (n *Nbxplorer) GetFeeRate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetFeeRate(ps.ByName("blockCount"))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

// Events

func (n *Nbxplorer) GetEventStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetEventStream(
		n.helper(strconv.Atoi(n.query(r, "lastEventID"))),
		n.state(r, "longPolling", false),
		n.query(r, "limit"),
	)
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

func (n *Nbxplorer) GetRecentEventStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.GetRecentEventStream(n.helper(strconv.Atoi(n.query(r, "limit"))))
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

// Psbt

func (n *Nbxplorer) CreatePSBT(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m nbxplorer.PSBT

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.CreatePSBT(ps.ByName("derivationScheme"), m)
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}

type updatePSBT struct {
	psbt                        string
	derivationScheme            *string
	rebaseKeyPaths              []nbxplorer.KeyPath
	alwaysIncludeNonWitnessUTXO bool
}

func (n *Nbxplorer) UpdatePSBT(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m updatePSBT

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(ps.ByName("cryptoCode")))
	response, err := client.UpdatePSBT(m.psbt, m.derivationScheme, m.rebaseKeyPaths, m.alwaysIncludeNonWitnessUTXO)
	if err != nil {
		n.r.Writer().WriteError(w, r, helper.WithStack(err))
		return
	}

	n.r.Writer().Write(w, r, response)
}
