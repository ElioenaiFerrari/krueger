/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
)

type Header struct {
	Key   string `yaml:"key" validate:"required"`
	Value string `yaml:"value" validate:"required"`
}

type Service struct {
	ExpectedStatusCode int           `yaml:"expected_status_code" validate:"required,number"`
	Headers            []Header      `yaml:"headers" validate:"-"`
	Method             string        `yaml:"method" validate:"required,oneof='GET' 'HEAD' 'OPTIONS'"`
	Retries            int           `yaml:"retries" validate:"required,number"`
	Timeout            time.Duration `yaml:"timeout" validate:"required"`
	Title              string        `yaml:"title" validate:"required,min=3"`
	URL                string        `yaml:"url" validate:"required,url"`
}

type Schedule struct {
	Interval time.Duration `yaml:"interval" validate:"required"`
}

type Config struct {
	Services []Service `yaml:"services" validate:"required"`
	Schedule Schedule  `yaml:"schedule" validate:"required"`
}

var (
	errorLogger   = color.New(color.FgBlack, color.BgRed, color.Bold)
	history       = map[string]int{}
	successLogger = color.New(color.FgBlack, color.BgGreen, color.Bold)
	warningLogger = color.New(color.FgBlack, color.BgYellow, color.Bold)
)

func fetchApp(service Service, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, service.Timeout)
	defer cancel()

	validate := validator.New()
	if err := validate.StructCtx(ctx, service); err != nil {
		errorLogger.Println(err)
		return
	}

	req, err := http.NewRequestWithContext(ctx, service.Method, service.URL, nil)
	if err != nil {
		errorLogger.Println(err)
		return
	}

	for _, header := range service.Headers {
		req.Header.Set(header.Key, header.Value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errorLogger.Println(err)
		return
	}

	if service.ExpectedStatusCode != res.StatusCode {
		history[service.URL]++
		if history[service.URL] >= service.Retries {
			errorLogger.Printf("%s exceeded %d retries errors", service.URL, service.Retries)
			fmt.Println()
			return
		}
		warningLogger.Printf(
			"%s %s %d -- expected %d received %d -- %d retries from %d",
			service.Method,
			service.URL,
			res.StatusCode,
			service.ExpectedStatusCode,
			res.StatusCode,
			history[service.URL],
			service.Retries,
		)
		fmt.Println()
	} else {
		history[service.URL] = 0
		successLogger.Printf("%s %s %d", service.Method, service.URL, res.StatusCode)
		fmt.Println()
	}
}

func FetchApps(config Config) {
	wg := new(sync.WaitGroup)

	for range time.Tick(config.Schedule.Interval) {
		for _, service := range config.Services {
			wg.Add(1)
			go fetchApp(service, wg)

			wg.Wait()
		}
	}
}
