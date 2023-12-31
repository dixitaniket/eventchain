package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/dixitaniket/eventchain/observer/client"
	"github.com/dixitaniket/eventchain/observer/event"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	envVariablePass = "password"
)

var rootCmd = &cobra.Command{
	Use:   "observer [config-file]",
	Args:  cobra.ExactArgs(1),
	Short: "start observing events",
	RunE:  observerHandler,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func observerHandler(cmd *cobra.Command, args []string) error {
	cfg, err := ParseConfig(args[0])
	if err != nil {
		return err
	}

	logFormat := zerolog.ConsoleWriter{Out: os.Stderr}
	logger := zerolog.New(logFormat).Level(zerolog.InfoLevel).With().Timestamp().Logger()
	rpcTimeout, err := time.ParseDuration(cfg.RPC.RPCTimeout)
	if err != nil {
		return fmt.Errorf("failed to parse RPC timeout: %w", err)
	}

	// Gather pass via env variable || std input
	keyringPass, err := getKeyringPassword()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(cmd.Context())
	g, ctx := errgroup.WithContext(ctx)
	trapSignal(cancel, logger)

	// create oracle client for sending tx
	oc, err := client.NewOracleClient(
		ctx,
		logger,
		cfg.Account.ChainID,
		cfg.Keyring.Backend,
		cfg.Keyring.Dir,
		keyringPass,
		cfg.RPC.TMRPCEndpoint,
		rpcTimeout,
		cfg.Account.Address,
		cfg.GasConfig.GasPrice,
		cfg.GasConfig.GasAdjustment,
	)

	// create event observer
	observer := event.NewObserver(
		logger,
		cfg.Account.Address,
		cfg.EthConfig.ContractAddress,
		cfg.EthConfig.RPC,
		cfg.TimeoutHeight,
		oc,
	)

	g.Go(func() error {
		return observer.Start(ctx)
	})

	return g.Wait()
}

func trapSignal(cancel context.CancelFunc, logger zerolog.Logger) {
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		sig := <-sigCh
		logger.Info().Str("signal", sig.String()).Msg("caught signal; shutting down...")
		cancel()
	}()
}

func getKeyringPassword() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	pass := os.Getenv(envVariablePass)
	if pass == "" {
		return input.GetString("Enter keyring password", reader)
	}
	return pass, nil
}
