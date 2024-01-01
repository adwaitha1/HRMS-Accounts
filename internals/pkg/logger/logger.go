package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func GetLogger(ctx context.Context) *logrus.Entry {
	// Create a new lumberjack logger for log rotation
	logRotation := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    5, // in megabytes
		MaxBackups: 3, // maximum number of old log files to retain
		MaxAge:     7, // maximum number of days to retain old log files
		Compress:   true,
	}

	// Create a new logrus logger instance
	logger := logrus.New()

	// Set the desired log level (e.g., Debug, Info, Warn, Error)
	logger.SetLevel(logrus.DebugLevel)

	// Set the desired log format (e.g., JSON, Text)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set the output to the lumberjack logger for log rotation
	logger.SetOutput(logRotation)

	// Create an entry with context fields
	entry := logger.WithFields(logrus.Fields{
		"requestID": getRequestID(ctx),
	})

	return entry
}

// Example function to extract request ID from context
func getRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		// If request ID is not present in context, generate a new one
		requestID = generateRequestID()
		ctx = context.WithValue(ctx, "requestID", requestID)
	}

	return requestID
}

// Example function to generate a random request ID
func generateRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
