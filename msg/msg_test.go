package msg_test

import (
	"fmt"
	"testing"

	"github.com/projetodorcas/control/msg"
)

func TestIfHiReturnHiDorcas(t *testing.T) {
	msg := msg.Hi()

	if msg != "Hi Dorcas!" {
		t.Log(fmt.Sprintf("Wrong msg, should be %s", "Hi Dorcas!"))
		t.Fail()
	}
}
