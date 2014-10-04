package config

import "fmt"

// Config runs the config subcommand
func Config(root string, targets []string) error {
	client, err := getEtcdClient()
	if err != nil {
		return err
	}
	rootPath := "/deis/" + root + "/"
	// vals, err := doConfigGet(client, rootPath, targets)
	_, err = doConfigGet(client, rootPath, targets)
	if err != nil {
		return err
	}
	// print results
	// for _, v := range vals {
	// 	fmt.Printf("%s GET: %v\n", rootPath, v)
	// }
	return nil
	// usage := `Deis Cluster Configuration

	//    Usage:
	//    deisctl config <target> get [<key>...] [options]
	//    deisctl config <target> set <key=val>... [options]

	//    Options:
	//    --verbose                   print out the request bodies [default: false]
	//    `
	// // parse command-line arguments
	// args, err := docopt.Parse(usage, nil, true, "", true)
	// if err != nil {
	// 	return err
	// }
	// err = setConfigFlags(args)
	// if err != nil {
	// 	return err
	// }
	// return doConfig(args)
}

func ConfigSet(root string, key string, target string) error {
	client, err := getEtcdClient()
	if err != nil {
		return err
	}
	rootPath := "/deis/" + root + "/" + key
	if _, err := doConfigSet(client, rootPath, target); err != nil {
		return err
	}
	return nil
}

// Flags for config package
var Flags struct {
}

func setConfigFlags(args map[string]interface{}) error {
	return nil
}

func doConfigSet(client *etcdClient, root string, target string) ([]string, error) {
	var result []string
	val, err := client.Set(root, target)
	if err != nil {
		return result, err
	}
	fmt.Printf("%s SET: %v\n", root, val)
	result = append(result, val)
	return result, nil
}

func doConfigGet(client *etcdClient, root string, keys []string) ([]string, error) {
	var result []string
	for _, k := range keys {
		val, err := client.Get(root + k)
		if err != nil {
			return result, err
		}
		fmt.Printf("%s%s GET: %v\n", root, k, val)
		result = append(result, val)
	}
	return result, nil
}
