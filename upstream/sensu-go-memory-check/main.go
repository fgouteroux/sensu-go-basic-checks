package main

//Import the packages we need
import (
	"fmt"
	"os"
	"io"
	"strconv"

	"github.com/sensu/sensu-go/types"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

//Set up some variables. Most notably, warning and critical as time durations
var (
	warning, critical string
	stdin   *os.File
)

//Start our main function
func main() {
	rootCmd := configureRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

//Set up our flags for the command. Note that we have time duration defaults for warning & critical
func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sensu-go-memory-checks",
		Short: "The Sensu Go check for system memory usage",
		RunE:  run,
	}

	cmd.Flags().StringVarP(&warning,
		"warning",
		"w",
		"75",
		"Warning used percentage for system memory.")

	cmd.Flags().StringVarP(&critical,
		"critical",
		"c",
		"90",
		"Critical used percentage for system memory")
		
	return cmd
}

func run(cmd *cobra.Command, args []string) error {

	if len(args) != 0 {
		_ = cmd.Help()
		return fmt.Errorf("invalid argument(s) received")
	}

	if stdin == nil {
		stdin = os.Stdin
	}
	
	event := &types.Event{}
	
	return checkMem(event)
}

//Here we start the meat of what we do.
func checkMem(event *types.Event) error {
	
	//Setting "CheckUptime" as a constant
	const checkName = "CheckMem"
	const metricName = "system_memory_used"
	
	warn, err := strconv.ParseFloat(warning, 64)
	crit, err := strconv.ParseFloat(critical, 64)
	memStat, _ := mem.VirtualMemory()	
	memUsedPct := memStat.UsedPercent
	
	//Let's set up some error handling
	if err != nil {
		msg := fmt.Sprintf("Failed to determine mem %s", err.Error())
		io.WriteString(os.Stdout, msg)
		os.Exit(3)
	}
	
	//Set up our comparison
	if memUsedPct > crit {
		msg := fmt.Sprintf("%s CRITICAL - value = %.2f | %s=%.2f\n", checkName, memUsedPct, metricName, memUsedPct)
		io.WriteString(os.Stdout, msg)
		os.Exit(2)
	} else if memUsedPct >= warn && memUsedPct <= crit {
		msg := fmt.Sprintf("%s WARNING - value = %.2f | %s=%.2f\n", checkName, memUsedPct, metricName, memUsedPct)
		io.WriteString(os.Stdout, msg)
		os.Exit(1)
	} else {
		msg := fmt.Sprintf("%s OK - value = %.2f | %s=%.2f\n", checkName, memUsedPct, metricName, memUsedPct)
		io.WriteString(os.Stdout, msg)
		os.Exit(0)
	}
	
	return nil
}