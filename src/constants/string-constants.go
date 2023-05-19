package constants

func GetLowerPool() []string {
	return []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z",
	}
}

func GetUpperPool() []string {
	return []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z",
	}
}

func GetNumberPool() []string {
	return []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
}

func GetSymbolPool() []string {
	return []string{
		"~",
		"`",
		"!",
		"@",
		"#",
		"$",
		"%",
		"^",
		"&",
		"*",
		"(",
		")",
		"-",
		"_",
		"=",
		"+",
		"[",
		"{",
		"]",
		"}",
		"\\",
		"|",
		":",
		";",
		"\"",
		"'",
		",",
		"<",
		".",
		">",
		"?",
		"/",		
	}
}

func GetAllPool() []string {
	var ret []string
	ret = append(ret, GetLowerPool()...)
	ret = append(ret, GetUpperPool()...)
	ret = append(ret, GetNumberPool()...)
	ret = append(ret, GetSymbolPool()...)
	return ret
}