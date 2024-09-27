package cli

type JSONStorage struct {
    FilePath string
}

type CLI struct {
    TaskManager *task.TaskManager
}
