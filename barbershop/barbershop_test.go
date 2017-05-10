package barbershop

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestOpenClose(t *testing.T) {
	expectedLogs := "Closing shop\nShop closed\n"

	output := captureOutput(func() {
		shop := New(3, time.Millisecond*100)
		shop.Close()
	})

	assertEqual(t, output, expectedLogs, "Unexpected output")
}

func TestAddBarber(t *testing.T) {
	expectedLogs := "Gerard is sleeping: Zzzzzzzzzzzzz ...\n" +
		"Closing shop\n" +
		"Gerard is going home\n" +
		"Shop closed\n"

	output := captureOutput(func() {
		shop := New(3, time.Millisecond*100)
		shop.AddBarber("Gerard")
		time.Sleep(time.Millisecond * 100)
		shop.Close()
	})

	assertEqual(t, output, expectedLogs, "Unexpected output")
}

func TestIdleBarber(t *testing.T) {
	expectedLogs := "Adding Client 1\n" +
		"Gerard is cutting client 1's hair\n" +
		"Gerard is finished cutting client 1's hair\n" +
		"Gerard is sleeping: Zzzzzzzzzzzzz ...\n" +
		"Closing shop\n" +
		"Gerard is going home\n" +
		"Shop closed\n"

	output := captureOutput(func() {
		shop := New(3, time.Millisecond*170)
		shop.AddClient("1")
		time.Sleep(time.Millisecond * 10)
		shop.AddBarber("Gerard")
		time.Sleep(time.Millisecond * 1000)
		shop.Close()
	})

	assertEqual(t, output, expectedLogs, "Unexpected output")
}

func TestAll(t *testing.T) {
	expectedLogs := "Adding Client 1\n" +
		"Gerard is cutting client 1's hair\n" +
		"Adding Client 2\n" +
		"Gerard is finished cutting client 1's hair\n" +
		"Gerard is cutting client 2's hair\n" +
		"Adding Client 3\n" +
		"Adding Client 4\n" +
		"Gerard is finished cutting client 2's hair\n" +
		"Gerard is cutting client 3's hair\n" +
		"Gerard is finished cutting client 3's hair\n" +
		"Gerard is cutting client 4's hair\n" +
		"Gerard is finished cutting client 4's hair\n" +
		"Gerard is sleeping: Zzzzzzzzzzzzz ...\n" +
		"Closing shop\n" +
		"Gerard is going home\n" +
		"Shop closed\n"

	output := captureOutput(func() {
		shop := New(3, time.Millisecond*170)
		shop.AddBarber("Gerard")
		for i := 1; i < 5; i++ {
			shop.AddClient(strconv.Itoa(i))
			time.Sleep(time.Millisecond * 100)
		}
		time.Sleep(time.Millisecond * 1000)
		shop.Close()
	})

	assertEqual(t, output, expectedLogs, "Unexpected output")
}

// captureOutput capture stdout into a pipe and output a string of the output
func captureOutput(source func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	source()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	return out
}

func assertEqual(t *testing.T, a string, b string, message string) {
	if strings.Compare(a, b) == 0 {
		return
	}
	message = fmt.Sprintf("%v != %v \n%s", a, b, message)
	t.Error(message)
}
