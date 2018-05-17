package mocks

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "Decrypter" can be found in crypto
*/
import (
	crypto "crypto"
	io "io"
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"

	testify_assert "github.com/stretchr/testify/assert"
)

//DecrypterMock implements crypto.Decrypter
type DecrypterMock struct {
	t minimock.Tester

	DecryptFunc       func(p io.Reader, p1 []byte, p2 crypto.DecrypterOpts) (r []byte, r1 error)
	DecryptCounter    uint64
	DecryptPreCounter uint64
	DecryptMock       mDecrypterMockDecrypt

	PublicFunc       func() (r crypto.PublicKey)
	PublicCounter    uint64
	PublicPreCounter uint64
	PublicMock       mDecrypterMockPublic
}

//NewDecrypterMock returns a mock for crypto.Decrypter
func NewDecrypterMock(t minimock.Tester) *DecrypterMock {
	m := &DecrypterMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DecryptMock = mDecrypterMockDecrypt{mock: m}
	m.PublicMock = mDecrypterMockPublic{mock: m}

	return m
}

type mDecrypterMockDecrypt struct {
	mock             *DecrypterMock
	mockExpectations *DecrypterMockDecryptParams
}

//DecrypterMockDecryptParams represents input parameters of the Decrypter.Decrypt
type DecrypterMockDecryptParams struct {
	p  io.Reader
	p1 []byte
	p2 crypto.DecrypterOpts
}

//Expect sets up expected params for the Decrypter.Decrypt
func (m *mDecrypterMockDecrypt) Expect(p io.Reader, p1 []byte, p2 crypto.DecrypterOpts) *mDecrypterMockDecrypt {
	m.mockExpectations = &DecrypterMockDecryptParams{p, p1, p2}
	return m
}

//Return sets up a mock for Decrypter.Decrypt to return Return's arguments
func (m *mDecrypterMockDecrypt) Return(r []byte, r1 error) *DecrypterMock {
	m.mock.DecryptFunc = func(p io.Reader, p1 []byte, p2 crypto.DecrypterOpts) ([]byte, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Decrypter.Decrypt method
func (m *mDecrypterMockDecrypt) Set(f func(p io.Reader, p1 []byte, p2 crypto.DecrypterOpts) (r []byte, r1 error)) *DecrypterMock {
	m.mock.DecryptFunc = f
	return m.mock
}

//Decrypt implements crypto.Decrypter interface
func (m *DecrypterMock) Decrypt(p io.Reader, p1 []byte, p2 crypto.DecrypterOpts) (r []byte, r1 error) {
	atomic.AddUint64(&m.DecryptPreCounter, 1)
	defer atomic.AddUint64(&m.DecryptCounter, 1)

	if m.DecryptMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.DecryptMock.mockExpectations, DecrypterMockDecryptParams{p, p1, p2},
			"Decrypter.Decrypt got unexpected parameters")

		if m.DecryptFunc == nil {

			m.t.Fatal("No results are set for the DecrypterMock.Decrypt")

			return
		}
	}

	if m.DecryptFunc == nil {
		m.t.Fatal("Unexpected call to DecrypterMock.Decrypt")
		return
	}

	return m.DecryptFunc(p, p1, p2)
}

//DecryptMinimockCounter returns a count of DecrypterMock.DecryptFunc invocations
func (m *DecrypterMock) DecryptMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.DecryptCounter)
}

//DecryptMinimockPreCounter returns the value of DecrypterMock.Decrypt invocations
func (m *DecrypterMock) DecryptMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.DecryptPreCounter)
}

type mDecrypterMockPublic struct {
	mock *DecrypterMock
}

//Return sets up a mock for Decrypter.Public to return Return's arguments
func (m *mDecrypterMockPublic) Return(r crypto.PublicKey) *DecrypterMock {
	m.mock.PublicFunc = func() crypto.PublicKey {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Decrypter.Public method
func (m *mDecrypterMockPublic) Set(f func() (r crypto.PublicKey)) *DecrypterMock {
	m.mock.PublicFunc = f
	return m.mock
}

//Public implements crypto.Decrypter interface
func (m *DecrypterMock) Public() (r crypto.PublicKey) {
	atomic.AddUint64(&m.PublicPreCounter, 1)
	defer atomic.AddUint64(&m.PublicCounter, 1)

	if m.PublicFunc == nil {
		m.t.Fatal("Unexpected call to DecrypterMock.Public")
		return
	}

	return m.PublicFunc()
}

//PublicMinimockCounter returns a count of DecrypterMock.PublicFunc invocations
func (m *DecrypterMock) PublicMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PublicCounter)
}

//PublicMinimockPreCounter returns the value of DecrypterMock.Public invocations
func (m *DecrypterMock) PublicMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.PublicPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *DecrypterMock) ValidateCallCounters() {

	if m.DecryptFunc != nil && atomic.LoadUint64(&m.DecryptCounter) == 0 {
		m.t.Fatal("Expected call to DecrypterMock.Decrypt")
	}

	if m.PublicFunc != nil && atomic.LoadUint64(&m.PublicCounter) == 0 {
		m.t.Fatal("Expected call to DecrypterMock.Public")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *DecrypterMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *DecrypterMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *DecrypterMock) MinimockFinish() {

	if m.DecryptFunc != nil && atomic.LoadUint64(&m.DecryptCounter) == 0 {
		m.t.Fatal("Expected call to DecrypterMock.Decrypt")
	}

	if m.PublicFunc != nil && atomic.LoadUint64(&m.PublicCounter) == 0 {
		m.t.Fatal("Expected call to DecrypterMock.Public")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *DecrypterMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *DecrypterMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.DecryptFunc == nil || atomic.LoadUint64(&m.DecryptCounter) > 0)
		ok = ok && (m.PublicFunc == nil || atomic.LoadUint64(&m.PublicCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.DecryptFunc != nil && atomic.LoadUint64(&m.DecryptCounter) == 0 {
				m.t.Error("Expected call to DecrypterMock.Decrypt")
			}

			if m.PublicFunc != nil && atomic.LoadUint64(&m.PublicCounter) == 0 {
				m.t.Error("Expected call to DecrypterMock.Public")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *DecrypterMock) AllMocksCalled() bool {

	if m.DecryptFunc != nil && atomic.LoadUint64(&m.DecryptCounter) == 0 {
		return false
	}

	if m.PublicFunc != nil && atomic.LoadUint64(&m.PublicCounter) == 0 {
		return false
	}

	return true
}
