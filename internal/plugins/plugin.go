package plugins

// Plugin .
type Plugin interface {
	OnStart() error
	OnStop() error
}
