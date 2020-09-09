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
	"errors"
	"io/ioutil"
	"strings"
)

/**
 * 项
 */
type Item struct {
	Key   string
	Value string
}

/**
 * 分部
 */
type Section struct {
	Key   string
	Items map[string]*Item
}

/**
 * 配置信息
 */
type Config struct {
	allSections map[string]*Section
	allItems    []*Item
	FilePath    string
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
	if strings.Index(content, "\r") > -1 {
		if strings.Index(content, "\n") < 0 {
			lineSeparator = "\r" // mac
		}
	} else if strings.Index(content, "\n") > -1 {
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

		idx := strings.Index(each, "#")
		if idx == 0 {
			continue
		}

		// is section
		if strings.Index(each, "[") == 0 && strings.LastIndex(each, "]") == length-1 {
			// 注释位置错误
			if idx > 0 && idx != length-1 {
				return nil, errors.New("语法错误：第" + string(rune(i+1)) + "行的注释符位置错误")
			}
			curSection = new(Section)
			curSection.Items = make(map[string]*Item)
			curSection.Key = strings.Replace(strings.Replace(each, "[", "", 1), "]", "", 1)
			allSections[curSection.Key] = curSection
		} else {
			// may be is item
			idxEq := strings.Index(each, "=")
			if idxEq < 0 {
				continue
			}
			if idx > -1 {
				if idxEq > idx {
					return nil, errors.New("语法错误：第" + string(rune(i+1)) + "行的注释符位置错误")
				}
				// 剪掉注释部分
				each = each[0:idx]
			}
			keyValue := strings.SplitN(each, "=", 2)
			key := keyValue[0]
			value := keyValue[1]

			curItem := new(Item)
			curItem.Key = key
			curItem.Value = value

			if curSection != nil {
				curSection.Items[curItem.Key] = curItem
			}

			allItems = append(allItems, curItem)
		}
	}

	c := new(Config)
	c.FilePath = filePath
	c.allSections = allSections
	c.allItems = allItems

	return c, nil
}

/**
 * getValue By section and item key
 */
func (receiver *Config) GetValueBySection(section string, item string) string {
	curSection := receiver.allSections[section]
	if curSection != nil {
		curItem := curSection.Items[item]
		if curItem != nil {
			return curItem.Value
		}
	}
	return ""
}

/**
 * getValue by item
 */
func (receiver *Config) GetValueByItem(item string) string {
	for i := range receiver.allItems {
		if strings.EqualFold(receiver.allItems[i].Key, item) {
			return receiver.allItems[i].Value
		}
	}
	return ""
}

/**
 * 获取section下所有items
 */
func (receiver *Config) GetAllItemInSection(section string) map[string]*Item {
	items := receiver.allSections[section]
	if items != nil {
		return items.Items
	}
	return nil
}
