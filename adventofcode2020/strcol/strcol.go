package strcol

func Remove(list []string, v string) []string {
	position := IndexOf(list, v)
	if position != -1 {
		return append(list[:position], list[position+1:]...)
	}
	return list
}

func Contains(list []string, item string) bool {
	return IndexOf(list, item) != -1
}

func IndexOf(list []string, item string) int {
	for i, s := range list {
		if s == item {
			return i
		}
	}
	return -1
}

func Eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Values(mapping map[string]string) []string {
	values := make([]string, 0, len(mapping))
	for _, v := range mapping {
		values = append(values, v)
	}
	return values
}

func Keys(mapping map[string]string) []string {
	keys := make([]string, 0, len(mapping))
	for k := range mapping {
		keys = append(keys, k)
	}
	return keys
}
