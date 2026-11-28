package permission

import "testing"

func TestYoloAllowedRequiresOptIn(t *testing.T) {
	t.Setenv("PASCAL_ALLOW_YOLO", "")
	if YoloAllowed() {
		t.Fatal("yolo must be disabled by default")
	}
	t.Setenv("PASCAL_ALLOW_YOLO", "1")
	if !YoloAllowed() {
		t.Fatal("yolo must be enabled when PASCAL_ALLOW_YOLO=1")
	}
}
