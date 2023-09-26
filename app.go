package main

import (
	"context"

	"github.com/gen2brain/beeep"
)

// App struct
type App struct {
	ctx context.Context
}

type ListPasswords struct {
	Count      int64
	TotalPages int
	Passwords  []PasswordGenerated
	Err        error
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	ConnectDb()
	a.ctx = ctx
}
func (a *App) AddPassword(val PasswordGenerated) string {
	result := DB.Db.Create(&val)
	if result.Error != nil {
		return result.Error.Error()
	}
	return string("")
}

func (a *App) ListPasswords(page int, pageSize int) ListPasswords {
	var count int64
	var locations []PasswordGenerated
	var result ListPasswords
	if err := DB.Db.Model(&locations).Count(&count).Error; err != nil {
		result.Err = err
		return result
	}
	if err := DB.Db.Scopes(Paginate(page, pageSize)).Find(&locations).Error; err != nil {
		result.Err = err
		return result
	}

	if len(locations) == 0 {
		result.TotalPages = 0
		result.Count = 0
		result.Passwords = locations
		result.Err = nil
		return result
	}
	totalPage := int(count/int64(pageSize)) + 1
	result.TotalPages = totalPage
	result.Count = count
	result.Passwords = locations
	return result

}
func (a *App) RemovePassword(id int64) error {
	result := DB.Db.Delete(&PasswordGenerated{}, id)
	return result.Error
}
func (a *App) ShowInfo(title string, message string) {
	err := beeep.Notify(title, message, "assets/information.png")
	if err != nil {
		panic(err)
	}
}
