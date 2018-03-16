package aci

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_USERNAME", nil),
				Description: "Username for ACI account.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_PASSWORD", nil),
				Description: "Password for ACI account.",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_URL", nil),
				Description: "Domain for ACI account.",
			},
			"domain": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_DOMAIN", nil),
				Description: "Domain for ACI account.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"aci_filter": resourceAciFilter(),
		},

		ConfigureFunc: configureClient,
	}
}

func configureClient(d *schema.ResourceData) (interface{}, error) {
	config := AciConfig{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Domain:   d.Get("domain").(string),
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	client, err := config.getAciClient()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c AciConfig) validate() error {
	var err *multierror.Error

	if c.Username == "" {
		err = multierror.Append(err, fmt.Errorf("Username must be configured for the ACI provider"))
	}

	if c.Password == "" {
		err = multierror.Append(err, fmt.Errorf("Password must be configured for the ACI provider"))
	}

	if c.Url == "" {
		err = multierror.Append(err, fmt.Errorf("URL must be configured for the ACI provider"))
	}

	// TODO: validate domain if present to match domain specification

	return err.ErrorOrNil()
}

func (c AciConfig) getAciClient() (interface{}, error) {
	// *cage.Client
	//client := aci.NewClient(c.Username, c.Password, c.Domain)
	//return client,nil
	return nil, nil
}

type AciConfig struct {
	Username string
	Password string
	Domain   string
	Url      string
}
