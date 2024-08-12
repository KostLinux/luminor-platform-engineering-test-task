package controller

import (
	"context"
	"log"
	appConfig "messager-web-app/config"
	"messager-web-app/model"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func SendMessage(cfg *appConfig.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload model.Payload
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			BadRequestError(ctx, err)
			return
		}

		// Extract the IP address from the request context
		clientIP := ctx.ClientIP()
		payload.SentFromIP = clientIP

		if err := model.ValidatePayload(payload); err != nil {
			ValidationError(ctx, err)
			return
		}

		data, err := jsoniter.Marshal(payload)
		if err != nil {
			InternalServerError(ctx, err)
			return
		}

		// Set a timeout for the request
		timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Minute)
		defer cancel()

		if err := sendMessageOrWriteToFile(timeoutCtx, cfg, data); err != nil {
			InternalServerError(ctx, err)
			return
		}

		StatusOK(ctx, "Message sent successfully.")
	}
}

func sendMessageOrWriteToFile(ctx context.Context, cfg *appConfig.Application, data []byte) error {
	if cfg.Queue != "" {
		log.Printf("Queue is set: %s", cfg.Queue)

		// Load the AWS configuration
		awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.Region))
		if err != nil {
			log.Printf("Failed to load AWS configuration: %v", err)
			return err
		}

		svc := sqs.NewFromConfig(awsCfg)

		// Log the message details
		log.Printf("Sending message to SQS: %s", string(data))

		// Send the message
		result, err := svc.SendMessage(ctx, &sqs.SendMessageInput{
			QueueUrl:    aws.String(cfg.Queue),
			MessageBody: aws.String(string(data)),
		})
		if err != nil {
			log.Printf("Failed to send message to SQS: %v", err)
			return err
		}

		log.Printf("Message sent to SQS successfully: %v", result)

	} else {
		log.Printf("Queue is not set, writing to file: %s", cfg.FilePath)

		file, err := os.Create(cfg.FilePath)
		if err != nil {
			log.Printf("Failed to create file: %v", err)
			return err
		}
		defer file.Close()

		if _, err := file.Write(data); err != nil {
			log.Printf("Failed to write to file: %v", err)
			return err
		}
	}

	return nil
}
