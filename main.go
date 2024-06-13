package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd.Execute()
}

var useOrigin bool

var rootCmd = &cobra.Command{
	Use:   "tagmaker",
	Short: "A CLI tool for creating and pushing git tags",
	Long:  `This tool helps you create and push git tags to a remote repository using a simple CLI interface.`,
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&useOrigin, "origin", false, "Create tag from origin branch")
	rootCmd.AddCommand(tagCmdDevelopment)
	rootCmd.AddCommand(tagCmdStaging)
	rootCmd.AddCommand(tagCmdRegress)
	rootCmd.AddCommand(tagCmdAll)
}

func getBranchName(branchName string) string {
	if useOrigin {
		return fmt.Sprintf("origin/%s", branchName)
	}
	return branchName
}

// tagCmdDevelopment represents the development tag command
var tagCmdDevelopment = &cobra.Command{
	Use:   "dev [branch-name]",
	Short: "Create and push a new git tag to development",
	Long: `This command allows you to create a new git tag and push it to the remote repository for development.
Usage:
    tagmaker dev <branch-name>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := getBranchName(args[0])
		tagName := fmt.Sprintf("develop_%s", time.Now().Format("0601021504"))

		if useOrigin {
			// Fetch the latest updates from the remote repository
			err := runCommand("git", "fetch", "origin")
			if err != nil {
				fmt.Println("Error fetching from remote:", err)
				return
			}
		}

		// Checkout the branch
		err := runCommand("git", "checkout", branchName)
		if err != nil {
			fmt.Println("Error checking out branch:", err)
			return
		}

		// Create the tag
		err = runCommand("git", "tag", tagName)
		if err != nil {
			fmt.Println("Error creating tag:", err)
			return
		}

		// Push the tag to the remote repository
		err = runCommand("git", "push", "origin", tagName)
		if err != nil {
			fmt.Println("Error pushing tag to remote:", err)
			return
		}

		fmt.Println("Tag pushed successfully:", tagName)
	},
}

// tagCmdStaging represents the staging tag command
var tagCmdStaging = &cobra.Command{
	Use:   "stg [branch-name]",
	Short: "Create and push a new git tag to staging",
	Long: `This command allows you to create a new git tag and push it to the remote repository for staging.
Usage:
    tagmaker stg <branch-name>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := getBranchName(args[0])
		tagName := fmt.Sprintf("staging_%s", time.Now().Format("0601021504"))

		if useOrigin {
			// Fetch the latest updates from the remote repository
			err := runCommand("git", "fetch", "origin")
			if err != nil {
				fmt.Println("Error fetching from remote:", err)
				return
			}
		}

		// Checkout the branch
		err := runCommand("git", "checkout", branchName)
		if err != nil {
			fmt.Println("Error checking out branch:", err)
			return
		}

		// Create the tag
		err = runCommand("git", "tag", tagName)
		if err != nil {
			fmt.Println("Error creating tag:", err)
			return
		}

		// Push the tag to the remote repository
		err = runCommand("git", "push", "origin", tagName)
		if err != nil {
			fmt.Println("Error pushing tag to remote:", err)
			return
		}

		fmt.Println("Tag pushed successfully:", tagName)
	},
}

// tagCmdRegress represents the regress tag command
var tagCmdRegress = &cobra.Command{
	Use:   "rgr [branch-name]",
	Short: "Create and push a new git tag to regress",
	Long: `This command allows you to create a new git tag and push it to the remote repository for regress.
Usage:
    tagmaker rgr <branch-name>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := getBranchName(args[0])
		tagName := fmt.Sprintf("regress_%s", time.Now().Format("0601021504"))

		if useOrigin {
			// Fetch the latest updates from the remote repository
			err := runCommand("git", "fetch", "origin")
			if err != nil {
				fmt.Println("Error fetching from remote:", err)
				return
			}
		}

		// Checkout the branch
		err := runCommand("git", "checkout", branchName)
		if err != nil {
			fmt.Println("Error checking out branch:", err)
			return
		}

		// Create the tag
		err = runCommand("git", "tag", tagName)
		if err != nil {
			fmt.Println("Error creating tag:", err)
			return
		}

		// Push the tag to the remote repository
		err = runCommand("git", "push", "origin", tagName)
		if err != nil {
			fmt.Println("Error pushing tag to remote:", err)
			return
		}

		fmt.Println("Tag pushed successfully:", tagName)
	},
}

// tagCmdAll represents the all tag command
var tagCmdAll = &cobra.Command{
	Use:   "all [branch-name]",
	Short: "Create and push new git tags for all environments",
	Long: `This command allows you to create new git tags and push them to the remote repository for all environments (development, staging, regress).
Usage:
    tagmaker all <branch-name>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := getBranchName(args[0])
		envs := []string{"develop", "staging", "regress"}
		for _, env := range envs {
			tagName := fmt.Sprintf("%s_%s", env, time.Now().Format("0601021504"))

			if useOrigin {
				// Fetch the latest updates from the remote repository
				err := runCommand("git", "fetch", "origin")
				if err != nil {
					fmt.Println("Error fetching from remote:", err)
					return
				}
			}

			// Checkout the branch
			err := runCommand("git", "checkout", branchName)
			if err != nil {
				fmt.Println("Error checking out branch:", err)
				return
			}

			// Create the tag
			err = runCommand("git", "tag", tagName)
			if err != nil {
				fmt.Println("Error creating tag:", err)
				return
			}

			// Push the tag to the remote repository
			err = runCommand("git", "push", "origin", tagName)
			if err != nil {
				fmt.Println("Error pushing tag to remote:", err)
				return
			}

			fmt.Printf("Tag pushed successfully for %s: %s\n", env, tagName)
		}
	},
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
