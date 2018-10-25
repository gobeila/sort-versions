package versioning

import (
	"fmt"
	"reflect"
)

const listItemType int = 2

type ListItem struct {
	value    []Item
	itemType int
}

func (o *ListItem) Add(i Item) []Item {
	o.value = append(o.value, i)
	return o.value
}

func (o ListItem) CompareTo(i Item) int {
	//fmt.Println("list_item.CompareTo")

	if i == nil {
		if len(o.value) == 0 {
			return 0
		}
		first := o.value[0]
		return first.CompareTo(nil)
	}

	switch i.GetType() {
	case integerItemType:
		return -1
	case stringItemType:
		return 1
	case listItemType:
		ii := i.(*ListItem)
		left := 0
		right := 0

		//fmt.Printf("o.value --> %v i.value --> %v\n", o.value, ii.value)

		for left < len(o.value) || right < len(ii.value) {
			var l Item
			if left < len(o.value) {
				l = o.value[left]
				//fmt.Printf("l: %v\n", l)
				left++
			}
			var r Item
			if right < len(ii.value) {
				r = ii.value[right]
				//fmt.Printf("r: %v\n", r)
				right++
			}

			//fmt.Println("listItemType")

			result := 0
			if l == nil {
				if r == nil {
					result = 0
					//fmt.Println("result = 0")
				} else {
					result = -1 * r.CompareTo(l)
					//fmt.Printf("%v = -1 * r.CompareTo(l)\n", result)
				}
			} else {
				result = l.CompareTo(r)
				//fmt.Printf("%v = l.CompareTo(r)\n", result)
			}

			//fmt.Printf("%v < %v || %v < %v\n", left, len(o.value), right, len(ii.value))
			//fmt.Printf("result: %v\n", result)

			if result != 0 {
				return result
			}
		}

		return 0
	default:
		panic(fmt.Errorf("invalid item: %T", i))
	}
}

func (o ListItem) GetType() int {
	return listItemType
}

func (o ListItem) IsNull() bool {
	return len(o.value) == 0
}

func (o *ListItem) normalize() {
	for i := len(o.value) - 1; i >= 0; i-- {
		lastItem := o.value[i]

		if lastItem.IsNull() {
			o.value = o.value[:len(o.value)-1]
		} else if reflect.TypeOf(o.value[i]).Name() != "ListItem" {
			break
		}
	}
}

func (o *ListItem) String() string {
	str := ""
	for i := 0; i < len(o.value); i++ {
		if str != "" {
			if o.value[i].GetType() == listItemType {
				str += "-"
			} else {
				str += "."
			}
		}
		str += fmt.Sprintf("%v", o.value[i])
	}
	return str
}

// NewListItem constructor
func NewListItem() *ListItem {
	return &ListItem{
		value:    make([]Item, 0),
		itemType: listItemType,
	}
}
