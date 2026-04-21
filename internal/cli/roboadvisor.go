package cli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	wallbitroboadvisor "github.com/jeremyjsx/wallbit-go/services/roboadvisor"
	"github.com/spf13/cobra"
)

var (
	roboAdvisorID     int
	roboAdvisorAmount float64
	roboAdvisorFrom   string
	roboAdvisorTo     string
)

var roboadvisorCmd = &cobra.Command{
	Use:   "roboadvisor",
	Short: "Robo-advisor portfolio operations",
}

var roboadvisorBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "List robo-advisor portfolios and balances",
	RunE:  runRoboadvisorBalance,
}

var roboadvisorDepositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Deposit into a robo-advisor portfolio",
	RunE:  runRoboadvisorDeposit,
}

var roboadvisorWithdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "Withdraw from a robo-advisor portfolio",
	RunE:  runRoboadvisorWithdraw,
}

func init() {
	roboadvisorDepositCmd.Flags().IntVar(&roboAdvisorID, "id", 0, "Robo-advisor portfolio ID")
	roboadvisorDepositCmd.Flags().Float64Var(&roboAdvisorAmount, "amount", 0, "Amount to deposit")
	roboadvisorDepositCmd.Flags().StringVar(&roboAdvisorFrom, "from", "", "Source account: DEFAULT or INVESTMENT")
	_ = roboadvisorDepositCmd.MarkFlagRequired("id")
	_ = roboadvisorDepositCmd.MarkFlagRequired("amount")
	_ = roboadvisorDepositCmd.MarkFlagRequired("from")

	roboadvisorWithdrawCmd.Flags().IntVar(&roboAdvisorID, "id", 0, "Robo-advisor portfolio ID")
	roboadvisorWithdrawCmd.Flags().Float64Var(&roboAdvisorAmount, "amount", 0, "Amount to withdraw")
	roboadvisorWithdrawCmd.Flags().StringVar(&roboAdvisorTo, "to", "", "Destination account: DEFAULT or INVESTMENT")
	_ = roboadvisorWithdrawCmd.MarkFlagRequired("id")
	_ = roboadvisorWithdrawCmd.MarkFlagRequired("amount")
	_ = roboadvisorWithdrawCmd.MarkFlagRequired("to")

	roboadvisorCmd.AddCommand(roboadvisorBalanceCmd, roboadvisorDepositCmd, roboadvisorWithdrawCmd)
	rootCmd.AddCommand(roboadvisorCmd)
}

func parseRoboAccountType(raw string) (wallbitroboadvisor.AccountType, error) {
	s := strings.ToUpper(strings.TrimSpace(raw))
	switch s {
	case string(wallbitroboadvisor.AccountTypeDefault):
		return wallbitroboadvisor.AccountTypeDefault, nil
	case string(wallbitroboadvisor.AccountTypeInvestment):
		return wallbitroboadvisor.AccountTypeInvestment, nil
	default:
		return "", fmt.Errorf("invalid account type %q, expected DEFAULT or INVESTMENT", raw)
	}
}

func runRoboadvisorBalance(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.RoboAdvisor.GetBalance(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

func runRoboadvisorDeposit(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	if roboAdvisorID <= 0 {
		return errors.New("--id must be a positive integer")
	}
	if roboAdvisorAmount <= 0 {
		return errors.New("--amount must be positive")
	}
	from, err := parseRoboAccountType(roboAdvisorFrom)
	if err != nil {
		return err
	}

	req := wallbitroboadvisor.DepositRequest{
		RoboAdvisorID: roboAdvisorID,
		Amount:        roboAdvisorAmount,
		From:          from,
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.RoboAdvisor.Deposit(ctx, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

func runRoboadvisorWithdraw(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), app.Timeout())
	defer cancel()

	if roboAdvisorID <= 0 {
		return errors.New("--id must be a positive integer")
	}
	if roboAdvisorAmount <= 0 {
		return errors.New("--amount must be positive")
	}
	to, err := parseRoboAccountType(roboAdvisorTo)
	if err != nil {
		return err
	}

	req := wallbitroboadvisor.WithdrawRequest{
		RoboAdvisorID: roboAdvisorID,
		Amount:        roboAdvisorAmount,
		To:            to,
	}

	svc, err := app.Services()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	out, err := svc.RoboAdvisor.Withdraw(ctx, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
