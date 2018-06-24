package plugin

import (
	"errors"
	"sync"
)

var pool = make(map[string]Holder)
var mux sync.Mutex

// Register Register Holder with name.
func Register(name string, h Holder) {
	mux.Lock()
	defer mux.Unlock()
	if h == nil {
		panic("Register Holder is nil")
	}
	if _, dup := pool[name]; dup {
		panic("Register called twice for Holder " + name)
	}
	pool[name] = h
}

// Registered Return registered holder names.
func Registered() (names []string) {
	mux.Lock()
	defer mux.Unlock()

	for key := range pool {
		names = append(names, key)
	}
	return names
}

// PickHolder Pick up holder by name.
func PickHolder(name string) Holder {
	mux.Lock()
	defer mux.Unlock()

	for key := range pool {
		if key == name {
			return pool[key]
		}
	}
	return nil
}

const (
	// DefHolder 默认Holder名
	DefHolder = "github"
	// DefRepo 默认仓库名
	DefRepo = "awesome-star"
	// DefBranch 默认分支名
	DefBranch = "master"
)

// Configuration 配置信息
type Configuration struct {
	Repo   string
	Branch string
}

// Authentication 登录认证信息
type Authentication struct {
	Token    string
	Username string
	Password string
}

// User 用户信息
type User struct {
	Login   string `json:"login"`
	Name    string `json:"name"`
	HTMLURL string `json:"html_url"`
	Email   string `json:"email"`
}

// Pagination 分页参数
type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

// StarredRepo 标星的仓库
type StarredRepo struct {
	FullName    string `json:"full_name"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
}

// File 文件
type File struct {
	Type     string
	Encoding string
	Size     int64
	Name     string
	Path     string
	Content  string
	SHA      string
}

var (
	// ErrBadCredentials 无效的认证信息
	ErrBadCredentials = errors.New("bad credentials")
	// ErrFileNotFound 文件不存在
	ErrFileNotFound = errors.New("file not found")
	// ErrCreateFile 创建文件失败
	ErrCreateFile = errors.New("failed to create file")
	// ErrUpdateFile 更新文件失败
	ErrUpdateFile = errors.New("failed to update file")
)

// Holder holder
type Holder interface {
	Whoami() (holder string)
	SetConfiguration(*Configuration)
	Login(auth *Authentication) (*User, error)
	FetchAllStarredRepos() (stars []*StarredRepo, err error)
	FetchStarredRepos(pg *Pagination) (stars []*StarredRepo, err error)
	GetFile(path string) (*File, error)
	CreateFile(path string, content []byte) (*File, error)
	UpdateFile(path, sha string, content []byte) (*File, error)
}
