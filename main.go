package main

import "reporter/types"

func main() {
	plugin := Plugin{
		Config:  types.Config{},
		Context: types.DroneContext{},
	}

	plugin.Run()
}
