package cmd

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/spf13/cobra"
)

var (
	tenancyID string

	describeCmd = &cobra.Command{
		Use:   "describe",
		Short: "Describe ? JCT",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	// rootCmd.AddCommand(describeCmd)
}

func run() {
	client, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// The OCID of the tenancy containing the compartment.
	_, err = common.DefaultConfigProvider().TenancyOCID()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	saasjcTenancyID := "ocid1.tenancy.oc4..aaaaaaaagvhpzf2a3kjaobar5pcvszd6sqpstadijytnx36ni4tssrreyrfq"

	request := identity.GetCompartmentRequest{
		CompartmentId: &saasjcTenancyID,
	}

	r, err := client.GetCompartment(context.Background(), request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("List of available domains: %v", r)

	// request = identity.ListAvailabilityDomainsRequest{
	// 	CompartmentId: &saasjcTenancyID,
	// }

	// r, err := client.ListAvailabilityDomains(context.Background(), request)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Printf("List of available domains: %v", r.Items)
}
