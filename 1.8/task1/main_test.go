package main

import "testing"

func Test_divide(t *testing.T) {
	tests := []struct {
		name string 
		a       float64
		b       float64
		want    float64
		wantErr bool
	}{
		{name: "Обычное деление", a: 5, b: 2, want: 2.5, wantErr: false},
		{name: "Деление на ноль", a: 1, b: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := divide(tt.a, tt.b)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("divide() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("divide() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
