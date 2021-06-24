package tx

// . "nc_tools"

type Item struct {
	key    string
	weight int
	value  string
}

type ItemList map[string]*Item

func (i ItemList) add(item *Item) {
	if item == nil {
		return
	}
	if old, ok := i.items[item.key]; ok {
		if old.weight < item.weight {
			i.items[item.key] = item
		}
	} else {
		i.items[item.key] = item
	}
}

func (i ItemList) get() []Item {
	if len(i.items) == 0 {
		return nil
	}
	// sort.Slice(i, less func(i, j int) bool)
}

// item的并发更新： Item{key, weight, value}. add(item *Item)添加item。 get()[]*Item 获取按照weight降序的itemList， key为唯一键值，同一个key只输出weight最大对应的value。 add {1,2,3}, {2, 2, 4}, {1, 1, 2}, {2, 3, 5} -> {1, 2, 3}, {2, 3, 5}
