package pkg

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ExtractTarGz(tempDir string, gzipStream io.Reader) error {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			dir := filepath.Join(tempDir, header.Name) // #nosec G305
			if err := os.Mkdir(dir, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			file := filepath.Join(tempDir, header.Name) // #nosec G305
			outFile, err := os.Create(file)
			if err != nil {
				return err
			}
			for {
				_, err := io.CopyN(outFile, tarReader, 1024)
				if err != nil {
					if err == io.EOF {
						break
					}
					return err
				}
			}
			outFile.Close()
		default:
			return fmt.Errorf(
				"ExtractTarGz: uknown type: %q in %s",
				header.Typeflag,
				header.Name)
		}
	}
	return nil
}
