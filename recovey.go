package rollbar

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/stvp/roll"
)

// Recovery middleware for rollbar crash reporting
func Recovery(token, environment string) gin.HandlerFunc {
	roll.Token = token
	roll.Environment = environment // defaults to "development"

	return func(c *gin.Context) {
		defer func() {
			custom := map[string]string{
				"endpoint": c.Request.RequestURI,
			}

			if rval := recover(); rval != nil {
				debug.PrintStack()

				if _, err := roll.CriticalStack(errors.New(fmt.Sprint(rval)), getCallers(3), custom); err != nil {
					fmt.Fprintln(gin.DefaultErrorWriter, err)
				}

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}

func getCallers(skip int) (pc []uintptr) {
	pc = make([]uintptr, 1000)
	i := runtime.Callers(skip+1, pc)
	return pc[0:i]
}
