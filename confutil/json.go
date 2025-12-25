package confutil

import (
	"encoding/json"
)

type JsonConfigBase struct{}

// ToJson 将配置转换为JSON格式
func (jc *JsonConfigBase) ToJson(c interface{}) (string, error) {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// LoadJson 从JSON文件加载配置
func (jc *JsonConfigBase) LoadJson(c interface{}, data string) error {
	if err := json.Unmarshal([]byte(data), c); err != nil {
		return err
	}
	return nil
}
