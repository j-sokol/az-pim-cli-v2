package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/j-sokol/az-pim-cli/internal/utils"
	azpim "github.com/j-sokol/az-pim-cli/pkg/az-pim"
)

func main() {
	role := flag.String("r", "", "Role to activate")
	scope := flag.String("s", "", "Scope of activation")
	list := flag.Bool("l", false, "List roles available for activation")
	activate := flag.Bool("a", false, "Activate role")
	flag.Parse()

	accessToken, err := utils.GetAccessToken()
	if err != nil {
		log.Fatal(err)
	}

	userId, err := utils.GetCurrentUserObjectId()
	if err != nil {
		log.Fatal(err)
	}
	result, err := azpim.ListEligibleRoleAssignments(accessToken)
	if err != nil {
		log.Fatal(err)
	}
	if !*list && !*activate {
		fmt.Println("No action specified, -l or -a required")
		flag.Usage()
		return
	}

	if *list {
		fmt.Println("Eligible roles:")
		for _, v := range result.Role {
			fmt.Printf("Role=%s\tScope=%s\n", v.Properties.ExpandedProperties.RoleDefinition.DisplayName, v.Properties.ExpandedProperties.Scope.DisplayName)
		}
		return
	}

	if *activate {
		fmt.Printf("Activating role %s on %s\n", *role, *scope)
		r, err := azpim.GetEligibleRoleAssignment(accessToken, *scope, *role)
		if err != nil {
			log.Fatal(err)
		}
		err = azpim.ActivateRoleAssignment(accessToken, userId, r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Role %s activated for %s\n", r.Properties.ExpandedProperties.RoleDefinition.DisplayName, r.Properties.ExpandedProperties.Principal.DisplayName)
	}
}
