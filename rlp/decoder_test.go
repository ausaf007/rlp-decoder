package rlp

import (
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
	"testing"
)

type stringPair struct {
	input, want string
}

func TestDecode(t *testing.T) {
	log.SetLevel(log.WarnLevel)

	// Test Cases taken from: https://github.com/ethereum/tests/blob/develop/RLPTests/rlptest.json

	input1 := "e5922034342e38313538393735343033373334319132302e3435343733343334343535353435"
	want1 := "List {\n  String \" 44.81589754037341\"\n  String \"20.45473434455545\"\n}"

	input2 := "ea8d4976616e20426a656c616a61638e4d616c697361205075736f6e6a618c536c61766b6f204a656e6963"
	want2 := "List {\n  String \"Ivan Bjelajac\"\n  String \"Malisa Pusonja\"\n  String \"Slavko Jenic\"\n}"

	input3 := "80"
	want3 := "String \"\""

	input4 := "83646f67"
	want4 := "String \"dog\""

	input5 := "b74c6f72656d20697073756d20646f6c6f722073697420616d65742c20636f6e736563746574757220616469706973696369" +
		"6e6720656c69"
	want5 := "String \"Lorem ipsum dolor sit amet, consectetur adipisicing eli\""

	input6 := "b8384c6f72656d20697073756d20646f6c6f722073697420616d65742c20636f6e73656374657475722061646970697369636" +
		"96e6720656c6974"
	want6 := "String \"Lorem ipsum dolor sit amet, consectetur adipisicing elit\""

	input7 := "b904004c6f72656d20697073756d20646f6c6f722073697420616d65742c20636f6e736563746574757220616469706973636" +
		"96e6720656c69742e20437572616269747572206d6175726973206d61676e612c20737573636970697420736564207665686963756c" +
		"61206e6f6e2c20696163756c697320666175636962757320746f72746f722e2050726f696e20737573636970697420756c747269636" +
		"96573206d616c6573756164612e204475697320746f72746f7220656c69742c2064696374756d207175697320747269737469717565" +
		"2065752c20756c7472696365732061742072697375732e204d6f72626920612065737420696d70657264696574206d6920756c6c616" +
		"d636f7270657220616c6971756574207375736369706974206e6563206c6f72656d2e2041656e65616e2071756973206c656f206d6f" +
		"6c6c69732c2076756c70757461746520656c6974207661726975732c20636f6e73657175617420656e696d2e204e756c6c6120756c7" +
		"4726963657320747572706973206a7573746f2c20657420706f73756572652075726e6120636f6e7365637465747572206e65632e20" +
		"50726f696e206e6f6e20636f6e76616c6c6973206d657475732e20446f6e65632074656d706f7220697073756d20696e206d6175726" +
		"97320636f6e67756520736f6c6c696369747564696e2e20566573746962756c756d20616e746520697073756d207072696d69732069" +
		"6e206661756369627573206f726369206c756374757320657420756c74726963657320706f737565726520637562696c69612043757" +
		"261653b2053757370656e646973736520636f6e76616c6c69732073656d2076656c206d617373612066617563696275732c20656765" +
		"74206c6163696e6961206c616375732074656d706f722e204e756c6c61207175697320756c747269636965732070757275732e20507" +
		"26f696e20617563746f722072686f6e637573206e69626820636f6e64696d656e74756d206d6f6c6c69732e20416c697175616d2063" +
		"6f6e73657175617420656e696d206174206d65747573206c75637475732c206120656c656966656e642070757275732065676573746" +
		"1732e20437572616269747572206174206e696268206d657475732e204e616d20626962656e64756d2c206e65717565206174206175" +
		"63746f72207472697374697175652c206c6f72656d206c696265726f20616c697175657420617263752c206e6f6e20696e746572647" +
		"56d2074656c6c7573206c65637475732073697420616d65742065726f732e20437261732072686f6e6375732c206d65747573206163" +
		"206f726e617265206375727375732c20646f6c6f72206a7573746f20756c747269636573206d657475732c20617420756c6c616d636" +
		"f7270657220766f6c7574706174"
	want7 := "String \"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur mauris magna, suscipit sed" +
		" vehicula non, iaculis faucibus tortor. Proin suscipit ultricies malesuada. Duis tortor elit, dictum quis t" +
		"ristique eu, ultrices at risus. Morbi a est imperdiet mi ullamcorper aliquet suscipit nec lorem. Aenean qui" +
		"s leo mollis, vulputate elit varius, consequat enim. Nulla ultrices turpis justo, et posuere urna consectet" +
		"ur nec. Proin non convallis metus. Donec tempor ipsum in mauris congue sollicitudin. Vestibulum ante ipsum " +
		"primis in faucibus orci luctus et ultrices posuere cubilia Curae; Suspendisse convallis sem vel massa fauci" +
		"bus, eget lacinia lacus tempor. Nulla quis ultricies purus. Proin auctor rhoncus nibh condimentum mollis. A" +
		"liquam consequat enim at metus luctus, a eleifend purus egestas. Curabitur at nibh metus. Nam bibendum, neq" +
		"ue at auctor tristique, lorem libero aliquet arcu, non interdum tellus lectus sit amet eros. Cras rhoncus, " +
		"metus ac ornare cursus, dolor justo ultrices metus, at ullamcorper volutpat\""

	input8 := "c0"
	want8 := "List {\n}"

	input9 := "cc83646f6783676f6483636174"
	want9 := "List {\n  String \"dog\"\n  String \"god\"\n  String \"cat\"\n}"

	input10 := "c6827a77c10401"
	want10 := "List {\n  String \"zw\"\n  List {\n    String \"\u0004\"\n  }\n  String \"\u0001\"\n}"

	input11 := "f784617364668471776572847a78637684617364668471776572847a78637684617364668471776572847a78637684617364" +
		"668471776572"
	want11 := "List {\n  String \"asdf\"\n  String \"qwer\"\n  String \"zxcv\"\n  String \"asdf\"\n  String \"qwer\"" +
		"\n  String \"zxcv\"\n  String \"asdf\"\n  String \"qwer\"\n  String \"zxcv\"\n  String \"asdf\"\n  String " +
		"\"qwer\"\n}"

	input12 := "f840cf84617364668471776572847a786376cf84617364668471776572847a786376cf84617364668471776572847a7863" +
		"76cf84617364668471776572847a786376"
	want12 := "List {\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    " +
		"String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String " +
		"\"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"" +
		"\n  }\n}"

	input13 := "f90200cf84617364668471776572847a786376cf84617364668471776572847a786376cf84617364668471776572847a7863" +
		"76cf84617364668471776572847a786376cf84617364668471776572847a786376cf84617364668471776572847a786376cf8461736" +
		"4668471776572847a786376cf84617364668471776572847a786376cf84617364668471776572847a786376cf846173646684717765" +
		"72847a786376cf84617364668471776572847a786376cf84617364668471776572847a786376cf84617364668471776572847a78637" +
		"6cf84617364668471776572847a786376cf84617364668471776572847a786376cf84617364668471776572847a786376cf84617364" +
		"668471776572847a786376cf84617364668471776572847a786376cf84617364668471776572847a786376cf8461736466847177657" +
		"2847a786376cf84617364668471776572847a786376cf84617364668471776572847a786376cf84617364668471776572847a786376" +
		"cf84617364668471776572847a786376cf84617364668471776572847a786376cf84617364668471776572847a786376cf846173646" +
		"68471776572847a786376cf84617364668471776572847a786376cf84617364668471776572847a786376cf84617364668471776572" +
		"847a786376cf84617364668471776572847a786376cf84617364668471776572847a786376"
	want13 := "List {\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    " +
		"String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String " +
		"\"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"" +
		"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String " +
		"\"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"" +
		"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n " +
		" List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"" +
		"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    " +
		"String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List " +
		"{\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    " +
		"String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String " +
		"\"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n   " +
		" String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String" +
		" \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"" +
		"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String" +
		" \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String " +
		"\"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String " +
		"\"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    " +
		"String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    " +
		"String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    " +
		"String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  " +
		"List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n   " +
		" String \"qwer\"\n    String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n   " +
		" String \"zxcv\"\n  }\n  List {\n    String \"asdf\"\n    String \"qwer\"\n    String \"zxcv\"\n  }\n}"

	input14 := "c4c2c0c0c0"
	want14 := "List {\n  List {\n    List {\n    }\n    List {\n    }\n  }\n  List {\n  }\n}"

	input15 := "c7c0c1c0c3c0c1c0"
	want15 := "List {\n  List {\n  }\n  List {\n    List {\n    }\n  }\n  List {\n    List {\n    }\n    " +
		"List {\n      List {\n      }\n    }\n  }\n}"

	input16 := "ecca846b6579318476616c31ca846b6579328476616c32ca846b6579338476616c33ca846b6579348476616c34"
	want16 := "List {\n  List {\n    String \"key1\"\n    String \"val1\"\n  }\n  List {\n    String \"key2\"\n" +
		"    String \"val2\"\n  }\n  List {\n    String \"key3\"\n    String \"val3\"\n  }\n  List {\n    String " +
		"\"key4\"\n    String \"val4\"\n  }\n}"

	var stringPairs = []stringPair{
		{input1, want1},
		{input2, want2},
		{input3, want3},
		{input4, want4},
		{input5, want5},
		{input6, want6},
		{input7, want7},
		{input8, want8},
		{input9, want9},
		{input10, want10},
		{input11, want11},
		{input12, want12},
		{input13, want13},
		{input14, want14},
		{input15, want15},
		{input16, want16},
	}

	for _, test := range stringPairs {
		got, err := Decode(test.input)
		if err != nil || strings.Compare(got, test.want) != 0 {
			t.Errorf("Expected '%q', but got '%q'\n", test.want, got)
		}
	}
}

func TestDecodeErrorCase(t *testing.T) {
	input := "83646f6767"
	_, err := Decode(input)

	if reflect.DeepEqual(err, ItemNotEnclosedInsideList) {
		//passed
	} else {
		t.Errorf("Error not caught")
	}
}

func TestDecodeErrorCase2(t *testing.T) {
	input := "ea8d4976616e20426a656c616a61638e4d616c697361205075736f6e6a618c536c61766b6f204a656e696363"
	_, err := Decode(input)

	if reflect.DeepEqual(err, ItemNotEnclosedInsideList) {
		//passed
	} else {
		t.Errorf("Error not caught")
	}
}
