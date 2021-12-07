package celeritas

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const version = "1.0.0"

type Celeritas struct {
	AppName string
	Debug bool
	Version string
	ErrorLog *log.Logger
	InfoLog *log.Logger
	RootPath string
	config config
}

// holds configuration for this package
type config struct {
	port string
	renderer string
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		// list of folders to create
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := c.Init(pathConfig)

	if err != nil {
		return err
	}

	err = c.checkDotEnv(rootPath)

	if err != nil {
		return err
	}

	// read .env file values into the environment of our app
	err = godotenv.Load(rootPath + "/.env")

	if err != nil {
		return nil
	}

	// create loggers
	infoLog, errorLog := c.startLoggers()

	c.InfoLog = infoLog
	c.ErrorLog = errorLog
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
	c.RootPath = rootPath

	c.config = config{
		port: os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folder if it doesn't exist
		// root contains path to current directory
		err := c.CreateDirIfNotExist(root + "/" + path)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Celeritas) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) startLoggers() (*log.Logger, *log.Logger)  {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
	// "short file" gives info about where the error occurred
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)

	return infoLog, errorLog
}