package db

import (
	"sync"
	"time"
)

const (
	maxRetries = 5
	delay      = 5 * time.Second
)

var (
	once     sync.Once
	readOnce sync.Once
)
