package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func search(path string) ([]fileinfo, string) {
	var info []fileinfo
	folder, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	files, err := folder.Readdir(-1)
	if err != nil {
		log.Println(err)
	}
	for _, file := range files {
		var infocache = fileinfo{
			name:   file.Name(),
			year:   file.ModTime().Year(),
			month:  int(file.ModTime().Month()),
			day:    file.ModTime().Day(),
			hour:   file.ModTime().Hour(),
			minute: file.ModTime().Minute(),
			second: file.ModTime().Second(),
		}
		info = append(info, infocache)
	}
	return info, path
}

func delete(info []fileinfo, path string) {
	var timestamp = fileinfo{
		name:   "timestamp",
		year:   time.Now().Local().Year(),
		month:  int(time.Now().Local().Month()),
		day:    time.Now().Local().Day(),
		hour:   time.Now().Local().Hour(),
		minute: time.Now().Local().Minute(),
		second: time.Now().Local().Second(),
	}
	for _, file := range info {
		if timestamp.year >= file.year && filepath.Ext(file.name) == ".bak" {
			if timestamp.month >= (file.month + 2) {
				os.Remove(path + "/" + file.name)
			} else if timestamp.month > file.month && timestamp.day >= file.day {
				os.Remove(path + "/" + file.name)
			}
		}
		if timestamp.year >= file.year && filepath.Ext(file.name) == ".dif" {
			if timestamp.month >= (file.month + 2) {
				os.Remove(path + "/" + file.name)
			} else if timestamp.month > file.month && timestamp.day >= file.day {
				os.Remove(path + "/" + file.name)
			}
		}
		if timestamp.year >= file.year && filepath.Ext(file.name) == ".trn" {
			if timestamp.month >= (file.month + 2) {
				os.Remove(path + "/" + file.name)
			} else if timestamp.month > file.month && timestamp.day >= file.day {
				os.Remove(path + "/" + file.name)
			}
		}
	}
}
