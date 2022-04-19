package utilities

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	ex = filepath.ToSlash(ex)
	file, err := os.OpenFile(path.Join(path.Dir(ex)+"/input/"+filename), os.O_RDWR, 0o660)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteLines(filename string, line string) error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	ex = filepath.ToSlash(ex)
	f, err := os.OpenFile(path.Join(path.Dir(ex)+"/input/"+filename), os.O_RDWR|os.O_APPEND, 0o660)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	_, err2 := f.WriteString(line + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}
	return nil
}

func TruncateLines(filename string, line []string) error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	ex = filepath.ToSlash(ex)
	f, err := os.OpenFile(path.Join(path.Dir(ex)+"/input/"+filename), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	for i := 0; i < len(line); i++ {
		_, err2 := f.WriteString(line[i] + "\n")
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	return nil
}

func ProcessAvatar(av string, memberid string) error {
	if strings.Contains(av, "a_") {
		// Nitro Avatar
		return nil
	}
	link := "https://cdn.discordapp.com/avatars/" + memberid + "/" + av + ".png"
	nameFile := "input/pfps/" + av + ".png"

	err := processFiles(link, nameFile)
	if err != nil {
		return err
	}

	return nil
}

func processFiles(url string, nameFile string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected http status code while downloading avatar%d", resp.StatusCode)
	}
	file, err := os.Create(nameFile)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func Append(filename string, items []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, item := range items {
		if _, err = file.WriteString(item + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func Truncate(filename string, items []string) error {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, item := range items {
		if _, err = file.WriteString(item + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func WriteLine(filename string, line string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(line + "\n"); err != nil {
		return err
	}

	return nil
}

func WriteFile(filename string, items []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, item := range items {
		if _, err = file.WriteString(item + "\n"); err != nil {
			return err
		}
	}

	return nil
}
