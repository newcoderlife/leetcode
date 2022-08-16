package nested_list_weight_sum

type NestedInteger struct{}

func (n NestedInteger) IsInteger() bool           { return true }
func (n NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) SetInteger(value int)     {}
func (n *NestedInteger) Add(elem NestedInteger)   {}
func (n NestedInteger) GetList() []*NestedInteger { return nil }

func dfs(current []*NestedInteger, depth int) int {
	ans := 0
	for _, ni := range current {
		if ni.IsInteger() {
			ans += ni.GetInteger() * depth
		} else {
			ans += dfs(ni.GetList(), depth+1)
		}
	}
	return ans
}

func depthSum(nestedList []*NestedInteger) int {
	return dfs(nestedList, 1)
}
