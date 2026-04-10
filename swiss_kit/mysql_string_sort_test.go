package swiss_kit

import (
	"slices"
	"testing"
)

func TestMysqlGbkComparePinyinSort(t *testing.T) {
	input := []string{
		"途牛机票预定", "汽车站_phrase", "神马推广计划", "中国", "骑马-草原",
		"中国", "汽车站_exact", "祈祷", "天才", "天涯海角", "起来", "漆",
		"英语考试", "爱生活", "海南旅游", "旅游", "驴友", "种郭", "黄山旅游",
		"黄山旅游-秋季", "海尔",
	}
	expected := []string{
		"爱生活",
		"海尔",
		"海南旅游",
		"黄山旅游",
		"黄山旅游-秋季",
		"驴友",
		"旅游",
		"漆",
		"祈祷",
		"骑马-草原",
		"起来",
		"汽车站_exact",
		"汽车站_phrase",
		"神马推广计划",
		"天才",
		"天涯海角",
		"途牛机票预定",
		"英语考试",
		"中国",
		"中国",
		"种郭",
	}
	slices.SortFunc(input, MysqlGbkCompare)
	for i := range input {
		if input[i] != expected[i] {
			t.Errorf("position %d: got %q, want %q", i, input[i], expected[i])
		}
	}
}

func TestMysqlGbkCompareCaseInsensitive(t *testing.T) {
	if r := MysqlGbkCompare("abc", "ABC"); r != 0 {
		t.Errorf("abc vs ABC: got %d, want 0", r)
	}
	if r := MysqlGbkCompare("Hello", "hello"); r != 0 {
		t.Errorf("Hello vs hello: got %d, want 0", r)
	}
}

func TestMysqlGbkCompareTrailingSpaces(t *testing.T) {
	if r := MysqlGbkCompare("abc", "abc   "); r != 0 {
		t.Errorf("abc vs 'abc   ': got %d, want 0", r)
	}
	if r := MysqlGbkCompare("abc   ", "abc"); r != 0 {
		t.Errorf("'abc   ' vs abc: got %d, want 0", r)
	}
}

func TestMysqlGbkCompareEmpty(t *testing.T) {
	if r := MysqlGbkCompare("", ""); r != 0 {
		t.Errorf("empty vs empty: got %d, want 0", r)
	}
	if r := MysqlGbkCompare("", " "); r != 0 {
		t.Errorf("empty vs space: got %d, want 0", r)
	}
	if r := MysqlGbkCompare("", "a"); r >= 0 {
		t.Errorf("empty vs 'a': got %d, want negative", r)
	}
}

func TestMysqlGbkCompareBasic(t *testing.T) {
	if r := MysqlGbkCompare("a", "b"); r >= 0 {
		t.Errorf("a vs b: got %d, want negative", r)
	}
	if r := MysqlGbkCompare("b", "a"); r <= 0 {
		t.Errorf("b vs a: got %d, want positive", r)
	}
	if r := MysqlGbkCompare("abc", "abc"); r != 0 {
		t.Errorf("abc vs abc: got %d, want 0", r)
	}
}

func TestMysqlGbkCompareMixed(t *testing.T) {
	// Chinese + ASCII mixed strings
	if r := MysqlGbkCompare("爱abc", "爱abd"); r >= 0 {
		t.Errorf("爱abc vs 爱abd: got %d, want negative", r)
	}
	if r := MysqlGbkCompare("爱abc", "爱abc"); r != 0 {
		t.Errorf("爱abc vs 爱abc: got %d, want 0", r)
	}
}

func TestMysqlGbkComparePrefix(t *testing.T) {
	if r := MysqlGbkCompare("abc", "abcd"); r >= 0 {
		t.Errorf("abc vs abcd: got %d, want negative", r)
	}
	if r := MysqlGbkCompare("黄山旅游", "黄山旅游-秋季"); r >= 0 {
		t.Errorf("黄山旅游 vs 黄山旅游-秋季: got %d, want negative", r)
	}
}
