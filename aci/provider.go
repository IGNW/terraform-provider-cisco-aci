package aci

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_USER", nil),
				Description: "Username for ACI account.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_PASS", nil),
				Description: "Password for ACI account.",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_URL", nil),
				Description: "Domain for ACI account.",
			},
			"allow_insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_ALLOW_INSECURE", nil),
				Description: "Allow insecure connections for ACI endpoint.",
			},
			"domain": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_DOMAIN", nil),
				Description: "Domain for ACI account.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"aci_tenant":      resourceAciTenant(),
			"aci_app_profile": resourceAciAppProfile(),
			// "aci_filter": resourceAciFilter(),
		},

		ConfigureFunc: configureClient,
	}
}

func configureClient(d *schema.ResourceData) (interface{}, error) {
	config := AciConfig{
		Username:      d.Get("username").(string),
		Password:      d.Get("password").(string),
		Url:           d.Get("url").(string),
		Domain:        d.Get("domain").(string),
		AllowInsecure: d.Get("allow_insecure").(bool),
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
	client := cage.InitializeClient(c.Url, c.Username, c.Password, c.AllowInsecure)
	return client, nil
}

type AciConfig struct {
	Username      string
	Password      string
	Domain        string
	Url           string
	AllowInsecure bool
}
