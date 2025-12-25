package confutil

import "gopkg.in/yaml.v3"

type YamlConfigBase struct{}

// ToYaml 将配置转换为YAML格式
func (yc *YamlConfigBase) ToYaml(c interface{}) (string, error) {
	yamlData, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(yamlData), nil
}

// LoadYaml 从YAML文件加载配置
func (yc *YamlConfigBase) LoadYaml(c interface{}, data string) error {
	if err := yaml.Unmarshal([]byte(data), c); err != nil {
		return err
	}
	return nil
}
