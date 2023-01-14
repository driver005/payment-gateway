package handler

import (
	"strconv"

	"github.com/driver005/gateway/gateway/nbxplorer"
	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
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

func (n *Nbxplorer) state(context *fiber.Ctx, name string, defaultState bool) bool {
	if context.Query(name) == "false" {
		return false
	} else if context.Query(name) != "true" {
		return true
	} else {
		return defaultState
	}
}

func (n *Nbxplorer) helper(m int, _ error) int {
	return m
}

// Basic operations

func (n *Nbxplorer) GetStatus(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetStatus()
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) RescanTransactions(context *fiber.Ctx) error {
	var m []nbxplorer.TransactionID

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	if err := client.RescanTransactions(m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusNoContent).JSON(nil)
}

// Derivation

func (n *Nbxplorer) CreateWallet(context *fiber.Ctx) error {
	var m nbxplorer.CreateWalletOptions

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.CreateWallet(m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) TrackDerivationScheme(context *fiber.Ctx) error {
	var m nbxplorer.TrackDerivationSchemeOptions

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	if err := client.TrackDerivationScheme(context.Params("derivationScheme"), &m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusNoContent).JSON(nil)
}

func (n *Nbxplorer) GetDerivationSchemeTransactions(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetDerivationSchemeTransactions(context.Params("derivationScheme"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) GetDerivationSchemeTransaction(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetDerivationSchemeTransaction(context.Params("derivationScheme"), context.Params("txId"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) GetCurrentBalance(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetCurrentBalance(context.Params("derivationScheme"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

type newUnusedAddress struct {
	feature nbxplorer.Feature
	skip    int
	reserve bool
}

func (n *Nbxplorer) NewUnusedAddress(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.NewUnusedAddress(
		context.Params("derivationScheme"),
		nbxplorer.Feature(context.Query("feature")),
		n.helper(strconv.Atoi(context.Query("skip"))),
		n.state(context, "reserve", false),
	)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) GetScriptPubKeyInfos(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetScriptPubKeyInfos(context.Params("derivationScheme"), context.Params("script"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) GetDerivationSchemeUTXOs(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetDerivationSchemeUTXOs(context.Params("derivationScheme"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) ScanUTXOSet(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	if err := client.ScanUTXOSet(
		context.Params("address"),
		n.helper(strconv.Atoi(context.Query("batchSize"))),
		n.helper(strconv.Atoi(context.Query("gapLimit"))),
		n.helper(strconv.Atoi(context.Query("from"))),
	); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusNoContent).JSON(nil)
}

type metadataScheme struct {
	metadata interface{}
}

func (n *Nbxplorer) AttachDerivationSchemeMetadata(context *fiber.Ctx) error {
	var m metadataScheme

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	if err := client.AttachDerivationSchemeMetadata(context.Params("derivationScheme"), context.Params("key"), m.metadata); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusNoContent).JSON(nil)
}

func (n *Nbxplorer) DetachDerivationSchemeMetadata(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	if err := client.DetachDerivationSchemeMetadata(context.Params("derivationScheme"), context.Params("key")); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusNoContent).JSON(nil)
}

func (n *Nbxplorer) GetDerivationSchemeMetadata(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetDerivationSchemeMetadata(context.Params("derivationScheme"), context.Params("key"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) PruneUTXOSet(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.PruneUTXOSet(context.Params("derivationScheme"), n.helper(strconv.Atoi(context.Query("daysToKeep"))))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

// Address

func (n *Nbxplorer) TrackAddress(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	if err := client.TrackAddress(context.Params("address")); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusNoContent).JSON(nil)
}

func (n *Nbxplorer) GetAddressTransactions(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetAddressTransactions(context.Params("address"), n.state(context, "includeTransaction", true))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) GetAddressTransaction(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetAddressTransaction(context.Params("address"), context.Params("txId"), n.state(context, "includeTransaction", true))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) GetAddressUTXOs(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetAddressUTXOs(context.Params("address"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

// Transaction

func (n *Nbxplorer) GetTransaction(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetTransaction(context.Params("txId"), n.state(context, "includeTransaction", true))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

type Broadcast struct {
	tx []byte
}

func (n *Nbxplorer) BroadcastTransaction(context *fiber.Ctx) error {
	var m Broadcast

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	if err := client.BroadcastTransaction(m.tx, n.state(context, "testMempoolAccept", false)); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusNoContent).JSON(nil)
}

// Fees

func (n *Nbxplorer) GetFeeRate(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetFeeRate(context.Params("blockCount"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

// Events

func (n *Nbxplorer) GetEventStream(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetEventStream(
		n.helper(strconv.Atoi(context.Query("lastEventID"))),
		n.state(context, "longPolling", false),
		context.Query("limit"),
	)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

func (n *Nbxplorer) GetRecentEventStream(context *fiber.Ctx) error {
	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.GetRecentEventStream(n.helper(strconv.Atoi(context.Query("limit"))))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

// Psbt

func (n *Nbxplorer) CreatePSBT(context *fiber.Ctx) error {
	var m nbxplorer.PSBT

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.CreatePSBT(context.Params("derivationScheme"), m)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}

type updatePSBT struct {
	psbt                        string
	derivationScheme            *string
	rebaseKeyPaths              []nbxplorer.KeyPath
	alwaysIncludeNonWitnessUTXO bool
}

func (n *Nbxplorer) UpdatePSBT(context *fiber.Ctx) error {
	var m updatePSBT

	if err := context.BodyParser(&m); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	client := nbxplorer.NewClient(n.host, nbxplorer.Chain(context.Params("cryptoCode")))
	response, err := client.UpdatePSBT(m.psbt, m.derivationScheme, m.rebaseKeyPaths, m.alwaysIncludeNonWitnessUTXO)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(helper.WithStack(err))
	}

	return context.Status(fiber.StatusOK).JSON(response)
}
