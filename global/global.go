package global

import (
	"config_service/pkgs/settings"

	"github.com/segmentio/kafka-go"
)

var (
	Config        *settings.Config
	KafkaProducer *kafka.Writer
)
