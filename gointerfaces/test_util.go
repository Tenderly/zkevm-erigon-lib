package gointerfaces

//go:generate moq -stub -out ./zkevmsentry/mocks.go ./zkevmsentry SentryServer SentryClient
//go:generate moq -stub -out ./zkevmremote/mocks.go ./zkevmremote KVClient KV_StateChangesClient
