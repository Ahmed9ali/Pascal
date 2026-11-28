package permission

import (
	"os"
	"strconv"
)

// YoloAllowed reports whether prompt-skipping modes (yolo and sysadmin) have
// been explicitly enabled. They are off unless PASCAL_ALLOW_YOLO=1 is set.
func YoloAllowed() bool {
	v, _ := strconv.ParseBool(os.Getenv("PASCAL_ALLOW_YOLO"))
	return v
}
