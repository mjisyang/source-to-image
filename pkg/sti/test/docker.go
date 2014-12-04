package test

import (
	"sync"

	"github.com/openshift/source-to-image/pkg/sti/docker"
)

// FakeDocker provides a fake docker interface
type FakeDocker struct {
	LocalRegistryImage           string
	LocalRegistryResult          bool
	LocalRegistryError           error
	RemoveContainerID            string
	RemoveContainerError         error
	DefaultUrlImage              string
	DefaultUrlResult             string
	DefaultUrlError              error
	RunContainerOpts             docker.RunContainerOptions
	RunContainerError            error
	RunContainerErrorBeforeStart bool
	RunContainerContainerID      string
	RunContainerCmd              []string
	GetImageIDImage              string
	GetImageIDResult             string
	GetImageIDError              error
	CommitContainerOpts          docker.CommitContainerOptions
	CommitContainerResult        string
	CommitContainerError         error
	RemoveImageName              string
	RemoveImageError             error

	mutex sync.Mutex
}

// IsImageInLocalRegistry checks if the image exists in the fake local registry
func (f *FakeDocker) IsImageInLocalRegistry(imageName string) (bool, error) {
	f.LocalRegistryImage = imageName
	return f.LocalRegistryResult, f.LocalRegistryError
}

// RemoveContainer removes a fake Docker container
func (f *FakeDocker) RemoveContainer(id string) error {
	f.RemoveContainerID = id
	return f.RemoveContainerError
}

// GetDefaultScriptsURL returns a default STI scripts URL
func (f *FakeDocker) GetDefaultScriptsURL(image string) (string, error) {
	f.DefaultUrlImage = image
	return f.DefaultUrlResult, f.DefaultUrlError
}

// RunContainer runs a fake Docker container
func (f *FakeDocker) RunContainer(opts docker.RunContainerOptions) error {
	f.RunContainerOpts = opts
	if f.RunContainerErrorBeforeStart {
		return f.RunContainerError
	}
	if opts.OnStart != nil {
		if err := opts.OnStart(); err != nil {
			return err
		}
	}
	if opts.PostExec != nil {
		opts.PostExec.PostExecute(f.RunContainerContainerID, append(f.RunContainerCmd, opts.Command))
	}
	return f.RunContainerError
}

// GetImageID returns a fake Docker image ID
func (f *FakeDocker) GetImageID(image string) (string, error) {
	f.GetImageIDImage = image
	return f.GetImageIDResult, f.GetImageIDError
}

// CommitContainer commits a fake Docker container
func (f *FakeDocker) CommitContainer(opts docker.CommitContainerOptions) (string, error) {
	f.CommitContainerOpts = opts
	return f.CommitContainerResult, f.CommitContainerError
}

// RemoveImage removes a fake Docker image
func (f *FakeDocker) RemoveImage(name string) error {
	f.RemoveImageName = name
	return f.RemoveImageError
}

// PullImage pulls a fake docker image
func (f *FakeDocker) PullImage(imageName string) error {
	return nil
}