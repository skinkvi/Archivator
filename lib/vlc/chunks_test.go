package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}

	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "Проверка разбивается ли слитая строка на отдельные буквы",
			args: args{
				bStr:      "001000100110100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "Проверка конвертируется ли двоичное в шестадцатириное",
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHexChankst(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want HexChunks
	}{
		{
			name: "Проверка разбивается ли строка на слайс",
			str:  "20 30 3C 18",
			want: HexChunks{"20", "30", "3C", "18"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunk() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hc   HexChunk
		want BinaryChunk
	}{
		{
			name: "Проверка на перевод из бинарного вида в шестанацитиричный",
			hc:   HexChunk("2F"),
			want: BinaryChunk("00101111"),
		},
		{
			name: "Проверка на перевод из бинарного вида в шестанацитиричный",
			hc:   HexChunk("80"),
			want: BinaryChunk("10000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hc.ToBinary(); got != tt.want {
				t.Errorf("TpBinary() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hcs  HexChunks
		want BinaryChunks
	}{
		{
			name: "Проверка на перевод из бинарного вида в шестанацитиричный(много чанков)",
			hcs:  HexChunks{"2F", "80"},
			want: BinaryChunks{"00101111", "10000000"},
		},
		{
			name: "Проверка на перевод из бинарного вида в шестанацитиричный(много чанков)",
			hcs:  HexChunks{"00", "20", "40", "00"},
			want: BinaryChunks{"00000000", "00100000", "01000000", "00000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcs.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TpBinary() = %v, want = %v", got, tt.want)
			}
		})
	}
}
func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		hcs  BinaryChunks
		want string
	}{
		{
			name: "Объединение строковых чисел",
			hcs:  BinaryChunks{"01001111", "1000000"},
			want: "010011111000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcs.Join(); got != tt.want {
				t.Errorf("Join() = %v, want = %v", got, tt.want)
			}
		})
	}
}
