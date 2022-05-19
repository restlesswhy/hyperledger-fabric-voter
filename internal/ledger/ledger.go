package ledger

import (
	"context"
	"encoding/json"
	"errors"
	"fabric-voter/internal"
	"fabric-voter/internal/models"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	gwproto "github.com/hyperledger/fabric-protos-go/gateway"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

type ledger struct {
	client *client.Contract
}

func NewLedger(client *client.Contract) internal.Ledger {
	return &ledger{
		client: client,
	}
}

// Создать голосование
func (l *ledger) CreateThread(params *models.ThreadParams) error {

	args := make([]string, 0)
	args = append(args, params.ID, params.Category, params.Theme, params.Description)
	args = append(args, params.Options...)

	_, err := l.client.Submit("CreateThread", client.WithArguments(args...))
	if err != nil {
		return errorHandling(err)
	}

	return nil
}

// Создать голос к определенному голосованию
func (l *ledger) CreateVote(threadID string, userID string) (string, error) {

	txid, err := l.client.SubmitTransaction("CreateVote", threadID, userID)
	if err != nil {
		return "", errorHandling(err)
	}

	return string(txid), nil
}

// Использовать голос
func (l *ledger) UseVote(vote *models.Vote) error {

	_, err := l.client.SubmitTransaction("UseVote", vote.ThreadID, vote.VoteID, vote.Option)
	if err != nil {
		return errorHandling(err)
	}

	return nil
}

// Завершить голосвание
func (l *ledger) EndThread(threadID string) error {

	_, err := l.client.SubmitTransaction("EndThread", threadID)
	if err != nil {
		return errorHandling(err)
	}

	return nil
}

// Посмотреть текущее состояние голосования
func (l *ledger) GetThread(threadID string) (*models.Thread, error) {

	res, err := l.client.SubmitTransaction("QueryThread", threadID)
	if err != nil {
		return nil, errorHandling(err)
	}

	thread := &models.Thread{}
	err = json.Unmarshal(res, &thread)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

// Submit transaction, passing in the wrong number of arguments ,expected to throw an error containing details of any error responses from the smart contract.
func errorHandling(err error) error {

	switch err := err.(type) {
	case *client.EndorseError:
		logrus.Errorf("endorse error with gRPC status %v: %s", status.Code(err), err)
	case *client.SubmitError:
		logrus.Errorf("submit error with gRPC status %v: %s", status.Code(err), err)
	case *client.CommitStatusError:
		if errors.Is(err, context.DeadlineExceeded) {
			logrus.Errorf("timeout waiting for transaction %s commit status: %s", err.TransactionID, err)
		} else {
			logrus.Errorf("error obtaining commit status with gRPC status %v: %s", status.Code(err), err)
		}
	case *client.CommitError:
		logrus.Errorf("transaction %s failed to commit with status %d: %s", err.TransactionID, int32(err.Code), err)
	}
	/*
	 Any error that originates from a peer or orderer node external to the gateway will have its details
	 embedded within the gRPC status error. The following code shows how to extract that.
	*/
	statusErr := status.Convert(err)
	for _, detail := range statusErr.Details() {
		errDetail := detail.(*gwproto.ErrorDetail)
		logrus.Errorf("error from endpoint: %s, mspId: %s, message: %s", errDetail.Address, errDetail.MspId, errDetail.Message)
	}

	return err
}

// Создать анонимное голосование
func (l *ledger) CreateAnonThread(params *models.ThreadParams) error {

	args := make([]string, 0)
	args = append(args, params.ID, params.Category, params.Theme, params.Description)
	args = append(args, params.Options...)

	_, err := l.client.Submit("CreateAnonThread", client.WithArguments(args...))
	if err != nil {
		return errorHandling(err)
	}

	return nil
}

// Посмотреть текущее состояние анонимного голосования
func (l *ledger) GetAnonThread(threadID string) (*models.AnonThread, error) {

	res, err := l.client.SubmitTransaction("QueryAnonThread", threadID)
	if err != nil {
		return nil, errorHandling(err)
	}

	thread := &models.AnonThread{}
	err = json.Unmarshal(res, &thread)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

// Использовать голос анонимно
func (l *ledger) UseAnonVote(vote *models.AnonVote) error {

	j, _ := json.Marshal(vote)

	_, err := l.client.Submit("UseAnonVote", client.WithTransient(map[string][]byte{
		"option": j,
	}))
	if err != nil {
		return errorHandling(err)
	}

	return nil
}

// Завершить анонимное голосование и подвести итоги
func (l *ledger) EndAnonThread(data *models.EndAnonData) error {

	j, _ := json.Marshal(data)

	_, err := l.client.Submit("EndAnonThread", client.WithTransient(map[string][]byte{
		"option": j,
	}))
	if err != nil {
		return errorHandling(err)
	}

	return nil
}
