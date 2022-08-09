// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"cmd/asm/internal/lex"
	"cmd/internal/obj"
	"cmd/internal/objabi"
)

// An end-to-end test for the assembler: Do we print what we parse?
// Output is generated by, in effect, turning on -S and comparing the
// result against a golden file.

func testEndToEnd(t *testing.T, goarch, file string) {
	input := filepath.Join("testdata", file+".s")
	architecture, ctxt := setArch(goarch)
	architecture.Init(ctxt)
	lexer := lex.NewLexer(input)
	parser := NewParser(ctxt, architecture, lexer)
	pList := new(obj.Plist)
	var ok bool
	testOut = new(bytes.Buffer) // The assembler writes test output to this buffer.
	ctxt.Bso = bufio.NewWriter(os.Stdout)
	defer ctxt.Bso.Flush()
	failed := false
	ctxt.DiagFunc = func(format string, args ...interface{}) {
		failed = true
		t.Errorf(format, args...)
	}
	pList.Firstpc, ok = parser.Parse()
	if !ok || failed {
		t.Errorf("asm: %s assembly failed", goarch)
		return
	}
	output := strings.Split(testOut.String(), "\n")

	// Reconstruct expected output by independently "parsing" the input.
	data, err := ioutil.ReadFile(input)
	if err != nil {
		t.Error(err)
		return
	}
	lineno := 0
	seq := 0
	hexByLine := map[string]string{}
	lines := strings.SplitAfter(string(data), "\n")
Diff:
	for _, line := range lines {
		lineno++

		// Ignore include of textflag.h.
		if strings.HasPrefix(line, "#include ") {
			continue
		}

		// The general form of a test input line is:
		//	// comment
		//	INST args [// printed form] [// hex encoding]
		parts := strings.Split(line, "//")
		printed := strings.TrimSpace(parts[0])
		if printed == "" || strings.HasSuffix(printed, ":") { // empty or label
			continue
		}
		seq++

		var hexes string
		switch len(parts) {
		default:
			t.Errorf("%s:%d: unable to understand comments: %s", input, lineno, line)
		case 1:
			// no comment
		case 2:
			// might be printed form or hex
			note := strings.TrimSpace(parts[1])
			if isHexes(note) {
				hexes = note
			} else {
				printed = note
			}
		case 3:
			// printed form, then hex
			printed = strings.TrimSpace(parts[1])
			hexes = strings.TrimSpace(parts[2])
			if !isHexes(hexes) {
				t.Errorf("%s:%d: malformed hex instruction encoding: %s", input, lineno, line)
			}
		}

		if hexes != "" {
			hexByLine[fmt.Sprintf("%s:%d", input, lineno)] = hexes
		}

		// Canonicalize spacing in printed form.
		// First field is opcode, then tab, then arguments separated by spaces.
		// Canonicalize spaces after commas first.
		// Comma to separate argument gets a space; comma within does not.
		var buf []byte
		nest := 0
		for i := 0; i < len(printed); i++ {
			c := printed[i]
			switch c {
			case '{', '[':
				nest++
			case '}', ']':
				nest--
			case ',':
				buf = append(buf, ',')
				if nest == 0 {
					buf = append(buf, ' ')
				}
				for i+1 < len(printed) && (printed[i+1] == ' ' || printed[i+1] == '\t') {
					i++
				}
				continue
			}
			buf = append(buf, c)
		}

		f := strings.Fields(string(buf))

		// Turn relative (PC) into absolute (PC) automatically,
		// so that most branch instructions don't need comments
		// giving the absolute form.
		if len(f) > 0 && strings.HasSuffix(printed, "(PC)") {
			last := f[len(f)-1]
			n, err := strconv.Atoi(last[:len(last)-len("(PC)")])
			if err == nil {
				f[len(f)-1] = fmt.Sprintf("%d(PC)", seq+n)
			}
		}

		if len(f) == 1 {
			printed = f[0]
		} else {
			printed = f[0] + "\t" + strings.Join(f[1:], " ")
		}

		want := fmt.Sprintf("%05d (%s:%d)\t%s", seq, input, lineno, printed)
		for len(output) > 0 && (output[0] < want || output[0] != want && len(output[0]) >= 5 && output[0][:5] == want[:5]) {
			if len(output[0]) >= 5 && output[0][:5] == want[:5] {
				t.Errorf("mismatched output:\nhave %s\nwant %s", output[0], want)
				output = output[1:]
				continue Diff
			}
			t.Errorf("unexpected output: %q", output[0])
			output = output[1:]
		}
		if len(output) > 0 && output[0] == want {
			output = output[1:]
		} else {
			t.Errorf("missing output: %q", want)
		}
	}
	for len(output) > 0 {
		if output[0] == "" {
			// spurious blank caused by Split on "\n"
			output = output[1:]
			continue
		}
		t.Errorf("unexpected output: %q", output[0])
		output = output[1:]
	}

	// Checked printing.
	// Now check machine code layout.

	top := pList.Firstpc
	var text *obj.LSym
	ok = true
	ctxt.DiagFunc = func(format string, args ...interface{}) {
		t.Errorf(format, args...)
		ok = false
	}
	obj.Flushplist(ctxt, pList, nil, "")

	for p := top; p != nil; p = p.Link {
		if p.As == obj.ATEXT {
			text = p.From.Sym
		}
		hexes := hexByLine[p.Line()]
		if hexes == "" {
			continue
		}
		delete(hexByLine, p.Line())
		if text == nil {
			t.Errorf("%s: instruction outside TEXT", p)
		}
		size := int64(len(text.P)) - p.Pc
		if p.Link != nil {
			size = p.Link.Pc - p.Pc
		} else if p.Isize != 0 {
			size = int64(p.Isize)
		}
		var code []byte
		if p.Pc < int64(len(text.P)) {
			code = text.P[p.Pc:]
			if size < int64(len(code)) {
				code = code[:size]
			}
		}
		codeHex := fmt.Sprintf("%x", code)
		if codeHex == "" {
			codeHex = "empty"
		}
		ok := false
		for _, hex := range strings.Split(hexes, " or ") {
			if codeHex == hex {
				ok = true
				break
			}
		}
		if !ok {
			t.Errorf("%s: have encoding %s, want %s", p, codeHex, hexes)
		}
	}

	if len(hexByLine) > 0 {
		var missing []string
		for key := range hexByLine {
			missing = append(missing, key)
		}
		sort.Strings(missing)
		for _, line := range missing {
			t.Errorf("%s: did not find instruction encoding", line)
		}
	}

}

func isHexes(s string) bool {
	if s == "" {
		return false
	}
	if s == "empty" {
		return true
	}
	for _, f := range strings.Split(s, " or ") {
		if f == "" || len(f)%2 != 0 || strings.TrimLeft(f, "0123456789abcdef") != "" {
			return false
		}
	}
	return true
}

// It would be nice if the error messages began with
// the standard file:line: prefix,
// but that's not where we are today.
// It might be at the beginning but it might be in the middle of the printed instruction.
var fileLineRE = regexp.MustCompile(`(?:^|\()(testdata[/\\][0-9a-z]+\.s:[0-9]+)(?:$|\))`)

// Same as in test/run.go
var (
	errRE       = regexp.MustCompile(`// ERROR ?(.*)`)
	errQuotesRE = regexp.MustCompile(`"([^"]*)"`)
)

func testErrors(t *testing.T, goarch, file string) {
	input := filepath.Join("testdata", file+".s")
	architecture, ctxt := setArch(goarch)
	lexer := lex.NewLexer(input)
	parser := NewParser(ctxt, architecture, lexer)
	pList := new(obj.Plist)
	var ok bool
	testOut = new(bytes.Buffer) // The assembler writes test output to this buffer.
	ctxt.Bso = bufio.NewWriter(os.Stdout)
	defer ctxt.Bso.Flush()
	failed := false
	var errBuf bytes.Buffer
	ctxt.DiagFunc = func(format string, args ...interface{}) {
		failed = true
		s := fmt.Sprintf(format, args...)
		if !strings.HasSuffix(s, "\n") {
			s += "\n"
		}
		errBuf.WriteString(s)
	}
	pList.Firstpc, ok = parser.Parse()
	obj.Flushplist(ctxt, pList, nil, "")
	if ok && !failed {
		t.Errorf("asm: %s had no errors", goarch)
	}

	errors := map[string]string{}
	for _, line := range strings.Split(errBuf.String(), "\n") {
		if line == "" || strings.HasPrefix(line, "\t") {
			continue
		}
		m := fileLineRE.FindStringSubmatch(line)
		if m == nil {
			t.Errorf("unexpected error: %v", line)
			continue
		}
		fileline := m[1]
		if errors[fileline] != "" && errors[fileline] != line {
			t.Errorf("multiple errors on %s:\n\t%s\n\t%s", fileline, errors[fileline], line)
			continue
		}
		errors[fileline] = line
	}

	// Reconstruct expected errors by independently "parsing" the input.
	data, err := ioutil.ReadFile(input)
	if err != nil {
		t.Error(err)
		return
	}
	lineno := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		lineno++

		fileline := fmt.Sprintf("%s:%d", input, lineno)
		if m := errRE.FindStringSubmatch(line); m != nil {
			all := m[1]
			mm := errQuotesRE.FindAllStringSubmatch(all, -1)
			if len(mm) != 1 {
				t.Errorf("%s: invalid errorcheck line:\n%s", fileline, line)
			} else if err := errors[fileline]; err == "" {
				t.Errorf("%s: missing error, want %s", fileline, all)
			} else if !strings.Contains(err, mm[0][1]) {
				t.Errorf("%s: wrong error for %s:\n%s", fileline, all, err)
			}
		} else {
			if errors[fileline] != "" {
				t.Errorf("unexpected error on %s: %v", fileline, errors[fileline])
			}
		}
		delete(errors, fileline)
	}
	var extra []string
	for key := range errors {
		extra = append(extra, key)
	}
	sort.Strings(extra)
	for _, fileline := range extra {
		t.Errorf("unexpected error on %s: %v", fileline, errors[fileline])
	}
}

func Test386EndToEnd(t *testing.T) {
	defer func(old string) { objabi.GO386 = old }(objabi.GO386)
	for _, go386 := range []string{"387", "sse2"} {
		t.Logf("GO386=%v", go386)
		objabi.GO386 = go386
		testEndToEnd(t, "386", "386")
	}
}

func TestARMEndToEnd(t *testing.T) {
	defer func(old int) { objabi.GOARM = old }(objabi.GOARM)
	for _, goarm := range []int{5, 6, 7} {
		t.Logf("GOARM=%d", goarm)
		objabi.GOARM = goarm
		testEndToEnd(t, "arm", "arm")
		if goarm == 6 {
			testEndToEnd(t, "arm", "armv6")
		}
	}
}

func TestARMErrors(t *testing.T) {
	testErrors(t, "arm", "armerror")
}

func TestARM64EndToEnd(t *testing.T) {
	testEndToEnd(t, "arm64", "arm64")
}

func TestARM64Encoder(t *testing.T) {
	testEndToEnd(t, "arm64", "arm64enc")
}

func TestARM64Errors(t *testing.T) {
	testErrors(t, "arm64", "arm64error")
}

func TestAMD64EndToEnd(t *testing.T) {
	defer func(old string) { objabi.GOAMD64 = old }(objabi.GOAMD64)
	for _, goamd64 := range []string{"normaljumps", "alignedjumps"} {
		t.Logf("GOAMD64=%s", goamd64)
		objabi.GOAMD64 = goamd64
		testEndToEnd(t, "amd64", "amd64")
	}
}

func Test386Encoder(t *testing.T) {
	testEndToEnd(t, "386", "386enc")
}

func TestAMD64Encoder(t *testing.T) {
	filenames := [...]string{
		"amd64enc",
		"amd64enc_extra",
		"avx512enc/aes_avx512f",
		"avx512enc/gfni_avx512f",
		"avx512enc/vpclmulqdq_avx512f",
		"avx512enc/avx512bw",
		"avx512enc/avx512cd",
		"avx512enc/avx512dq",
		"avx512enc/avx512er",
		"avx512enc/avx512f",
		"avx512enc/avx512pf",
		"avx512enc/avx512_4fmaps",
		"avx512enc/avx512_4vnniw",
		"avx512enc/avx512_bitalg",
		"avx512enc/avx512_ifma",
		"avx512enc/avx512_vbmi",
		"avx512enc/avx512_vbmi2",
		"avx512enc/avx512_vnni",
		"avx512enc/avx512_vpopcntdq",
	}
	for _, name := range filenames {
		testEndToEnd(t, "amd64", name)
	}
}

func TestAMD64Errors(t *testing.T) {
	testErrors(t, "amd64", "amd64error")
}

func TestMIPSEndToEnd(t *testing.T) {
	testEndToEnd(t, "mips", "mips")
	testEndToEnd(t, "mips64", "mips64")
}

func TestLOONG64Encoder(t *testing.T) {
	testEndToEnd(t, "loong64", "loong64enc1")
	testEndToEnd(t, "loong64", "loong64enc2")
	testEndToEnd(t, "loong64", "loong64enc3")
	testEndToEnd(t, "loong64", "loong64")
}

func TestPPC64EndToEnd(t *testing.T) {
	testEndToEnd(t, "ppc64", "ppc64")
}

func TestPPC64Encoder(t *testing.T) {
	testEndToEnd(t, "ppc64", "ppc64enc")
}

func TestRISCVEncoder(t *testing.T) {
	testEndToEnd(t, "riscv64", "riscvenc")
}

func TestS390XEndToEnd(t *testing.T) {
	testEndToEnd(t, "s390x", "s390x")
}
