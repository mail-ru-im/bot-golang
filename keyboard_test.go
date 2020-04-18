package botgolang

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
	btnArray1 = []Button{{Text: "test"}, {Text: "test2"}}
	btnArray2 = []Button{{Text: "tes123t"}, {Text: "test123"}, {Text: "test123"}, {Text: "test2231"}}
	btnArray3 = []Button{{Text: "ew"}}
)

type fields struct {
	Rows [][]Button
}

func TestKeyboard_AddButton(t *testing.T) {
	type args struct {
		rowIndex int
		button   Button
	}

	btn := NewURLButton("test", "mail.ru")
	btn2 := NewCallbackButton("test2", "asdfww34gsw35")
	btnRow := []Button{btn, btn2}

	newBtn := NewCallbackButton("newBtn", "sdtw234")

	expected := []Button{btn, btn2, newBtn}

	tests := []struct {
		name    string
		fields  fields
		args    args
		exp     fields
		wantErr bool
	}{
		{
			name:    "OK",
			fields:  fields{[][]Button{btnRow}},
			args:    args{0, newBtn},
			exp:     fields{[][]Button{expected}},
			wantErr: false,
		},
		{
			name:    "Error",
			fields:  fields{[][]Button{btnRow}},
			args:    args{4, newBtn},
			exp:     fields{[][]Button{btnRow}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keyboard{
				Rows: tt.fields.Rows,
			}
			if err := k.AddButton(tt.args.rowIndex, tt.args.button); (err != nil) != tt.wantErr {
				t.Errorf("AddButton() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, reflect.DeepEqual(tt.exp.Rows, k.GetKeyboard()), true)
		})
	}
}

func TestKeyboard_AddRow(t *testing.T) {
	type args struct {
		row []Button
	}
	btn := NewURLButton("test", "mail.ru")
	btn2 := NewCallbackButton("test2", "asdfww34gsw35")
	btnRow := []Button{btn}
	btnRow2 := []Button{btn2}

	tests := []struct {
		name   string
		fields fields
		args   args
		exp    fields
	}{
		{
			name:   "OK_First",
			fields: fields{nil},
			args:   args{row: btnRow},
			exp:    fields{[][]Button{btnRow}},
		},
		{
			name:   "OK_Add",
			fields: fields{[][]Button{btnRow}},
			args:   args{row: btnRow2},
			exp:    fields{[][]Button{btnRow, btnRow2}},
		},
		{
			name:   "Nil",
			fields: fields{nil},
			args:   args{row: nil},
			exp:    fields{[][]Button{[]Button(nil)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keyboard{
				Rows: tt.fields.Rows,
			}
			k.AddRow(tt.args.row...)
			assert.Equal(t, tt.exp.Rows, k.GetKeyboard())
		})
	}
}

func TestKeyboard_ChangeButton(t *testing.T) {
	type args struct {
		rowIndex    int
		buttonIndex int
		newButton   Button
	}
	newBtn := Button{Text: "TestButton"}

	array1 := make([]Button, len(btnArray1))
	copy(array1, btnArray1)

	expArray := make([]Button, len(array1))
	expArray[0] = newBtn
	expArray[1] = btnArray1[1]

	tests := []struct {
		name    string
		fields  fields
		args    args
		exp     fields
		wantErr bool
	}{
		{
			name:    "OK",
			fields:  fields{[][]Button{array1}},
			args:    args{0, 0, newBtn},
			exp:     fields{[][]Button{expArray}},
			wantErr: false,
		},
		{
			name:    "Error",
			fields:  fields{[][]Button{array1}},
			args:    args{0, 5, newBtn},
			exp:     fields{[][]Button{expArray}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keyboard{
				Rows: tt.fields.Rows,
			}
			if err := k.ChangeButton(tt.args.rowIndex, tt.args.buttonIndex, tt.args.newButton); (err != nil) != tt.wantErr {
				t.Errorf("ChangeButton() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, reflect.DeepEqual(tt.exp.Rows, k.GetKeyboard()), true)
		})
	}
}

func TestKeyboard_DeleteButton(t *testing.T) {
	type args struct {
		rowIndex    int
		buttonIndex int
	}

	array1 := [][]Button{btnArray1}
	exp1 := [][]Button{{btnArray1[1]}}

	tests := []struct {
		name    string
		fields  fields
		args    args
		exp     fields
		wantErr bool
	}{
		{
			name:    "OK",
			fields:  fields{array1},
			args:    args{0, 0},
			exp:     fields{exp1},
			wantErr: false,
		},
		{
			name:    "Error",
			fields:  fields{array1},
			args:    args{2, 13},
			exp:     fields{exp1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keyboard{
				Rows: tt.fields.Rows,
			}
			if err := k.DeleteButton(tt.args.rowIndex, tt.args.buttonIndex); (err != nil) != tt.wantErr {
				t.Errorf("DeleteButton() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.exp.Rows, k.GetKeyboard())
		})
	}
}

func TestKeyboard_DeleteRow(t *testing.T) {
	type args struct {
		index int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		exp     fields
		wantErr bool
	}{
		{
			name:    "OK",
			fields:  fields{[][]Button{btnArray1}},
			args:    args{0},
			exp:     fields{[][]Button{}},
			wantErr: false,
		},
		{
			name:    "Error",
			fields:  fields{[][]Button{btnArray1}},
			args:    args{1},
			exp:     fields{[][]Button{btnArray1}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keyboard{
				Rows: tt.fields.Rows,
			}
			if err := k.DeleteRow(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRow() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.exp.Rows, k.GetKeyboard())
		})
	}
}

func TestKeyboard_RowSize(t *testing.T) {
	type args struct {
		row int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "OK",
			fields: fields{[][]Button{btnArray2}},
			args:   args{0},
			want:   len(btnArray2),
		},
		{
			name:   "OK",
			fields: fields{[][]Button{btnArray2, btnArray3}},
			args:   args{1},
			want:   len(btnArray3),
		},
		{
			name:   "NO_ROW",
			fields: fields{[][]Button{btnArray2, btnArray3}},
			args:   args{4},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keyboard{
				Rows: tt.fields.Rows,
			}
			if got := k.RowSize(tt.args.row); got != tt.want {
				t.Errorf("RowSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyboard_RowsCount(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "OK_3",
			fields: fields{[][]Button{btnArray1, btnArray2, btnArray3}},
			want:   3,
		},
		{
			name:   "OK_2",
			fields: fields{[][]Button{btnArray1, btnArray3}},
			want:   2,
		},
		{
			name:   "OK_0",
			fields: fields{[][]Button{}},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keyboard{
				Rows: tt.fields.Rows,
			}
			if got := k.RowsCount(); got != tt.want {
				t.Errorf("RowsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyboard_SwapRows(t *testing.T) {
	type args struct {
		first  int
		second int
	}

	array1 := [][]Button{btnArray1, btnArray2, btnArray3}
	array2 := [][]Button{btnArray3, btnArray2, btnArray1}

	tests := []struct {
		name    string
		fields  fields
		args    args
		exp     fields
		wantErr bool
	}{
		{
			name:    "OK",
			fields:  fields{array1},
			args:    args{0, 2},
			exp:     fields{array2},
			wantErr: false,
		},
		{
			name:    "Error",
			fields:  fields{array1},
			args:    args{0, 6},
			exp:     fields{array1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keyboard{
				Rows: tt.fields.Rows,
			}
			if err := k.SwapRows(tt.args.first, tt.args.second); (err != nil) != tt.wantErr {
				t.Errorf("SwapRows() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.exp.Rows, k.GetKeyboard())
		})
	}
}
