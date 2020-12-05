package Logger

import "log"

func LogAllFatal(errors []error){
	for _, v := range errors{
		if v != nil {
			log.Fatal(v)
		}
	}
}
