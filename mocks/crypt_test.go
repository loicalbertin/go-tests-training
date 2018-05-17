package mocks

import (
	"crypto"
	"crypto/rand"
	"errors"
	"io"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

type decryptMock struct {
	publicCalled  int
	publicKey     interface{}
	decryptCalled int
	decryptMsg    []byte
	decryptError  error
	decryptResult []byte
}

func (m *decryptMock) Public() crypto.PublicKey {
	m.publicCalled++
	return m.publicKey
}

func (m *decryptMock) Decrypt(rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error) {
	m.decryptCalled++
	return m.decryptResult, m.decryptError
}

func Test_message_decrypt_basic(t *testing.T) {
	type fields struct {
		msg []byte
	}
	type args struct {
		d *decryptMock
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		want              string
		wantErr           bool
		wantPublicCalled  int
		wantDecryptCalled int
	}{
		{"NoError", fields{[]byte("test")}, args{&decryptMock{decryptResult: []byte("result_no_error")}}, "result_no_error", false, 0, 1},
		{"DecyptError", fields{[]byte("test")}, args{&decryptMock{decryptError: errors.New("This is an expected error")}}, "", true, 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{
				msg: tt.fields.msg,
			}
			got, err := m.decrypt(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("message.decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("message.decrypt() = %v, want %v", got, tt.want)
			}
			if tt.args.d.decryptCalled != tt.wantDecryptCalled {
				t.Errorf("message.decrypt() mock decrypt called= %d, want %d", tt.args.d.decryptCalled, tt.wantDecryptCalled)
			}
			if tt.args.d.publicCalled != tt.wantPublicCalled {
				t.Errorf("message.decrypt() mock public called= %d, want %d", tt.args.d.publicCalled, tt.wantPublicCalled)
			}
		})
	}
}

func Test_message_decrypt_minimock(t *testing.T) {
	type fields struct {
		msg []byte
	}
	tests := []struct {
		name              string
		fields            fields
		setup             func(mc *minimock.Controller) *DecrypterMock
		want              string
		wantErr           bool
		wantDecryptCalled uint64
		wantPublicCalled  uint64
	}{
		{
			name:   "NoError",
			fields: fields{[]byte("input")},
			setup: func(mc *minimock.Controller) *DecrypterMock {

				return NewDecrypterMock(mc).DecryptMock.Expect(rand.Reader, []byte("input"), nil).Return([]byte("decrypted"), nil)
			},
			want:              "decrypted",
			wantErr:           false,
			wantDecryptCalled: 1,
			wantPublicCalled:  0,
		},
		{
			name:   "NoError",
			fields: fields{[]byte("input2")},
			setup: func(mc *minimock.Controller) *DecrypterMock {
				return NewDecrypterMock(mc).DecryptMock.Expect(rand.Reader, []byte("input2"), nil).Return(nil, errors.New("expected error"))
			},
			want:              "",
			wantErr:           true,
			wantDecryptCalled: 1,
			wantPublicCalled:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			defer mc.Wait(time.Second)
			d := tt.setup(mc)
			m := &message{
				msg: tt.fields.msg,
			}
			got, err := m.decrypt(d)
			if (err != nil) != tt.wantErr {
				t.Errorf("message.decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("message.decrypt() = %v, want %v", got, tt.want)
			}
			dc := atomic.LoadUint64(&d.DecryptCounter)
			if dc != tt.wantDecryptCalled {
				t.Errorf("message.decrypt() mock decrypt called= %d, want %d", dc, tt.wantDecryptCalled)
			}
			pc := atomic.LoadUint64(&d.PublicCounter)
			if pc != tt.wantPublicCalled {
				t.Errorf("message.decrypt() mock public called= %d, want %d", pc, tt.wantPublicCalled)
			}
		})
	}
}
