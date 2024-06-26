// Copyright 2023 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encodings

// Dec8 represents the `dec8` character set encoding.
var Dec8 Encoder = &RangeMap{
	inputEntries: [][]rangeMapEntry{
		{
			{
				inputRange:  rangeBounds{{0, 127}},
				outputRange: rangeBounds{{0, 127}},
				inputMults:  []int{1},
				outputMults: []int{1},
			},
			{
				inputRange:  rangeBounds{{128, 163}},
				outputRange: rangeBounds{{194, 194}, {128, 163}},
				inputMults:  []int{1},
				outputMults: []int{36, 1},
			},
			{
				inputRange:  rangeBounds{{165, 165}},
				outputRange: rangeBounds{{194, 194}, {165, 165}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{167, 167}},
				outputRange: rangeBounds{{194, 194}, {167, 167}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{168, 168}},
				outputRange: rangeBounds{{194, 194}, {164, 164}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{169, 171}},
				outputRange: rangeBounds{{194, 194}, {169, 171}},
				inputMults:  []int{1},
				outputMults: []int{3, 1},
			},
			{
				inputRange:  rangeBounds{{176, 179}},
				outputRange: rangeBounds{{194, 194}, {176, 179}},
				inputMults:  []int{1},
				outputMults: []int{4, 1},
			},
			{
				inputRange:  rangeBounds{{181, 183}},
				outputRange: rangeBounds{{194, 194}, {181, 183}},
				inputMults:  []int{1},
				outputMults: []int{3, 1},
			},
			{
				inputRange:  rangeBounds{{185, 189}},
				outputRange: rangeBounds{{194, 194}, {185, 189}},
				inputMults:  []int{1},
				outputMults: []int{5, 1},
			},
			{
				inputRange:  rangeBounds{{191, 191}},
				outputRange: rangeBounds{{194, 194}, {191, 191}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{192, 207}},
				outputRange: rangeBounds{{195, 195}, {128, 143}},
				inputMults:  []int{1},
				outputMults: []int{16, 1},
			},
			{
				inputRange:  rangeBounds{{209, 214}},
				outputRange: rangeBounds{{195, 195}, {145, 150}},
				inputMults:  []int{1},
				outputMults: []int{6, 1},
			},
			{
				inputRange:  rangeBounds{{215, 215}},
				outputRange: rangeBounds{{197, 197}, {146, 146}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{216, 220}},
				outputRange: rangeBounds{{195, 195}, {152, 156}},
				inputMults:  []int{1},
				outputMults: []int{5, 1},
			},
			{
				inputRange:  rangeBounds{{221, 221}},
				outputRange: rangeBounds{{197, 197}, {184, 184}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{223, 239}},
				outputRange: rangeBounds{{195, 195}, {159, 175}},
				inputMults:  []int{1},
				outputMults: []int{17, 1},
			},
			{
				inputRange:  rangeBounds{{241, 246}},
				outputRange: rangeBounds{{195, 195}, {177, 182}},
				inputMults:  []int{1},
				outputMults: []int{6, 1},
			},
			{
				inputRange:  rangeBounds{{247, 247}},
				outputRange: rangeBounds{{197, 197}, {147, 147}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{248, 252}},
				outputRange: rangeBounds{{195, 195}, {184, 188}},
				inputMults:  []int{1},
				outputMults: []int{5, 1},
			},
			{
				inputRange:  rangeBounds{{253, 253}},
				outputRange: rangeBounds{{195, 195}, {191, 191}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
		},
		nil,
		nil,
		nil,
	},
	outputEntries: [][]rangeMapEntry{
		{
			{
				inputRange:  rangeBounds{{0, 127}},
				outputRange: rangeBounds{{0, 127}},
				inputMults:  []int{1},
				outputMults: []int{1},
			},
		},
		{
			{
				inputRange:  rangeBounds{{128, 163}},
				outputRange: rangeBounds{{194, 194}, {128, 163}},
				inputMults:  []int{1},
				outputMults: []int{36, 1},
			},
			{
				inputRange:  rangeBounds{{165, 165}},
				outputRange: rangeBounds{{194, 194}, {165, 165}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{167, 167}},
				outputRange: rangeBounds{{194, 194}, {167, 167}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{168, 168}},
				outputRange: rangeBounds{{194, 194}, {164, 164}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{169, 171}},
				outputRange: rangeBounds{{194, 194}, {169, 171}},
				inputMults:  []int{1},
				outputMults: []int{3, 1},
			},
			{
				inputRange:  rangeBounds{{176, 179}},
				outputRange: rangeBounds{{194, 194}, {176, 179}},
				inputMults:  []int{1},
				outputMults: []int{4, 1},
			},
			{
				inputRange:  rangeBounds{{181, 183}},
				outputRange: rangeBounds{{194, 194}, {181, 183}},
				inputMults:  []int{1},
				outputMults: []int{3, 1},
			},
			{
				inputRange:  rangeBounds{{185, 189}},
				outputRange: rangeBounds{{194, 194}, {185, 189}},
				inputMults:  []int{1},
				outputMults: []int{5, 1},
			},
			{
				inputRange:  rangeBounds{{191, 191}},
				outputRange: rangeBounds{{194, 194}, {191, 191}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{192, 207}},
				outputRange: rangeBounds{{195, 195}, {128, 143}},
				inputMults:  []int{1},
				outputMults: []int{16, 1},
			},
			{
				inputRange:  rangeBounds{{209, 214}},
				outputRange: rangeBounds{{195, 195}, {145, 150}},
				inputMults:  []int{1},
				outputMults: []int{6, 1},
			},
			{
				inputRange:  rangeBounds{{215, 215}},
				outputRange: rangeBounds{{197, 197}, {146, 146}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{216, 220}},
				outputRange: rangeBounds{{195, 195}, {152, 156}},
				inputMults:  []int{1},
				outputMults: []int{5, 1},
			},
			{
				inputRange:  rangeBounds{{221, 221}},
				outputRange: rangeBounds{{197, 197}, {184, 184}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{223, 239}},
				outputRange: rangeBounds{{195, 195}, {159, 175}},
				inputMults:  []int{1},
				outputMults: []int{17, 1},
			},
			{
				inputRange:  rangeBounds{{241, 246}},
				outputRange: rangeBounds{{195, 195}, {177, 182}},
				inputMults:  []int{1},
				outputMults: []int{6, 1},
			},
			{
				inputRange:  rangeBounds{{247, 247}},
				outputRange: rangeBounds{{197, 197}, {147, 147}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
			{
				inputRange:  rangeBounds{{248, 252}},
				outputRange: rangeBounds{{195, 195}, {184, 188}},
				inputMults:  []int{1},
				outputMults: []int{5, 1},
			},
			{
				inputRange:  rangeBounds{{253, 253}},
				outputRange: rangeBounds{{195, 195}, {191, 191}},
				inputMults:  []int{1},
				outputMults: []int{1, 1},
			},
		},
		nil,
		nil,
	},
	toUpper: map[rune]rune{
		97:  65,
		98:  66,
		99:  67,
		100: 68,
		101: 69,
		102: 70,
		103: 71,
		104: 72,
		105: 73,
		106: 74,
		107: 75,
		108: 76,
		109: 77,
		110: 78,
		111: 79,
		112: 80,
		113: 81,
		114: 82,
		115: 83,
		116: 84,
		117: 85,
		118: 86,
		119: 87,
		120: 88,
		121: 89,
		122: 90,
		224: 192,
		225: 193,
		226: 194,
		227: 195,
		228: 196,
		229: 197,
		230: 198,
		231: 199,
		232: 200,
		233: 201,
		234: 202,
		235: 203,
		236: 204,
		237: 205,
		238: 206,
		239: 207,
		241: 209,
		242: 210,
		243: 211,
		244: 212,
		245: 213,
		246: 214,
		248: 216,
		249: 217,
		250: 218,
		251: 219,
		252: 220,
		255: 376,
	},
	toLower: map[rune]rune{
		65:  97,
		66:  98,
		67:  99,
		68:  100,
		69:  101,
		70:  102,
		71:  103,
		72:  104,
		73:  105,
		74:  106,
		75:  107,
		76:  108,
		77:  109,
		78:  110,
		79:  111,
		80:  112,
		81:  113,
		82:  114,
		83:  115,
		84:  116,
		85:  117,
		86:  118,
		87:  119,
		88:  120,
		89:  121,
		90:  122,
		192: 224,
		193: 225,
		194: 226,
		195: 227,
		196: 228,
		197: 229,
		198: 230,
		199: 231,
		200: 232,
		201: 233,
		202: 234,
		203: 235,
		204: 236,
		205: 237,
		206: 238,
		207: 239,
		209: 241,
		210: 242,
		211: 243,
		212: 244,
		213: 245,
		214: 246,
		216: 248,
		217: 249,
		218: 250,
		219: 251,
		220: 252,
		376: 255,
	},
}
