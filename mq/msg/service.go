package msg

import "github.com/streadway/amqp"

type Service struct {
	name string
	command []IServiceCommand
	queueConfig *QueueConfig
	responseQueueConfig *QueueConfig
}

type QueueConfig struct {
	name string
	durable bool
	autoDelete bool
	exclusive bool
	noWait bool
	args amqp.Table
	bindKey string
}

type IServiceCommand interface {
	GetCommand() ServiceCommand
}

type ServiceCommand struct {
	name string
}

func NewCommand(commandName string) ServiceCommand {
	return ServiceCommand{name: commandName}
}

func (c *ServiceCommand) GetCommand() string {
	return c.name
}

var (
	User Service = Service{
		name: "user",
		queueConfig: &QueueConfig{
			name: "cg-user",
			durable: true,
			autoDelete: false,
			exclusive: false,
			noWait: false,
			bindKey: "user.*",
		},
		responseQueueConfig: &QueueConfig{
			name: "cg-user-pipeline-response",
			durable: true,
			autoDelete: false,
			exclusive: false,
			noWait: false,
			bindKey: "user-res.*",
		},
	}

	Orders Service = Service{
		name: "orders",
		queueConfig: &QueueConfig{
			name: "cg-orders",
			durable: true,
			autoDelete: false,
			exclusive: false,
			noWait: false,
			bindKey: "orders.*",
		},
		responseQueueConfig: &QueueConfig{
			name: "cg-orders-pipeline-response",
			durable: true,
			autoDelete: false,
			exclusive: false,
			noWait: false,
			bindKey: "orders-res.*",
		},
	}

	Wallet Service = Service{
		name: "wallet",
		queueConfig: &QueueConfig{
			name: "cg-wallet",
			durable: true,
			autoDelete: false,
			exclusive: false,
			noWait: false,
			bindKey: "orders.*",
		},
		responseQueueConfig: &QueueConfig{
			name: "cg-wallet-pipeline-response",
			durable: true,
			autoDelete: false,
			exclusive: false,
			noWait: false,
			bindKey: "orders-res.*",
		},
	}

	Logger Service = Service{
		name: "logger",
		queueConfig: &QueueConfig{
			name: "cg-logger",
			durable: false,
			autoDelete: false,
			exclusive: false,
			noWait: false,
		},
		responseQueueConfig: &QueueConfig{
			name: "cg-logger-pipeline-response",
			durable: false,
			autoDelete: false,
			exclusive: false,
			noWait: false,
		},
	}
)

func (s *Service) Name() string {
	return s.name
}


func (s *Service) QueueName() string {
	return s.queueConfig.name
}

func (s *Service) ResponseQueueName() string {
	return s.responseQueueConfig.name
}

func (s *Service) GetQueueConfig() QueueConfig {
	return *s.queueConfig
}

func (s *Service) GetQueueBind() string {
	return s.queueConfig.bindKey
}

func (s *Service) GetResponseQueueBind() string {
	return s.responseQueueConfig.bindKey
}

func (s *Service) GetResponseQueueConfig() QueueConfig {
	return *s.queueConfig
}

func (s *Service) GetQueueConfigSpread() (bool, bool, bool, bool, amqp.Table) {
	cfg := s.queueConfig
	return cfg.durable, cfg.autoDelete, cfg.exclusive, cfg.noWait, cfg.args
}

func (cfg *QueueConfig) Spread() (string, bool, bool, bool, bool, amqp.Table) {
	return cfg.name, cfg.durable, cfg.autoDelete, cfg.exclusive, cfg.noWait, cfg.args
}
