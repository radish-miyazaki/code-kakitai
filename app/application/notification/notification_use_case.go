package notification

import (
	"context"
	"log"
	"sync"

	userDomain "github.com/radish-miyazaki/code-kakitai/domain/user"
	"go.uber.org/multierr"
)

const emailBatchSize = 1000

type SendSystemMailUseCase struct {
	userRepo     userDomain.UserRepository
	MailNotifier MailNotifier
}

func NewSendSystemMailUseCase(userRepo userDomain.UserRepository, mailNotifier MailNotifier) *SendSystemMailUseCase {
	return &SendSystemMailUseCase{
		userRepo:     userRepo,
		MailNotifier: mailNotifier,
	}
}

func (uc *SendSystemMailUseCase) Run(ctx context.Context) error {
	users, err := uc.userRepo.FindAll(ctx)
	if err != nil {
		return err
	}

	var chunkUsers [][]*userDomain.User
	for i := 0; i < len(users); i += emailBatchSize {
		end := i + emailBatchSize
		if end > len(users) {
			end = len(users)
		}
		chunkUsers = append(chunkUsers, users[i:end])
	}

	var allContents [][]MailContent
	for _, chunkUser := range chunkUsers {
		var contents []MailContent{}
		for _, user := range chunkUser {
			contents = append(contents, MailContent{
				To:      user.Email(),
				Subject: "件名",
				Body:    "本文",
			})
		}

		allContents = append(allContents, contents)
	}

	var errs error
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, v := range allContents {
		wg.Add(1)
		go func(v []MailContent) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Printf("panic: %v", rec)
				}
			}()

			defer wg.Done()
			if err := uc.MailNotifier.Send(ctx, v); err != nil {
				mu.Lock()
				errs = multierr.Append(errs, err)
				mu.Unlock()
				return
			}
		}(v)
	}
	wg.Wait()

	return errs
}
