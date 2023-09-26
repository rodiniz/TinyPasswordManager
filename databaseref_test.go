package main

import (
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestConnectDb(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConnectDb()
		})
	}
}

func TestPaginate(t *testing.T) {
	type args struct {
		page     int
		pageSize int
	}
	tests := []struct {
		name string
		args args
		want func(db *gorm.DB) *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Paginate(tt.args.page, tt.args.pageSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Paginate() = %v, want %v", got, tt.want)
			}
		})
	}
}
