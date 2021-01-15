package todo_test

import (
	"github.com/oleg/incubator/go-playground/todo"
	"io/ioutil"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %v got %v instead", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("New task should not be completed %v", l[0])
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	l.Add("New Task")

	err := l.Complete(1)

	if err != nil {
		t.Errorf("Should complete without errors, got %v", err)
	}
	if !l[0].Done {
		t.Errorf("Task should be completed, actual %v", l[0])
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}
	for _, v := range tasks {
		l.Add(v)
	}

	err := l.Delete(2)

	if err != nil {
		t.Errorf("Should delete without error, got %v", err)
	}
	if len(l) != 2 {
		t.Errorf("Should have size %v got %v", 2, len(l))
	}
	if l[1].Task != tasks[2] {
		t.Errorf("Second task should have name %v got %v", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "Save me"
	l1.Add(taskName)

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Failed to create tmp file %v", err)
	}
	defer os.Remove(tf.Name())

	if err = l1.Save(tf.Name()); err != nil {
		t.Fatalf("Failed to Save todo list %v", err)
	}
	if err = l2.Get(tf.Name()); err != nil {
		t.Fatalf("Failed to Get todo list %v", err)
	}
	if len(l1) != len(l2) {
		t.Errorf("Size of loaded todo list %v is not equal to size of saved %v", len(l2), len(l1))
	}
	if l1[0].Task != l2[0].Task ||
		l1[0].Done != l2[0].Done ||
		!l1[0].CreatedAt.Equal(l2[0].CreatedAt) {
		t.Errorf("Loaded todo list %v is not equal to saved %v", l2, l1)
	}
}
