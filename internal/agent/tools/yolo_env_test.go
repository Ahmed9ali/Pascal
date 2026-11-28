package tools

import "os"

// Tests in this package exercise yolo mode, which is opt-in in Pascal.
func init() { os.Setenv("PASCAL_ALLOW_YOLO", "1") }
