package main

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type UpfConfig struct {
	InterfaceName  string `mapstructure:"interface_name"`
	ApiAddress     string `mapstructure:"api_address"`
	PfcpAddress    string `mapstructure:"pfcp_address"`
	PfcpNodeId     string `mapstructure:"pfcp_node_id"`
	MetricsAddress string `mapstructure:"metrics_address"`
}

var config UpfConfig

func LoadConfig() error {
	pflag.String("iface", "lo", "Interface to bind XDP program to")
	pflag.String("aaddr", ":8080", "Address to bind api server to")
	pflag.String("paddr", ":8805", "Address to bind PFCP server to")
	pflag.String("nodeid", "localhost", "PFCP Server Node ID")
	pflag.String("maddr", ":9090", "Address to bind metrics server to")
	pflag.Parse()

	viper.BindPFlag("interface_name", pflag.Lookup("iface"))
	viper.BindPFlag("api_address", pflag.Lookup("aaddr"))
	viper.BindPFlag("pfcp_address", pflag.Lookup("paddr"))
	viper.BindPFlag("pfcp_node_id", pflag.Lookup("nodeid"))
	viper.BindPFlag("metrics_address", pflag.Lookup("maddr"))

	viper.SetDefault("interface_name", "lo")
	viper.SetDefault("api_address", ":8080")
	viper.SetDefault("pfcp_address", ":8805")
	viper.SetDefault("pfcp_node_id", "localhost")
	viper.SetDefault("metrics_address", ":9090")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("upf")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Unable to read config file, %v", err)
	}

	log.Println(viper.AllSettings())
	var c UpfConfig
	if err := viper.UnmarshalExact(&c); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return err
	}
	log.Println(c)
	return nil
}
