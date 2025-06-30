package remote

import (
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
)

type github struct{}

func (g *github) Pull(opts PullOptions) error {
	repo, err := git.PlainClone(opts.Out, &git.CloneOptions{
		URL: opts.Url,
	})
	if err != nil {
		return err
	}

	if opts.Commit == nil {
		return nil
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	return worktree.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(*opts.Commit),
	})
}
