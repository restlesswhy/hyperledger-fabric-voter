package service

import (
	"fabric-voter/internal"
	"fabric-voter/internal/models"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type service struct {
	repo   internal.Repository
	ledger internal.Ledger
}

func NewService(repo internal.Repository, ledger internal.Ledger) internal.Service {
	return &service{
		repo:   repo,
		ledger: ledger,
	}
}

func (s *service) CreateThread(params *models.ThreadParams) (string, error) {

	var now = time.Now()
	var threadID = fmt.Sprintf("thread%d", now.Unix()*1e3+int64(now.Nanosecond())/1e6)
	params.ID = threadID

	err := s.ledger.CreateThread(params)
	if err != nil {
		return "", errors.Wrap(err, "s.ledger.CreateThread()")
	}

	return threadID, err
}

func (s *service) GetThread(threadID string) (*models.Thread, error) {

	thread, err := s.repo.GetThread(threadID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			thread, err = s.ledger.GetThread(threadID)
			if err != nil {
				return nil, errors.Wrap(err, "s.ledger.GetThread()")
			}
			return thread, nil
		}
		return nil, errors.Wrap(err, "s.repo.GetThread()")
	}

	return thread, nil
}

func (s *service) CreateVote(threadID string, userID string) (*models.Vote, error) {

	vote, err := s.ledger.CreateVote(threadID, userID)
	if err != nil {
		return nil, errors.Wrap(err, "s.ledger.CreateVote()")
	}

	return vote, nil
}

func (s *service) UseVote(vote *models.Vote) error {

	err := s.ledger.UseVote(vote)
	if err != nil {
		return errors.Wrap(err, "s.ledger.UseVote()")
	}

	return err
}

func (s *service) EndThread(threadID string) error {

	err := s.ledger.EndThread(threadID)
	if err != nil {
		return errors.Wrap(err, "s.ledger.EndThread()")
	}

	return nil
}

func (s *service) CreateAnonThread(params *models.ThreadParams) (string, error) {

	var now = time.Now()
	var threadID = fmt.Sprintf("thread%d", now.Unix()*1e3+int64(now.Nanosecond())/1e6)
	params.ID = threadID

	err := s.ledger.CreateAnonThread(params)
	if err != nil {
		return "", errors.Wrap(err, "s.ledger.CreateThread()")
	}

	return threadID, nil
}

func (s *service) GetAnonThread(threadID string) (*models.AnonThread, error) {

	// thread, err := s.ledger.GetAnonThread(threadID)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "s.ledger.GetThread()")
	// }

	thread, err := s.repo.GetAnonThread(threadID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			thread, err = s.ledger.GetAnonThread(threadID)
			if err != nil {
				return nil, errors.Wrap(err, "s.ledger.GetThread()")
			}
			return thread, nil
		}
		return nil, errors.Wrap(err, "s.repo.GetThread()")
	}

	return thread, nil
}

func (s *service) UseAnonVote(vote *models.AnonVote) error {

	err := s.ledger.UseAnonVote(vote)
	if err != nil {
		return errors.Wrap(err, "s.ledger.UseVote()")
	}

	return nil
}

func (s *service) EndAnonThread(data *models.EndAnonData) error {

	err := s.ledger.EndAnonThread(data)
	if err != nil {
		return errors.Wrap(err, "s.ledger.EndThread()")
	}

	return nil
}
