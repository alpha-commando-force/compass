// package impl get, insert, delete operations for log.
package log

import (
	"os"
	"testing"
)

func TestNewLogReadWriter(t *testing.T) {
	t.Run("noraml", func(t *testing.T) {
		if _, err := NewReadWriter("./foobar"); err != nil {
			t.Error("NewLogReadWriter() get err =", err)
			return
		}

		if err := os.RemoveAll("./foobar"); err != nil {
			t.Error("delete dir err", err)
		}
	})
}

func Test_createDirIfNotExist(t *testing.T) {
	tests := []struct {
		name    string
		pre     func(t *testing.T)
		post    func(t *testing.T) // Verification and post-processing.
		path    string
		wantErr bool
	}{
		{
			name: "no exist",
			path: "./foobar",
			post: func(t *testing.T) {
				t.Helper()

				s, err := os.Stat("./foobar")
				if err != nil {
					t.Error("vierfy resutlt error", err)
					return
				}

				if !s.IsDir() {
					t.Error("verify result error, not dir")
					return
				}

				if err := os.Remove("./foobar"); err != nil {
					t.Fatal(err)
				}
			},
			wantErr: false,
		},
		{
			name: "already exist",
			path: "./foobar",
			pre: func(t *testing.T) {
				if err := os.Mkdir("./foobar", os.ModePerm); err != nil {
					t.Fatal("prepare error", err)
				}
			},
			post: func(t *testing.T) {
				t.Helper()

				s, err := os.Stat("./foobar")
				if err != nil {
					t.Error("vierfy resutlt error", err)
				}

				if !s.IsDir() {
					t.Error("verify result error, not dir")
				}

				if err := os.Remove("./foobar"); err != nil {
					t.Fatal(err)
				}
			},
			wantErr: false,
		},
		{
			name: "already exist as file",
			path: "./foobar",
			pre: func(t *testing.T) {
				if _, err := os.Create("./foobar"); err != nil {
					t.Fatal("prepare error", err)
				}
			},
			post: func(t *testing.T) {
				t.Helper()

				if err := os.Remove("./foobar"); err != nil {
					t.Fatal(err)
				}
			},
			wantErr: true,
		},
		{
			name:    "error path",
			path:    "/etc/hosts",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pre != nil {
				tt.pre(t)
			}

			if err := createDirIfNotExist(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("createDirIfNotExist() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.post != nil {
				tt.post(t)
			}
		})
	}
}
