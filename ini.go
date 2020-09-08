/**
 * Read the configuration file
 *
 * @copyright           (C) 2020  lroyia
 * @lastModify          2020-9-8
 * @website		https://blog.lroyia.top
 *
 */
package goini

import (
	"io/ioutil"
	"strings"
)

/**
 * 项
 */
type Item struct {
	key string
	value string
}

/**
 * 分部
 */
type Section struct {
	key string
	items map[string]*Item
}

/**
 * 配置信息
 */
type Config struct {
	allSections map[string]*Section
	allItems []*Item
	filePath string
}

/**
 * 读取文件
 */
func Read(filePath string) (*Config, error) {
	// 读取目标文件
	bytes, err := ioutil.ReadFile(filePath)

	// 异常处理
	if err != nil {
		return nil, err
	}

	// 读取内容并进行拆分
	content := string(bytes)

	// 换行符判定
	lineSeparator := "\r\n" // window
	if strings.Index(content, "\r") > -1{
		if strings.Index(content, "\n") < 0{
			lineSeparator = "\r" // mac
		}
	}else if strings.Index(content, "\n") > -1{
		lineSeparator = "\n" // linux
	}

	lines := strings.Split(content, lineSeparator)

	var allSections = make(map[string]*Section)
	var allItems []*Item
	var curSection *Section

	for i := range lines {
		each := strings.Trim(lines[i], " ")
		length := strings.Count(each, "") - 1

		if length == 0 {
			continue
		}

		// is section
		if strings.Index(each, "[") == 0 && strings.LastIndex(each, "]") == length-1{
			curSection = new(Section)
			curSection.items = make(map[string]*Item)
			curSection.key = strings.Replace(strings.Replace(each, "[", "", 1), "]", "", 1)
			allSections[curSection.key] = curSection
		}else{
			// may be is item
			if strings.Index(each, "=") < 0{
				continue
			}
			keyValue := strings.SplitN(each, "=", 2)
			key := keyValue[0]
			value := keyValue[1]

			curItem := new(Item)
			curItem.key = key
			curItem.value = value

			if curSection != nil{
				curSection.items[curItem.key] = curItem
			}

			allItems = append(allItems, curItem)
		}
	}

	c := new(Config)
	c.filePath = filePath
	c.allSections = allSections
	c.allItems = allItems

	return c, nil
}

/**
 * getValue By section and item key
 */
func (receiver *Config) GetValueBySection(section string, item string) string {
	curSection := receiver.allSections[section]
	if curSection != nil{
		curItem := curSection.items[item]
		if curItem != nil{
			return curItem.value
		}
	}
	return ""
}

/**
 * getValue by item
 */
func (receiver *Config) GetValueByItem(item string) string{
	for i := range receiver.allItems {
		if strings.EqualFold(receiver.allItems[i].key, item){
			return receiver.allItems[i].value
		}
	}
	return ""
}

/**
 * 获取section下所有items
 */
func (receiver *Config) GetAllItemInSection(section string) map[string]*Item {
	items := receiver.allSections[section]
	if items != nil{
		return items.items
	}
	return nil
}
