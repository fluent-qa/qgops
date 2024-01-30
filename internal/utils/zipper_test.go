package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnzip(t *testing.T) {
	type args struct {
		zipFile string
		destDir string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unzip(tt.args.zipFile, tt.args.destDir)
			if !tt.wantErr(t, err, fmt.Sprintf("Unzip(%v, %v)", tt.args.zipFile, tt.args.destDir)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Unzip(%v, %v)", tt.args.zipFile, tt.args.destDir)
		})
	}
}

func TestZipFiles(t *testing.T) {

}
