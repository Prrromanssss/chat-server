package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Prrromanssss/platform_common/pkg/closer"

	"github.com/gofiber/fiber/v2/log"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/Prrromanssss/chat-server/config"
	pb "github.com/Prrromanssss/chat-server/pkg/chat_v1"
)

type App struct {
	cfg             *config.Config
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	log.Info("Config loaded")

	a.cfg = cfg

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.cfg)

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	pb.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatAPI(ctx))

	return nil
}

func (a *App) Run(ctx context.Context, cancel context.CancelFunc) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	// Starting gRPC server
	go func() {
		err := a.runGRPCServer()
		if err != nil {
			log.Panic(err)
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-ctx.Done():
		log.Info("Context cancelled, initiating graceful shutdown...")
		a.grpcServer.GracefulStop()
	case <-quit:
		log.Info("Received termination signal, initiating graceful shutdown...")
		a.grpcServer.GracefulStop()
	}

	log.Info("gRPC server shut down gracefully")
	cancel()

	return nil
}

func (a *App) runGRPCServer() error {
	listener, err := net.Listen(
		"tcp",
		a.cfg.GRPC.Address(),
	)
	if err != nil {
		return errors.Wrapf(err, "Error starting listener")
	}

	log.Infof("Starting gRPC server on %s", a.cfg.GRPC.Address())

	if err := a.grpcServer.Serve(listener); err != nil {
		return errors.Wrapf(err, "Error starting gRPC server")
	}

	return nil
}
