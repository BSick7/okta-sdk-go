package api

type Apps struct {
	executor *ClientExecutor
}

func NewApps(executor *ClientExecutor) *Apps {
	return &Apps{
		executor: executor,
	}
}

type App struct {
}

func (a *Apps) List() ([]*App, error) {
	return []*App{}, nil
}
