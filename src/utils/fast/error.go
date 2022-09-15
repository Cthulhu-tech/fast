package Fast

import (
	"fmt"
	"os"
	"runtime"
)

func errorPrint(_error string) {

	_, file, line, _ := runtime.Caller(2)

	fmt.Fprintf(os.Stdout, "\033[0;31m ERROR! \033[33m File: \033[0m%s \033[34m Type: \033[0m %s \033[35m Line: \033[0m%d", file, _error, line)

}
