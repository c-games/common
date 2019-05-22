package msg

type Service struct {
	name string
	queueName string
	queueResponseName string
	command []IServiceCommand
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
		queueName: "cg-user",
		queueResponseName: "cg-user-pipeline-response",
	}

	Order Service = Service{
		name: "orders",
		queueName: "cg-orders",
		queueResponseName: "cg-orders-pipeline-response",
	}

	Wallet Service = Service{
		name: "wallet",
		queueName: "cg-wallet",
		queueResponseName: "cg-wallet-pipeline-response",
	}

	Logger Service = Service{
		name: "logger",
		queueName: "cg-logger",
		queueResponseName: "cg-logger-pipeline-response",
	}
)

func (s *Service) Name() string {
	return s.name
}


func (s *Service) QueueName() string {
	return s.queueName
}

func (s *Service) ResponseQueueName() string {
	return s.queueResponseName
}
