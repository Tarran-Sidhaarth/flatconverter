package remote

import "errors"

type RemoteType string

const (
	Github RemoteType = "github"
)

type PullOptions struct {
	Out    string
	Url    string
	Commit *string
}

type Remote interface {
	Pull(opts PullOptions) error
}

func NewRemote(remote RemoteType) (Remote, error) {
	switch remote {
	case Github:
		return &github{}, nil

	default:
		return nil, errors.New("unsupported Remote Type")
	}
}

type Manager struct {
	remote map[RemoteType]Remote
}

func NewManager() *Manager {
	return &Manager{
		remote: make(map[RemoteType]Remote),
	}
}

func (m *Manager) RegisterRemote(remoteType RemoteType, r Remote) {
	m.remote[remoteType] = r
}

func (m *Manager) GetRemote(remoteType RemoteType) Remote {
	return m.remote[remoteType]
}
