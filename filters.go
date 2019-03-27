package goform

import (
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
	uuid "github.com/satori/go.uuid"
	"github.com/semihs/goform/slugify"
)

type FilterInterface interface {
	Apply() error
	SetValue(string)
	SetValues([]string)
	SetFile(file *File)
}

type Filter struct {
	Options map[string]interface{}
	Value   string
	Values  []string
	File    *File
}

func (filter *Filter) SetValue(s string) {
	filter.Value = s
}

func (filter *Filter) SetValues(s []string) {
	filter.Values = s
}

func (filter *Filter) SetFile(f *File) {
	filter.File = f
}

func (filter *Filter) SetOptions(o map[string]interface{}) {
	filter.Options = o
}

type RenameFilter struct {
	Filter
}

func (filter *RenameFilter) Apply() error {
	if filter.File == nil {
		return nil
	}
	path := filter.Options["path"].(string)
	randomize := filter.Options["randomize"].(bool)

	fileName := path + strings.Replace(strings.ToLower(slugify.Marshal(filter.File.Name)), "."+filter.File.Extension, "", -1)

	if randomize {
		uid, _ := uuid.NewV4()
		fileName = fileName + "-" + uid.String()
	}
	fileName = strings.ToLower(fileName)

	filter.File.Location = fileName

	return nil
}

type ResizeConversion struct {
	Width  int
	Height int
	Layout string
}

type ImageResizeFilter struct {
	Filter
}

func (filter *ImageResizeFilter) Apply() error {
	if filter.File == nil {
		return nil
	}
	conversions := filter.Options["conversions"].([]ResizeConversion)

	filter.File.Binary.Seek(0, 0)
	img, format, err := image.Decode(filter.File.Binary)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, conversion := range conversions {
		location := fmt.Sprintf(conversion.Layout, filter.File.Location, conversion.Width, conversion.Height, filter.File.Extension)
		fmt.Println("Resizing to", location)

		m := resize.Resize(uint(conversion.Width), uint(conversion.Height), img, resize.Lanczos3)
		out, err := os.Create(location)
		if err != nil {
			return err
		}
		defer out.Close()
		if format == "jpeg" {
			jpeg.Encode(out, m, nil)
		}
		if format == "png" {
			png.Encode(out, m)
		}
		if format == "gif" {
			gif.Encode(out, m, nil)
		}
	}
	return nil
}
