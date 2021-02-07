package main

import (
	"bufio"
	"os"
	"time"
)

type TimestampRepo interface {
	LoadAll() ([]time.Time, error)
	AppendOne(time.Time) error
	StoreAll([]time.Time) error
}

type FileTimestampsRepo struct {
	file   *os.File
	format string
}

func NewFileTimestampsRepo(filename string) (*FileTimestampsRepo, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	repo := &FileTimestampsRepo{
		file:   file,
		format: time.RFC3339Nano,
	}
	return repo, nil
}

func (r *FileTimestampsRepo) LoadAll() ([]time.Time, error) {
	stamps := make([]time.Time, 0)
	scanner := bufio.NewScanner(r.file)
	for scanner.Scan() {
		t, err := time.Parse(r.format, scanner.Text())
		if err != nil {
			return nil, err
		}
		stamps = append(stamps, t)
	}
	return stamps, nil
}

func (r *FileTimestampsRepo) AppendOne(t time.Time) error {
	_, err := r.file.WriteString(r.formatTime(t))
	return err
}

func (r *FileTimestampsRepo) StoreAll(times []time.Time) error {
	err := r.file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = r.file.Seek(0, 0)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(r.file)
	for _, v := range times {
		_, err := writer.WriteString(v.Format(r.format) + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func (r *FileTimestampsRepo) formatTime(t time.Time) string {
	return t.Format(r.format) + "\n"
}

func (r *FileTimestampsRepo) Close() error {
	return r.file.Close()
}
