package tag

import (
	"encoding/json"
	"sync"
)

const (
	// Sep 标签层级分隔符
	Sep = "/"
)

var (
	tmux sync.Mutex
	tags = []*Tag{
		Default(),
	}
)

// Tag 标签
type Tag struct {
	// ID      string // 标签ID
	// PID     string // 父ID。若此属性为空，则表示该标签为根标签。
	Path    string // 标签路径，如'/计算机/网络/tcp_ip'
	Name    string // 标签名
	Desc    string
	Builtin bool // 内置标签不可修改与删除
}

// Default 返回默认标签
func Default() *Tag {
	return &Tag{
		// ID:      "764272fd0ae0422d8d1541a55c46ec0c",
		Path:    "/Misc",
		Name:    "Misc",
		Builtin: true,
	}
}

// CreateTag 创建标签
func CreateTag(name string) error {
	return nil
}

// RemoveTag 移除标签
func RemoveTag(name string) error {
	return nil
}

// RenameTag 重命名标签
func RenameTag(name string) error {
	return nil
}

// ListTags 返回当前的标签列表
func ListTags() (items []Tag) {
	tmux.Lock()
	defer tmux.Unlock()

	items = make([]Tag, 0, len(tags))
	for i := range tags {
		items = append(items, *tags[i])
	}
	return items
}

// FindTag 根据标签名称查找标签
func FindTag(name string) *Tag {
	// name = strings.TrimPrefix(strings.TrimSpace(name), Sep)
	// elements := strings.Split(name, Sep)

	// tmux.Lock()
	// defer tmux.Unlock()

	// for i := range elements{

	// }

	// for i := range tags {
	// 	// if tags[]
	// }
	return nil
}

// Rebuild 依据序列化后的JSON数据重建Tag列表
func Rebuild(data []byte) (err error) {
	var ts []*Tag
	if err = json.Unmarshal(data, &ts); err != nil {
		return err
	}
	tmux.Lock()
	tags = ts
	tmux.Unlock()
	return nil
}

// MarshalJSON 将当前标签列表序列化成JSON
func MarshalJSON() []byte {
	tmux.Lock()
	defer tmux.Unlock()

	b, _ := json.MarshalIndent(tags, "  ", "  ")
	return b
}
