package options

import (
	"os"
	"reflect"
	"testing"
)

func TestOptions(t *testing.T) {
	if j := os.Getenv(EnvGoogleAuthJSON); j != "" {
		os.Unsetenv(EnvGoogleAuthJSON)
		defer os.Setenv(EnvGoogleAuthJSON, j)
	}

	for _, c := range []struct {
		name string
		envs map[string]string
		args []string
		want *Options
	}{
		{
			"multiple hosts with comma separated",
			map[string]string{},
			[]string{
				"-host", "a.com,b.com",
			},
			&Options{
				AllowedHosts: []string{
					"a.com",
					"b.com",
				},
				Port:       80,
				Enviroment: "production",
			},
		},
		{
			"multiple hosts with specified multiple times",
			map[string]string{},
			[]string{
				"-host", "a.com",
				"-host", "b.com",
			},
			&Options{
				AllowedHosts: []string{
					"a.com",
					"b.com",
				},
				Port:       80,
				Enviroment: "production",
			},
		},
		{
			"multiple hosts with both way",
			map[string]string{},
			[]string{
				"-host", "a.com,b.com",
				"-host", "c.com",
			},
			&Options{
				AllowedHosts: []string{
					"a.com",
					"b.com",
					"c.com",
				},
				Port:       80,
				Enviroment: "production",
			},
		},
		{
			"only env",
			map[string]string{
				EnvBucket: "foo",
			},
			[]string{},
			&Options{
				Bucket:     "foo",
				Port:       80,
				Enviroment: "production",
			},
		},
		{
			"only args",
			map[string]string{},
			[]string{
				"-bucket", "bar",
			},
			&Options{
				Bucket:     "bar",
				Port:       80,
				Enviroment: "production",
			},
		},
		{
			"envs and args",
			map[string]string{
				EnvBucket: "foo",
			},
			[]string{
				"-bucket", "bar",
			},
			&Options{
				Bucket:     "bar",
				Port:       80,
				Enviroment: "production",
			},
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			got := &Options{}

			for _, k := range Envs {
				os.Setenv(k, "")
			}
			for k, v := range c.envs {
				os.Setenv(k, v)
			}
			if err := got.parse(c.args); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Error("ENVS:")
				for _, k := range Envs {
					t.Errorf("%s: %s\n", k, os.Getenv(k))
				}
				t.Errorf("\ngot:\n%+v\nwant:\n%+v", got, c.want)
			}
		})
	}
}
